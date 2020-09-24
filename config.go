package main

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

type Config struct {
	TelegramBotKey string
	FinhubApiKey   string
}

func GetConfig() Config {
	configuration := Config{}

	err := gonfig.GetConf("./config.json", &configuration)
	if err != nil {
		fmt.Printf("Config file not found")
	}

	return configuration
}
