package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Token string `json:"Token"`
}

func LoadConfigFile(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
