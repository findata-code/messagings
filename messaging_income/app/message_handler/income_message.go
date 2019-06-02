package message_handler

import (
	"context"
	_ "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/line/line-bot-sdk-go/linebot"
	"messaging_income/app"
	"messaging_income/app/model"
	"regexp"
	"strconv"
	"strings"
	"time"
)



/*
	Entry point
*/
func IncomeMessage(ctx context.Context, psm model.PubSubMessage) error {
	m, err := getMessage(psm)
	if err != nil {
		return err
	}

	if !isIncomePattern(m.Message) {
		return nil
	}

	value, err := extractValue(m.Message)
	if err != nil {
		return err
	}

	i := model.Income{
		UserId:      m.UserId,
		Value:       value,
		FullMessage: m.Message,
		UnixNano:    m.Timestamp,
		Timestamp:   time.Now(),
	}

	err = app.Db.Create(&i).Error
	if err != nil {
		return err
	}

	_, err = app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("รับทราบจ้า บันทึกกก")).Do()
	if err != nil {
		return err
	}

	return nil
}


func extractValue(s string) (float64, error) {
	var v float64
	re := regexp.MustCompile("([+][ ]?[0-9]*[kKmM]?)")
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

func isIncomePattern(s string) bool {
	re := regexp.MustCompile("([+][ ]?[0-9]*[kKmM]?)")
	return re.Match([]byte(s))
}

func getMessage(psm model.PubSubMessage) (model.Message, error) {
	var message model.Message
	err := json.Unmarshal(psm.Data, &message)
	if err != nil {
		return model.Message{}, err
	}

	return message, nil
}