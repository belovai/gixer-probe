package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig      `mapstructure:"app"`
	Api ApiConfig      `mapstructure:"api"`
	Rmq RabbitMQConfig `mapstructure:"rabbitmq"`
}

type AppConfig struct {
	Uuid string `mapstructure:"uuid"`
}

type ApiConfig struct {
	Host  string `mapstructure:"host"`
	Token string `mapstructure:"token"`
}

type RabbitMQConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Vhost    string `mapstructure:"vhost"`
	Queue    string `mapstructure:"queue"`
}

func NewConfig() (config Config, err error) {
	_ = godotenv.Load()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	// replace . (dot) with _ (underscore) in env vars, so you can overwrite config values with env vars.
	// E.g., RABBITMQ_HOST becomes rabbitmq.host
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, err
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return config, nil
}
