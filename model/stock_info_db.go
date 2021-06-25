package model

import (
	"time"
)

// DayTimeStockInfo [...]
type DayTimeStockInfo struct {
	ID                int       `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Time              time.Time `gorm:"column:time;type:datetime" json:"time"`
	StockName         string    `gorm:"column:stockName;type:varchar(50)" json:"stock_name"`
	StockCode         string    `gorm:"column:stockCode;type:varchar(10)" json:"stock_code"`
	StockKind         string    `gorm:"column:stockKind;type:varchar(30)" json:"stock_kind"`
	DayInfo           string    `gorm:"column:dayInfo;type:varchar(10)" json:"day_info"`
	DayStockValue     int       `gorm:"column:dayStockValue;type:int" json:"day_stock_value"`
	CurrentStockValue int       `gorm:"column:currentStockValue;type:int" json:"current_stock_value"`
}