package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)


func LoadConfig(config interface{}, resourceRoot string, profiles ...string) error {

	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	log.Info().Str("executable", path).Send()

	envPrefix := "APP"

	v := viper.New()

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath(filepath.Join(resourceRoot, "config"))

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %v", err)
	}

	for _, profile := range profiles {
		profileConfig := filepath.Join(resourceRoot, "config", "app-"+profile+".yml")
		if fileExists(profileConfig) {
			pf, err := os.Open(profileConfig)

			if err != nil {
				log.Error().Err(err).Str("configPath", profileConfig).Msg("failed to read profile config file")
				return fmt.Errorf("failed to read config file: %s, error: %v", profileConfig, err)
			}

			if err := v.MergeConfig(pf); err != nil {
				return fmt.Errorf("failed to read the configuration file: %s, error: %v", profileConfig, err)
			}
			log.Info().Str("configPath", profileConfig).Msg("successfully loaded config")
		} else {
			log.Error().Str("configPath", profileConfig).Msg("config does not exist")
			return fmt.Errorf("config file %s does not exist", profileConfig)
		}
	}

	if err := v.Unmarshal(config); err != nil {
		return fmt.Errorf("failed to unmarshall config file: %v", err)
	}
	log.Info().Interface("config", config).Msg("successfully loaded configuration")
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
