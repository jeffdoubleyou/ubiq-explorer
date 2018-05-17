package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"
	"ubiq-explorer/daemon/core"
	"ubiq-explorer/daemon/tokens"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
	"ubiq-explorer/util"
)

type balanceChange struct{ Incoming, Outgoing, Contract, Blocks, Uncles int }

func main() {
	_, err := util.InitializeMongoDB()
	if err != nil {
		panic(err)
	}
	tokenTool := tokens.NewTokenUtils(nil)
	blocks := core.NewBlocks()

	minerDAO := daos.NewMinerDAO()
	uncleDAO := daos.NewUncleDAO()
	blockDAO := daos.NewBlockDAO()
	transactionDAO := daos.NewTransactionDAO()
	tokenDAO := daos.NewTokenDAO()
	balanceDAO := daos.NewBalanceDAO()
	tokenBalanceDAO := daos.NewTokenBalanceDAO()
	historyDAO := daos.NewHistoryDAO()
	addressDAO := daos.NewAddressDAO()

	tokenBalanceService := services.NewTokenBalanceService(*tokenBalanceDAO)
	tokenService := services.NewTokenService(*tokenDAO)
	balanceService := services.NewBalanceService(*balanceDAO)
	statsService := services.NewStatsService()

	statsSmoothing, err := beego.AppConfig.Int("stats::smoothing")
	statsWindow, err := beego.AppConfig.Int64("stats::history_window")

	if err != nil {
		panic(err)
	}

	currentBlock, err := blocks.GetCurrentBlock()
	if err != nil {
		panic(err)
	}
	lastBlock, err := blockDAO.LastImportedBlock()
	if err != nil {
		panic(err)
	}
	log.Printf("Current block: %d - Last block imported: %d", currentBlock, lastBlock.Int64())

	run := 1
	syncing := 1
	inStatsWindow := big.NewInt(0)
	importCount := float64(0)
	importSeconds := float64(0)
	remaining := float64(0)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signals
		log.Printf("Caught signal %s - shutting down", sig)
		run = 0
	}()

	// Commence spaghetti
	for run == 1 {
		for lastBlock.Cmp(currentBlock) < 1 {
			start := time.Now()
			log.Printf("Going to get block #%d out of %d - %f hours remaining until complete sync", lastBlock, currentBlock, remaining)
			_, err := blocks.GetBlock(lastBlock)
			if err != nil {
				log.Fatalf("%s\n", err)
				run = 0
			}
			var balance *models.Balance
			var balanceChanges = make(map[string]*balanceChange)
			miner := blocks.Miner()
			_, err = minerDAO.Insert(*miner)
			if err != nil {
				panic(err)
			}
			if _, ok := balanceChanges[miner.Miner]; !ok {
				balanceChanges[miner.Miner] = &balanceChange{0, 0, 0, 0, 0}
			}
			balanceChanges[miner.Miner].Blocks++

			uncles := blocks.Uncles()
			for _, u := range uncles {
				_, err = uncleDAO.Insert(*u)
				if err != nil {
					panic(err)
				}
				if _, ok := balanceChanges[u.Miner]; !ok {
					balanceChanges[u.Miner] = &balanceChange{0, 0, 0, 0, 0}
				}
				balanceChanges[u.Miner].Uncles++
			}

			for _, t := range blocks.Transactions() {
				go func() {
					_, err = transactionDAO.Insert(*t)
					if err != nil {
						panic(err)
					}
				}()

				if t.To.String() != "0x0000000000000000000000000000000000000000" {
					if _, ok := balanceChanges[t.To.String()]; !ok {
						balanceChanges[t.To.String()] = &balanceChange{0, 0, 0, 0, 0}
					}
					balanceChanges[t.To.String()].Incoming++
				}

				if _, ok := balanceChanges[t.From.String()]; !ok {
					balanceChanges[t.From.String()] = &balanceChange{0, 0, 0, 0, 0}
				}

				if t.To.String() == "0x0000000000000000000000000000000000000000" {
					balanceChanges[t.From.String()].Contract++
				} else {
					balanceChanges[t.From.String()].Outgoing++
				}

				token, isNewToken, _ := tokenTool.GetTokenInfo(t.To)
				if token != nil {
					if isNewToken == true {
						tokenInfo, _ := tokenService.GetTokenByAddress(token.Address)
						if tokenInfo.Name == "" {
							_, err := tokenDAO.Insert(*token)
							if err != nil {
								log.Fatalf("Failed to insert new token: %s", err)
								run = 0
							}
							// Add a known address for this token
							_, err = addressDAO.Insert(models.AddressInfo{t.To, token.Name})
							if err != nil {
								log.Fatalf("Failed to insert new token address info: %s", err)
								run = 0
							}
						}
					}
					tokenTransactions, err := tokenTool.GetTransactionInfo(t)
					if err != nil {
						log.Fatalf("Failed to retrieve token transactions: %s", err)
						run = 0
					}
					var tokenAddresses = make(map[string]common.Address)
					for _, tokenTransaction := range tokenTransactions {
						go func() {
							_, err := tokenDAO.InsertTokenTransaction(*tokenTransaction)
							if err != nil {
								log.Fatalf("Failed to insert token transaction: %s", err)
								run = 0
							}
						}()
						if _, e := tokenAddresses[tokenTransaction.From.String()]; !e {
							tokenAddresses[tokenTransaction.From.String()] = tokenTransaction.From
						}
						if _, e := tokenAddresses[tokenTransaction.To.String()]; !e {
							tokenAddresses[tokenTransaction.To.String()] = tokenTransaction.To
						}
					}
					if _, e := tokenAddresses[t.To.String()]; !e {
						tokenAddresses[t.To.String()] = t.To
					}
					for tokenAddress := range tokenAddresses {
						tokenBalance, err := tokenTool.GetTokenBalance(t.To, tokenAddresses[tokenAddress])
						if err != nil {
							log.Fatalf("Failed to get token balance for %s : Error: %s", tokenAddresses[tokenAddress].String(), err)
							run = 0
						}
						balance := &models.TokenBalance{
							common.HexToAddress(tokenAddress),
							t.To,
							token.Name,
							token.Symbol,
							tokenBalance,
						}

						_, err = tokenBalanceService.Update(balance)
						if err != nil {
							log.Fatalf("Unable to update token balance: %s", err)
							run = 0
						}
					}
				}
			}
			for a, b := range balanceChanges {
				change := fmt.Sprintf("Changes at block: %d:", lastBlock)
				if b.Blocks > 0 {
					change = change + fmt.Sprintf(" %c Mined block", 0x1F528)
				}
				if b.Uncles > 0 {
					change = change + fmt.Sprintf(" %c Mined uncle", 0x1F528)
				}
				if b.Incoming > 0 {
					change = change + fmt.Sprintf(" %c %d Incoming transactions", 0x2199, b.Incoming)
				}
				if b.Outgoing > 0 {
					change = change + fmt.Sprintf(" %c %d Outgoing transactions", 0x2198, b.Outgoing)
				}
				if b.Contract > 0 {
					change = change + fmt.Sprintf(" %c %d Contract executions", 0x2699, b.Contract)
				}
				balance, err = blocks.Balance(common.HexToAddress(a), lastBlock)
				if err != nil {
					panic(err)
				}
				balance.ChangedBy = change
				_, err = balanceDAO.Insert(*balance)
				if err != nil {
					panic(err)
				}
				_, err = balanceService.UpdateCurrentBalance(balance.Address, balance.Balance)
			}
			inStatsWindow = inStatsWindow.Add(lastBlock, big.NewInt(statsWindow))

			if syncing == 0 || inStatsWindow.Cmp(currentBlock) == 1 {
				log.Printf("Getting stats at block %d - window is %d", lastBlock, inStatsWindow)
				stats, err := statsService.Get(statsSmoothing)

				if err != nil {
					log.Fatalf("Could not generate statistics: %s", err)
					run = 0
				}

				if stats != nil {
					historyDAO.Insert("blockTimeHistory", struct{ Value float64 }{stats.BlockTime})
					historyDAO.Insert("uncleRateHistory", struct{ Value float64 }{stats.UncleRate})
					historyDAO.Insert("difficultyHistory", struct{ Value string }{stats.Difficulty})
					historyDAO.Insert("hashRateHistory", struct{ Value string }{stats.HashRate})
				}
			}

			lastBlock = lastBlock.Add(lastBlock, big.NewInt(1))
			end := time.Now()
			elapsed := end.Sub(start)
			importCount++
			importSeconds += elapsed.Seconds()
			remaining = float64(currentBlock.Int64()-lastBlock.Int64()) * (importSeconds / importCount) / 60 / 60
			if run == 0 {
				log.Println("Exiting because we have been asked to stop running")
				os.Exit(1)
			}
		}
		if syncing == 1 {
			log.Printf("Current Block: %d Pending Block: %d", currentBlock, lastBlock)
			syncing = 0
		}
		time.Sleep(2000 * time.Millisecond)
		currentBlock, err = blocks.GetCurrentBlock()
	}
}
