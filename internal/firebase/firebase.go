package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type Client struct {
	app *firebase.App
}

func Init(keyPath string, projectID string) (*Client, error) {
	opt := option.WithCredentialsFile(keyPath)
	config := &firebase.Config{
		ProjectID: projectID,
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %w", err)
	}
	client := &Client{}
	client.app = app
	return client, nil
}

func (client *Client) SendToTopic(topic string, title string) error {
	ctx := context.Background()
	messageClient, err := client.app.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("error getting messaging client: %w", err)
	}
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
		},
		Topic: topic,
	}

	_, err = messageClient.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("error sending message to topic: %w", err)
	}
	return nil
}
