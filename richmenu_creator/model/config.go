package model

import (
	"errors"
	"flag"
)

type Config struct {
	Width       int
	Height      int
	Selected    bool
	Name        string
	ChatBarText string
	AreaFile    string
	ImageFile   string
}

func (config *Config) Read(args []string) error {
	fs := flag.NewFlagSet("Richmenu Configuration", flag.ContinueOnError)
	fs.IntVar(&config.Width, "width", -1, "")
	fs.IntVar(&config.Height, "height", -1, "")
	fs.BoolVar(&config.Selected, "selected", false, "")
	fs.StringVar(&config.Name, "name", "", "")
	fs.StringVar(&config.ChatBarText, "chatBarText", "", "")
	fs.StringVar(&config.AreaFile, "areaFile", "", "")
	fs.StringVar(&config.ImageFile, "imageFile", "", "")
	err := fs.Parse(args[1:])
	if err != nil {
		return err
	}

	return checkRequiredProgramArgument(*config)
}

func checkRequiredProgramArgument(config Config) error {
	if config.Width == -1 ||
		config.Height == -1 ||
		config.Name == "" ||
		config.ChatBarText == "" ||
		config.AreaFile == "" ||
		config.ImageFile == "" {
		return errors.New("required field are missing")
	}
	return nil
}
