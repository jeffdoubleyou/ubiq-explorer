package main

import (
	"log"
	"ubiq-explorer/daos"
	"ubiq-explorer/services"
)

func main() {
	exchange := services.NewExchangeService()
	tokenDAO := daos.NewTokenDAO()
	tokenService := services.NewTokenService(*tokenDAO)

	limit := 10
	cursor := ""
	symbolList := make(map[string]int)
	symbolList["BTC"] = 1
	symbolList["UBQ"] = 1

	tokenList, _ := tokenService.TokenList(limit, cursor)

	for tokenList.Total > 0 && tokenList.Start != "" {
		for _, token := range tokenList.Tokens {
			symbolList[token.Symbol] = 1
		}
		cursor = tokenList.End
		tokenList, _ = tokenService.TokenList(limit, cursor)
	}

	for symbol := range symbolList {
		log.Printf("Get exchange for %s", symbol)
		exchangeRate, err := exchange.UpdateExchangeRate(symbol)
		if err != nil {
			log.Printf("Failed to retrieve exchange data for %s: %s", symbol, err)
		} else {
			log.Printf("Exchange rate for %s - BTC: %f USD: %f", symbol, exchangeRate.Btc, exchangeRate.Usd)
		}
	}
}
