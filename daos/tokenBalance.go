package daos

import (
	"gopkg.in/mgo.v2/bson"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
    "strings"
)

type TokenBalanceDAO struct {
}

func NewTokenBalanceDAO() *TokenBalanceDAO {
	return &TokenBalanceDAO{}
}

func (dao *TokenBalanceDAO) Insert(balance *models.TokenBalance) (bool, error) {
	conn := db.Conn()
	defer conn.Close()
	_, err := conn.DB("").C("tokenBalance").Upsert(bson.M{"address": strings.ToLower(balance.Address.String()),"tokenAddress": strings.ToLower(balance.TokenAddress.String())},balance)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *TokenBalanceDAO) Find(query bson.M) ([]models.TokenBalance, error) {
	conn := db.Conn()
	defer conn.Close()

	var balances []models.TokenBalance
	err := conn.DB("").C("tokenBalance").Find(query).All(&balances)
	if err != nil {
		return balances, err
	}
	return balances, nil
}

func (dao *TokenBalanceDAO) Delete(query bson.M) (bool, error) {
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("tokenBalance").Remove(query)
	if err != nil {
		return false, err
	}
	return true, nil
}
