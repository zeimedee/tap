package services

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func SendPushNotifs(token string) error {
	ctx := context.Background()

	opt := option.WithCredentialsFile("internal/services/awesomenotif.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Printf("firebase-init: %v", err)
		return err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		fmt.Printf("msg-init: %v", err)
		return err
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "😘",
			Body:  "I thought of you today",
		},

		Token: token,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		fmt.Printf("client-send: %v", err)
		return err
	}

	log.Printf("Successfully sent message: %v", response)
	return nil
}
