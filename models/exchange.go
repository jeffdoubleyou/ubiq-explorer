package models

import ()

type ExchangeRate struct {
	Symbol    string  `json:"symbol"`
	Btc       float64 `json:"btc"`
	Usd       float64 `json:"usd"`
	Timestamp int64   `json:"timestamp"`
	Source    string  `json:"source"`
}

type ExchangeSource struct {
	Symbol    string `json:"symbol"`
	Exchange  string `json:"exchange"`
	Timestamp int64  `json:"timestamp"`
	Blacklist bool   `json:"blacklist"`
}

type ExchangeAuth struct {
	Exchange  string `json:"exchange"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

type ExchangeRateList struct {
	Symbols []*ExchangeRate
}

type ExchangeRateHistory struct {
	History []*ExchangeRate
}
