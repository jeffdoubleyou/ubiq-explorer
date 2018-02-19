package daos

import (
	"gopkg.in/mgo.v2/bson"
	"ubiq-explorer/models/db"
)

type HistoryDAO struct {
}

func NewHistoryDAO() *HistoryDAO {
	return &HistoryDAO{}
}

func (dao *HistoryDAO) Insert(collection string, history interface{}) (bool, error) {
	conn := db.Conn()
	defer conn.Close()
	err := conn.DB("").C(collection).Insert(history)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *HistoryDAO) Find(collection string) ([]interface{}, error) {
	conn := db.Conn()
	defer conn.Close()

	var history []interface{}
	err := conn.DB("").C(collection + "History").Find(bson.M{}).All(&history)
	return history, err
}
