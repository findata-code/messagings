package message_handler

import (
	"context"
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

func DeleteIncomeMessage(ctx context.Context, psm model.PubSubMessage) error {
	m, err := getMessage(psm)
	if err != nil {
		return err
	}

	if !isDeleteIncomeMessagePattern(m.Message) {
		return nil
	}

	id, err := utils.GetEntityIdentifier(m.Message)
	if err != nil {
		return err
	}

	if id == utils.LatestExpense {
		incomes, err := api.GetLastNIncome(m.UserId, 1)
		if err != nil {
			return err
		}

		if len(incomes) != 1 {
			app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("หารายการล่าสุดไม่เจอจ้า")).Do()
			return errors.New(fmt.Sprintf("cannot find last income for user %s", m.UserId))
		}

		id = int64(incomes[0].ID)
	}

	log.Println("deleting income id", id)

	err = api.DeleteIncome(m.UserId, int(id))
	if err != nil {
		app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("น้องกำลังงง โปรดลองใหม่อีกครั้งจ้า")).Do()
		return err
	}

	_, err = app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage("ลบสำเร็จ")).Do()

	return nil
}

func isDeleteIncomeMessagePattern(s string) bool {
	return strings.Contains(s, "ลบ") && strings.Contains(s, "รายรับ")
}
