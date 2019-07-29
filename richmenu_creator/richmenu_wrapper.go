package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

type BotWrapper interface {
	CreateRichMenu(richMenu linebot.RichMenu) (*linebot.RichMenuIDResponse, error)
	UploadRichMenuImage(richMenuId, imgPath string) (*linebot.BasicResponse, error)
	SetDefaultRichMenu(richMenuId string) (*linebot.BasicResponse, error)
}

type BotWrapperImpl struct {
	bot *linebot.Client
}

func NewBotWrapper(bot *linebot.Client) BotWrapper {
	return BotWrapperImpl{
		bot: bot,
	}
}

func (rc BotWrapperImpl) CreateRichMenu(richMenu linebot.RichMenu) (*linebot.RichMenuIDResponse, error) {
	return rc.bot.CreateRichMenu(richMenu).Do()
}

func (rc BotWrapperImpl) UploadRichMenuImage(richMenuId, imgPath string) (*linebot.BasicResponse, error) {
	return rc.bot.UploadRichMenuImage(richMenuId, imgPath).Do()
}

func (rc BotWrapperImpl) SetDefaultRichMenu(richMenuId string) (*linebot.BasicResponse, error) {
	return rc.bot.SetDefaultRichMenu(richMenuId).Do()
}
