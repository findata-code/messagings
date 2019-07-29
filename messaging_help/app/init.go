package app

import (
	"encoding/base64"
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

const (
	CONFIG = "CONFIG"
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
	encodedConfigurationValue := os.Getenv(CONFIG)
	b, err := base64.StdEncoding.DecodeString(encodedConfigurationValue)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &Config)
	return err
}

func newLinebot() *linebot.Client {
	client, err := linebot.New(Config.LineBot.Secret, Config.LineBot.Token)
	if err != nil {
		log.Panic(err)
	}

	return client
}
