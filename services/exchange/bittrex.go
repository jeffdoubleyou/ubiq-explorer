package exchange

import (
	"github.com/toorop/go-bittrex"
	"time"
	"ubiq-explorer/models"
)

type BittrexExchange struct {
	API bittrex.Bittrex
}

func NewBittrexExchange(apiKey string, apiSecret string) *BittrexExchange {
	API := bittrex.New(apiKey, apiSecret)
	return &BittrexExchange{*API}
}

func (exchange *BittrexExchange) GetTicker(symbol string) (*models.ExchangeRate, error) {
	ticker := "BTC-" + symbol
	if symbol == "BTC" {
		ticker = "USDT-BTC"
	}
	market, err := exchange.API.GetTicker(ticker)
	exchangeRate := &models.ExchangeRate{
		Symbol:    symbol,
		Source:    "bittrex",
		Timestamp: time.Now().Unix(),
	}
	if symbol == "BTC" {
		exchangeRate.Usd, _ = market.Last.Float64()
	} else {
		exchangeRate.Btc, _ = market.Last.Float64()
	}
	return exchangeRate, err
}
