package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DB dbConfig `yaml:"db"`
}

type dbConfig struct {
	URL string `yaml:"url"`
}

func LoadConfig(configPaths ...string) Config {
	envPrefix := "GO_BASE"

	v := viper.New()

	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	v.AddConfigPath("./resources/config")

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	var config = Config{}
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("failed to read the configuration file: %v", err))
	}
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("failed to unmarshall config file: %v", err))
	}
	logrus.WithField("config", config).Info("Successfully loaded configuration")
	return config
}
