package app

import (
	_ "database/sql"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/findata-code/fastvault-client-go"
	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"messaging_income/app/model"
	"os"
)

const (
	FastvaultLocation = "http://128.199.147.139:9800"
	EnvFastvaultToken = "FV_TOKEN"
)

var (
	Config Configuration
	Db     *gorm.DB
	Bot    *linebot.Client
)

func init() {
	err := getConfiguration()

	url := fmt.Sprintf("%s:%s@cloudsql(%s)/%s?charset=utf8&parseTime=True&loc=Local", Config.DB.User, Config.DB.Password, Config.DB.Location, Config.DB.Database)
	Db, err = gorm.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	Db.AutoMigrate(&model.Income{})

	Bot = newLineBot()
}


func getConfiguration() error {
	token := os.Getenv(EnvFastvaultToken)
	if token == "" {
		log.Fatal("Could not read fastvault token from env variable")
	}
	fv := fastvault_client_go.New(FastvaultLocation)
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

