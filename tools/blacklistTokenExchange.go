package main

import (
	"fmt"
	"os"
	"ubiq-explorer/daos"
)

func main() {
	input := os.Args
	dao := daos.NewExchangeDAO()
	token, err := dao.GetExchangeSource(input[1])
	if err != nil {
		fmt.Printf("Failed to find a token with symbol %s", input[1])
		os.Exit(1)
	}
	token.Blacklist = true
	_, err = dao.InsertExchangeSource(token)

	if err != nil {
		fmt.Printf("Failed to update token: %s\n", err)
	} else {
		fmt.Println("Updated token")
	}

    _, err = dao.DeleteExchangeRate(input[1])

    if err != nil {
        fmt.Printf("Failed to remove exchange rate for %s: %s", input[1], err)
    } else {
        fmt.Println("Removed exchange rate")
    }
}
