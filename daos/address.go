package daos

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
)

type AddressDAO struct {
}

func NewAddressDAO() *AddressDAO {
	return &AddressDAO{}
}

func (dao *AddressDAO) Insert(address models.AddressInfo) (bool, error) {
	err := db.Upsert("addresses", &bson.M{"address": strings.ToLower(address.Address.String())}, address)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dao *AddressDAO) Find(query bson.M) ([]models.AddressInfo, error) {
	conn := db.Conn()
	defer conn.Close()

	var addresses []models.AddressInfo
	err := conn.DB("").C("addresses").Find(query).All(&addresses)
	if err != nil {
		return addresses, err
	}
	return addresses, nil
}
