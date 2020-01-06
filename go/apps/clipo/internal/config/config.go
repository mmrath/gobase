package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mmrath/gobase/go/pkg/auth"
	"github.com/mmrath/gobase/go/pkg/db"
	"github.com/mmrath/gobase/go/pkg/email"

	"os"
)

type Config struct {
	DevMode bool             `envconfig:"" yaml:"devMode"`
	Web     WebConfig        `yaml:"web"`
	DB      db.Config        `yaml:"db"`
	SMTP    email.SMTPConfig `yaml:"smtp"`
	JWT     auth.JWTConfig   `yaml:"jwt"`
}

type WebConfig struct {
	URL         string `envconfig:"optional" yaml:"url"`
	Port        string `default:"9010" yaml:"port"`
	CorsEnabled bool   `default:"false" yaml:"corsEnabled"`
}

func LoadConfig(cfg *Config) error {
	envconfig.Process("", cfg)
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
