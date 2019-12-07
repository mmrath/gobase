package config

import (
	"fmt"
	"github.com/mmrath/gobase/model"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type Config struct {
	DB    model.DBConfig `mapstructure:"db"`
	Web   WebConfig      `mapstructure:"web"`
	Hydra HydraConfig    `mapstructure:"hydra"`
}

type WebConfig struct {
	ExternalURL string `mapstructure:"externalUrl"`
	URL         string `mapstructure:"url"`
	Port        string `mapstructure:"port"`
	ContextPath string `mapstructure:"contextPath"`
	CorsEnabled bool   `mapstructure:"corsEnabled"`
	TemplateDir string `mapstructure:"templateDir"`
}

type HydraConfig struct {
	Host     string `mapstructure:"host"`
	BasePath string `mapstructure:"basePath"`
}

func LoadConfig(files ...string) (*Config, error) {
	dir := os.Getenv("UAA_DB_URL")
	envPrefix := "UAA"
	fmt.Println(dir)

	cfg := &Config{
		model.DBConfig{},
		WebConfig{
			Port:        "9020",
			TemplateDir: "uaa/uaa-web-app/build",
		},
		HydraConfig{Host: "127.0.0.1:9001"},
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
