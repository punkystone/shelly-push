package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type Message struct {
	Topic string
	Value string
}

type Client struct {
	client    mqtt.Client
	Messages  chan Message
	OnConnect func()
}

func NewClient(connectUrl string, clientID string) (*Client, error) {
	client := &Client{client: nil, Messages: make(chan Message)}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectUrl)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(client.publishHandler)
	opts.OnConnect = client.onConnect
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Info().Msgf("connection lost %v", err)
	}
	opts.AutoReconnect = true
	opts.ConnectRetry = true
	client.client = mqtt.NewClient(opts)
	return client, nil
}

func (c *Client) publishHandler(_ mqtt.Client, msg mqtt.Message) {
	c.Messages <- Message{
		Topic: msg.Topic(),
		Value: string(msg.Payload()),
	}
}

func (c *Client) onConnect(_ mqtt.Client) {
	c.OnConnect()
}

func (c *Client) Connect() error {
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Client) Subscribe(topic string) {
	const qos = 2
	token := c.client.Subscribe(topic, qos, nil)
	token.Wait()
}
