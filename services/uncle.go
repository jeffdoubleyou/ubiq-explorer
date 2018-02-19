package services

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	//	"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
)

type UncleService struct {
	dao daos.UncleDAO
}

func NewUncleService(dao daos.UncleDAO) *UncleService {
	return &UncleService{dao}
}

// This is for getting recent blocks, so the cursor is really just the most recent block to start from
func (s *UncleService) List(limit int, cursor string) (models.UnclePage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	return s.dao.Find(bson.M{}, limit, cursor)
}

func (s *UncleService) Miner(address common.Address, limit int, cursor string) (models.UnclePage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	if address.String() == "0x0000000000000000000000000000000000000000" || len(address.String()) != 42 {
		return models.UnclePage{}, fmt.Errorf("Invalid miner address %s", address.String())
	}
	return s.dao.Find(bson.M{"miner": strings.ToLower(address.String())}, limit, cursor)
}

func (s *UncleService) Block(block big.Int, limit int, cursor string) (models.UnclePage, error) {
	if limit > 1000 {
		limit = 1000
	}
	return s.dao.Find(bson.M{"block": block.Uint64()}, limit, cursor)
}
