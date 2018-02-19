package models

import (
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
)

type Uncle struct {
	Block      *big.Int `json:"block"`
	Height     *big.Int `json:"height"`
	Difficulty *big.Int `json:"difficulty"`
	Timestamp  uint64   `json:"timestamp"`
	Gas        uint64   `json:"gas"`
	Miner      string   `json:"miner"`
	Uncle      int      `json:"uncle"`
}

type UnclePage struct {
	Start  string
	End    string
	Total  int
	Uncles []*Uncle
}

// GetBSON implements bson.Getter.
func (u Uncle) GetBSON() (interface{}, error) {
	difficulty := u.Difficulty.String()

	return struct {
		Block      int64  `json:"block" bson:"block"`
		Height     int64  `json:"block" bson:"height"`
		Difficulty string `json:"difficulty"`
		Timestamp  uint64 `json:"timestamp"`
		Gas        uint64 `json:"gas"`
		Miner      string `json:"miner"`
		Uncle      int    `json:"uncle"`
	}{
		Block:      u.Block.Int64(),
		Height:     u.Height.Int64(),
		Difficulty: difficulty,
		Timestamp:  u.Timestamp,
		Gas:        u.Gas,
		Miner:      strings.ToLower(u.Miner),
		Uncle:      u.Uncle,
	}, nil
}

// SetBSON implements bson.Setter.
func (u *Uncle) SetBSON(raw bson.Raw) error {

	decoded := new(struct {
		Block      int64  `json:"block" bson:"block"`
		Height     int64  `json:"block" bson:"height"`
		Difficulty string `json:"difficulty"`
		Timestamp  uint64 `json:"timestamp"`
		Gas        uint64 `json:"gas"`
		Miner      string `json:"miner"`
		Uncle      int    `json:"uncle"`
	})

	bsonErr := raw.Unmarshal(decoded)

	block := big.NewInt(decoded.Block)
	height := big.NewInt(decoded.Height)
	diff := new(big.Int)
	diff.SetString(decoded.Difficulty, 10)

	if bsonErr == nil {
		u.Block = block
		u.Height = height
		u.Difficulty = diff
		u.Timestamp = decoded.Timestamp
		u.Gas = decoded.Gas
		u.Miner = decoded.Miner
		u.Uncle = decoded.Uncle
		return nil
	} else {
		return bsonErr
	}
}
