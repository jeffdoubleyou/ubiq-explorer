package main

import (
	"fmt"
	"ubiq-explorer/daos"
    "ubiq-explorer/services"
    "os"
)

func main() {
	dao := daos.NewAddressDAO()
    addressService := services.NewAddressService(*dao)
    addresses, err := addressService.List()
    if err != nil {
        fmt.Printf("Failed to get list %s", err)
    } else {
        out, err := os.Create("tools/addresses.sh")
        if err != nil {
            panic(fmt.Sprintf("Failed to open file: %s\n", err))
        }
        defer out.Close()
        for _, address := range(addresses) {
            out.WriteString(fmt.Sprintf("go run tools/addAddress.go %s '%s'\n", address.Address.String(), address.Name))
        }
    }
    fmt.Printf("Wrote addresses to tools/addresses.sh")
}
