package main

import (
	"go_test/internal/firebase"
	"go_test/internal/info"
	"go_test/internal/mqtt"
	"go_test/internal/util"

	"github.com/rs/zerolog/log"
)

const openState = "open"
const closeState = "close"

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
	var previousState bool
	for message := range mqttClient.Messages {
		processMessage(&previousState, &connectMessage, message, env.MqttTopic, env.FirebaseTopic, firebaseClient)
	}
}

func processMessage(previousState *bool, connectMessage *bool, message mqtt.Message, mqttTopic string, firebaseTopic string, firebaseClient *firebase.Client) {
	if message.Topic != mqttTopic {
		return
	}
	state, err := info.ParseStatus(message.Value)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse status")
		return
	}
	if *connectMessage {
		*connectMessage = false
		if state == openState {
			*previousState = true
		}
		if state == closeState {
			*previousState = false
		}
		return
	}
	if previousState != nil && ((state == openState && *previousState) || (state == closeState && !*previousState)) {
		return
	}
	if state == openState {
		*previousState = true
	}
	if state == closeState {
		*previousState = false
	}

	log.Info().Msgf("Received state: %s", state)
	if state == openState {
		sendPushNotification(firebaseClient, firebaseTopic, "Door was opened")
	} else if state == closeState {
		sendPushNotification(firebaseClient, firebaseTopic, "Door was closed")
	}
}

func sendPushNotification(firebaseClient *firebase.Client, topic string, title string) {
	err := firebaseClient.SendToTopic(topic, title)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send message")
	}
}
