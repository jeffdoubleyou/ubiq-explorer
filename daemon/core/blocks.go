package core

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"ubiq-explorer/models"
	"ubiq-explorer/node"
)

type Blocks struct {
	Context context.Context
	Block   *types.Block
}

func NewBlocks() *Blocks {
	return &Blocks{
		context.TODO(), // TODO
		&types.Block{},
	}
}

func (b *Blocks) GetCurrentBlock() (*big.Int, error) {
	block, err := node.Client().HeaderByNumber(b.Context, nil)
	if err != nil {
		return nil, err
	}
	n := block.Number
	return n, nil
}

func (b *Blocks) GetBlock(blockNumber *big.Int) (*types.Block, error) {
	block, err := node.Client().BlockByNumber(b.Context, blockNumber)
	b.Block = block
	return block, err
}

func (b *Blocks) Miner() *models.Miner {
	miner := &models.Miner{
		Block:      b.Block.Number(),
		Difficulty: b.Block.Difficulty(),
		Timestamp:  b.Block.Time().Uint64(),
		Gas:        b.Block.GasUsed(),
		Miner:      b.Block.Coinbase().String(),
	}
	return miner
}

func (b *Blocks) Uncles() map[int]*models.Uncle {
	uncles := make(map[int]*models.Uncle)
	for i, u := range b.Block.Uncles() {
		uncle := &models.Uncle{
			b.Block.Number(),
			u.Number,
			u.Difficulty,
			u.Time.Uint64(),
			u.GasUsed,
			u.Coinbase.String(),
			i,
		}
		uncles[i] = uncle
	}
	return uncles
}

func (b *Blocks) Transactions() map[int]*models.Transaction {
	var txnList = make(map[int]*models.Transaction)

	for idx, txn := range b.Block.Transactions() {
		sender, _ := node.Client().TransactionSender(context.Background(), txn, b.Block.Hash(), uint(idx))
		var isContract int = 0
		to := common.Address{}
		if txn.To() == nil {
			isContract = 1
		} else {
			to = *txn.To()
			if txn.Value().Int64() == 0 {
				code, _ := node.Client().CodeAt(b.Context, *txn.To(), b.Block.Number())
				if code != nil && len(code) > 0 {
					isContract = 1
				}
			}
		}
		t := &models.Transaction{
			Id:        "",
			Hash:      txn.Hash(),
			Timestamp: b.Block.Time(),
			Value:     txn.Value(),
			From:      sender,
			To:        to,
			Number:    b.Block.Number(),
			Contract:  isContract,
			Gas:       txn.Gas(),
			GasPrice:  txn.GasPrice(),
		}
		txnList[idx] = t
	}

	return txnList
}

func (b *Blocks) Balance(address common.Address, block *big.Int) (*models.Balance, error) {
	balance, err := node.Client().BalanceAt(context.TODO(), address, block)
	if err != nil {
		return nil, err
	}
	addressBalance := &models.Balance{
		Block:   block,
		Address: address,
		Balance: balance,
	}
	return addressBalance, nil
}
