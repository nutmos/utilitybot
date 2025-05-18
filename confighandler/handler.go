package confighandler

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (config ConfigStruct, err error) {
	viper.SetConfigName("secrets.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Config Read Error: File Not Found")
			return ConfigStruct{}, err
		}
		fmt.Printf("%v\n", err)
	}

	if err = viper.Unmarshal(&config); err != nil {
		fmt.Println("Config Unmarshal Error")
		return ConfigStruct{}, err
	}

	return config, nil
}

func init() {
	fmt.Println("Now Reading Config File")
	var err error
	Config, err = LoadConfig("./secrets")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Printf("Telegram API Key: %s\n", Config.ApiKey.Telegram)
}
