package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
    "ubiq-explorer/models/db"
)

func main() {
	input := os.Args
	dao := daos.NewAddressDAO()
    defer db.Close()
	_, err := dao.Insert(models.AddressInfo{common.HexToAddress(input[1]), input[2]})
	if err != nil {
		fmt.Printf("Failed to insert address: %s\n", err)
	} else {
		fmt.Println("Inserted address")
	}
}
