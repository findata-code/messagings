package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"io/ioutil"
	"log"
	"os"
)

const (
	SECRET = "SECRET"
	TOKEN  = "TOKEN"
)

func main() {
	fmt.Println(SECRET, os.Getenv(SECRET))
	fmt.Println(TOKEN, os.Getenv(TOKEN))

	var (
		width       = flag.Int("width", 2500, "")
		height      = flag.Int("height", 1686, "")
		selected    = flag.Bool("selected", false, "")
		name        = flag.String("name", "", "")
		chatBarText = flag.String("chatBarText", "", "")
		filename    = flag.String("area", "", "")
		image       = flag.String("image", "", "")
	)

	flag.Parse()

	fmt.Println(*width, *height, *selected, *name, *chatBarText, *filename, *image)

	if 	width == nil ||
		height == nil ||
		selected == nil ||
		name == nil ||
		chatBarText == nil ||
		filename == nil ||
		image == nil {

		panic("required field are missing")
	}

	areas, err := GetAreaDetail(*filename)
	if err != nil {
		log.Println(*filename, "not found")
		panic(err)
	}

	bot, err := linebot.New(os.Getenv(SECRET), os.Getenv(TOKEN))
	if err != nil {
		panic(err)
	}

	richMenu := linebot.RichMenu{
		Size:        linebot.RichMenuSize{Width: *width, Height: *height},
		Selected:    *selected,
		Name:        *name,
		ChatBarText: *chatBarText,
		Areas:       areas,
	}

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

func GetAreaDetail(filename string) (result []linebot.AreaDetail, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &result)

	return result, err
}
