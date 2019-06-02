package message_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"messaging_summary/app"
	"messaging_summary/app/api"
	"messaging_summary/app/model"
	"strings"
)

func GetSummaryMessage(ctx context.Context, psm model.PubSubMessage) error {
	m, err := getMessage(psm)
	if err != nil {
		ctx.Err()
		return err
	}

	if !isSummaryMessage(m.Message) {
		ctx.Done()
		return nil
	}

	latestReset, err := api.GetLatestReset(m.UserId)
	if err != nil {
		ctx.Err()
		return err
	}

	incomes, err := api.GetIncomes(m.UserId, latestReset.UnixNano)
	if err != nil {
		ctx.Err()
		return err
	}

	expenses, err := api.GetExpenses(m.UserId, latestReset.UnixNano)
	if err != nil {
		ctx.Err()
		return err
	}

	sum := 0.0
	for _, in := range incomes {
		sum += in.Value
	}

	for _, ex := range expenses {
		sum += ex.Value
	}

	_, err = app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage(CreateMessage(incomes, expenses, sum))).Do()
	if err != nil {
		return err
	}

	return nil
}

func CreateMessage(ins []model.Income, exs []model.Expense, sum float64) string {
	totalIncome := 0.0
	totalExpense := 0.0

	for _, in := range ins {
		totalIncome += in.Value
	}

	for _, ex := range exs {
		totalExpense += ex.Value
	}

	return fmt.Sprintf(`รวมรายรับ %.2f
รวมรายจ่าย %.2f
คงเหลือ %.2f`, totalIncome, totalExpense, sum)
}

func isSummaryMessage(s string) bool {
	resetKeyWord := []string{
		"สรุป",
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
