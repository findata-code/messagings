package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"io/ioutil"
	"os"
)

const (
	SECRET = "SECRET"
	TOKEN  = "TOKEN"
)

func main() {
	Exec()
}

func Exec() {
	fmt.Println(SECRET, os.Getenv(SECRET))
	fmt.Println(TOKEN, os.Getenv(TOKEN))

	var (
		width       = flag.Int("width", -1, "")
		height      = flag.Int("height", -1, "")
		selected    = flag.Bool("selected", false, "")
		name        = flag.String("name", "", "")
		chatBarText = flag.String("chatBarText", "", "")
		area        = flag.String("area", "", "")
		image       = flag.String("image", "", "")
	)

	flag.Parse()

	if err := checkRequiredProgramArgument(width, height, name, chatBarText, area, image); err != nil {
		panic(err)
	}

	areas, err := getAreaDetail(*area)
	if err != nil {
		panic(err)
	}

	bot, err := linebot.New(os.Getenv(SECRET), os.Getenv(TOKEN))
	if err != nil {
		panic(err)
	}

	richMenu := createRichMenu(width, height, selected, name, chatBarText, areas)

	res, err := bot.CreateRichMenu(richMenu).Do()
	if err != nil {
		panic(err)
	}

	_, err = bot.UploadRichMenuImage(res.RichMenuID, *image).Do()
	if err != nil {
		panic(err)
	}

	_, err = bot.SetDefaultRichMenu(res.RichMenuID).Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(res.RichMenuID)
}

func checkRequiredProgramArgument(width *int, height *int, name *string, chatBarText *string, area *string, image *string) error{
	if *width == -1 ||
		*height == -1 ||
		*name == "" ||
		*chatBarText == "" ||
		*area == "" ||
		*image == "" {
		fmt.Println(*width, *height, *name, *chatBarText, *area, *image)
		return errors.New("required field are missing")
	}

	return nil
}

func createRichMenu(width *int, height *int, selected *bool, name *string, chatBarText *string, areas []linebot.AreaDetail) linebot.RichMenu {
	richMenu := linebot.RichMenu{
		Size:        linebot.RichMenuSize{Width: *width, Height: *height},
		Selected:    *selected,
		Name:        *name,
		ChatBarText: *chatBarText,
		Areas:       areas,
	}
	return richMenu
}


func getAreaDetail(filename string) (result []linebot.AreaDetail, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &result)

	return result, err
}
