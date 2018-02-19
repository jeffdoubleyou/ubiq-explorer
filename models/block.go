package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
)

func init() {
	//TransactionList = make(map[string]*Transaction)
	//u := Transaction{"user_11111", "astaxie", "11111", "literally nothing"}
	//TransactionList["user_11111"] = &u
}

type BlockPage struct {
	Start  string
	End    string
	Total  int
	Blocks []*Miner
}

type RecentBlock struct {
	Block     string `json:"block"`
	Miner     string `json:"miner"`
	Timestamp string `json:"timestamp"`
}

type Header struct {
	ParentHash  string             `json:"parentHash"       gencodec:"required"`
	Hash        common.Hash        `json:"hash"       gencodec:"required"`
	UncleHash   string             `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    string             `json:"miner"            gencodec:"required"`
	Root        string             `json:"stateRoot"        gencodec:"required"`
	TxHash      string             `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash string             `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom        `json:"logsBloom"        gencodec:"required"`
	Difficulty  string             `json:"difficulty"       gencodec:"required"`
	Number      string             `json:"number"           gencodec:"required"`
	GasLimit    uint64             `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64             `json:"gasUsed"          gencodec:"required"`
	Time        uint64             `json:"timestamp"        gencodec:"required"`
	Extra       string             `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash        `json:"mixHash"          gencodec:"required"`
	Nonce       uint64             `json:"nonce"            gencodec:"required"`
	Size        common.StorageSize `json:"size"`
}

func NewBlockHeader(header *types.Block) Header {
	return Header{
		ParentHash:  header.ParentHash().String(),
		Hash:        header.Hash(),
		UncleHash:   header.UncleHash().String(),
		Coinbase:    strings.ToLower(header.Coinbase().String()),
		Root:        header.Root().String(),
		TxHash:      header.TxHash().String(),
		ReceiptHash: header.ReceiptHash().String(),
		Bloom:       header.Bloom(),
		Difficulty:  header.Difficulty().String(),
		Number:      header.Number().String(),
		GasLimit:    header.GasLimit(),
		GasUsed:     header.GasUsed(),
		Time:        header.Time().Uint64(),
		Extra:       hexutil.Encode(header.Extra()),
		MixDigest:   header.MixDigest(),
		Nonce:       header.Nonce(),
		Size:        header.Size(),
	}
}

/*func (h Header) MarshalJSON() ([]byte, error) {
	return json.Marshal(NewBlockHeader(h))
}*/
