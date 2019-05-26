package app

import (
	_ "database/sql"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
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
	config Configuration
	bot    *linebot.Client
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func init() {
	err := getConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	bot = newLineBot()
}

func getConfiguration() error {
	token := os.Getenv(ENV_FASTVAULT_TOKEN)
	if token == "" {
		log.Fatal("Could not read fastvault token from env variable")
	}
	fv := fastvault_client_go.New(FASTVAULT_LOCATION)
	err := fv.GetJson(token, &config)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func newLineBot() *linebot.Client {
	client, err := linebot.New(config.LineBot.Secret, config.LineBot.Token)
	if err != nil {
		log.Panic(err)
	}

	return client
}
