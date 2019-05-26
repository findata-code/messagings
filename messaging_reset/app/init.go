package app

import (
	"fmt"
	"github.com/findata-code/fastvault-client-go"
	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	_ "database/sql"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"os"
)

const (
	FASTVAULT_LOCATION  = "http://128.199.147.139:9800"
	ENV_FASTVAULT_TOKEN = "FV_TOKEN"
)

var (
	config Configuration
	db     *gorm.DB
	bot    *linebot.Client
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func init() {
	err := getConfiguration()

	url := fmt.Sprintf("%s:%s@cloudsql(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DB.User, config.DB.Password, config.DB.Location, config.DB.Database)
	db, err = gorm.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Reset{})

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
