package info

import (
	"encoding/json"
	"strings"
)

type infoJSON struct {
	Sensor struct {
		State string `json:"state"`
	} `json:"sensor"`
}

func ParseStatus(info string) (string, error) {
	decoded := json.NewDecoder(strings.NewReader(info))
	infoJSON := &infoJSON{}
	err := decoded.Decode(&infoJSON)
	if err != nil {
		return "", err
	}
	return infoJSON.Sensor.State, nil
}
