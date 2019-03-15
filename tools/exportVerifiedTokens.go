package main

import (
	"fmt"
	"ubiq-explorer/daos"
	"gopkg.in/mgo.v2/bson"
    "os"
    "strings"
)

func main() {
	dao := daos.NewTokenDAO()
    tokens, err := dao.FindTokens(bson.M{}, "-_id", 1000, "")
    if err != nil {
        fmt.Printf("Failed to get list %s", err)
    } else {
        out, err := os.Create("scripts/verified_tokens.sh")
        if err != nil {
            panic(fmt.Sprintf("Failed to open file: %s\n", err))
        }
        defer out.Close()
        for _, token := range(tokens.Tokens) {
            if token.Verified == true {
                out.WriteString(fmt.Sprintf("go run tools/verifyToken.go %s\n", strings.ToLower(token.Address.String())))
            }
        }
    }
    fmt.Printf("Wrote tokenes to scripts/verified_tokens.sh\n")
}
