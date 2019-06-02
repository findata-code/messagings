package http_handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"messaging_delete/app"
	"messaging_delete/app/api"
	"messaging_delete/app/model"
	"strconv"
	"strings"
)

const (
	LatestExpense = -1
)

func DeleteExpenseMessage(ctx context.Context, psm model.PubSubMessage) error {
	m, err := getMessage(psm)
	if err != nil {
		return err
	}

	if !isDeleteExpenseMessagePattern(m.Message) {
		return nil
	}

	id, err := getEntityIdentifier(m.Message)
	if err != nil {
		return err
	}

	if id == LatestExpense {
		expenses, err := api.GetLastNExpense(m.UserId, 1)
		if err != nil {
			return err
		}

		if len(expenses) != 1 {
			app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("หารายการล่าสุดไม่เจอจ้า")).Do()
			return errors.New(fmt.Sprintf("cannot find last expense for user %s", m.UserId))
		}
		id = int64(expenses[0].ID)
	}

	err = api.DeleteExpense(m.UserId, int(id))
	if err != nil {
		app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("น้องกำลังงง โปรลองใหม่อีกครั้งจ้า")).Do()
		return err
	}

	_, err = app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("ลบสำเร็จจ้า")).Do()

	return err
}

func getEntityIdentifier(s string) (int64, error) {
	if strings.HasSuffix(s, "ล่าสุด") {
		return -1, nil
	}

	ss := strings.Split(s, " ")
	return strconv.ParseInt(ss[len(ss) - 1], 10, 64)
}

func isDeleteExpenseMessagePattern(s string) bool {
	return strings.Contains(s, "ลบ") && strings.Contains(s, "รายจ่าย")
}

func getMessage(psm model.PubSubMessage) (model.Message, error) {
	var message model.Message
	err := json.Unmarshal(psm.Data, &message)
	if err != nil {
		return model.Message{}, err
	}

	return message, nil
}
