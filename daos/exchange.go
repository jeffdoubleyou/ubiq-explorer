package daos

import (
	"gopkg.in/mgo.v2/bson"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
)

type ExchangeDAO struct {
}

func NewExchangeDAO() *ExchangeDAO {
	return &ExchangeDAO{}
}

func (dao *ExchangeDAO) Insert(exchange *models.ExchangeRate) (bool, error) {
	_, err := db.Upsert("exchangeRate", bson.M{"symbol": exchange.Symbol}, exchange)
	if err != nil {
		return false, err
	}
	err := db.Insert("exchangeRate_"+exchange.Symbol, exchange)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *ExchangeDAO) Get(symbol string) *models.ExchangeRate {
	conn := db.Conn()
	defer conn.Close()
	exchange := models.ExhangeRate{}
	err := conn.DB("").C("exchangeRate").Find(bson.M{"symbol": symbol}).Sort("-timestamp").One(&exchange)
	if err != nil {
		return exchange, err
	}
	return exchange, nil

}

func (dao *ExchangeDAO) List() *models.ExchangeList {

}

func (dao *ExchangeDAO) History(name string, symbol string) *models.ExchangeHistory {

}

func (dao *ExchangeDAO) GetExchangeSource(symbol string) (*models.ExchangeSource, error) {
	source := models.ExchangeSource{}
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("exchangeSource").Find(bson.M{"symbol": symbol}).One(&source)
	if err != nil {
		return source, err
	}
	return source, nil
}

func (dao *ExchangeDAO) InsertExchangeSource(source *models.ExchangeSource) (bool, error) {
	_, err := db.Upsert("exchangeSource", bson.M{"symbol": source.Symbol}, source)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *ExhangeDAO) GetExhangeAuth(exchange string) (*models.ExchangeAuth, error) {
	auth := models.ExchangeAuth{}
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("exchangeAuth").Find(bson.M{"exchange": exchange}).One(&auth)
	if err != nil {
		return auth, err
	}
	return auth, nil
}

func (dao *ExchangeDAO) InsertExchangeAuth(auth *models.ExchangeAuth) (bool, error) {
	_, err := db.Upsert("exchangeAuth", bson.M{"exchange": auth.Exchange}, auth)
	if err != nil {
		return false, err
	}
	return true, nil
}
