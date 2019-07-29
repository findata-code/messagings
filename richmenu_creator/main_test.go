package main_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	. "richmenu_creator"
	"richmenu_creator/model"
	"testing"

	"github.com/line/line-bot-sdk-go/linebot"
)

const AREA = `
[
  {
	"bounds": {
	  "x": 0,
	  "y": 0,
	  "width": 833,
	  "height": 843
	},
	"action": {
	  "type": "message",
	  "text": "settings"
	}
  },
  {
	"bounds": {
	  "x": 834,
	  "y": 0,
	  "width": 833,
	  "height": 843
	},
	"action": {
	  "type": "message",
	  "text": "summary"
	}
  },
  {
	"bounds": {
	  "x": 1667,
	  "y": 0,
	  "width": 833,
	  "height": 843
	},
	"action": {
	  "type": "message",
	  "text": "income"
	}
  },
  {
	"bounds": {
	  "x": 0,
	  "y": 844,
	  "width": 833,
	  "height": 843
	},
	"action": {
	  "type": "message",
	  "text": "help"
	}
  },
  {
	"bounds": {
	  "x": 1667,
	  "y": 844,
	  "width": 833,
	  "height": 843
	},
	"action": {
	  "type": "message",
	  "text": "expense"
	}
  }
]
`

func TestCreateRichMenuShouldReturnCorrectRichMenu(t *testing.T) {
	var areas []linebot.AreaDetail
	w := 1
	h := 2
	s := true
	n := "name"
	cbt := "chatBarText"

	err := json.Unmarshal([]byte(AREA), &areas)
	if err != nil {
		t.Error("TestCreateRichMenu should not fail parsing AREA")
	}

	richMenu := CreateRichMenu(w, h, s, n, cbt, areas)

	if richMenu.Size.Width != w {
		t.Errorf("expect %d, actual %d", w, richMenu.Size.Width)
	}

	if richMenu.Size.Height != h {
		t.Errorf("expect %d, actual %d", h, richMenu.Size.Height)
	}

	if richMenu.Selected != s {
		t.Errorf("expect %t, actual %t", richMenu.Selected, s)
	}

	if richMenu.Name != n {
		t.Errorf("expect %s, actual %s", n, richMenu.Name)
	}

	if richMenu.ChatBarText != cbt {
		t.Errorf("expect %s, actual %s", cbt, richMenu.ChatBarText)
	}

	if !reflect.DeepEqual(richMenu.Areas, areas) {
		t.Errorf("expect richMenu.Area equals areas, but not")
	}
}

func TestGetAreaShouldReturnErrorIfFileDoesNotExists(t *testing.T) {
	filename := "testdata/notFound.json"
	errMessage := fmt.Sprintf("open %s: no such file or directory", filename)

	_, err := GetArea(filename)

	if err == nil {
		t.Fatal("Expect to have error occurred")
	}
	if err.Error() != errMessage {
		t.Error("Expect", errMessage, "actual", err.Error())
	}
}

func TestGetAreaShouldReturnCorrectValueOfArrayOfAreaDetail(t *testing.T) {
	var expectedAreaDetail []linebot.AreaDetail
	filename := "testdata/area.json"
	json.Unmarshal([]byte(AREA), &expectedAreaDetail)

	area, _ := GetArea(filename)

	if !reflect.DeepEqual(expectedAreaDetail, area) {
		t.Error("expect equal to area deeply but false")
	}
}

func TestExecShouldCallLineBotProperlyAndCreateRichMenu(t *testing.T) {
	var expectedAreas []linebot.AreaDetail
	if err := json.Unmarshal([]byte(AREA), &expectedAreas); err != nil {
		t.Errorf("expect can parse area, actual failed to parse area")
	}


	config := model.Config{
		Width:       1,
		Height:      2,
		ChatBarText: "ChatBarText",
		Name:        "Name",
		Selected:    true,
		ImageFile:   "image.png",
		AreaFile:    "testdata/area.json",
	}

	expectedRichMenu := linebot.RichMenu{
		Size: linebot.RichMenuSize{
			Width:  config.Width,
			Height: config.Height,
		},
		Selected:    config.Selected,
		Name:        config.Name,
		ChatBarText: config.ChatBarText,
		Areas:       expectedAreas,
	}

	botMock := NewBotWrapperMock()

	expectedRichMenuResponseId := "1234"

	botMock.SetupCreateRichMenu = func() (response *linebot.RichMenuIDResponse, e error) {
		return &linebot.RichMenuIDResponse{RichMenuID: expectedRichMenuResponseId}, nil
	}

	Exec(botMock, config)

	if !reflect.DeepEqual(expectedRichMenu, botMock.Calls["CreateRichMenu"][0][0]) {
		t.Errorf("expect passed %v but actual passed %v", expectedRichMenu, botMock.Calls["CreateRichMenu"][0][0])
	}

	if expectedRichMenuResponseId != botMock.Calls["UploadRichMenuImage"][0][0]{
		t.Errorf("expect %s, actual %s", expectedRichMenuResponseId, botMock.Calls["UploadRichMenuImage"][0][0])
	}

	if config.ImageFile != botMock.Calls["UploadRichMenuImage"][0][1]{
		t.Errorf("expect %s, actual %s", expectedRichMenuResponseId, botMock.Calls["UploadRichMenuImage"][0][0])
	}

	if expectedRichMenuResponseId != botMock.Calls["SetDefaultRichMenu"][0][0]{
		t.Errorf("expect %s, actual %s", expectedRichMenuResponseId, botMock.Calls["SetDefaultRichMenu"][0][0])
	}
}
