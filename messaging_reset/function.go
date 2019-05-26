package messaging_reset

import (
	"context"
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/findata-code/fastvault-client-go"
	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
	"strings"
	"time"
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

func ResetMessage(ctx context.Context, psm PubSubMessage) error {
	m, err := getMessage(psm)
	if err != nil {
		return err
	}

	if !isResetMessage(m.Message) {
		return nil
	}

	r := Reset{
		UserId:m.UserId,
		UnixNano:m.Timestamp,
		Timestamp:time.Now(),
		FullMessage:m.Message,
	}

	err = db.Create(&r).Error
	if err != nil {
		return err
	}

	_, err = bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("รับทราบ มาเริ่มต้นใหม่กันเลย!!")).Do()
	if err != nil {
		return err
	}

	return nil
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

func isResetMessage(s string) bool {
	resetKeyWord := []string{
		"เริ่มใหม่",
		"reset",
		"รีเซ็ต",
	}

	for _, w := range resetKeyWord {
		if strings.Contains(s, w) {
			return true
		}
	}

	return false
}

func getMessage(psm PubSubMessage) (Message, error) {
	var message Message
	err := json.Unmarshal(psm.Data, &message)
	if err != nil {
		return Message{}, err
	}

	return message, nil
}

func newLineBot() *linebot.Client {
	client, err := linebot.New(config.LineBot.Secret, config.LineBot.Token)
	if err != nil {
		log.Panic(err)
	}

	return client
}
