package services

import (
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
)

type StatsService struct {
	blocks       *daos.BlockDAO
	uncles       *daos.UncleDAO
	transactions *daos.TransactionDAO
	history      *daos.HistoryDAO
	address      *daos.AddressDAO
}

func NewStatsService() *StatsService {
	return &StatsService{
		daos.NewBlockDAO(),
		daos.NewUncleDAO(),
		daos.NewTransactionDAO(),
		daos.NewHistoryDAO(),
		daos.NewAddressDAO(),
	}
}

func (s *StatsService) Get(blocks int) (*models.Stats, error) {
	if blocks > 500 {
		blocks = 500
	}
	blockList, err := s.blocks.Find(bson.M{}, blocks, "")
	if err != nil {
		return nil, err
	}
	if len(blockList.Blocks) == blocks {
		lastBlock := blockList.Blocks[0].Block
		firstBlock := blockList.Blocks[len(blockList.Blocks)-1].Block
		totalBlockTime := blockList.Blocks[0].Timestamp - blockList.Blocks[len(blockList.Blocks)-1].Timestamp
		totalDifficulty := big.NewFloat(0)
		for _, block := range blockList.Blocks {
			totalDifficulty = totalDifficulty.Add(totalDifficulty, big.NewFloat(0).SetInt64(block.Difficulty.Int64()))
		}

		avgBlockTime := float64(totalBlockTime) / float64(blocks-1)
		var avgDifficulty, blocksFloat big.Float
		blocksFloat.SetFloat64(float64(blocks - 1))
		avgDifficulty.Quo(totalDifficulty, &blocksFloat)
		uncles, err := s.uncles.Count(bson.M{"block": bson.M{"$gte": firstBlock.Uint64()}})
		if err != nil {
			return nil, err
		}
		uncleRate := float64(0)
		if uncles > 0 {
			uncleRate = float64(uncles) / float64(blocks) * 100
		}
		var hashRate big.Float
		hashRate.Quo(&avgDifficulty, big.NewFloat(avgBlockTime))
		stats := &models.Stats{
			LastBlock:  lastBlock,
			Difficulty: avgDifficulty.Text('f', 16),
			HashRate:   hashRate.Text('f', 16),
			UncleRate:  uncleRate,
			BlockTime:  avgBlockTime,
		}
		return stats, nil
	} else {
		// Because we might not have enough blocks yet
		if len(blockList.Blocks) > 0 {
			return nil, nil
		}
		return nil, err
	}
}

func (s *StatsService) HashRateHistory() ([]interface{}, error) {
	return s.history.Find("hashRate")
}

func (s *StatsService) UncleRateHistory() ([]interface{}, error) {
	return s.history.Find("uncleRate")
}

func (s *StatsService) BlockTimeHistory() ([]interface{}, error) {
	return s.history.Find("blockTime")
}

func (s *StatsService) DifficultyHistory() ([]interface{}, error) {
	return s.history.Find("difficulty")
}

func (s *StatsService) MinerList(blocks int) (map[string]*models.MinerList, error) {
	addressList, err := s.address.Find(bson.M{})
	if err != nil {
		return nil, err
	}
	addresses := make(map[string]string)
	for _, address := range addressList {
		addresses[strings.ToLower(address.Address.String())] = address.Name
	}
	blockList, err := s.blocks.Find(bson.M{}, blocks, "")
	if err != nil {
		return nil, err
	}
	minedBlocks := make(map[string]*models.MinerList)
	for _, miner := range blockList.Blocks {
		name := addresses[miner.Miner]
		if name == "" {
			name = "Unknown"
		}
		if minedBlocks[name] == nil {
			minedBlocks[name] = &models.MinerList{
				Name:    name,
				Address: miner.Miner,
				Count:   1,
			}
		} else {
			minedBlocks[name].Count++
		}
	}
	for _, miner := range minedBlocks {
		miner.Percent = float64(miner.Count) / float64(blocks) * 100
	}
	return minedBlocks, nil
}
