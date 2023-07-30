package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	goLoggerHTTPClient "github.com/pobyzaarif/go-logger/http/client"
	"github.com/pobyzaarif/telegram-bot/config"
)

type BodyPayload struct {
	ChatID              string `json:"chat_id"`
	Text                string `json:"text"`
	DisableNotification bool   `json:"disable_notification"`
}

func main() {
	messagePtr := flag.String("message", "this is test message", "message that will send to bot")
	flag.Parse()

	conf := config.LoadConfig("./config.json")

	toURL := conf.AppSetting.TelegramBotPostURL

	bodyPayload := BodyPayload{
		ChatID:              conf.AppSetting.ChatID,
		DisableNotification: conf.AppSetting.DisableNotification,
		Text:                *messagePtr,
	}

	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(bodyPayload)
	if err != nil {
		fmt.Println(err)
		return
	}

	request, err := http.NewRequest(http.MethodPost, toURL, payload)
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/json")
	timeout := time.Second * 5
	var response string

	httpCode, err := goLoggerHTTPClient.Call(
		context.Background(),
		request,
		timeout,
		goLoggerHTTPClient.RawResponseBodyFormat,
		&response,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	if httpCode != 200 {
		fmt.Println("error response HTTP code is not 200 " + toURL)
		return
	}

	fmt.Println("OK")
}