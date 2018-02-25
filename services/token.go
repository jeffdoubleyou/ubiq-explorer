package services

import (
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
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

type TokenService struct {
	dao daos.TokenDAO
}

func NewTokenService(dao daos.TokenDAO) *TokenService {
	return &TokenService{dao}
}

func (s *TokenService) From(address common.Address, limit int, cursor string) (models.TokenTransactionPage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	return s.dao.FindTransactions(bson.M{"from": strings.ToLower(address.String())}, "-_id", limit, cursor)
}

func (s *TokenService) To(address common.Address, limit int, cursor string) (models.TokenTransactionPage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	return s.dao.FindTransactions(bson.M{"to": strings.ToLower(address.String())}, "-_id", limit, cursor)
}

func (s *TokenService) TransactionList(limit int, cursor string) (models.TokenTransactionPage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	return s.dao.FindTransactions(bson.M{}, "-_id", limit, cursor)
}

func (s *TokenService) GetTokenByAddress(address common.Address) (models.TokenInfo, error) {
	return s.dao.GetTokenByAddress(strings.ToLower(address.String()))
}

func (s *TokenService) GetTokenBySymbol(symbol string) (models.TokenInfo, error) {
	return s.dao.GetTokenBySymbol(symbol)
}

func (s *TokenService) TokenList(limit int, cursor string) (models.TokenInfoPage, error) {
	if limit > 100 {
		limit = 100
	}
	return s.dao.FindTokens(bson.M{}, "-_id", limit, cursor)
}
