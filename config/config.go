package config

import (
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type ApiKey struct {
	Aviationstack string `yaml:"aviationstack"`
	Telegram      string `yaml:"telegram"`
}

type ConfigStruct struct {
	ApiKey ApiKey `yaml:"apiKey"`
}

var (
	Config ConfigStruct
)

func init() {
	filename, _ := filepath.Abs("./secrets/secrets.yaml")
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
