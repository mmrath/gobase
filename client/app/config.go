package app

import (
	"github.com/mmrath/gobase/common/auth"
	"github.com/mmrath/gobase/common/config"
	"github.com/mmrath/gobase/common/email"
	"github.com/mmrath/gobase/model"
)

type Config struct {
	Web  WebConfig        `yaml:"web"`
	DB   model.DBConfig   `yaml:"db"`
	SMTP email.SMTPConfig `yaml:"smtp"`
	JWT  auth.JWTConfig   `yaml:"jwt"`
}

type WebConfig struct {
	URL         string `yaml:"url"`
	Port        string `yaml:"port"`
	CorsEnabled bool   `yaml:"corsEnabled"`
}

func LoadConfig(profiles ...string) Config {
	cfg := Config{}
	err := config.LoadConfig(&cfg, "./resources", profiles...)
	if err != nil {
		panic(err)
	}
	return cfg
}
