package services

import (
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
)

type TokenBalanceService struct {
	dao daos.TokenBalanceDAO
}

func NewTokenBalanceService(dao daos.TokenBalanceDAO) *TokenBalanceService {
	return &TokenBalanceService{dao}
}

func (s *TokenBalanceService) Update(tokenBalance *models.TokenBalance) (bool, error) {
	query := bson.M{
		"address":      strings.ToLower(tokenBalance.Address.String()),
		"tokenAddress": strings.ToLower(tokenBalance.TokenAddress.String()),
	}

	balance, _ := s.dao.Find(query)

	if len(balance) > 0 {
		ok, err := s.dao.Delete(query)
		if !ok {
			panic(err)
			return false, err
		}
	}
	return s.dao.Insert(tokenBalance)
}

func (s *TokenBalanceService) Address(address common.Address) ([]models.TokenBalance, error) {
	return s.dao.Find(bson.M{"address": strings.ToLower(address.String())})
}
