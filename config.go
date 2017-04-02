package main

import (
	"encoding/json"
	"log"
	"os"
	_ "strings"
)

type ConfigDatabase struct {
	User, Ip, Port, SslCertLocation, SslKeyLocation string
}

type Config struct {
	Port, OsUser string
	Database     ConfigDatabase
}

var config Config

func InitializeConfiguration() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening the configuration file: ", err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatalf("Error reading the configuration file: ", err)
	}
}
