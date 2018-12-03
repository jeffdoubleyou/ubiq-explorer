package exchange

import (
	"github.com/jeffdoubleyou/go-cryptopia"
	"time"
	"ubiq-explorer/models"
)

type CryptopiaExchange struct {
	API cryptopia.Cryptopia
}

func NewCryptopiaExchange(apiKey string, apiSecret string) *CryptopiaExchange {
	API := cryptopia.New(apiKey, apiSecret)
	return &CryptopiaExchange{*API}
}

func (exchange *CryptopiaExchange) GetTicker(symbol string) (*models.ExchangeRate, error) {
	ticker := symbol + "_BTC"
	if symbol == "BTC" {
		ticker = "USDT_BTC"
	}
	market, err := exchange.API.GetMarket(ticker, 24)
	exchangeRate := &models.ExchangeRate{
		Symbol:    symbol,
		Source:    "cryptopia",
		Timestamp: time.Now().Unix(),
	}
	exchangeRate.Btc = market.LastPrice
	return exchangeRate, err
}
