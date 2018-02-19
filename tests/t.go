package main

import (
	"fmt"
	//"github.com/ethereum/go-ethereum/ethclient"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//	"math/big"
	//"ubiq-explorer/daos"
	//"ubiq-explorer/models/db"
	//"ubiq-explorer/services"
	"ubiq-explorer/services"
)

func main() {
	s := services.NewStatsService()
	x, err := s.MinerList(100)
	fmt.Printf("ERR: %s, MINED: %v", err, x)
	for _, m := range x {
		fmt.Printf("CNT: %d, P: %v\n", m.Count, m.Percent)
	}
}
