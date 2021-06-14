package database

import (
	"fmt"
	"realStock/config"
)

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