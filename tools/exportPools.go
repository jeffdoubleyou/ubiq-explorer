package main

import (
	"fmt"
	"ubiq-explorer/daos"
    "ubiq-explorer/services"
    "os"
)

func main() {
	dao := daos.NewPoolsDAO()
    poolService := services.NewPoolsService(*dao)
    pools, err := poolService.List()
    if err != nil {
        fmt.Printf("Failed to get list %s", err)
    } else {
        out, err := os.Create("scripts/pools.sh")
        if err != nil {
            panic(fmt.Sprintf("Failed to open file: %s\n", err))
        }
        defer out.Close()
        for _, pool := range(pools) {
            fmt.Printf("go run tools/addPool.go '%s' '%s' '%s' %f %f %f %s %d %d\n", pool.Name, pool.Url, pool.StatsUrl, pool.Miners, pool.Hashrate, pool.Uptime, pool.Software, pool.Count, pool.OnlineCount)
            out.WriteString(fmt.Sprintf("go run tools/addPool.go '%s' '%s' '%s' %f %f %f %s %d %d\n", pool.Name, pool.Url, pool.StatsUrl, pool.Miners, pool.Hashrate, pool.Uptime, pool.Software, pool.Count, pool.OnlineCount))
        }
    }
    fmt.Printf("Wrote pools to scripts/pools.sh")
}

/*
go run tools/addPool.go 'ubiqpool.fr' 'https://ubiqpool.fr' 'https://ubiqpool.fr/api/stats' 0 0 99.5111352525801 open-ethereum 55230 54960


Name        string  `json:"name"`
Url         string  `json:"url"`
StatsUrl    string  `json:"statsUrl"`
Miners      float64 `json:"miners"`
Hashrate    float64 `json:"hashrate"`
Uptime      float64 `json:"uptime"`
Software    string  `json:"software"`
Count       uint64  `json:"count"`
OnlineCount uint64  `json:"onlineCount"
*/
