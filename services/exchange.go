package services

import (
	"errors"
	"fmt"
	"time"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services/exchange"
)

type Exchange struct {
	API interface {
		GetTicker(symbol string) (*models.ExchangeRate, error)
	}
}

type ExchangeService struct {
	dao       *daos.ExchangeDAO
	exchanges map[string]*Exchange
}

func NewExchangeService() *ExchangeService {
	dao := daos.NewExchangeDAO()
	exchanges := make(map[string]*Exchange)
	exchanges["bittrex"] = &Exchange{exchange.NewBittrexExchange("", "")}
	exchanges["cryptopia"] = &Exchange{exchange.NewCryptopiaExchange("", "")}
	return &ExchangeService{
		dao,
		exchanges,
	}
}

func (s *ExchangeService) Get(symbol string) (*models.ExchangeRate, error) {
	if symbol == "" {
		return nil, errors.New("Symbol is required")
	}
	return s.dao.GetExchangeRate(symbol)

}

func (s *ExchangeService) List() ([]*models.ExchangeRate, error) {
	return s.dao.ExchangeRateList()
}

func (s *ExchangeService) History(symbol string) ([]*models.ExchangeRate, error) {
	if symbol == "" {
		return nil, errors.New("Symbol is required")
	}
	return s.dao.ExchangeRateHistory(symbol)
}

func (s *ExchangeService) UpdateExchangeRate(symbol string) (*models.ExchangeRate, error) {
	exchangeSource, _ := s.dao.GetExchangeSource(symbol)
	if exchangeSource.Exchange != "" {
		market, err := s.exchanges[exchangeSource.Exchange].API.GetTicker(symbol)
		if err != nil || market.Btc == 0 {
			fmt.Printf("Failed to retrieve %s from existing exchange %s", symbol, exchangeSource.Exchange)
		} else {
			s.dao.InsertExchangeRate(market, 0)
			return market, nil
		}
	}
	if exchangeSource.Timestamp > 0 && time.Now().Unix()-exchangeSource.Timestamp < 86400 {
		return nil, fmt.Errorf("Symbol %s does not exist at any exchange in the last 24 hours", symbol)
	}
	exchangeSource.Timestamp = time.Now().Unix()
	exchangeSource.Symbol = symbol
	for exchange := range s.exchanges {
		market, _ := s.exchanges[exchange].API.GetTicker(symbol)
		if market.Btc != 0 {
			exchangeSource.Exchange = exchange
			s.dao.InsertExchangeSource(exchangeSource)
			s.dao.InsertExchangeRate(market, 0)
			return market, nil
		}
	}
	s.dao.InsertExchangeSource(exchangeSource)
	return nil, fmt.Errorf("Unable to find an exchange which supports %s", symbol)
}
