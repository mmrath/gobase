package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"mmrath.com/gobase/common/auth"
	"mmrath.com/gobase/common/email"
	"mmrath.com/gobase/common/log"
	"mmrath.com/gobase/model"
)

type Config struct {
	Server Server           `yaml:"server"`
	DB     model.DBConfig   `yaml:"db"`
	SMTP   email.SMTPConfig `yaml:"smtp"`
	JWT    auth.JWTConfig   `yaml:"jwt"`
}

type Server struct {
	URL         string `yaml:"url"`
	Port        string `yaml:"port"`
	CorsEnabled bool   `yaml:"corsEnabled"`
}

func LoadConfig(resourceRoot string, profiles ...string) Config {

	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	log.Info("Executable location:", path)

	envPrefix := "ARA"

	v := viper.New()

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath(filepath.Join(resourceRoot, "config"))

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("failed to read the configuration file: %v", err))
	}

	for _, profile := range profiles {
		profileConfig := filepath.Join(resourceRoot, "config", "app-"+profile+".yml")
		if fileExists(profileConfig) {
			pf, err := os.Open(profileConfig)

			if err != nil {
				log.Error("Failed to read profile config file", profileConfig)
				panic(fmt.Sprintf("failed to config file: %s, error: %v", profileConfig, err))
			}

			if err := v.MergeConfig(pf); err != nil {
				panic(fmt.Sprintf("Failed to read the configuration file: %s, error: %v", profileConfig, err))
			}
			log.Info("Loaded config from ", profileConfig)
		} else {
			log.Warn("Config does not exist ", profileConfig)
		}
	}

	var config = Config{}
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("failed to unmarshall config file: %v", err))
	}
	log.WithField("config", config).Info("Successfully loaded configuration")
	return config
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
