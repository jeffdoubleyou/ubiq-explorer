package daos

import (
	"gopkg.in/mgo.v2/bson"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
)

type PoolsDAO struct {
}

func NewPoolsDAO() *PoolsDAO {
	return &PoolsDAO{}
}

func (dao *PoolsDAO) Insert(pool *models.Pool) (bool, error) {
	conn := db.Conn()
	defer conn.Close()
	_, err := conn.DB("").C("pools").Upsert(bson.M{"url": pool.Url}, pool)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *PoolsDAO) Find(query bson.M) ([]models.Pool, error) {
	conn := db.Conn()
	defer conn.Close()

	var pools []models.Pool
	err := conn.DB("").C("pools").Find(query).All(&pools)
	if err != nil {
		return pools, err
	}
	return pools, nil
}
