package model

import (
	"fmt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

type _DayTimeStockInfoMgr struct {
	*_BaseMgr
}

func DayTimeStockInfoMgr(db *gorm.DB) *_DayTimeStockInfoMgr {
	if db == nil {
		panic(fmt.Errorf("DayTimeStockInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DayTimeStockInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("dayTimeStockInfo"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_DayTimeStockInfoMgr) GetTableName() string {
	return "dayTimeStockInfo"
}


func (obj *_DayTimeStockInfoMgr) Get() (result DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}


func (obj *_DayTimeStockInfoMgr) Gets() (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_DayTimeStockInfoMgr) GetID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

func (obj *_DayTimeStockInfoMgr) GetTime(time time.Time) Option {
	return optionFunc(func(o *options) { o.query["time"] = time })
}


func (obj *_DayTimeStockInfoMgr) GetStockName(stockName string) Option {
	return optionFunc(func(o *options) { o.query["stockName"] = stockName })
}


func (obj *_DayTimeStockInfoMgr) GetStockCode(stockCode string) Option {
	return optionFunc(func(o *options) { o.query["stockCode"] = stockCode })
}

func (obj *_DayTimeStockInfoMgr) GetStockKind(stockKind string) Option {
	return optionFunc(func(o *options) { o.query["stockKind"] = stockKind })
}

func (obj *_DayTimeStockInfoMgr) GetDayInfo(dayInfo string) Option {
	return optionFunc(func(o *options) { o.query["dayInfo"] = dayInfo })
}

// WithDayStockValue dayStockValue获取
func (obj *_DayTimeStockInfoMgr) WithDayStockValue(dayStockValue int) Option {
	return optionFunc(func(o *options) { o.query["dayStockValue"] = dayStockValue })
}

func (obj *_DayTimeStockInfoMgr) GetCurrentStockValue(currentStockValue int) Option {
	return optionFunc(func(o *options) { o.query["currentStockValue"] = currentStockValue })
}

func (obj *_DayTimeStockInfoMgr) GetByOption(opts ...Option) (result DayTimeStockInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}


func (obj *_DayTimeStockInfoMgr) GetByOptions(opts ...Option) (results []* DayTimeStockInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_DayTimeStockInfoMgr) GetFromID(id int) (result DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

func (obj *_DayTimeStockInfoMgr) GetBatchFromID(ids []int) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	return
}

func (obj *_DayTimeStockInfoMgr) GetFromTime(time time.Time) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time = ?", time).Find(&results).Error

	return
}

func (obj *_DayTimeStockInfoMgr) GetBatchFromTime(times []time.Time) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time IN (?)", times).Find(&results).Error

	return
}


func (obj *_DayTimeStockInfoMgr) GetFromStockName(stockName string) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stockName = ?", stockName).Find(&results).Error

	return
}


func (obj *_DayTimeStockInfoMgr) GetBatchFromStockName(stockNames []string) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stockName IN (?)", stockNames).Find(&results).Error
	return
}


func (obj *_DayTimeStockInfoMgr) GetFromStockCode(stockCode string) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stockCode = ?", stockCode).Find(&results).Error

	return
}


func (obj *_DayTimeStockInfoMgr) GetBatchFromStockCode(stockCodes []string) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stockCode IN (?)", stockCodes).Find(&results).Error

	return
}


func (obj *_DayTimeStockInfoMgr) GetFromStockKind(stockKind string) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stockKind = ?", stockKind).Find(&results).Error

	return
}


func (obj *_DayTimeStockInfoMgr) GetBatchFromStockKind(stockKinds []string) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stockKind IN (?)", stockKinds).Find(&results).Error
	return
}


func (obj *_DayTimeStockInfoMgr) GetFromDayInfo(dayInfo string) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dayInfo = ?", dayInfo).Find(&results).Error

	return
}


func (obj *_DayTimeStockInfoMgr) GetBatchFromDayInfo(dayInfos []string) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dayInfo IN (?)", dayInfos).Find(&results).Error

	return
}

func (obj *_DayTimeStockInfoMgr) GetFromDayStockValue(dayStockValue int) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dayStockValue = ?", dayStockValue).Find(&results).Error

	return
}

func (obj *_DayTimeStockInfoMgr) GetBatchFromDayStockValue(dayStockValues []int) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("dayStockValue IN (?)", dayStockValues).Find(&results).Error

	return
}

func (obj *_DayTimeStockInfoMgr) GetFromCurrentStockValue(currentStockValue int) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("currentStockValue = ?", currentStockValue).Find(&results).Error
	return
}

func (obj *_DayTimeStockInfoMgr) GetBatchFromCurrentStockValue(currentStockValues []int) (results []* DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("currentStockValue IN (?)", currentStockValues).Find(&results).Error
	return
}

func (obj *_DayTimeStockInfoMgr) FetchByPrimaryKey(id int) (result DayTimeStockInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	return
}

func (obj *_DayTimeStockInfoMgr)GetCurrentStockInfo(code, date string)(result DayTimeStockInfo, err error) {
	selectQuery := fmt.Sprintf("SELECT currentStockValue from dayTimeStockInfo where 1=1 AND id = (SELECT Max(id) from dayTimeStockInfo where stockCode = '%s' AND time like  '%%%s%%')", code, date)
	err = obj.DB.Raw(selectQuery).Scan(&result).Error
	//err = obj.DB.Exec(selectQuery).Scan(&result).Error
	return
}