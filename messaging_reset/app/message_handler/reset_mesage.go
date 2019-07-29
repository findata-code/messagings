package message_handler

import (
	"context"
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"messaging_reset/app"
	"messaging_reset/app/model"
	"strings"
	"time"
)

func ResetMessage(ctx context.Context, psm model.PubSubMessage) error {
	m, err := getMessage(psm)
	if err != nil {
		return err
	}

	if !isResetMessage(m.Message) {
		return nil
	}

	r := model.Reset{
		UserId:      m.UserId,
		UnixNano:    m.Timestamp,
		Timestamp:   time.Now(),
		FullMessage: m.Message,
	}

	err = app.Db.Create(&r).Error
	if err != nil {
		return err
	}

	_, err = app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("รับทราบ มาเริ่มต้นใหม่กันเลย!!")).Do()
	if err != nil {
		return err
	}

	return nil
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

func getMessage(psm model.PubSubMessage) (model.Message, error) {
	var message model.Message
	err := json.Unmarshal(psm.Data, &message)
	if err != nil {
		return model.Message{}, err
	}

	return message, nil
}
