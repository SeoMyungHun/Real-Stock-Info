package main

import (
	"realStock/database"
	"realStock/stockcode"
	"time"
)

func main() {

	database.InitDB()

	tradeInfo := stockcode.TradeInfo{}
	tradeInfo.SetTradeInfo()

	duration, _ := time.ParseDuration("5s")
	ticker := time.NewTicker(duration)

	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:

			tradeInfo.GetTradeInfo()
		}
	}
}
