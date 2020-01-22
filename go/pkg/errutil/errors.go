package errutil

import (
	stdError "errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type FieldError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func New(msg string) error {
	return errors.New(msg)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

func Is(err, target error) bool {
	return stdError.Is(err, target)
}

func Cause(err error) error {
	return errors.Cause(err)
}

func RenderError(w http.ResponseWriter, r *http.Request, err error) {

	var ce *clientError
	if errors.As(err, &ce) {
		if ce.Code == 0 {
			log.Error().Err(err).Msg("error code was zero for client error. This will be sent as internal error")
			render.Status(r, http.StatusInternalServerError)
			return
		}
		render.Status(r, ce.Code)
		render.JSON(w, r, ce)
		return
	}

	log.Error().Err(err).Send()
	render.Status(r, 500)
	render.PlainText(w, r, http.StatusText(http.StatusInternalServerError))
}
