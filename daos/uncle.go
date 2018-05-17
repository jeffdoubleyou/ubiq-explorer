package daos

import (
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
)

type UncleDAO struct {
}

func NewUncleDAO() *UncleDAO {
	return &UncleDAO{}
}

func (dao *UncleDAO) Insert(uncle models.Uncle) (bool, error) {
	err := db.Insert("minedUncles", uncle)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *UncleDAO) Find(query bson.M, limit int, cursor string) (models.UnclePage, error) {
	conn := db.Conn()
	defer conn.Close()

	if cursor != "" {
		var bigCursor = new(big.Int)
		bigCursor.SetString(cursor, 10)
		query = bson.M{"block": bson.M{"$lt": bigCursor.Int64()}}
	}

	count, err := conn.DB("").C("minedUncles").Find(query).Count()
	var blocks []*models.Uncle
	var page = models.UnclePage{Total: count, Start: "", End: ""}
	err = conn.DB("").C("minedUncles").Find(query).Sort("-block").Limit(limit).All(&blocks)
	if err != nil {
		return page, err
	}
	if len(blocks) > 0 {
		page.Start = blocks[0].Block.String()
		page.End = blocks[len(blocks)-1].Block.String()
	}
	page.Uncles = blocks
	return page, nil
}

func (dao *UncleDAO) Count(query bson.M) (int, error) {
	conn := db.Conn()
	defer conn.Close()

	count, err := conn.DB("").C("minedUncles").Find(query).Count()

	if err != nil {
		return 0, err
	}
	return count, nil
}
