package app

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"time"
)

func getTextMessage(event *linebot.Event) (string, error) {
	if event.Type == linebot.EventTypeMessage {
		switch msg := event.Message.(type) {
		case *linebot.TextMessage:
			return msg.Text, nil
		}
	}

	return "", errors.New("event type is not message type")
}

func publishToPubSub(msg Message) error {
	pubSubMessage, err := toPubSubMessage(msg)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < PUBSUB_RETRY; i++ {
		if _, processing := <-pubSubClient.Topic(config.PubSub.Topic).Publish(ctx, pubSubMessage).Ready(); processing {
			time.Sleep(1 * time.Second)
		} else {
			return nil
		}
	}

	cancel()

	return errors.New(fmt.Sprintf("Could not process message %s from user %s", msg.Message, msg.UserId))
}

func toPubSubMessage(msg Message) (*pubsub.Message, error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return &pubsub.Message{
		Data: b,
	}, nil
}
