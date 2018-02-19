package models

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
)

type Balance struct {
	Id        bson.ObjectId  `json:"id,omitempty" bson:"_id,omitempty"`
	Block     *big.Int       `json:"block"`
	Address   common.Address `json:"address"`
	Balance   *big.Int       `json:"balance"`
	ChangedBy string         `json:"changedBy"`
}

type CurrentBalance struct {
	Id      bson.ObjectId  `json:"id,omitempty" bson:"_id,omitempty"`
	Address common.Address `json:"address"`
	Balance *big.Int       `json:"balance"`
}

type BalancePage struct {
	Start    string
	End      string
	Total    int
	Balances []*Balance
}

type CurrentBalancePage struct {
	Start    string
	End      string
	Total    int
	Balances []*CurrentBalance
}

// GetBSON implements bson.Getter.
func (b Balance) GetBSON() (interface{}, error) {
	return struct {
		Id        string `json:"id,omitempty" bson:"_id,omitempty"`
		Block     int64  `json:"block" bson:"block"`
		Address   string `json:"address"`
		Balance   string `json:"balance"`
		ChangedBy string `json:"changedBy"`
	}{
		Block:     b.Block.Int64(),
		Address:   strings.ToLower(b.Address.String()),
		Balance:   fmt.Sprintf("%036s", b.Balance.String()),
		ChangedBy: b.ChangedBy,
	}, nil
}

// SetBSON implements bson.Setter.
func (b *Balance) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Block     int64         `json:"block" bson:"block"`
		Address   string        `json:"address"`
		Balance   string        `json:"balance"`
		ChangedBy string        `json:"changedBy"`
	})

	bsonErr := raw.Unmarshal(decoded)

	block := big.NewInt(decoded.Block)
	balance := new(big.Int)
	balance.SetString(strings.TrimLeft(decoded.Balance, "0"), 10)
	if bsonErr == nil {
		b.Id = decoded.Id
		b.Block = block
		b.Address = common.HexToAddress(decoded.Address)
		b.Balance = balance
		b.ChangedBy = decoded.ChangedBy
		return nil
	} else {
		return bsonErr
	}
}

func (b CurrentBalance) GetBSON() (interface{}, error) {
	return struct {
		Id      string `json:"id,omitempty" bson:"_id,omitempty"`
		Address string `json:"address"`
		Balance string `json:"balance"`
	}{
		Address: strings.ToLower(b.Address.String()),
		Balance: fmt.Sprintf("%036s", b.Balance.String()),
	}, nil
}

func (b *CurrentBalance) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Id      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Address string        `json:"address"`
		Balance string        `json:"balance"`
	})

	bsonErr := raw.Unmarshal(decoded)

	if bsonErr == nil {
		balance := new(big.Int)
		balance.SetString(strings.TrimLeft(decoded.Balance, "0"), 10)
		b.Id = decoded.Id
		b.Address = common.HexToAddress(decoded.Address)
		b.Balance = balance
		return nil
	} else {
		return bsonErr
	}
}
