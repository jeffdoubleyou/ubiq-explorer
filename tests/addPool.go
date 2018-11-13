package main

import (
	"fmt"
	"os"
	"strconv"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/models/db"
)

func main() {
    defer db.Close()
	input := os.Args
	dao := daos.NewPoolsDAO()
	miners, _ := strconv.ParseFloat(input[4], 10)
	hashrate, _ := strconv.ParseFloat(input[5], 1)
	uptime, _ := strconv.ParseFloat(input[6], 10)
	count, _ := strconv.ParseUint(input[8], 10, 64)
	onlineCount, _ := strconv.ParseUint(input[9], 10, 64)

	pool := models.Pool{
		Name:        input[1],
		Url:         input[2],
		StatsUrl:    input[3],
		Miners:      miners,
		Hashrate:    hashrate,
		Uptime:      uptime,
		Software:    input[7],
		Count:       count,
		OnlineCount: onlineCount,
	}
	_, err := dao.Insert(&pool)
	if err != nil {
		fmt.Printf("Failed to insert pool: %s\n", err)
	} else {
		fmt.Println("Inserted pool")
	}
}
