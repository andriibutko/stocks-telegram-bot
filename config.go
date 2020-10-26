package main

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

// Config contains config props.
type Config struct {
	TelegramBotKey string
	FinhubAPIKey   string
}

// GetConfig Gets config.
func GetConfig() Config {
	configuration := Config{}

	err := gonfig.GetConf("./config.json", &configuration)
	if err != nil {
		fmt.Printf("Config file not found")
	}

	return configuration
}
