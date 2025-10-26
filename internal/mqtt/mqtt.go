package mqtt

import (
	"context"
	"go_test/internal/tailscale"
	"net"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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

const connectWait = 20 * time.Second

func NewClient(connectUrl string, clientID string, tailscaleServer *tailscale.Server) (*Client, error) {
	client := &Client{client: nil, Messages: make(chan Message)}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectUrl)
	opts.SetClientID(clientID)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetResumeSubs(true)
	opts.SetConnectRetryInterval(5 * time.Second)
	opts.SetMaxReconnectInterval(2 * time.Minute)
	opts.SetDefaultPublishHandler(client.publishHandler)
	opts.SetCustomOpenConnectionFn(func(uri *url.URL, options mqtt.ClientOptions) (net.Conn, error) {
		return tailscaleServer.TSServer.Dial(context.Background(), "tcp", uri.Host)
	})
	opts.OnConnect = client.onConnect
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
	if token := c.client.Connect(); token.WaitTimeout(connectWait) && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Client) Subscribe(topic string) {
	const qos = 2
	token := c.client.Subscribe(topic, qos, nil)
	token.Wait()
}
