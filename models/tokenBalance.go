package models

import (
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
	"math/big"

	"strings"
)

type TokenBalance struct {
	Address      common.Address `json:"address"`
	TokenAddress common.Address `json:"tokenAddress" bson:"tokenAddress"`
	Token        string         `json:"token"`
	Symbol       string         `json:"symbol"`
	Balance      *big.Float     `json:"balance"`
}

// GetBSON implements bson.Getter.
func (b TokenBalance) GetBSON() (interface{}, error) {
	balance, _ := b.Balance.Float64()
	return struct {
		Address      string  `json:"address"`
		TokenAddress string  `json:"tokenAddress" bson:"tokenAddress"`
		Token        string  `json:"token"`
		Symbol       string  `json:"symbol"`
		Balance      float64 `json:"balance"`
	}{
		Address:      strings.ToLower(b.Address.String()),
		TokenAddress: strings.ToLower(b.TokenAddress.String()),
		Token:        b.Token,
		Symbol:       b.Symbol,
		Balance:      balance,
	}, nil
}

// SetBSON implements bson.Setter.
func (b *TokenBalance) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Address      string  `json:"address"`
		TokenAddress string  `json:"tokenAddress" bson:"tokenAddress"`
		Token        string  `json:"token"`
		Symbol       string  `json:"symbol"`
		Balance      float64 `json:"balance"`
	})

	bsonErr := raw.Unmarshal(decoded)
	if bsonErr == nil {
		b.Address = common.HexToAddress(decoded.Address)
		b.TokenAddress = common.HexToAddress(decoded.TokenAddress)
		b.Token = decoded.Token
		b.Symbol = decoded.Symbol
		b.Balance = big.NewFloat(decoded.Balance)
		return nil
	} else {
		return bsonErr
	}
}
