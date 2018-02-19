package services

import (
	"github.com/ethereum/go-ethereum/common"
	//	"github.com/ethereum/go-ethereum/common/hexutil"
	//	"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
)

/*
type TransactionDAO interface {
	From(address string, limit int, cursor string) (*models.TransactionPage, error)
	To(address string, limit int, cursor string) (*models.TransactionPage, error)
	Block(address string, limit int, cursor string) (*models.TransactionPage, error)
}*/

type TransactionService struct {
	dao daos.TransactionDAO
}

func NewTransactionService(dao daos.TransactionDAO) *TransactionService {
	return &TransactionService{dao}
}

func (s *TransactionService) From(address common.Address, limit int, cursor string) (models.TransactionPage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	return s.dao.Find(bson.M{"from": strings.ToLower(address.String())}, "-_id", limit, cursor)
}

func (s *TransactionService) To(address common.Address, limit int, cursor string) (models.TransactionPage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	return s.dao.Find(bson.M{"to": strings.ToLower(address.String())}, "-_id", limit, cursor)
}

func (s *TransactionService) List(limit int, cursor string) (models.TransactionPage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	return s.dao.Find(bson.M{}, "-_id", limit, cursor)
}

// Just putting limit of 1000 here, there should not be that many transactions
func (s *TransactionService) Block(block big.Int, limit int, cursor string) (models.TransactionPage, error) {
	if limit > 1000 {
		limit = 1000
	}
	return s.dao.Find(bson.M{"number": block.String()}, "-_id", limit, cursor)
}

func (s *TransactionService) Get(hashString string) (*models.RpcTransaction, error) {
	hash := common.HexToHash(hashString)
	t, err := s.dao.GetFromRPC(hash)
	if err != nil {
		return nil, err
	}
	t.FormatJSON()
	return t, err
}

func (s *TransactionService) Receipt(hashString string) (*models.Receipt, error) {
	hash := common.HexToHash(hashString)
	t, err := s.dao.Receipt(hash)
	if err != nil {
		return nil, err
	}
	t.FormatJSON()
	return t, err
}
