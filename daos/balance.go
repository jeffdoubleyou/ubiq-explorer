package daos

import (
	"gopkg.in/mgo.v2/bson"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
)

type BalanceDAO struct {
}

func NewBalanceDAO() *BalanceDAO {
	return &BalanceDAO{}
}

func (dao *BalanceDAO) Insert(balance models.Balance) (bool, error) {
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("balances").Insert(balance)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *BalanceDAO) Find(query bson.M, limit int, cursor string, sort string) (models.BalancePage, error) {
	conn := db.Conn()
	defer conn.Close()

	count, err := conn.DB("").C("balances").Find(query).Count()

	if cursor != "" {
		query["_id"] = bson.M{"$gt": bson.ObjectIdHex(cursor)}
	}

	if sort == "" {
		sort = "-block"
	}

	var balances []*models.Balance
	var page = models.BalancePage{Total: count, Start: "", End: ""}
	err = conn.DB("").C("balances").Find(query).Sort(sort).Limit(limit).All(&balances)
	if err != nil {
		return page, err
	}
	if len(balances) > 0 {
		page.Start = balances[0].Id.Hex()
		page.End = balances[len(balances)-1].Id.Hex()
	}
	page.Balances = balances
	return page, nil
}

func (dao *BalanceDAO) InsertCurrentBalance(balance models.CurrentBalance) (bool, error) {
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("currentBalance").Insert(balance)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *BalanceDAO) FindCurrentBalance(query bson.M, limit int, cursor string, sort string) (models.CurrentBalancePage, error) {
	conn := db.Conn()
	defer conn.Close()

	count, err := conn.DB("").C("currentBalance").Find(query).Count()

	if cursor != "" {
		query["_id"] = bson.M{"$gt": bson.ObjectIdHex(cursor)}
	}

	if sort == "" {
		sort = "-block"
	}

	var balances []*models.CurrentBalance
	var page = models.CurrentBalancePage{Total: count, Start: "", End: ""}
	err = conn.DB("").C("currentBalance").Find(query).Sort(sort).Limit(limit).All(&balances)
	if err != nil {
		return page, err
	}
	if len(balances) > 0 {
		page.Start = balances[0].Id.Hex()
		page.End = balances[len(balances)-1].Id.Hex()
	}
	page.Balances = balances
	return page, nil

}

func (dao *BalanceDAO) DeleteCurrentBalance(query bson.M) (bool, error) {
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("currentBalance").Remove(query)
	if err != nil {
		return false, err
	}
	return true, nil
}
