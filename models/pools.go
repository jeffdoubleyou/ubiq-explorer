package models

type Pool struct {
	Name        string  `json:"name"`
	Url         string  `json:"url"`
	StatsUrl    string  `json:"statsUrl"`
	Miners      float64 `json:"miners"`
	Hashrate    float64 `json:"hashrate"`
	Uptime      float64 `json:"uptime"`
	Software    string  `json:"software"`
	Count       uint64  `json:"count"`
	OnlineCount uint64  `json:"onlineCount"`
}
