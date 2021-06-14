package stockcode

import (
	"bufio"
	"log"
	"os"
	"realStock/config"
	"strings"
)

type StockInfo struct {
	code, codeName string
}

func (info* StockInfo) GetCode() string {
	return info.code
}

func (info* StockInfo) GetCodeName() string  {
	return info.codeName
}

func GetStockInfo() [] StockInfo {
	var stockCodeArray [] StockInfo

	readFile, err := os.Open(config.TradeInfoFile)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		data := strings.Split(fileScanner.Text(), ",")

		stockCode := StockInfo{
			code:     data[0],
			codeName: data[1],
		}

		stockCodeArray = append(stockCodeArray, stockCode)
	}

	readFile.Close()

	return stockCodeArray
}