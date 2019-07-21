package main

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/line/line-bot-sdk-go/linebot"
	"io/ioutil"
	"os"
)

func main() {
	Exec()
}

func Exec() {
	fs := flag.NewFlagSet("Rich Menu Uploader", flag.ContinueOnError)
	var (
		width       = fs.Int("width", -1, "")
		height      = fs.Int("height", -1, "")
		_           = fs.Bool("selected", false, "")
		name        = fs.String("name", "", "")
		chatBarText = fs.String("chatBarText", "", "")
		areaFile    = fs.String("areaFile", "", "")
		image       = fs.String("image", "", "")
	)

	fs.Parse(os.Args[1:])

	if err := checkRequiredProgramArgument(width, height, name, chatBarText, areaFile, image); err != nil {
		panic(err)
	}
}

func GetArea(areaFile *string) ([]linebot.AreaDetail, error) {
	b, err := ioutil.ReadFile(*areaFile)
	if err != nil {
		return nil, err
	}

	var ad []linebot.AreaDetail

	err = json.Unmarshal(b, &ad)

	return ad, err
}

func checkRequiredProgramArgument(width *int, height *int, name *string, chatBarText *string, areaFile *string, image *string) error {
	if *width == -1 ||
		*height == -1 ||
		*name == "" ||
		*chatBarText == "" ||
		*areaFile == "" ||
		*image == "" {
		return errors.New("required field are missing")
	}

	return nil
}
