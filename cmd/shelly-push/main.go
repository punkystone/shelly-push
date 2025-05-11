package main

import (
	"go_test/internal/firebase"
	"go_test/internal/info"
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
	mqttClient.Subscribe(env.MqttTopic)
	log.Info().Msgf("Subscribed to topic: %s", env.MqttTopic)
	connectMessage := true
	for message := range mqttClient.Messages {
		if message.Topic == env.MqttTopic {
			if connectMessage {
				connectMessage = false
				continue
			}
			state, err := info.ParseStatus(message.Value)
			if err != nil {
				log.Error().Err(err).Msg("Failed to parse status")
				continue
			}
			log.Info().Msgf("Received state: %s", state)
			if state == "open" {
				sendPushNotification(firebaseClient, env.FirebaseTopic, "Door was opened")
			} else if state == "close" {
				sendPushNotification(firebaseClient, env.FirebaseTopic, "Door was closed")
			}
		}
	}
}

func sendPushNotification(firebaseClient *firebase.Client, topic string, title string) {
	err := firebaseClient.SendToTopic(topic, title)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send message")
	}
}
