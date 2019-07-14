package message_handler

import (
	"context"
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"messaging_help/app"
	"messaging_help/app/model"
	"strings"
)

const message = `คำสั่งทั้งหมด
- เริ่มต้นรอบใหม่
 รีเซ็ต
 reset
- บันทึกรายรับ
 เงินเดือน +50000
- บันทึกรายจ่าย
 กินข้าว -40
- ลบรายรับที่ผิดพลาดล่าสุด
 ลบรายรับ ล่าสุด
- ลบรายจ่ายที่ผิดพลาดล่าสุด
 ลบรายจ่าย ล่าสุด`


func HelpMessage(ctx context.Context, psm model.PubSubMessage) error {
	m, err := getMessage(psm)
	if err != nil {
		return err
	}

	if !isHelpMessage(m.Message) {
		return nil
	}

	_, err = app.Bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage(message)).Do()
	if err != nil {
		return err
	}

	return nil
}

func isHelpMessage(s string) bool {
	result := false
	for _, hm := range getAllHelpRequestMessage() {
		if strings.Contains(s, hm) {
			result = true
			break
		}
	}

	return result
}

func getAllHelpRequestMessage() []string {
	return []string {
		"help",
		"คำสั่ง",
		"ช่วยเหลือ",
	}
}

func getMessage(psm model.PubSubMessage) (model.Message, error) {
	var message model.Message
	err := json.Unmarshal(psm.Data, &message)
	if err != nil {
		return model.Message{}, err
	}

	return message, nil
}
