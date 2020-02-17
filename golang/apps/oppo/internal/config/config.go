package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/mmrath/gobase/golang/pkg/db"
	"github.com/mmrath/gobase/golang/pkg/errutil"
)

type Config struct {
	DB  db.Config `yaml:"db"`
	Web WebConfig `yaml:"web"`
}

type WebConfig struct {
	URL         string `yaml:"url"`
	Port        string `yaml:"port"`
	CorsEnabled bool   `yaml:"corsEnabled"`
	TemplateDir string `yaml:"templateDir"`
}

func LoadConfig(cfg *Config) error {
	err := envconfig.Process("", cfg)
	return errutil.Wrap(err, "failed to load config")
}
