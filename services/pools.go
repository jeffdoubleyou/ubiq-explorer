package services

import (
	"gopkg.in/mgo.v2/bson"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
)

type PoolsService struct {
	dao daos.PoolsDAO
}

func NewPoolsService(dao daos.PoolsDAO) *PoolsService {
	return &PoolsService{dao}
}

func (s *PoolsService) List() ([]models.Pool, error) {
	return s.dao.Find(bson.M{})
}

func (s *PoolsService) Insert(pool *models.Pool) (bool, error) {
	return s.dao.Insert(pool)
}
