package account

import (
	"github.com/google/wire"

	"github.com/mmrath/gobase/common/email"
	"github.com/mmrath/gobase/model"
	"github.com/mmrath/gobase/uaa-server/pkg/config"
)

var Provider = wire.NewSet(
	wire.Bind(new(Service), new(*service)),
	wire.Bind(new(Notifier), new(*notifier)),
	NewHandler,
	NewService,
	NewNotifier,
)

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func NewService(db *model.DB, notifier Notifier) *service {
	return &service{
		notifier: notifier,
		db:       db,
	}
}

func NewNotifier(config *config.Config, mailer email.Mailer) *notifier {
	return &notifier{
		sender:  mailer,
		baseUrl: config.Web.ExternalURL,
	}
}
