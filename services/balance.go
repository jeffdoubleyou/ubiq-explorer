package services

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
)

type BalanceService struct {
	dao daos.BalanceDAO
}

func NewBalanceService(dao daos.BalanceDAO) *BalanceService {
	return &BalanceService{dao}
}

func (s *BalanceService) Get(address common.Address, blockNumber int64) (*models.Balance, error) {

	if address.String() == "0x0000000000000000000000000000000000000000" || len(address.String()) != 42 {
		return &models.Balance{}, fmt.Errorf("Invalid address %s", address.String())
	}

	query := bson.M{"address": strings.ToLower(address.String())}

	if blockNumber > 0 {
		bigCursor := big.NewInt(blockNumber)
		query["block"] = bson.M{"$lte": bigCursor.Int64()}
	}

	page, err := s.dao.Find(query, 1, "", "")

	if err != nil {
		return &models.Balance{}, err
	}

	if len(page.Balances) > 0 {
		return page.Balances[0], err
	} else {
		return &models.Balance{}, fmt.Errorf("No balance for %s at block %d", address.String(), blockNumber)
	}
}

func (s *BalanceService) History(address common.Address, limit int, cursor string) (models.BalancePage, error) {
	if limit > 100 {
		limit = 100
	}
	if limit == 0 {
		limit = 10
	}
	if address.String() == "0x0000000000000000000000000000000000000000" || len(address.String()) != 42 {
		return models.BalancePage{}, fmt.Errorf("Invalid address %s", address.String())
	}
	return s.dao.Find(bson.M{"address": strings.ToLower(address.String())}, limit, cursor, "")
}

func (s *BalanceService) RichList(count int) (models.CurrentBalancePage, error) {
	return s.dao.FindCurrentBalance(bson.M{}, count, "", "-balance")
}

func (s *BalanceService) UpdateCurrentBalance(address common.Address, balance *big.Int) (bool, error) {
	query := bson.M{
		"address": strings.ToLower(address.String()),
	}

	currentBalance, _ := s.dao.FindCurrentBalance(query, 1, "", "")

	if len(currentBalance.Balances) > 0 {
		ok, err := s.dao.DeleteCurrentBalance(query)
		if !ok {
			return false, err
		}
	}
	return s.dao.InsertCurrentBalance(models.CurrentBalance{Address: address, Balance: balance})
}
