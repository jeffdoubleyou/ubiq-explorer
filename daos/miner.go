package daos

import (
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
)

type MinerDAO struct {
}

func NewMinerDAO() *MinerDAO {
	return &MinerDAO{}
}

func (dao *MinerDAO) Insert(miner models.Miner) (bool, error) {
	err := db.Insert("minedBlocks", miner)
	if err != nil {
		return false, err
	}
	return true, nil
}
