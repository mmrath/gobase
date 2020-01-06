package pkg

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mmrath/gobase/go/pkg/db"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	DB           db.Config
	MigrationDir string
}

func LoadConfig() Config {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Error().Err(err).Msg("failed to load env config")
		os.Exit(1)
	}
	return cfg
}
