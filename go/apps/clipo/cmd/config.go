package cmd

import (
	"github.com/mmrath/gobase/go/apps/clipo/internal/config"
)

func LoadConfig(profiles ...string) config.Config {
	cfg := config.Config{
		Web: config.WebConfig{
			Port: ":6010",
		},
	}

	err := config.LoadConfig(&cfg, profiles...)
	if err != nil {
		panic(err)
	}
	return cfg
}
