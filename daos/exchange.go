package daos

import (
    "errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
)

type ExchangeDAO struct {
}

func NewExchangeDAO() *ExchangeDAO {
	return &ExchangeDAO{}
}

func (dao *ExchangeDAO) InsertExchangeRate(exchange *models.ExchangeRate, cap int) (bool, error) {
	err := db.Upsert("exchangeRate", &bson.M{"symbol": exchange.Symbol}, exchange)
	if err != nil {
		return false, err
	}
	// We need a limit on the table
	c := db.Conn()
	defer c.Close()
	symbolCollection := "exchangeRate_" + exchange.Symbol
	if cnt, _ := c.DB("").C(symbolCollection).Count(); cnt == 0 {
		if cap == 0 {
			cap = 288
		}
		dbOptions := &mgo.CollectionInfo{
			Capped:   true,
			MaxDocs:  cap,
			MaxBytes: 500000,
		}
		c.DB("").C(symbolCollection).DropCollection()
		err = c.DB("").C(symbolCollection).Create(dbOptions)
		if err != nil {
			return false, err
		}
	}

	err = db.Insert(symbolCollection, exchange)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *ExchangeDAO) GetExchangeRate(symbol string) (*models.ExchangeRate, error) {
	conn := db.Conn()
	defer conn.Close()
	exchange := &models.ExchangeRate{}
	err := conn.DB("").C("exchangeRate").Find(bson.M{"symbol": symbol}).Sort("-timestamp").One(&exchange)
	if err != nil {
		return exchange, err
	}
	return exchange, nil

}

func (dao *ExchangeDAO) DeleteExchangeRate(symbol string) (bool, error) {
	if symbol == "" {
		return false, errors.New("Symbol is required")
	}
    conn := db.Conn()
    defer conn.Close()
    err := conn.DB("").C("exchangeRate").Remove(&bson.M{"symbol": symbol})
    if err != nil {
        return false, err
    }
    return true, nil
}


func (dao *ExchangeDAO) ExchangeRateList() ([]*models.ExchangeRate, error) {
	conn := db.Conn()
	defer conn.Close()
	exchangeRates := []*models.ExchangeRate{}
	err := conn.DB("").C("exchangeRate").Find(bson.M{}).All(&exchangeRates)
	return exchangeRates, err
}

func (dao *ExchangeDAO) ExchangeRateHistory(symbol string) ([]*models.ExchangeRate, error) {
	conn := db.Conn()
	defer conn.Close()
	exchangeRates := []*models.ExchangeRate{}
	err := conn.DB("").C("exchangeRate_" + symbol).Find(bson.M{}).Sort("-timestamp").All(&exchangeRates)
	return exchangeRates, err
}

func (dao *ExchangeDAO) GetExchangeSource(symbol string) (*models.ExchangeSource, error) {
	source := &models.ExchangeSource{}
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("exchangeSource").Find(bson.M{"symbol": symbol}).One(&source)
	if err != nil {
		return source, err
	}
	return source, nil
}

func (dao *ExchangeDAO) InsertExchangeSource(source *models.ExchangeSource) (bool, error) {
	err := db.Upsert("exchangeSource", &bson.M{"symbol": source.Symbol}, source)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *ExchangeDAO) GetExhangeAuth(exchange string) (*models.ExchangeAuth, error) {
	auth := &models.ExchangeAuth{}
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C("exchangeAuth").Find(bson.M{"exchange": exchange}).One(&auth)
	if err != nil {
		return auth, err
	}
	return auth, nil
}

func (dao *ExchangeDAO) InsertExchangeAuth(auth *models.ExchangeAuth) (bool, error) {
	err := db.Upsert("exchangeAuth", &bson.M{"exchange": auth.Exchange}, auth)
	if err != nil {
		return false, err
	}
	return true, nil
}
