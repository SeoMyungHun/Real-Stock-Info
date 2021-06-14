package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	DB_ENGIN    string
}


func GetConfig() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf("config/config.json", &configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
