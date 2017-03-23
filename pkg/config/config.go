package config

import (
	"encoding/json"
	"os"
)

// ConfigurationInfo contains the config file's content
type ConfigurationInfo struct {
	MongoDBURL  string `json:"mongoDBURL"`
	MongoUseTLS bool   `json:"mongoUseTLS"`
	APIToken    string `json:"apiToken"`
	GitHubToken string `json:"gitHubToken"`
	AutoTLS     bool   `json:"autoTLS"`
	Bind        string `json:"bind"`
	Hostname    string `json:"hostname"`
	CertCache   string `json:"certCache"`
}

// GetConfiguration reads the confiruration from config.json and returns it
func GetConfiguration() ConfigurationInfo {
	returnConfig := ConfigurationInfo{}
	data, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(data)
	err = jsonParser.Decode(&returnConfig)
	if err != nil {
		panic(err)
	}
	return returnConfig
}
