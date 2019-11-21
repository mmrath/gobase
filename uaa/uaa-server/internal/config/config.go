package config

import (
	"fmt"

	"github.com/mmrath/gobase/common/config"
	"github.com/mmrath/gobase/model"
)

type Config struct {
	DB    model.DBConfig `yaml:"db"`
	Web   WebConfig      `yaml:"web"`
	Hydra HydraConfig    `yaml:"hydra"`
}

type WebConfig struct {
	ExternalURL string `yaml:"externalUrl"`
	URL         string `yaml:"url"`
	Port        string `yaml:"port"`
	ContextPath string `yaml:"contextPath"`
	CorsEnabled bool   `yaml:"corsEnabled"`
	TemplateDir string `yaml:"templateDir"`
}

type HydraConfig struct {
	Host     string
	BasePath string
}

func LoadConfig(resourceRoot string, profiles ...string) (*Config, error) {
	cfg := &Config{}
	err := config.LoadConfig(cfg, resourceRoot, profiles...)
	if err != nil {
		return nil, err
	}

	if cfg.Hydra.Host == "" {
		return nil, fmt.Errorf("hydra host must be specified")
	}

	return cfg, nil
}
