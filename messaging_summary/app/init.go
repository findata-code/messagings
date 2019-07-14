package app

import (
	_ "database/sql"
	"encoding/base64"
	"encoding/json"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
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
		log.Fatal(err)
	}

	Bot = newLineBot()
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

func newLineBot() *linebot.Client {
	client, err := linebot.New(Config.LineBot.Secret, Config.LineBot.Token)
	if err != nil {
		log.Panic(err)
	}

	return client
}
