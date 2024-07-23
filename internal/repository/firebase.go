package repository

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"

	"firebase.google.com/go/messaging"
)

type FirebaseMessaging interface {
	SendMessage(ctx context.Context, message messaging.Message) error
}

type FirebaseClient struct {
	pathToConfig string
}

func NewFirebaseClient(pathToConfig string) *FirebaseClient {
	return &FirebaseClient{pathToConfig: pathToConfig}
}

func (c *FirebaseClient) SendMessage(ctx context.Context, message messaging.Message) error {
	opt := option.WithCredentialsFile(c.pathToConfig)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	response, err := client.Send(ctx, &message)
	if err != nil {
		log.Fatalf("error sending message: %v\n", err)
	}
	log.Println("response from firebase send messaging: " + response)
	return nil
}
