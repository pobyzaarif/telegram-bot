package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type AppSetting struct {
	TelegramBotPostURL  string `json:"telegram_bot_post_url"`
	ChatID              string `json:"chat_id"`
	DisableNotification bool   `json:"disable_notification"`
}

type Config struct {
	AppSetting AppSetting `json:"app_setting"`
}

func LoadConfig(filename string) *Config {
	// Read the JSON file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read config file: %v", err))
	}

	// Unmarshal the JSON data into the config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to unmarshal config JSON: %v", err))
	}

	return &config
}
