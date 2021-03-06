package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"richmenu_creator/model"

	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	SECRET = "SECRET"
	TOKEN  = "TOKEN"
)

func main() {
	bot, err := linebot.New(os.Getenv(SECRET), os.Getenv(TOKEN))
	if err != nil {
		panic(err)
	}

	config := model.Config{}
	if err := config.Read(os.Args); err != nil {
		panic(err)
	}

	wrapper := NewBotWrapper(bot)

	Exec(wrapper, config)
}

func Exec(bot BotWrapper, config model.Config) {
	area, err := GetArea(config.AreaFile)
	if err != nil {
		panic(err)
	}

	richMenu := CreateRichMenu(
		config.Width,
		config.Height,
		config.Selected,
		config.Name,
		config.ChatBarText,
		area)

	res, err := bot.CreateRichMenu(richMenu)
	if err != nil {
		panic(err)
	}
	fmt.Println("Rich menu created name:", config.Name)
	fmt.Println("\tId:", res.RichMenuID)

	_, err = bot.UploadRichMenuImage(res.RichMenuID, config.ImageFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tImage uploaded")


	if config.Selected {
		_, err = bot.SetDefaultRichMenu(res.RichMenuID)
		if err != nil {
			panic(err)
		}
		fmt.Println("\tSet as a default menu")
	}

	fmt.Println("\tDone")
}

func CreateRichMenu(
	width int,
	height int,
	selected bool,
	name string,
	chatBarText string,
	areas []linebot.AreaDetail) linebot.RichMenu {

	richMenu := linebot.RichMenu{
		Size: linebot.RichMenuSize{
			Width:  width,
			Height: height},
		Selected:    selected,
		Name:        name,
		ChatBarText: chatBarText,
		Areas:       areas,
	}

	return richMenu
}

func GetArea(areaFile string) ([]linebot.AreaDetail, error) {
	b, err := ioutil.ReadFile(areaFile)
	if err != nil {
		return nil, err
	}
	var areaDetail []linebot.AreaDetail
	err = json.Unmarshal(b, &areaDetail)
	return areaDetail, err
}
