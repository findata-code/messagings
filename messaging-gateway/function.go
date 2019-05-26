package messaging_gateway

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"github.com/findata-code/fastvault-client-go"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	PUBSUB_RETRY = 10
	FASTVAULT_LOCATION  = "http://128.199.147.139:9800"
	ENV_FASTVAULT_TOKEN = "FV_TOKEN"
)

var (
	config        Configuration
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
	token := os.Getenv(ENV_FASTVAULT_TOKEN)
	if token == "" {
		log.Fatal("Could not read fastvault token from env variable")
	}
	fv := fastvault_client_go.New(FASTVAULT_LOCATION)
	err := fv.GetJson(token, &config)
	if err != nil {
		log.Fatal(err)
	}
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
