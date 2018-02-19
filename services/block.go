package services

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
)

type BlockService struct {
	dao daos.BlockDAO
}

func NewBlockService(dao daos.BlockDAO) *BlockService {
	return &BlockService{dao}
}

func (s *BlockService) Get(blockNumber int64) (models.Header, error) {
	blockNumberBig := big.NewInt(blockNumber)
	block, err := s.dao.Get(blockNumberBig)
	if err != nil {
		return models.Header{}, err
	}
	blockJSON := models.NewBlockHeader(block)
	return blockJSON, nil
}

// This is for getting recent blocks, so the cursor is really just the most recent block to start from
func (s *BlockService) List(limit int, cursor string) (models.BlockPage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	return s.dao.Find(bson.M{}, limit, cursor)
}

func (s *BlockService) Miner(address common.Address, limit int, cursor string) (models.BlockPage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	if address.String() == "0x0000000000000000000000000000000000000000" || len(address.String()) != 42 {
		return models.BlockPage{}, fmt.Errorf("Invalid miner address %s", address.String())
	}
	return s.dao.Find(bson.M{"miner": strings.ToLower(address.String())}, limit, cursor)
}
