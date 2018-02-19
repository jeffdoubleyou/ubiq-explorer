package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
)

func main() {
	input := os.Args
	dao := daos.NewAddressDAO()
	_, err := dao.Insert(models.AddressInfo{common.HexToAddress(input[1]), input[2]})
	if err != nil {
		fmt.Printf("Failed to insert address: %s\n", err)
	} else {
		fmt.Println("Inserted address")
	}
}
