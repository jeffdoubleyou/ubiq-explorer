package node

/*
	Instead of having to import ethclient everywhere and dial up a connection, just use this
*/

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var rc *rpc.Client
var node *ethclient.Client

func init() {
	url := beego.AppConfig.String("node::url")
	client, _ := ethclient.Dial(url)
	rpcClient, _ := rpc.Dial(url)

	node = client
	rc = rpcClient
}

func Client() *ethclient.Client {
	return node
}

func RPC() *rpc.Client {
	return rc
}
