package main

import (
	"fmt"
	"os"
	"ubiq-explorer/daos"
)

func main() {
	input := os.Args
	dao := daos.NewTokenDAO()
	token, err := dao.GetTokenByAddress(input[1])
	if err != nil {
		fmt.Printf("Failed to find a token at %s", input[1])
		os.Exit(1)
	}
	token.Verified = true
	_, err = dao.Insert(token)

	if err != nil {
		fmt.Printf("Failed to update token: %s\n", err)
	} else {
		fmt.Println("Updated token")
	}
}
