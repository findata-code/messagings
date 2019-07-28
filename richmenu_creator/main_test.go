package main_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	. "richmenu_creator"
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

func TestCreateRichMenu(t *testing.T) {
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

func TestExecShouldPanicStopIfRequiredProgramArgumentIsAreMissing(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r.(error).Error() != "required field are missing" {
				t.Error("expect", "required field are missing", "actual", r)
			}
		}
	}()

	os.Args = []string{
		"",
		"-width=2500",
		"-height=1686",
		"-selected=true",
	}

	Exec(nil)
}

