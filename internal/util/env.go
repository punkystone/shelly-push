package util

import (
	"errors"
	"os"
	"strconv"
)

type Env struct {
	ShellyMqttTopic      string
	BatteryMqttTopic     string
	FirebaseProjectID    string
	ShellyFirebaseTopic  string
	BatteryFirebaseTopic string
	FirebaseKeyPath      string
	MqttURL              string
	MqttClientID         string
	Debug                bool
}

func CheckEnv() (*Env, error) {
	shellyMqttTopic, exists := os.LookupEnv("SHELLY_MQTT_TOPIC")
	if !exists {
		return nil, errors.New("SHELLY_MQTT_TOPIC environment variable not set")
	}
	batteryMqttTopic, exists := os.LookupEnv("BATTERY_MQTT_TOPIC")
	if !exists {
		return nil, errors.New("BATTERY_MQTT_TOPIC environment variable not set")
	}
	firebaseProjectID, exists := os.LookupEnv("FIREBASE_PROJECT_ID")
	if !exists {
		return nil, errors.New("FIREBASE_PROJECT_ID environment variable not set")
	}
	shellyFirebaseTopic, exists := os.LookupEnv("SHELLY_FIREBASE_TOPIC")
	if !exists {
		return nil, errors.New("SHELLY_FIREBASE_TOPIC environment variable not set")
	}
	batteryFirebaseTopic, exists := os.LookupEnv("BATTERY_FIREBASE_TOPIC")
	if !exists {
		return nil, errors.New("BATTERY_FIREBASE_TOPIC environment variable not set")
	}
	firebaseKeyPath, exists := os.LookupEnv("FIREBASE_KEY_PATH")
	if !exists {
		return nil, errors.New("FIREBASE_KEY_PATH environment variable not set")
	}
	mqttURL, exists := os.LookupEnv("MQTT_URL")
	if !exists {
		return nil, errors.New("MQTT_URL environment variable not set")
	}
	mqttClientID, exists := os.LookupEnv("MQTT_CLIENT_ID")
	if !exists {
		return nil, errors.New("MQTT_CLIENT_ID environment variable not set")
	}
	debug, exists := os.LookupEnv("DEBUG")
	if !exists {
		return nil, errors.New("DEBUG environment variable not set")
	}
	debugParsed, err := strconv.ParseBool(debug)
	if err != nil {
		return nil, errors.New("DEBUG  environment variable not a bool")
	}

	env := &Env{
		ShellyMqttTopic:      shellyMqttTopic,
		BatteryMqttTopic:     batteryMqttTopic,
		FirebaseProjectID:    firebaseProjectID,
		ShellyFirebaseTopic:  shellyFirebaseTopic,
		BatteryFirebaseTopic: batteryFirebaseTopic,
		FirebaseKeyPath:      firebaseKeyPath,
		MqttURL:              mqttURL,
		MqttClientID:         mqttClientID,
		Debug:                debugParsed,
	}
	return env, nil
}
