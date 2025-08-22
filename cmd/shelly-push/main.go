package main

import (
	"go_test/internal/firebase"
	"go_test/internal/handlers"
	"go_test/internal/mqtt"
	"go_test/internal/util"

	"github.com/rs/zerolog/log"
)

func main() {
	env, err := util.CheckEnv()
	if err != nil {
		panic(err)
	}
	firebaseClient, err := firebase.Init(env.FirebaseKeyPath, env.FirebaseProjectID)
	if err != nil {
		panic(err)
	}
	log.Info().Msg("Firebase client initialized")
	mqttClient, err := mqtt.NewClient(env.MqttURL, env.MqttClientID, env.MqttUsername, env.MqttPassword, env.MqttCAPath, env.MqttClientCertPath, env.MqttClientKeyPath)
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
