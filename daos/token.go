package daos

import (
	"gopkg.in/mgo.v2/bson"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
)

type TokenDAO struct {
}

func NewTokenDAO() *TokenDAO {
	return &TokenDAO{}
}

func (dao *TokenDAO) Insert(t models.TokenInfo) (bool, error) {
	token, err := dao.GetTokenByAddress(t.Address.String())
	if err == nil && token.Address.String() != "" {
		return true, nil
	}
	conn := db.Conn()
	defer conn.Close()
	err = conn.DB("").C("tokens").Insert(t)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *TokenDAO) GetTokenBySymbol(symbol string) (models.TokenInfo, error) {
	token := models.TokenInfo{}
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("tokens").Find(bson.M{"symbol": symbol}).One(&token)
	if err != nil {
		return token, err
	}
	return token, nil
}

func (dao *TokenDAO) GetTokenByAddress(address string) (models.TokenInfo, error) {
	token := models.TokenInfo{}
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("tokens").Find(bson.M{"address": address}).One(&token)
	if err != nil {
		return token, err
	}
	return token, nil
}

func (dao *TokenDAO) InsertTokenTransaction(txn models.TokenTransaction) (bool, error) {
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("tokenTransactions").Insert(txn)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *TokenDAO) Find(query bson.M, sort string, limit int, cursor string) (models.TokenTransactionPage, error) {
	conn := db.Conn()
	defer conn.Close()

	count, err := conn.DB("").C("tokenTransactions").Find(query).Count()

	if cursor != "" {
		query["_id"] = bson.M{"$lt": bson.ObjectIdHex(cursor)}
	}

	var txns []*models.TokenTransaction
	var page = models.TokenTransactionPage{Total: count, Start: "", End: ""}
	err = conn.DB("").C("tokenTransactions").Find(query).Sort(sort).Limit(limit).All(&txns)
	if err != nil {
		return page, err
	}
	if len(txns) > 0 {
		page.Start = txns[0].Id.Hex()
		page.End = txns[len(txns)-1].Id.Hex()
	}
	page.Transactions = txns
	return page, nil
}
