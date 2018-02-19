package models

import (
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type AddressInfo struct {
	Address common.Address `json:"address"`
	Name    string         `json:"name"`
}

func (a AddressInfo) GetBSON() (interface{}, error) {
	return struct {
		Address string `json:"address"`
		Name    string `json:"name"`
	}{
		Address: strings.ToLower(a.Address.String()),
		Name:    a.Name,
	}, nil
}

func (a *AddressInfo) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Address string `json:"address"`
		Name    string `json:"name"`
	})

	bsonErr := raw.Unmarshal(decoded)

	if bsonErr == nil {
		a.Address = common.HexToAddress(decoded.Address)
		a.Name = decoded.Name
		return nil
	} else {
		return bsonErr
	}
}
