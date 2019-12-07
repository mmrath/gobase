package oauth2

import (
	"net/http"

	"github.com/ory/hydra/sdk/go/hydra/client"
	hydraAdmin "github.com/ory/hydra/sdk/go/hydra/client/admin"
	"github.com/ory/hydra/sdk/go/hydra/models"
	hydraModels "github.com/ory/hydra/sdk/go/hydra/models"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/apps/uaa/internal/helpers"
)

func ConsentGetHandler(hydraClient *client.OryHydra, templateProvider TemplateProvider) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["consent_challenge"]

		if !ok || len(keys[0]) < 1 || keys[0] == "" {
			log.Info().Msg("URL param 'consent_challenge' is missing")
			w.Header().Set("X-Status-Reason", "no consent_challenge")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		challenge := keys[0]

		consentRequestParams := hydraAdmin.NewGetConsentRequestParams()
		consentRequestParams.ConsentChallenge = challenge
		getConsentResponse, err := hydraClient.Admin.GetConsentRequest(consentRequestParams)

		if err != nil {
			w.Header().Set("X-Status-Reason", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if getConsentResponse.Payload.Skip || isInternalClient(getConsentResponse.Payload.Client) {
			// You can apply logic here, for example grant another scope, or do whatever...
			// ...

			acceptConsentRequestParams := hydraAdmin.NewAcceptConsentRequestParams()
			var handledConsentRequest hydraModels.HandledConsentRequest
			acceptConsentRequestParams.ConsentChallenge = challenge
			acceptConsentRequestParams.Body = &handledConsentRequest

			// Now it's time to grant the consent request. You could also deny the request if something went terribly wrong

			// We can grant all scopes that have been requested - hydra already checked for us that no additional scopes
			// are requested accidentally.
			acceptConsentRequestParams.Body.GrantedScope = getConsentResponse.Payload.RequestedScope

			// Grant the roles scope to every client, so they can get the users roles over the backchannel
			// ... is roles scope already in requestBody.GrantScope
			// the client app will be able to request /api/v0/roles. This endpoint will only return
			// data if a valid appHash is provided in the querystring. Returned information
			// is strictly limited to the app requesting the roles. You only receive your roles in the
			// requesting app

			for i := 0; i < len(AutoScopes); i++ {
				if helpers.SliceIndex(len(acceptConsentRequestParams.Body.GrantedScope), func(j int) bool {
					return acceptConsentRequestParams.Body.GrantedScope[j] == AutoScopes[i]
				}) == -1 {
					acceptConsentRequestParams.Body.GrantedScope = append(acceptConsentRequestParams.Body.GrantedScope, AutoScopes[i])
				}
			}

			// ORY Hydra checks if requested audiences are allowed by the client, so we can simply echo this.
			acceptConsentRequestParams.Body.GrantedAudience = getConsentResponse.Payload.RequestedAudience

			// This data will be available when introspecting the token. Try to avoid sensitive information here,
			// unless you limit who can introspect tokens.
			// access_token: { foo: 'bar' },

			// This data will be available in the ID token.
			// id_token: { baz: 'bar' },

			session := new(models.ConsentRequestSessionData)
			acceptConsentRequestParams.Body.Session = session

			acceptConsentResponse, err := hydraClient.Admin.AcceptConsentRequest(acceptConsentRequestParams)

			if err != nil {
				w.Header().Set("X-Status-Reason", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// All we need to do now is to redirect the user back to hydra!
			http.Redirect(w, r, acceptConsentResponse.Payload.RedirectTo, http.StatusTemporaryRedirect)
		} else {
			// If consent can't be skipped we MUST show the consent UI.

			data := map[string]interface{}{
				"title":          "consent",
				"csrfToken":      "",
				"challenge":      challenge,
				"requestedScope": getConsentResponse.Payload.RequestedScope,
				"user":           getConsentResponse.Payload.Subject,
				"client":         getConsentResponse.Payload.Client,
			}

			err = templateProvider.ConsentTemplate().Execute(w, data)

			return

		}
	}

}

func isInternalClient(c *hydraModels.Client) bool {
	if c != nil && c.Metadata["isInternal"] != nil {
		if isInternal, ok := c.Metadata["isInternal"].(bool); ok {
			return isInternal
		}
	}
	return false
}

func ConsentPostHandler(hydraClient *client.OryHydra, templateProvider TemplateProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		challenge := r.FormValue("challenge")

		if challenge == "" {
			w.Header().Set("X-Status-Reason", "no challenge")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		submitValue := r.FormValue("submit")

		if submitValue == "Deny access" {
			// Looks like the consent request was denied by the user

			rejectConsentRequestParams := hydraAdmin.NewRejectConsentRequestParams()
			var requestBody models.RequestDeniedError
			rejectConsentRequestParams.ConsentChallenge = challenge
			rejectConsentRequestParams.Body = &requestBody
			rejectConsentRequestParams.Body.Name = "access_denied"
			rejectConsentRequestParams.Body.Description = "The resource owner denied the request"

			rejectConsentResponse, err := hydraClient.Admin.RejectConsentRequest(rejectConsentRequestParams)

			if err != nil {
				w.Header().Set("X-Status-Reason", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// All we need to do now is to redirect the browser back to hydra!
			http.Redirect(w, r, rejectConsentResponse.Payload.RedirectTo, http.StatusTemporaryRedirect)
		} else {
			grantScope := r.PostForm["grant_scope"]

			rememberValue := r.PostForm.Get("remember")
			remember := rememberValue == "1"

			// Seems like the user authenticated! Let's tell hydra...
			getConsentRequestParams := hydraAdmin.NewGetConsentRequestParams()
			getConsentRequestParams.ConsentChallenge = challenge
			getConsentResponse, err := hydraClient.Admin.GetConsentRequest(getConsentRequestParams)

			if err != nil {
				w.Header().Set("X-Status-Reason", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// You can apply logic here, for example grant another scope, or do whatever...
			// ...

			// Now it's time to grant the consent request. You could also deny the request if something went terribly wrong

			acceptConsentRequestParams := hydraAdmin.NewAcceptConsentRequestParams()
			acceptConsentRequestParams.ConsentChallenge = challenge
			var handledConsentRequest hydraModels.HandledConsentRequest
			acceptConsentRequestParams.Body = &handledConsentRequest

			// We can grant all scopes that have been requested - hydra already checked for us that no additional scopes
			// are requested accidentally.
			acceptConsentRequestParams.Body.GrantedScope = grantScope

			// Grant the roles scope to every client, so they can get the users roles over the backchannel
			// ... is roles scope already in requestBody.GrantScope
			// the client app will be able to request /api/v0/roles. This endpoint will only return
			// data if a valid appHash is provided in the querystring. Returned information
			// is strictly limited to the app requesting the roles. You only receive your roles in the
			// requesting app

			for i := 0; i < len(AutoScopes); i++ {
				if helpers.SliceIndex(len(acceptConsentRequestParams.Body.GrantedScope), func(j int) bool {
					return acceptConsentRequestParams.Body.GrantedScope[j] == AutoScopes[i]
				}) == -1 {
					acceptConsentRequestParams.Body.GrantedScope = append(acceptConsentRequestParams.Body.GrantedScope, AutoScopes[i])
				}
			}

			// ORY Hydra checks if requested audiences are allowed by the client, so we can simply echo this.
			acceptConsentRequestParams.Body.GrantedAudience = getConsentResponse.Payload.RequestedAudience

			// This tells hydra to remember this consent request and allow the same client to request the same
			// scopes from the same user, without showing the UI, in the future.
			acceptConsentRequestParams.Body.Remember = remember

			// When this "remember" session expires, in seconds. Set this to 0 so it will never expire.
			acceptConsentRequestParams.Body.RememberFor = 3600

			// This data will be available when introspecting the token. Try to avoid sensitive information here,
			// unless you limit who can introspect tokens.
			// access_token: { foo: 'bar' },

			// This data will be available in the ID token.
			// id_token: { baz: 'bar' },
			session := new(models.ConsentRequestSessionData)
			//acceptConsentRequestParams.Body.Session = session

			login := getConsentResponse.Payload.Subject
			clientID := getConsentResponse.Payload.Client.ClientID

			log.Info().Str("login", login).Str("client", clientID).Msg("consent approved")

			permissions := [2]string{"user.edit", "user.read"}

			session.AccessToken = map[string]interface{}{
				"permissions": permissions,
			}

			session.IDToken = map[string]interface{}{
				"permissions": permissions,
			}

			acceptConsentRequestParams.Body.Session = session

			acceptConsentResponse, err := hydraClient.Admin.AcceptConsentRequest(acceptConsentRequestParams)

			if err != nil {
				w.Header().Set("X-Status-Reason", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// All we need to do now is to redirect the user back to hydra!
			http.Redirect(w, r, acceptConsentResponse.Payload.RedirectTo, http.StatusTemporaryRedirect)

		}
	}
}
