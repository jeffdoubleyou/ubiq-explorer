package main

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"time"
	"ubiq-explorer/daemon/core"
	"ubiq-explorer/daemon/tokens"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
	"ubiq-explorer/util"
)

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
	balanceService := services.NewBalanceService(*balanceDAO)
	statsService := services.NewStatsService()

	statsSmoothing, err := beego.AppConfig.Int("stats::smoothing")

	if err != nil {
		panic(err)
	}

	currentBlock, err := blocks.GetCurrentBlock()
	if err != nil {
		panic(err)
	}
	lastBlock, err := blockDAO.LastImportedBlock()
	if err != nil {
		//panic(err)
	}
	log.Printf("Current block: %d - Last block imported: %d", currentBlock, lastBlock.Int64())

	run := 1
	pending := 0

	// Commence spaghetti
	for run == 1 {
		for lastBlock.Cmp(currentBlock) < 1 {
			//for lastBlock.Int64() <= currentBlock.Int64() {
			log.Printf("Going to get block #%d out of %d", lastBlock, currentBlock)
			_, err := blocks.GetBlock(lastBlock)
			if err != nil {
				log.Fatalf("%s\n", err)
				run = 0
			}
			miner := blocks.Miner()
			_, err = minerDAO.Insert(*miner)
			if err != nil {
				panic(err)
			}
			balance, err := blocks.Balance(common.HexToAddress(miner.Miner), lastBlock)
			if err != nil {
				panic(err)
			}
			balance.ChangedBy = "Mined Block"
			_, err = balanceDAO.Insert(*balance)
			if err != nil {
				panic(err)
			}
			_, err = balanceService.UpdateCurrentBalance(balance.Address, balance.Balance)
			if err != nil {
				panic(err)
			}
			uncles := blocks.Uncles()
			for _, u := range uncles {
				_, err = uncleDAO.Insert(*u)
				if err != nil {
					panic(err)
				}
				balance, err = blocks.Balance(common.HexToAddress(u.Miner), lastBlock)
				if err != nil {
					panic(err)
				}
				balance.ChangedBy = "Mined Uncle"
				_, err = balanceDAO.Insert(*balance)
				if err != nil {
					panic(err)
				}
				_, err = balanceService.UpdateCurrentBalance(balance.Address, balance.Balance)
				if err != nil {
					panic(err)
				}
			}
			for _, t := range blocks.Transactions() {
				_, err = transactionDAO.Insert(*t)
				if err != nil {
					panic(err)
				}

				if t.To.String() != "0x0000000000000000000000000000000000000000" {
					balance, err = blocks.Balance(t.To, lastBlock)
					if err != nil {
						panic(err)
					}
					balance.ChangedBy = "Incoming Transaction"
					_, err = balanceDAO.Insert(*balance)
					if err != nil {
						panic(err)
					}
					_, err = balanceService.UpdateCurrentBalance(balance.Address, balance.Balance)
					if err != nil {
						panic(err)
					}
				}

				balance, err = blocks.Balance(t.From, lastBlock)
				if err != nil {
					panic(err)
				}

				if t.To.String() == "0x0000000000000000000000000000000000000000" {
					balance.ChangedBy = "Contract Execution"
				} else {
					balance.ChangedBy = "Outgoing Transaction"
				}
				_, err = balanceDAO.Insert(*balance)
				if err != nil {
					panic(err)
				}
				_, err = balanceService.UpdateCurrentBalance(balance.Address, balance.Balance)
				if err != nil {
					panic(err)
				}

				token, isNewToken, _ := tokenTool.GetTokenInfo(t.To)
				if token != nil {
					if isNewToken == true {
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
					tokenTransactions, err := tokenTool.GetTransactionInfo(t)
					if err != nil {
						log.Fatalf("Failed to retrieve token transactions: %s", err)
						run = 0
					}
					var tokenAddresses = make(map[string]common.Address)
					for _, tokenTransaction := range tokenTransactions {
						_, err := tokenDAO.InsertTokenTransaction(*tokenTransaction)
						if err != nil {
							log.Fatalf("Failed to insert token transaction: %s", err)
							run = 0
						}
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

			lastBlock = lastBlock.Add(lastBlock, big.NewInt(1))
		}
		if pending == 0 {
			log.Printf("Current Block: %d Pending Block: %d", currentBlock, lastBlock)
			pending = 1
		}
		time.Sleep(2000 * time.Millisecond)
		currentBlock, err = blocks.GetCurrentBlock()
	}
}
