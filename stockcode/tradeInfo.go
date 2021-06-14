package stockcode

import (
	"fmt"
	"realStock/config"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type TradeInfo struct {
	isRequestCall 	bool
	stockCodeArray [] StockInfo
}

func (info * TradeInfo) GetCallTradeInfo() bool  {
	return info.isRequestCall
}

func (info * TradeInfo) SetTradeInfo() {
	info.stockCodeArray = GetStockInfo()
}

func (info * TradeInfo)GetTradeInfo()  {

	info.isRequestCall = true
	for index := range info.stockCodeArray {
		stockInfo := info.stockCodeArray[index]

		info.requestTradeInfo(stockInfo)

		time.Sleep(500 * time.Microsecond)
	}

	info.isRequestCall = false

}

func (info * TradeInfo) requestTradeInfo(codeInfo StockInfo) {

	c := colly.NewCollector(
		colly.AllowedDomains("finance.naver.com", "https://finance.naver.com"),
	)
	infoCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("dl", func(e *colly.HTMLElement) {

		var tempInfo StockResponseValueInfo

		e.ForEach("dd", func(index int, dd *colly.HTMLElement) {
			switch index {
			case config.FirstTagIndex:
				tempValue := dd.Text
				if strings.Contains(tempValue, "종목명") {
					tempInfo.stockName = codeInfo.codeName
				}

			case config.SecondTagIndex:
				tempValue := dd.Text
				if strings.Contains(tempValue, "종목코드") {
					code := strings.Split(tempValue, " ")[1]
					tempInfo.stockCode = strings.Trim(code, " ")

					kind := strings.Split(tempValue, " ")[2]
					tempInfo.stockKind = strings.Trim(kind, " ")
				}

			case config.ThirdTagIndex:
				tempValue := dd.Text
				if strings.Contains(tempValue, "현재가") {
					var spliteValue = strings.Split(tempValue, " ")

					var valueInfo = strings.Replace(spliteValue[3], ",", "", -1)
					var value = strings.Replace(spliteValue[4], ",", "", -1)

					if valueInfo == "하락" {
						value = fmt.Sprintf("-%s", value)
					}

					tempTradeValue, _ := strconv.Atoi(strings.Trim(value, " "))
					currentValue, _ := strconv.Atoi(strings.Trim(strings.Replace(spliteValue[1], ",", "", -1), " "))

					if tempInfo.stockCode == codeInfo.code {

						var infoValue = StockResponseValueInfo{
							stockCode:         strings.Trim(codeInfo.code, " "),
							stockName:         strings.Trim(codeInfo.codeName, " "),
							dayInfo:           strings.Trim(valueInfo, " "),
							stockKind:         tempInfo.stockKind,
							dayStockValue:     tempTradeValue,
							currentStockValue: currentValue,
						}

						infoValue.TradeInfo(infoValue)
					}
				}
			}
		})
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Profile URL: ", r.URL.String())
	})

	url := fmt.Sprintf("https://finance.naver.com/item/main.nhn?code=%s", codeInfo.code)

	c.Visit(url)
}


