package main

import (
	"fmt"
	"realStock/stockcode"
	"time"
)

func main() {

	tradeInfo := stockcode.TradeInfo{}
	tradeInfo.SetTradeInfo()

	duration, _ := time.ParseDuration("5s")
	ticker := time.NewTicker(duration)

	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:

			fmt.Println(tradeInfo.GetCallTradeInfo())
			if !tradeInfo.GetCallTradeInfo() {
				tradeInfo.GetTradeInfo()
			}
		}
	}
}
