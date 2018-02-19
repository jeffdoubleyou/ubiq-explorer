package models

import (
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
)

type Miner struct {
	Block      *big.Int `json:"block"`
	Difficulty *big.Int `json:"difficulty"`
	Timestamp  uint64   `json:"timestamp"`
	Gas        uint64   `json:"gas"`
	Miner      string   `json:"miner"`
}

// GetBSON implements bson.Getter.
func (m Miner) GetBSON() (interface{}, error) {
	//blockNumber := m.Block.String()
	difficulty := m.Difficulty.String()

	return struct {
		Block      int64  `json:"block" bson:"block"`
		Difficulty string `json:"difficulty"`
		Timestamp  uint64 `json:"timestamp"`
		Gas        uint64 `json:"gas"`
		Miner      string `json:"miner"`
	}{
		Block: m.Block.Int64(),
		//Block:      blockNumber,
		Difficulty: difficulty,
		Timestamp:  m.Timestamp,
		Gas:        m.Gas,
		Miner:      strings.ToLower(m.Miner),
	}, nil
}

// SetBSON implements bson.Setter.
func (m *Miner) SetBSON(raw bson.Raw) error {

	decoded := new(struct {
		Block      int64  `json:"block" bson:"block"`
		Difficulty string `json:"difficulty"`
		Timestamp  uint64 `json:"timestamp"`
		Gas        uint64 `json:"gas"`
		Miner      string `json:"miner"`
	})

	bsonErr := raw.Unmarshal(decoded)

	block := big.NewInt(decoded.Block)
	diff := new(big.Int)
	//block.SetString(decoded.Block, 10)
	diff.SetString(decoded.Difficulty, 10)

	if bsonErr == nil {
		m.Block = block
		m.Difficulty = diff
		m.Timestamp = decoded.Timestamp
		m.Gas = decoded.Gas
		m.Miner = decoded.Miner
		return nil
	} else {
		return bsonErr
	}
}
