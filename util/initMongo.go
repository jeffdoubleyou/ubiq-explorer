package util

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"ubiq-explorer/models/db"
)

func InitializeMongoDB() (bool, error) {
	historyWindow, err := beego.AppConfig.Int("stats::history_window")
	if err != nil {
		return false, err
	}

	db := db.Conn()
	defer db.Close()

	dbOptions := &mgo.CollectionInfo{
		Capped:   true,
		MaxDocs:  historyWindow,
		MaxBytes: 500000,
	}

	if cnt, _ := db.DB("").C("blockTimeHistory").Count(); cnt == 0 {
		err = db.DB("").C("blockTimeHistory").Create(dbOptions)
		if err != nil {
			return false, err
		}
	}
	if cnt, _ := db.DB("").C("hashRateHistory").Count(); cnt == 0 {
		err = db.DB("").C("hashRateHistory").Create(dbOptions)
		if err != nil {
			return false, err
		}
	}
	if cnt, _ := db.DB("").C("difficultyHistory").Count(); cnt == 0 {
		err = db.DB("").C("difficultyHistory").Create(dbOptions)
		if err != nil {
			return false, err
		}
	}
	if cnt, _ := db.DB("").C("uncleRateHistory").Count(); cnt == 0 {
		err = db.DB("").C("uncleRateHistory").Create(dbOptions)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
