package stockcode

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"realStock/database"
	"realStock/model"
	"time"
)

type StockResponseValueInfo struct {
	stockCode, stockName, stockKind, dayInfo string
	dayStockValue, currentStockValue         int
}


/*
func (info *StockResponseValueInfo) TradeInfo(infoValue StockResponseValueInfo) {

	fmt.Println(infoValue)

	dataSource := database.GetDBInfo()
	db, err := sql.Open(database.GetDBEngin(), dataSource)

	if err != nil {
		log.Fatal(err.Error())
	}

	now := time.Now().Local()
	date := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
	selectQuery := fmt.Sprintf("SELECT currentStockValue from dayTimeStockInfo where 1=1 AND id = (SELECT Max(id) from dayTimeStockInfo where stockCode = '%s' AND time like  '%%%s%%')", infoValue.stockCode, date)

	rows, err := db.Query(selectQuery)

	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		var currentStockValue int
		err := rows.Scan(&currentStockValue)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if currentStockValue == infoValue.currentStockValue {
			rows.Close()
			db.Close()
			return
		}

		rows.Close()
	}

	tx, _ := db.Begin()

	log.Println("Inserting tradeInfo record ... start ")
	query := `INSERT INTO dayTimeStockInfo(time, stockName, stockCode, stockKind, dayInfo, dayStockValue, currentStockValue) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatalln(err.Error())
	}

	time := fmt.Sprintf("%d-%02d-%02d-%02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	_, err = statement.Exec(time,
		infoValue.stockName,
		infoValue.stockCode,
		infoValue.stockKind,
		infoValue.dayInfo,
		infoValue.dayStockValue,
		infoValue.currentStockValue,
	)

	if err != nil {
		log.Fatalln(err.Error())
	}

	tx.Commit()

	defer db.Close()
	log.Println("Inserting tradeInfo record ... end ")
}
 */


func (info *StockResponseValueInfo) TradeInfo(infoValue StockResponseValueInfo) {

	db := database.GetDB()
	now := time.Now().Local()
	date := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())

	stockInfo := model.DayTimeStockInfoMgr(db)
	dayStockInfo, err := stockInfo.GetCurrentStockInfo(infoValue.stockCode, date)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//result := fmt.Sprintf("%s-%d", dayStockInfo.Time.String(), dayStockInfo.CurrentStockValue)
	//fmt.Println(result)

	if dayStockInfo.CurrentStockValue == infoValue.currentStockValue {
		return
	}

	result := fmt.Sprintf("%d - %d", dayStockInfo.CurrentStockValue, infoValue.currentStockValue)
	fmt.Println(result)

	log.Println("Inserting tradeInfo record ... start ")

	currentTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(),
																	 now.Month(),
																	 now.Day(),
																	 now.Hour(),
																	 now.Minute(),
																	 now.Second())
	t, _ := time.Parse("2006-01-02 15:04:05", currentTime)

	//fmt.Println(t)

	stockValue := model.DayTimeStockInfo{
		Time: t,
		StockName: infoValue.stockName,
		StockCode: infoValue.stockCode,
		StockKind: infoValue.stockKind,
		DayInfo: infoValue.dayInfo,
		DayStockValue: infoValue.dayStockValue,
		CurrentStockValue: infoValue.currentStockValue,
	}

	//fmt.Println(stockValue.Time)

	tx := stockInfo.DB.Begin()
	if err := tx.Create(&stockValue).Error; err != nil {
		tx.Rollback()
		log.Println("Rollback")
	} else {
		tx.Commit()
		log.Println("Commit")
	}

	log.Println("Inserting tradeInfo record ... end ")
}
