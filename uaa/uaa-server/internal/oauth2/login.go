package oauth2

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/ory/hydra/sdk/go/hydra/client"
	"github.com/rs/zerolog/log"

	hydraAdmin "github.com/ory/hydra/sdk/go/hydra/client/admin"
	hydraModels "github.com/ory/hydra/sdk/go/hydra/models"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type TemplateProvider interface {
	LoginTemplate() *template.Template
	ConsentTemplate() *template.Template
}

func LoginGetHandler(hydra *client.OryHydra, templateProvider TemplateProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("received login")

		keys, ok := r.URL.Query()["login_challenge"]

		if !ok || len(keys[0]) < 1 || keys[0] == "" {
			log.Info().Msg("Url Param 'login_challenge' is missing")
			w.Header().Set("X-Status-Reason", "no challenge")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		challenge := keys[0]

		getLoginRequestParams := hydraAdmin.NewGetLoginRequestParams()
		getLoginRequestParams.LoginChallenge = challenge

		getLoginResponse, err := hydra.Admin.GetLoginRequest(getLoginRequestParams)

		if err != nil {
			w.Header().Set("X-Status-Reason", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if getLoginResponse.Payload.Skip {
			// You can apply logic here, for example grant another scope, or do whatever...
			// ...

			// Now it's time to grant the login request. You could also deny the request if something went terribly wrong
			acceptLoginRequestParams := hydraAdmin.NewAcceptLoginRequestParams()
			var handledLoginRequest hydraModels.HandledLoginRequest
			acceptLoginRequestParams.LoginChallenge = challenge
			acceptLoginRequestParams.Body = &handledLoginRequest
			acceptLoginRequestParams.Body.Subject = &getLoginResponse.Payload.Subject

			acceptLoginResponse, err := hydra.Admin.AcceptLoginRequest(acceptLoginRequestParams)

			if err != nil {
				w.Header().Set("X-Status-Reason", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Error().Msg("redirecting from get login")

			// All we need to do now is to redirect the user back to hydra!
			http.Redirect(w, r, acceptLoginResponse.Payload.RedirectTo, http.StatusTemporaryRedirect)
		} else {
			data := map[string]interface{}{
				"title":     "Login",
				"challenge": challenge,
				"client":    getLoginResponse.Payload.Client,
			}
			err = templateProvider.LoginTemplate().Execute(w, data)
			if err != nil {
				log.Error().Err(err).Msg("error rendering network page")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

func LoginPostHandler(hydra *client.OryHydra, templateProvider TemplateProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username := r.FormValue("username")
		password := r.FormValue("password")

		challenge := r.FormValue("challenge")

		if challenge == "" {
			w.Header().Set("X-Status-Reason", "no challenge")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := validatePassword(username, password)

		if err != nil {

			data := map[string]interface{}{
				"title":     "Login failed",
				"challenge": challenge,
				"err":       err.Error(),
			}

			err = templateProvider.LoginTemplate().Execute(w, data)
			if err != nil {
				log.Error().Err(err).Msg("error rendering network page")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		acceptLoginRequestParams := hydraAdmin.NewAcceptLoginRequestParams()
		var handledLoginRequest hydraModels.HandledLoginRequest
		acceptLoginRequestParams.Body = &handledLoginRequest
		acceptLoginRequestParams.LoginChallenge = challenge

		// Subject is an alias for user ID. A subject can be a random string, a UUID, an email address, ....
		acceptLoginRequestParams.Body.Subject = &username

		// This tells hydra to remember the browser and automatically authenticate the user in future requests. This will
		// set the "skip" parameter in the other route to true on subsequent requests!
		acceptLoginRequestParams.Body.Remember = true

		// When the session expires, in seconds. Set this to 0 so it will never expire.
		acceptLoginRequestParams.Body.RememberFor = 0

		// Sets which "level" (e.g. 2-factor authentication) of authentication the user has. The value is really arbitrary
		// and optional. In the context of OpenID Connect, a value of 0 indicates the lowest authorization level.
		acceptLoginRequestParams.Body.ACR = "0"

		// Seems like the user authenticated! Let's tell hydra...
		acceptLoginResponse, err := hydra.Admin.AcceptLoginRequest(acceptLoginRequestParams)

		if err != nil {
			log.Info().Str("username", username).Msg("hydra accept login failed")
			w.Header().Set("X-Status-Reason", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// All we need to do now is to redirect the browser back to hydra!
		http.Redirect(w, r, acceptLoginResponse.Payload.RedirectTo, http.StatusTemporaryRedirect)
	}
}

func validatePassword(username string, password string) error {
	if username == password {
		log.Info().Str("username", username).Msg("login passed")
		return nil
	}
	log.Info().Str("username", username).Msg("login failed")
	return ErrInvalidCredentials
}
