package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/mmrath/gobase/pkg/auth"
	"github.com/mmrath/gobase/pkg/db"
	"github.com/mmrath/gobase/pkg/email"
	"github.com/mmrath/gobase/pkg/errutil"
)

type Config struct {
	DevMode       bool             `yaml:"devMode" split_words:"true"`
	AppDomainName string           `required:"true" split_words:"true"`
	Web           WebConfig        `yaml:"web"`
	DB            db.Config        `yaml:"db"`
	SMTP          email.SMTPConfig `yaml:"smtp"`
	JWT           auth.JWTConfig   `yaml:"jwt"`
}

type WebConfig struct {
	Port        string `default:"9010" yaml:"port"`
	CorsEnabled bool   `default:"false" split_words:"true" yaml:"corsEnabled"`
}

func LoadConfig(cfg *Config) error {
	err := envconfig.Process("", cfg)
	return errutil.Wrap(err, "failed to load config")
}
