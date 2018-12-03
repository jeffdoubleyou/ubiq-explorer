package main

import (
	"fmt"
	"math/big"
)

func main() {
	blockNumber := big.NewInt(365472)
	fmt.Println(blockNumber.Text(16))
	bn, ok := blockNumber.SetString("0x593a0", 0)
	if ok != true {
		fmt.Printf("COULD NOT PARSE %s", "0x593a0")
	}
	fmt.Println(bn)
	fmt.Println(blockNumber)

}
