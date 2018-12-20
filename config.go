package main

import (
	"encoding/json"
	"os"
)

// Configuration struct for the global config
type Configuration struct {
	WorkingDir       string `json:"working_dir"`
	DockerAPIVersion string `json:"docker_api_version"`
}

var config Configuration

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	os.Setenv("DOCKER_API_VERSION", config.DockerAPIVersion)
}
