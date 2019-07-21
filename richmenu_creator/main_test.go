package main_test

import (
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"os"
	"reflect"
	"richmenu_creator"
	"testing"
)

func TestGetAreaShouldReturnCorrectValueOfArrayOfAreaDetail(t *testing.T) {
	filename := "testdata/AREA.json"
	var expectedAreaDetail []linebot.AreaDetail
	json.Unmarshal([]byte(AREA), &expectedAreaDetail)

	area, _ := main.GetArea(&filename)

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

	main.Exec()
}

func TestExecShouldNotPanicWithRequiredFieldIsAreMissingMessageIfAllArgumentIsAreProvidedCorrectly(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("expect", "no panic", "actual", r)
		}
	}()

	os.Args = []string{
		"",
		"-width=2500",
		"-height=1686",
		"-selected=true",
		"-name=Home",
		"-chatBarText=Home",
		"-areaFile=$(pwd)/areas/Home.json",
		"-image=$(pwd)/images/Home.png",
	}

	main.Exec()
}

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