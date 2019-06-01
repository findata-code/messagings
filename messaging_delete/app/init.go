package app

import "github.com/line/line-bot-sdk-go/linebot"

const (
	FastvaultLocation = "http://128.199.147.139:9800"
	EnvFastvaultToken = "FV_TOKEN"
)

var (
	config Configuration
	bot    *linebot.Client
)
