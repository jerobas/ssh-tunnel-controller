package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port    int    `json:"port"`
	Url     string `json:"url"`
	KeyPath string `json:"keyPath"`
}

var config = Config{
	Port:    8192,
	Url:     "ec2-user@ec2-15-228-220-157.sa-east-1.compute.amazonaws.com",
	KeyPath: "/home/alfred/keys/ssh.pem",
}

func LoadConfig() {
	const path = "/etc/territo/config.json"

	if _, err := os.Stat(path); err == nil {
		file, err := os.Open(path)
		if err == nil {
			defer file.Close()
			json.NewDecoder(file).Decode(&config) // be sure to add validation later
		}
	}
}

func GetConfig() Config {
	return config
}
