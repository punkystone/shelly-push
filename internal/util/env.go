package util

import (
	"errors"
	"os"
)

type Env struct {
	MqttTopic          string
	FirebaseProjectID  string
	FirebaseTopic      string
	FirebaseKeyPath    string
	MqttURL            string
	MqttClientID       string
	MqttUsername       string
	MqttPassword       string
	MqttCAPath         string
	MqttClientCertPath string
	MqttClientKeyPath  string
}

func CheckEnv() (*Env, error) {
	mqttTopic, exists := os.LookupEnv("MQTT_TOPIC")
	if !exists {
		return nil, errors.New("MQTT_TOPIC environment variable not set")
	}
	firebaseProjectID, exists := os.LookupEnv("FIREBASE_PROJECT_ID")
	if !exists {
		return nil, errors.New("FIREBASE_PROJECT_ID environment variable not set")
	}
	firebaseTopic, exists := os.LookupEnv("FIREBASE_TOPIC")
	if !exists {
		return nil, errors.New("FIREBASE_TOPIC environment variable not set")
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
	mqttUsername, exists := os.LookupEnv("MQTT_USERNAME")
	if !exists {
		return nil, errors.New("MQTT_USERNAME environment variable not set")
	}
	mqttPassword, exists := os.LookupEnv("MQTT_PASSWORD")
	if !exists {
		return nil, errors.New("MQTT_PASSWORD environment variable not set")
	}
	mqttCAPath, exists := os.LookupEnv("MQTT_CA_PATH")
	if !exists {
		return nil, errors.New("MQTT_CA_PATH environment variable not set")
	}
	mqttClientCertPath, exists := os.LookupEnv("MQTT_CLIENT_CERT_PATH")
	if !exists {
		return nil, errors.New("MQTT_CLIENT_CERT_PATH environment variable not set")
	}
	mqttClientKeyPath, exists := os.LookupEnv("MQTT_CLIENT_KEY_PATH")
	if !exists {
		return nil, errors.New("MQTT_CLIENT_KEY_PATH environment variable not set")
	}

	env := &Env{
		MqttTopic:          mqttTopic,
		FirebaseProjectID:  firebaseProjectID,
		FirebaseTopic:      firebaseTopic,
		FirebaseKeyPath:    firebaseKeyPath,
		MqttURL:            mqttURL,
		MqttClientID:       mqttClientID,
		MqttUsername:       mqttUsername,
		MqttPassword:       mqttPassword,
		MqttCAPath:         mqttCAPath,
		MqttClientCertPath: mqttClientCertPath,
		MqttClientKeyPath:  mqttClientKeyPath,
	}
	return env, nil
}
