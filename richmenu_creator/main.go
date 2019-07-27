package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"richmenu_creator/model"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	Exec()
}

func Exec() {
	config := model.Config{}
	if err := config.Read(os.Args); err != nil {
		panic(err)
	}
}

func GetArea(areaFile *string) ([]linebot.AreaDetail, error) {
	b, err := ioutil.ReadFile(*areaFile)
	if err != nil {
		return nil, err
	}
	var areaDetail []linebot.AreaDetail
	err = json.Unmarshal(b, &areaDetail)
	return areaDetail, err
}
