package app

import (
	"github.com/go-chi/chi"

	"github.com/mmrath/gobase/uaa-server/pkg/config"
)

type RouteHandler interface {
	RegisterHandlers(r chi.Router, config config.Config)
}

func Routes(account.) {

}
