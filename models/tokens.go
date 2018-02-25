package models

import (
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
)

type TokenInfo struct {
	Id       bson.ObjectId  `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string         `json:"name"`
	Address  common.Address `json:"address"`
	Symbol   string         `json:"symbol"`
	Decimals uint8          `json:"decimals"`
}

type TokenInfoPage struct {
	Start  string
	End    string
	Total  int
	Tokens []*TokenInfo
}

type TokenTransaction struct {
	Id        bson.ObjectId  `json:"id,omitempty" bson:"_id,omitempty"`
	Address   common.Address `json:"address"`
	Hash      common.Hash    `json:"hash"`
	Timestamp *big.Int       `json:"timestamp"`
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Value     *big.Float     `json:"value"`
	TokenInfo *TokenInfo     `json:"tokenInfo"`
}

type TokenTransactionPage struct {
	Start        string
	End          string
	Total        int
	Transactions []*TokenTransaction
}

func (t TokenInfo) GetBSON() (interface{}, error) {
	return struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Symbol   string `json:"symbol"`
		Decimals uint8  `json:"decimals"`
	}{
		Name:     t.Name,
		Address:  strings.ToLower(t.Address.String()),
		Symbol:   t.Symbol,
		Decimals: t.Decimals,
	}, nil
}

func (t *TokenInfo) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Name     string        `json:"name"`
		Address  string        `json:"address"`
		Symbol   string        `json:"symbol"`
		Decimals uint8         `json:"decimals"`
	})
	bsonErr := raw.Unmarshal(decoded)

	if bsonErr == nil {
		t.Id = decoded.Id
		t.Name = decoded.Name
		t.Address = common.HexToAddress(decoded.Address)
		t.Symbol = decoded.Symbol
		t.Decimals = decoded.Decimals
		return nil
	} else {
		return bsonErr
	}
}

// GetBSON implements bson.Getter.
func (t TokenTransaction) GetBSON() (interface{}, error) {
	return struct {
		Id        string     `json:"id,omitempty" bson:"_id,omitempty"`
		Address   string     `json:"address"`
		Hash      string     `json:"hash"`
		Timestamp uint64     `json:"timestamp"`
		From      string     `json:"from"`
		To        string     `json:"to"`
		Value     string     `json:"value"`
		TokenInfo *TokenInfo `json:"tokenInfo"`
	}{
		//Id:        t.Id.String(),
		Address:   strings.ToLower(t.Address.String()),
		Hash:      t.Hash.String(),
		Timestamp: t.Timestamp.Uint64(),
		From:      strings.ToLower(t.From.String()),
		To:        strings.ToLower(t.To.String()),
		Value:     t.Value.String(),
		TokenInfo: t.TokenInfo,
	}, nil
}

// SetBSON implements bson.Setter.
func (t *TokenTransaction) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Address   string        `json:"address"`
		Hash      string        `json:"hash"`
		Timestamp uint64        `json:"timestamp"`
		From      string        `json:"from"`
		To        string        `json:"to"`
		Value     string        `json:"value"`
		TokenInfo *TokenInfo    `json:"tokenInfo"`
	})

	bsonErr := raw.Unmarshal(decoded)

	value := new(big.Float)
	timestamp := big.NewInt(int64(decoded.Timestamp))

	value.SetString(decoded.Value)

	if bsonErr == nil {
		t.Id = decoded.Id
		t.Address = common.HexToAddress(decoded.Address)
		t.Hash = common.HexToHash(decoded.Hash)
		t.Timestamp = timestamp
		t.From = common.HexToAddress(decoded.From)
		t.To = common.HexToAddress(decoded.To)
		t.Value = value
		t.TokenInfo = decoded.TokenInfo
		return nil
	} else {
		return bsonErr
	}
}
