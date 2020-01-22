package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/mmrath/gobase/go/pkg/auth"
	"github.com/mmrath/gobase/go/pkg/db"
	"github.com/mmrath/gobase/go/pkg/email"
	"github.com/mmrath/gobase/go/pkg/errutil"
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
	err := envconfig.Process("", cfg)
	return errutil.Wrap(err, "failed to load config")
}
