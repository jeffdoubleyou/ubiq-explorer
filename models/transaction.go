package models

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"strings"
)

var (
	TransactionList map[string]*Transaction
)

type RpcTransaction struct {
	tx *types.Transaction
	TxExtraInfo
}

type TxTraceResult struct {
	Pc      uint64 `json:"pc"`
	Op      string `json:"op"`
	Gas     uint64 `json:"gas"`
	GasCost uint64 `json:"gasCost"`
	//Memory  []string `json:"memory"`
	//MemorySize int                         `json:"memSize"`
	Stack   []string                    `json:"stack"`
	Storage map[common.Hash]common.Hash `json:"-"`
	Depth   int                         `json:"depth"`
	Err     error                       `json:"-"`
}

type RpcTraceResult struct {
	Gas         uint64 `json:"gas"`
	ReturnValue string `json:"returnValue"`
	StructLogs  []*TxTraceResult
}

type Receipt struct {
	// Consensus fields
	Status            string       `json:"status"`
	CumulativeGasUsed string       `json:"cumulativeGasUsed"`
	Bloom             types.Bloom  `json:"logsBloom"`
	Logs              []*types.Log `json:"logs"`

	// Implementation fields (don't reorder!)
	TxHash          common.Hash `json:"transactionHash"`
	ContractAddress string      `json:"contractAddress"`
	GasUsed         string      `json:"gasUsed"`
}

type TxExtraInfo struct {
	BlockNumber      string         `json:"blockNumber" bson:"_id"`
	BlockHash        common.Hash    `json:"blockHash"`
	From             common.Address `json:"from"`
	Gas              string         `json:"gas"`
	GasPrice         string         `json:"gasPrice"`
	Hash             string         `json:"hash"`
	Input            string         `json:"input"`
	Nonce            string         `json:"nonce"`
	R                string         `json:"r"`
	S                string         `json:"s"`
	To               string         `json:"to,omitempty"`
	TransactionIndex string         `json:"transactionIndex"`
	V                string         `json:"v"`
	Value            string         `json:"value"`
}

type TransactionPage struct {
	Start        string
	End          string
	Total        int
	Transactions []*Transaction
}

type Transaction struct {
	Id        bson.ObjectId  `json:"id,omitempty" bson:"_id,omitempty"`
	Hash      common.Hash    `json:"hash"`
	Timestamp *big.Int       `json:"timestamp"`
	Value     *big.Int       `json:"value"`
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Number    *big.Int       `json:"number"`
	Contract  int            `json:"contract"`
	Gas       uint64         `json:"gas"`
	GasPrice  *big.Int       `json:"gas_price"`
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Hash      string         `json:"hash"`
		Timestamp *big.Int       `json:"timestamp"`
		Value     *big.Int       `json:"value"`
		From      common.Address `json:"from"`
		To        common.Address `json:"to"`
		Number    *big.Int       `json:"number"`
		Contract  int            `json:"contract"`
		Gas       uint64         `json:"gas"`
		GasPrice  *big.Int       `json:"gas_price"`
	}{
		Hash:      t.Hash.String(),
		Timestamp: t.Timestamp,
		Value:     t.Value,
		From:      t.From,
		To:        t.To,
		Number:    t.Number,
		Contract:  t.Contract,
		Gas:       t.Gas,
		GasPrice:  t.GasPrice,
	})
}

func (t *Transaction) UnmarshalJSON(b []byte) error {
	var txn map[string]string
	fmt.Println("ABOUT TO DO IT\n")
	err := json.Unmarshal(b, &txn)
	if err != nil {
		return err
	}
	for k, v := range txn {
		fmt.Printf("K: %s V: %s\n", k, v)
	}
	t.From = common.HexToAddress(txn["from"])
	t.To = common.HexToAddress(txn["to"])
	t.Hash = common.HexToHash(txn["hash"])
	if txn["to"] == "0x0000000000000000000000000000000000000000" {
		t.Contract = 1
	}
	t.Gas, _ = hexutil.DecodeUint64(txn["hash"])
	t.GasPrice = new(big.Int)
	t.GasPrice.SetString(txn["gasPrice"], 0)
	t.Timestamp = new(big.Int)
	t.Timestamp.SetString(txn["timestamp"], 0)
	t.Value = new(big.Int)
	t.Value.SetString(txn["value"], 0)
	return nil
}

// GetBSON implements bson.Getter.
func (t Transaction) GetBSON() (interface{}, error) {
	return struct {
		Id        string `json:"id,omitempty" bson:"_id,omitempty"`
		Hash      string `json:"hash"`
		Timestamp uint64 `json:"timestamp"`
		Value     string `json:"value"`
		From      string `json:"from"`
		To        string `json:"to"`
		Number    string `json:"number"`
		Contract  int    `json:"contract"`
		Gas       uint64 `json:"gas"`
		GasPrice  string `json:"gas_price"`
	}{
		//Id:        t.Id.String(),
		Hash:      t.Hash.String(),
		Timestamp: t.Timestamp.Uint64(),
		Value:     t.Value.String(),
		From:      strings.ToLower(t.From.String()),
		To:        strings.ToLower(t.To.String()),
		Number:    t.Number.String(),
		Contract:  t.Contract,
		Gas:       t.Gas,
		GasPrice:  t.GasPrice.String(),
	}, nil
}

// SetBSON implements bson.Setter.
func (t *Transaction) SetBSON(raw bson.Raw) error {
	decoded := new(struct {
		Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Hash      string        `json:"hash"`
		Timestamp uint64        `json:"timestamp"`
		Value     string        `json:"value"`
		From      string        `json:"from"`
		To        string        `json:"to"`
		Number    string        `json:"number"`
		Contract  int           `json:"contract"`
		Gas       uint64        `json:"gas"`
		GasPrice  string        `json:"gas_price"`
	})

	bsonErr := raw.Unmarshal(decoded)

	value := new(big.Int)
	number := new(big.Int)
	gasPrice := new(big.Int)
	timestamp := big.NewInt(int64(decoded.Timestamp))

	value.SetString(decoded.Value, 10)
	number.SetString(decoded.Number, 10)
	gasPrice.SetString(decoded.GasPrice, 10)

	if bsonErr == nil {
		t.Id = decoded.Id
		t.Hash = common.HexToHash(decoded.Hash)
		t.Timestamp = timestamp
		t.Value = value
		t.From = common.HexToAddress(decoded.From)
		t.To = common.HexToAddress(decoded.To)
		t.Number = number
		t.Contract = decoded.Contract
		t.Gas = decoded.Gas
		t.GasPrice = gasPrice
		return nil
	} else {
		return bsonErr
	}
}

// This mainly coverts hex strings to big int strings
// This isn't being stored in the DB so we don't need the whole BSON translation stuff
func (t *TxExtraInfo) FormatJSON() (*TxExtraInfo, error) {
	var blockNumber,
		gas,
		gasPrice,
		value,
		nonce,
		transactionIndex big.Int

	blockNumber.SetString(t.BlockNumber, 0)
	gas.SetString(t.Gas, 0)
	gasPrice.SetString(t.GasPrice, 0)
	value.SetString(t.Value, 0)
	nonce.SetString(t.Nonce, 0)
	transactionIndex.SetString(t.TransactionIndex, 0)

	t.BlockNumber = blockNumber.String()
	t.Gas = gas.String()
	t.GasPrice = gasPrice.String()
	t.Value = value.String()
	t.Nonce = nonce.String()
	t.TransactionIndex = transactionIndex.String()
	t.To = strings.ToLower(common.HexToAddress(t.To).String())
	return t, nil
}

func (t *Receipt) FormatJSON() (*Receipt, error) {
	var cumulativeGasUsed,
		gasUsed big.Int

	cumulativeGasUsed.SetString(t.CumulativeGasUsed, 0)
	gasUsed.SetString(t.GasUsed, 0)
	t.CumulativeGasUsed = cumulativeGasUsed.String()
	t.GasUsed = gasUsed.String()
	return t, nil
}
