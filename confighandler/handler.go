package confighandler

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, err
		}
	}

	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func init() {
	config, err := LoadConfig("./secrets/secret.yaml")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Printf("API Port: %s\n", config.API.Port)
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Database Port: %d\n", config.Database.Port)
	fmt.Printf("Database Username: %s\n", config.Database.Username)
	fmt.Printf("Database Password: %s\n", config.Database.Password)
	fmt.Printf("Database Name: %s\n", config.Database.Name)
}
