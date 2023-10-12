package config

import (
	"os"

	"github.com/spf13/viper"
)

func GetConfig(v *viper.Viper) (*Config, error) {
	var config Config
	args := os.Args
	for _, arg := range args {
		if arg == "--version" {
			config.Version = "1.0.0"
			return &config, nil
		}
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	// Попытка чтения конфига
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Попытка заполнение структуры Config полученными данными
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
