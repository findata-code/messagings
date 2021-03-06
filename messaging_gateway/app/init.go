package app

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/base64"
	"github.com/line/line-bot-sdk-go/linebot"
	"google.golang.org/api/option"
	"log"
)

func newLineBot() *linebot.Client {
	client, err := linebot.New(Config.LineBot.Secret, Config.LineBot.Token)
	if err != nil {
		log.Panic(err)
	}

	return client
}

func newPubSub() *pubsub.Client {
	ctx := context.Background()
	cred, err := base64.StdEncoding.DecodeString(Config.PubSub.Credential)
	if err != nil {
		log.Fatal(err)
	}

	client, err := pubsub.NewClient(
		ctx,
		Config.PubSub.ProjectId,
		option.WithCredentialsJSON(
			cred))

	if err != nil {
		log.Panic(err)
	}

	return client
}
