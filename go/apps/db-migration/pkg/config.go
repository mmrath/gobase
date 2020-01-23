package pkg

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/go/pkg/db"
)

type Config struct {
	DB           db.Config
	MigrationDir string `split_words:"true" required:"true"`
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
