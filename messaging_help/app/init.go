package app

import (
	"errors"
	"github.com/findata-code/fastvault-client-go"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

const (
	FastvaultLocation = "http://128.199.147.139:9800"
	EnvFastvaultToken = "FV_TOKEN"
)

var (
	Config Configuration
	Bot    *linebot.Client
)

func init() {
	err := getConfiguration()
	if err != nil {
		log.Panic(err)
	}

	Bot = newLinebot()
}

func getConfiguration() error {
	token := os.Getenv(EnvFastvaultToken)
	if token == ""{
		return errors.New("Could not read fastvault token from env variable")
	}

	fv := fastvault_client_go.New(FastvaultLocation)
	err := fv.GetJson(token, &Config)

	return err
}

func newLinebot() *linebot.Client {
	client, err := linebot.New(Config.LineBot.Secret, Config.LineBot.Token)
	if err != nil {
		log.Panic(err)
	}

	return client
}