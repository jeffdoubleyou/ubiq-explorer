package models

import ()

type ExchangeRate struct {
	Name      string  `json:"name"`
	Symbol    string  `json:"symbol"`
	Btc       float32 `json:"btc"`
	Usd       float32 `json:"usd"`
	Timestamp uint32  `json:"timestamp"`
	Source    string  `json:"source"`
}

type ExchangeSource struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Market   string `json:"market"`
}

type ExchangeAuth struct {
	Exchange  string `json:"exchange"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

type ExchangeList struct {
	Symbols []*ExchangeRate
}

type ExchangeRateHistory struct {
	History []*ExchangeRate
}
