package main

import (
	"encoding/json"
	"fmt"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadConfig() {
	fmt.Println("<<<	LOADING CONFIG	>>>")
	file, err := os.Open("config.json")
	check(err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	check(err)
	os.Setenv("DOCKER_API_VERSION", config.DockerAPIVersion)
}
