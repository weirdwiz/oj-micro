package main

// Config struct for the global config
type Config struct {
	WorkingDir       string `json:"working_dir"`
	DockerAPIVersion string `json:"docker_api_version"`
}

var config Config

func init() {
	config = loadConfig()
}

func loadConfig() Config {
	var parsed Config
	//load from memory
	return parsed
}
