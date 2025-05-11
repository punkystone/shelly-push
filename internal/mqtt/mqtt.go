package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type message struct {
	Topic string
	Value string
}

type Client struct {
	client    mqtt.Client
	Messages  chan message
	OnConnect func()
}

func NewClient(url string, clientID string, username string, password string, caPath string, clientCertPath string, clientKeyPath string) (*Client, error) {
	client := &Client{client: nil, Messages: make(chan message)}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(url)
	opts.SetClientID(clientID)
	tlsConfig, err := client.getTLSConfig(caPath, clientCertPath, clientKeyPath)
	if err != nil {
		return nil, err
	}
	opts.SetTLSConfig(tlsConfig)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(client.publishHandler)
	opts.OnConnect = client.onConnect
	opts.AutoReconnect = true
	client.client = mqtt.NewClient(opts)
	return client, nil
}

func (c *Client) getTLSConfig(caPath string, clientCertPath string, clientKeyPath string) (*tls.Config, error) {
	certpool := x509.NewCertPool()
	ca, err := os.ReadFile(caPath)
	if err != nil {
		return nil, err
	}
	certpool.AppendCertsFromPEM(ca)
	clientKeyPair, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		RootCAs:      certpool,
		ClientAuth:   tls.NoClientCert,
		ClientCAs:    nil,
		MinVersion:   tls.VersionTLS13,
		Certificates: []tls.Certificate{clientKeyPair},
	}, nil
}

func (c *Client) publishHandler(_ mqtt.Client, msg mqtt.Message) {
	c.Messages <- message{
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
