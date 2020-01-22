package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/mmrath/gobase/go/pkg/errutil"
)

func LoadConfig(cfg interface{}) error {
	err := envconfig.Process("", cfg)
	return errutil.Wrap(err, "failed to load config")
}
