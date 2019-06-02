package app

import (
	"github.com/findata-code/fastvault-client-go"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

const (
	FASTVAULT_LOCATION  = "http://128.199.147.139:9800"
	ENV_FASTVAULT_TOKEN = "FV_TOKEN"
)

var (
	Config Configuration
	Bot    *linebot.Client
)



func init() {
	err := getConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	Bot = newLineBot()
}

func getConfiguration() error {
	token := os.Getenv(ENV_FASTVAULT_TOKEN)
	if token == "" {
		log.Fatal("Could not read fastvault token from env variable")
	}
	fv := fastvault_client_go.New(FASTVAULT_LOCATION)
	err := fv.GetJson(token, &Config)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func newLineBot() *linebot.Client {
	client, err := linebot.New(Config.LineBot.Secret, Config.LineBot.Token)
	if err != nil {
		log.Panic(err)
	}

	return client
}
