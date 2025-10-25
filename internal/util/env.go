package util

import (
	"errors"
	"os"
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
	TSHostname           string
	TSControlURL         string
	TSAuthKey            string
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
	tsHostname, exists := os.LookupEnv("TS_HOSTNAME")
	if !exists {
		return nil, errors.New("TS_HOSTNAME environment variable not set")
	}
	tsControlURL, exists := os.LookupEnv("TS_CONTROL_URL")
	if !exists {
		return nil, errors.New("TS_CONTROL_URL environment variable not set")
	}
	tsAuthKey, exists := os.LookupEnv("TS_AUTH_KEY")
	if !exists {
		return nil, errors.New("TS_AUTH_KEY environment variable not set")
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
		TSHostname:           tsHostname,
		TSControlURL:         tsControlURL,
		TSAuthKey:            tsAuthKey,
	}
	return env, nil
}
