package account

import (
	"github.com/google/wire"

	"github.com/mmrath/gobase/common/email"
	"github.com/mmrath/gobase/common/template_util"
	"github.com/mmrath/gobase/model"
	"github.com/mmrath/gobase/uaa-server/internal/config"
)

var Provider = wire.NewSet(
	wire.Bind(new(Service), new(*service)),
	wire.Bind(new(Notifier), new(*notifier)),
	TemplateRegistry,
	NewHandler,
	NewService,
	NewNotifier,
)

func TemplateRegistry(config *config.Config) *template_util.Registry {
	t, err := template_util.BuildRegistry(config.Web.TemplateDir)
	if err != nil {
		panic(err)
	}
	return t
}
func NewHandler(s Service, registry *template_util.Registry) *Handler {
	return &Handler{service: s, templateRegistry: registry}
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
