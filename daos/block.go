package daos

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
	"ubiq-explorer/node"
)

type BlockDAO struct {
}

func NewBlockDAO() *BlockDAO {
	return &BlockDAO{}
}

func (dao *BlockDAO) LastImportedBlock() (*big.Int, error) {
	conn := db.Conn()
	defer conn.Close()
	b := models.Miner{}
	err := conn.DB("").C("minedBlocks").Find(bson.M{}).Sort("-timestamp").One(&b)
	if err != nil {
		return big.NewInt(0), nil
	}
	return b.Block, nil
}

// TODO Rename all this minedBlocks crap to just Blocks, mined is irrelevant since it's now used for more than just miner info
func (dao *BlockDAO) Find(query bson.M, limit int, cursor string) (models.BlockPage, error) {
	conn := db.Conn()
	defer conn.Close()

	count, err := conn.DB("").C("minedBlocks").Find(query).Count()

	if cursor != "" {
		var bigCursor big.Int
		bigCursor.SetString(cursor, 10)
		query = bson.M{"block": bson.M{"$lt": bigCursor.Int64()}}
	}

	var blocks []*models.Miner
	var page = models.BlockPage{Total: count, Start: "", End: ""}
	err = conn.DB("").C("minedBlocks").Find(query).Sort("-block").Limit(limit).All(&blocks)
	if err != nil {
		return page, err
	}
	if len(blocks) > 0 {
		page.Start = blocks[0].Block.String()
		page.End = blocks[len(blocks)-1].Block.String()
	}
	page.Blocks = blocks
	return page, nil
}

func (dao *BlockDAO) Get(blockNumber *big.Int) (*types.Block, error) {
	return node.Client().BlockByNumber(context.TODO(), blockNumber)
}
