package config

import (
	"fmt"
	"github.com/mmrath/gobase/pkg/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type Config struct {
	DevMode bool      `mapstructure:"devMode"`
	DB      db.Config `mapstructure:"db"`
	Web     WebConfig `mapstructure:"web"`
	SSO     SSOConfig `mapstructure:"ssoCookie"`
}

type SSOConfig struct {
	CookieName            string
	CookieDomain          string
	CookieValidityMinutes int64
	JwtPrivateKeyPath     string
}

type WebConfig struct {
	Port        string `mapstructure:"port"`
	CorsEnabled bool   `mapstructure:"corsEnabled"`
	ContextPath string `mapstructure:"contextPath"`
	SSLCertPath string `mapstructure:"sslCertPath"`
	SSLKeyPath  string `mapstructure:"sslKeyPath"`
}

func LoadConfig(files ...string) (*Config, error) {
	envPrefix := "UAA"
	configPathEnvVar := fmt.Sprintf("$%s_CONFIG_DIR/", strings.ToUpper(envPrefix))
	//appEnv := fmt.Sprintf("%s_ENV", envPrefix)

	cfg := &Config{
		Web: WebConfig{
			Port:        ":6010",
			SSLCertPath: "dist/ssl_certs/ssl_public.crt",
			SSLKeyPath:  "dist/ssl_certs/ssl_private.key",
		},
		SSO: SSOConfig{
			CookieName:            "SSO",
			CookieDomain:          "",
			CookieValidityMinutes: 60,
			JwtPrivateKeyPath:     "dist/key_pair/sso_private.key",
		},
	}
	v := viper.New()
	v.BindEnv("db.url")
	// Viper settings
	if dirExists("./config") {
		v.AddConfigPath("./config")
	}
	v.AddConfigPath(configPathEnvVar)
	v.SetConfigName("app")

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

	//envSpecificFileName := "app."++""

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

func dirExists(dirName string) bool {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
