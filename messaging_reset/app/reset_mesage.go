package app

import (
	"context"
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"strings"
	"time"
)

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