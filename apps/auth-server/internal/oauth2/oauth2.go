package oauth2

import (
	"net/http"
)

func RegisterHandlers() {
	// Set up oauth2 endpoints. You could also use gorilla/mux or any other router.
	http.HandleFunc("/oauth2/auth", authEndpoint)
	http.HandleFunc("/oauth2/token", tokenEndpoint)

	// revoke tokens
	http.HandleFunc("/oauth2/revoke", revokeEndpoint)
	http.HandleFunc("/oauth2/introspect", introspectionEndpoint)
}

func AuthHandler() http.HandlerFunc {
	return authEndpoint
}

func TokenHandler() http.HandlerFunc {
	return tokenEndpoint
}

func RevokeHandler() http.HandlerFunc {
	return revokeEndpoint
}

func IntrospectionHandler() http.HandlerFunc {
	return introspectionEndpoint
}
