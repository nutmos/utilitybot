package confighandler

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (config ConfigStruct, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return ConfigStruct{}, err
		}
	}

	if err = viper.Unmarshal(&config); err != nil {
		return ConfigStruct{}, err
	}

	return config, nil
}

func init() {
	config, err := LoadConfig("./secrets/secret.yaml")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Printf("Telegram API Key: %s\n", config.ApiKey.Telegram)
}
