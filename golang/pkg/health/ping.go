package health

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func PingHandlerFunc(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Error().Err(err).Msg("failed to reply to ping")
	}
}
