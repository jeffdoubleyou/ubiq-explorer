package services

import (
	"gopkg.in/mgo.v2/bson"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
)

type AddressService struct {
	dao daos.AddressDAO
}

func NewAddressService(dao daos.AddressDAO) *AddressService {
	return &AddressService{dao}
}

func (s *AddressService) List() ([]models.AddressInfo, error) {
	return s.dao.Find(bson.M{})
}
