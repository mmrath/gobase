package cmd

import (
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/go/apps/clipo/internal/config"
)

func LoadConfig() config.Config {
	cfg := config.Config{
		Web: config.WebConfig{
			Port: ":9010",
		},
	}

	err := config.LoadConfig(&cfg)
	if err != nil {
		panic(err)
	}

	log.Info().Interface("config", cfg).Msg("config loaded successfully")

	return cfg
}
