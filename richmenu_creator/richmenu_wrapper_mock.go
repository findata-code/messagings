package main

import "github.com/line/line-bot-sdk-go/linebot"

type BotWrapperMock struct {
	Calls map[string][][]interface{}
	SetupCreateRichMenu      func() (*linebot.RichMenuIDResponse, error)
	SetupUploadRichMenuImage func() (*linebot.BasicResponse, error)
	SetupSetDefaultRichMenu  func() (*linebot.BasicResponse, error)
}

func NewBotWrapperMock() *BotWrapperMock {
	return &BotWrapperMock{
		Calls:make(map[string][][]interface{}),
	}
}

func (rc *BotWrapperMock) call(method string, a ...interface{}) {
	args := make([]interface{}, 0)
	for _, v := range a {
		args = append(args, v)
	}

	calls := make([][]interface{}, 0)
	calls = append(calls, args)

	if rc.Calls[method] == nil {
		rc.Calls[method] = calls
	}
}

func (rc *BotWrapperMock) CreateRichMenu(richMenu linebot.RichMenu) (*linebot.RichMenuIDResponse, error) {
	mName := "CreateRichMenu"
	rc.call(mName, richMenu)

	if rc.SetupCreateRichMenu != nil {
		return rc.SetupCreateRichMenu()
	}

	return nil, nil
}

func (rc *BotWrapperMock) UploadRichMenuImage(richMenuId, imgPath string) (*linebot.BasicResponse, error) {
	rc.call("UploadRichMenuImage", richMenuId, imgPath)

	if rc.SetupUploadRichMenuImage != nil {
		return rc.SetupUploadRichMenuImage()
	}

	return nil, nil
}

func (rc *BotWrapperMock) SetDefaultRichMenu(richMenuId string) (*linebot.BasicResponse, error) {
	rc.call("SetDefaultRichMenu", richMenuId)

	if rc.SetupSetDefaultRichMenu != nil {
		return rc.SetupSetDefaultRichMenu()
	}

	return nil, nil
}
