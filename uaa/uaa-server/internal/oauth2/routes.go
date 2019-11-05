package oauth2

import (
	"github.com/go-chi/chi"
	"github.com/ory/hydra/sdk/go/hydra/client"

	"github.com/mmrath/gobase/uaa-server/internal/config"
)

func RegisterHandlers(r chi.Router, config *config.Config) {

	hydraTransportConfig := &client.TransportConfig{
		Schemes:  []string{"http"},
		Host:     config.Hydra.Host,
		BasePath: config.Hydra.BasePath,
	}

	hydraClient := client.NewHTTPClientWithConfig(nil, hydraTransportConfig)
	templateProvider,err  := loadTemplates(config.Web)
	if err!= nil {
		panic(err)
	}

	r.Get("/login", LoginGetHandler(hydraClient, templateProvider))
	r.Post("/login", LoginPostHandler(hydraClient, templateProvider))
	r.Get("/consent", ConsentGetHandler(hydraClient, templateProvider))
	r.Post("/consent", ConsentPostHandler(hydraClient, templateProvider))
}
