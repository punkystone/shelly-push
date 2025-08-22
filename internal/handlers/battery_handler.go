package handlers

import (
	"fmt"
	"go_test/internal/firebase"

	"github.com/rs/zerolog/log"
)

type BatteryHandler struct {
	connectMessage bool
	previousState  *string
	firebaseClient *firebase.Client
	fireBaseTopic  string
}

func NewBatteryHandler(firebaseClient *firebase.Client, fireBaseTopic string) *BatteryHandler {
	return &BatteryHandler{
		connectMessage: true,
		previousState:  nil,
		firebaseClient: firebaseClient,
		fireBaseTopic:  fireBaseTopic,
	}
}

func (batteryHandler *BatteryHandler) Handle(message string) error {
	if batteryHandler.connectMessage {
		batteryHandler.connectMessage = false
		batteryHandler.previousState = &message
		return nil
	}
	if batteryHandler.previousState != nil && message == *batteryHandler.previousState {
		return nil
	}
	batteryHandler.previousState = &message
	log.Info().Msgf("Received battery state: %s", message)
	err := batteryHandler.firebaseClient.SendToTopic(batteryHandler.fireBaseTopic, message)
	if err != nil {
		return fmt.Errorf("error sending message to topic: %w", err)
	}
	return nil
}
