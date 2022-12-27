package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
}

var AppConfig Config

func GetConfig() error {

	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}
	return nil
}
