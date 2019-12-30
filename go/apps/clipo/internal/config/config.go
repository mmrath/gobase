package config

import (
	"fmt"
	"github.com/mmrath/gobase/go/pkg/auth"
	"github.com/mmrath/gobase/go/pkg/db"
	"github.com/mmrath/gobase/go/pkg/email"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	DevMode bool             `mapstructure:"devMode" yaml:"devMode"`
	Web     WebConfig        `mapstructure:"web" yaml:"web" `
	DB      db.Config        `mapstructure:"db" yaml:"db"`
	SMTP    email.SMTPConfig `mapstructure:"smtp" yaml:"smtp"`
	JWT     auth.JWTConfig   `mapstructure:"jwt" yaml:"jwt"`
}

type WebConfig struct {
	URL         string `mapstructure:"url" yaml:"url"`
	Port        string `mapstructure:"port" yaml:"port"`
	CorsEnabled bool   `mapstructure:"corsEnabled" yaml:"corsEnabled"`
}

func LoadConfig(cfg interface{}, profiles ...string) error {
	envPrefix := "CLIPO"

	configPathEnvVar := fmt.Sprintf("%s_CONFIG_DIR", strings.ToUpper(envPrefix))
	configDir := os.Getenv(configPathEnvVar)


	if configDir == "" {
		configDir = "./config"
	}

	log.Info().Str("configDir", configDir).Msg("resolved config dir")

	v := viper.New()
	v.BindEnv("db.url")
	// Viper settings
	if dirExists(configDir) {
		log.Info().Str("configDir", configDir).Msg("dir exists")
		v.AddConfigPath(configDir)
	}

	v.SetConfigName("app")
	v.SetConfigType("yaml")

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

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %v", err)
	}

	for _, profile := range profiles {
		profileConfig := filepath.Join(configDir, "app-"+profile+".yml")
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

	if err := v.Unmarshal(cfg); err != nil {
		return fmt.Errorf("failed to unmarshall config file: %v", err)
	}
	log.Info().Interface("config", cfg).Msg("successfully loaded configuration")
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExists(dirName string) bool {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
