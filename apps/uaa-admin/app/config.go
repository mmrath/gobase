package app

import "github.com/mmrath/gobase/model"

type Config struct {
	DB  model.DBConfig `yaml:"db"`
	Web WebConfig      `yaml:"web"`
}

type WebConfig struct {
	URL         string `yaml:"url"`
	Port        string `yaml:"port"`
	CorsEnabled bool   `yaml:"corsEnabled"`
	TemplateDir string `yaml:"templateDir"`
}

