package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	HOST      string
	USER_NAME string
	PASSWORD  string
	DB_NAME   string
	PORT      int
}

func GetConfig() Configuration {
	conf := Configuration{}

	gonfig.GetConf("./config/config.json", &conf)

	return conf
}
