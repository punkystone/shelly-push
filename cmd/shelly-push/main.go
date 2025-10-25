package main

import (
	"fmt"
	"go_test/internal/firebase"
	"go_test/internal/handlers"
	"go_test/internal/mqtt"
	"go_test/internal/tailscale"
	"go_test/internal/util"
	"io"

	"github.com/rs/zerolog/log"
)

func main() {
	env, err := util.CheckEnv()
	if err != nil {
		panic(err)
	}
	tailScaleServer := tailscale.NewServer(env.TSHostname, env.TSAuthKey, env.TSControlURL)
	err = tailScaleServer.Connect()
	if err != nil {
		panic(err)
	}
	resp, err := tailScaleServer.TSServer.HTTPClient().Get("http://192.168.0.27:8000")
	if err != nil {
		panic(err)
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	firebaseClient, err := firebase.Init(env.FirebaseKeyPath, env.FirebaseProjectID)
	if err != nil {
		panic(err)
	}
	log.Info().Msg("Firebase client initialized")
	mqttClient, err := mqtt.NewClient(env.MqttURL, env.MqttClientID, tailScaleServer)
	if err != nil {
		panic(err)
	}
	mqttClient.OnConnect = func() {
		log.Info().Msg("Connected to MQTT broker")
	}
	err = mqttClient.Connect()
	if err != nil {
		panic(err)
	}
	mqttClient.Subscribe(env.ShellyMqttTopic)
	mqttClient.Subscribe(env.BatteryMqttTopic)

	doorHandler := handlers.NewDoorHandler(firebaseClient, env.ShellyFirebaseTopic)
	batteryHandler := handlers.NewBatteryHandler(firebaseClient, env.BatteryFirebaseTopic)
	for message := range mqttClient.Messages {
		fmt.Printf("Topic %s Message %s\n", message.Topic, message.Value)

		switch message.Topic {
		case env.ShellyMqttTopic:
			err := doorHandler.Handle(message.Value)
			if err != nil {
				log.Error().Err(err)
			}
		case env.BatteryMqttTopic:
			err := batteryHandler.Handle(message.Value)
			if err != nil {
				log.Error().Err(err)
			}
		}
	}
}
