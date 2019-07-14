package app

import (
	_ "database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"messaging_expense/app/model"
	"os"
)

const (
	CONFIG = "CONFIG"
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

	Db.AutoMigrate(&model.Expense{})

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
