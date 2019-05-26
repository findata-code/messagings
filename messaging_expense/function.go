package messaging_expense

import (
	"context"
	_ "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/findata-code/fastvault-client-go"
	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
	"regexp"
	"strconv"
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

	db.AutoMigrate(&Expense{})

	bot = newLineBot()
}

/*
	Entry point
*/
func ExpenseMessage(ctx context.Context, psm PubSubMessage) error {
	m, err := getMessage(psm)
	if err != nil {
		return err
	}

	if !isExpensePattern(m.Message) {
		return nil
	}

	value, err := extractValue(m.Message)
	if err != nil {
		return err
	}

	i := Expense{
		UserId:      m.UserId,
		Value:       value,
		FullMessage: m.Message,
		UnixNano:    m.Timestamp,
		Timestamp:   time.Now(),
	}

	err = db.Create(&i).Error
	if err != nil {
		return err
	}

	_, err = bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("รับทราบจ้า บันทึกกก")).Do()
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

func newLineBot() *linebot.Client {
	client, err := linebot.New(config.LineBot.Secret, config.LineBot.Token)
	if err != nil {
		log.Panic(err)
	}

	return client
}

func extractValue(s string) (float64, error) {
	var v float64
	re := regexp.MustCompile("([-][ ]?[0-9]*[kKmM]?)")
	gs := re.FindAllStringSubmatch(s, -1)
	if len(gs) != 1 {
		return v, errors.New(fmt.Sprintf("found %d groups in message %s", len(gs), s))
	}

	m := gs[0][0]
	m = strings.Replace(m, "+", "", -1)
	m = strings.Replace(m, "k", "000", -1)
	m = strings.Replace(m, "K", "000", -1)
	m = strings.Replace(m, "m", "000000", -1)
	m = strings.Replace(m, "M", "000000", -1)

	return strconv.ParseFloat(m, 64)
}

func isExpensePattern(s string) bool {
	re := regexp.MustCompile("([-][ ]?[0-9]*[kKmM]?)")
	return re.Match([]byte(s))
}

func getMessage(psm PubSubMessage) (Message, error) {
	var message Message
	err := json.Unmarshal(psm.Data, &message)
	if err != nil {
		return Message{}, err
	}

	return message, nil
}
