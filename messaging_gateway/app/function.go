package app

import (
	"cloud.google.com/go/pubsub"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	PUBSUB_RETRY = 10
	CONFIG = "CONFIG"
)

var (
	Config        Configuration
	pubSubClient  *pubsub.Client
	lineBotClient *linebot.Client
)

func init() {
	getConfiguration()
	pubSubClient = newPubSub()
	lineBotClient = newLineBot()
}

func LineToPubsub(w http.ResponseWriter, r *http.Request) {
	events, err := lineBotClient.ParseRequest(r)
	if err != nil {
		log.Println(err)
		response5xx(w)
		return
	}

	go sendEventToPubSub(events)

	responseOK(w)
}

func getConfiguration() error {
	encodedConfigurationValue := os.Getenv(CONFIG)
	b, err := base64.StdEncoding.DecodeString(encodedConfigurationValue)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &Config)
	return err
}

func sendEventToPubSub(events []*linebot.Event) {
	for _, event := range events {
		lineMsg, err := getTextMessage(event)
		if err != nil {
			log.Println(err)
			continue
		}

		msg := Message{
			Message:    lineMsg,
			ReplyToken: event.ReplyToken,
			UserId:     event.Source.UserID,
			Timestamp:  fmt.Sprintf("%d", time.Now().UnixNano()),
		}

		err = publishToPubSub(msg)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func responseOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "")
}

func response5xx(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "")
}
