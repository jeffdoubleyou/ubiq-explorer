package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"ubiq-explorer/services"
)

func main() {
	exchange := services.NewExchangeService()

	pools, err := poolService.List()

	if err != nil {
		panic(err)
	}

	for _, pool := range pools {
		pool.Count++
		log.Printf("Checking on %s: ", pool.Name)
		json, err := GetStats(pool)
		if err != nil {
			log.Printf("ERR: %s\n", err)
			pool.Uptime = float64(float64(pool.OnlineCount)/float64(pool.Count)) * 100
			pool.Hashrate = 0
			pool.Miners = 0
			_, err := poolService.Insert(&pool)
			if err != nil {
				log.Printf("FAILED TO UPDATE POOL: %s", err)
			}
		} else {
			pool.OnlineCount++
			pool.Uptime = float64(float64(pool.OnlineCount)/float64(pool.Count)) * 100
			pool.Miners = json.Miners
			pool.Hashrate = json.Hashrate
			log.Printf("ONLINE: %f UPTIME: %f\n", json.Miners, pool.Uptime)
			_, err := poolService.Insert(&pool)
			if err != nil {
				log.Printf("FAILED TO UPDATE POOL: %s", err)
			}
		}
	}
}

func GetStats(pool models.Pool) (models.Pool, error) {
	req := httplib.Get(pool.StatsUrl).SetTimeout(10*time.Second, 10*time.Second)
	req.Header("User-Agent", beego.AppConfig.String("pool_monitor::user_agent"))
	var tlsConfig = &tls.Config{VerifyPeerCertificate: nil, InsecureSkipVerify: true}
	req.SetTLSClientConfig(tlsConfig)

	switch pool.Software {
	case "open-ethereum":
		res := &OpenEthereum{}
		if err := req.ToJSON(&res); err != nil {
			return pool, err
		}
		pool.Hashrate = res.Hashrate
		pool.Miners = res.Miners
	case "mpos":
		res := &MPOS{}
		if err := req.ToJSON(&res); err != nil {
			return pool, err
		}
		pool.Hashrate = res.Getpoolstatus.Data.Hashrate * 1000
		pool.Miners = res.Getpoolstatus.Data.Miners
	case "kings":
		res := &King{}
		if err := req.ToJSON(&res); err != nil {
			return pool, err
		}
		pool.Hashrate = res.Totals.Hashrate
		pool.Miners = res.Totals.Miners
	case "minerall":
		res := &Minerall{}
		if err := req.ToJSON(&res); err != nil {
			return pool, err
		}
		pool.Hashrate = res.Hashrate
		pool.Miners = res.Miners
	default:
		return pool, fmt.Errorf("Invalid or undefined pool software")
	}

	return pool, nil
}

func UpdateStats(pools []*models.Pool) error {
	return nil
}

func SetDown(pool *models.Pool) error {
	return nil
}
