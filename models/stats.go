package models

import (
	"math/big"
)

type Stats struct {
	LastBlock  *big.Int `json:"lastBlock"`
	Difficulty string   `json:"difficulty"`
	BlockTime  float64  `json:"blockTime"`
	HashRate   string   `json:"hashRate"`
	UncleRate  float64  `json:"uncleRate"`
}

type MinerList struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Count   uint16  `json:"count"`
	Percent float64 `json:"percent"`
}
