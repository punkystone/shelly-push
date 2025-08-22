package handlers

import (
	"encoding/json"
	"fmt"
	"go_test/internal/firebase"
	"strings"

	"github.com/rs/zerolog/log"
)

const openState = "open"
const closeState = "close"

type infoJSON struct {
	Sensor struct {
		State string `json:"state"`
	} `json:"sensor"`
}

type DoorHandler struct {
	connectMessage bool
	previousState  *bool
	firebaseClient *firebase.Client
	fireBaseTopic  string
}

func NewDoorHandler(firebaseClient *firebase.Client, fireBaseTopic string) *DoorHandler {
	return &DoorHandler{
		connectMessage: true,
		previousState:  nil,
		firebaseClient: firebaseClient,
		fireBaseTopic:  fireBaseTopic,
	}
}

func (doorHandler *DoorHandler) Handle(message string) error {
	state, err := doorHandler.parseStatus(message)
	if err != nil {
		return fmt.Errorf("failed to parse status: %w", err)
	}
	if doorHandler.connectMessage {
		doorHandler.connectMessage = false
		if state == openState {
			doorHandler.previousState = new(bool)
			*doorHandler.previousState = true
		}
		if state == closeState {
			doorHandler.previousState = new(bool)
			*doorHandler.previousState = false
		}
		return nil
	}
	if doorHandler.previousState != nil && ((state == openState && *doorHandler.previousState) || (state == closeState && !*doorHandler.previousState)) {
		return nil
	}
	if state == openState {
		doorHandler.previousState = new(bool)
		*doorHandler.previousState = true
	}
	if state == closeState {
		doorHandler.previousState = new(bool)
		*doorHandler.previousState = false
	}
	log.Info().Msgf("Received door state: %s", state)
	switch state {
	case openState:
		err := doorHandler.firebaseClient.SendToTopic(doorHandler.fireBaseTopic, "Door was opened")
		if err != nil {
			return fmt.Errorf("error sending message to topic: %w", err)
		}
	case closeState:
		err := doorHandler.firebaseClient.SendToTopic(doorHandler.fireBaseTopic, "Door was closed")
		if err != nil {
			return fmt.Errorf("error sending message to topic: %w", err)
		}
	}
	return nil
}

func (doorHandler *DoorHandler) parseStatus(info string) (string, error) {
	decoded := json.NewDecoder(strings.NewReader(info))
	infoJSON := &infoJSON{}
	err := decoded.Decode(&infoJSON)
	if err != nil {
		return "", err
	}
	return infoJSON.Sensor.State, nil
}
