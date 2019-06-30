package message_handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"messaging_delete/app"
	"messaging_delete/app/api"
	"messaging_delete/app/model"
	"messaging_delete/app/utils"
	"strings"
)

func DeleteExpenseMessage(ctx context.Context, psm model.PubSubMessage) error {
	m, err := getMessage(psm)
	if err != nil {
		return err
	}

	if !isDeleteExpenseMessagePattern(m.Message) {
		return nil
	}

	id, err := utils.GetEntityIdentifier(m.Message)
	if err != nil {
		return err
	}

	if id == utils.LatestExpense {
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

	log.Println("deleting expense id", id)

	err = api.DeleteExpense(m.UserId, int(id))
	if err != nil {
		app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("น้องกำลังงง โปรดลองใหม่อีกครั้งจ้า")).Do()
		return err
	}

	_, err = app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("ลบสำเร็จจ้า")).Do()

	return err
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
