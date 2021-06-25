package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"realStock/config"
)


var (
	DB* gorm.DB
)


/*
func GetDBInfo() string {
	configuration := config.GetConfig()

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?allowNativePasswords=true",
													configuration.DB_USERNAME,
													configuration.DB_PASSWORD,
													configuration.DB_HOST,
													configuration.DB_NAME)
}

func GetDBEngin() string  {
	configuration := config.GetConfig()
	return configuration.DB_ENGIN
}
 */


func InitDB() *gorm.DB {
	configuration := config.GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowNativePasswords=true",
		configuration.DB_USERNAME,
		configuration.DB_PASSWORD,
		configuration.DB_HOST,
		configuration.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}