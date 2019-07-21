package main

import (
	"errors"
	"flag"
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
		area        = fs.String("area", "", "")
		image       = fs.String("image", "", "")
	)

	fs.Parse(os.Args[1:])

	if err := checkRequiredProgramArgument(width, height, name, chatBarText, area, image); err != nil {
		panic(err)
	}
}

func checkRequiredProgramArgument(width *int, height *int, name *string, chatBarText *string, area *string, image *string) error {
	if *width == -1 ||
		*height == -1 ||
		*name == "" ||
		*chatBarText == "" ||
		*area == "" ||
		*image == "" {
		return errors.New("required field are missing")
	}

	return nil
}
