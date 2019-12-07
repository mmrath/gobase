package config

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type Config struct {
	DB  DBConfig  `mapstructure:"db"`
	Web WebConfig `mapstructure:"web"`
}

type DBConfig struct {
	URL string `mapstructure:"url"`
}

type WebConfig struct {
	Port string `mapstructure:"port"`
	CorsEnabled bool `mapstructure:"corsEnabled"`
	ContextPath string `mapstructure:"contextPath"`
}

func LoadConfig(files ...string) (*Config, error) {
	envPrefix := "AUTH_SERVER"

	cfg := &Config{
		Web: WebConfig{Port: ":8020"},
	}
	v := viper.New()
	v.BindEnv("db.url")
	// Viper settings
	v.AddConfigPath(".")
	v.AddConfigPath(fmt.Sprintf("$%s_CONFIG_DIR/", strings.ToUpper(envPrefix)))

	// Environment variable settings
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	// Global configuration
	v.SetDefault("environment", "production")
	v.SetDefault("debug", false)
	v.SetDefault("shutdownTimeout", 15*time.Second)
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		v.SetDefault("no_color", true)
	}

	for _, configFile := range files {
		if fileExists(configFile) {
			pf, err := os.Open(configFile)

			if err != nil {
				log.Error().Err(err).Str("configPath", configFile).Msg("failed to read profile config file")
				return nil, fmt.Errorf("failed to read config file: %s, error: %v", configFile, err)
			}

			if err := v.MergeConfig(pf); err != nil {
				return nil, fmt.Errorf("failed to read the configuration file: %s, error: %v", configFile, err)
			}
			log.Info().Str("configPath", configFile).Msg("successfully loaded config")
		} else {
			log.Error().Str("configPath", configFile).Msg("config does not exist")
			return nil, fmt.Errorf("config file %s does not exist", configFile)
		}
	}

	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshall config file: %v", err)
	}
	log.Info().Interface("config", cfg).Msg("successfully loaded configuration")

	return cfg, nil

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
