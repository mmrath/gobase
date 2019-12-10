package app

import (
	"github.com/mmrath/gobase/pkg/db"
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
