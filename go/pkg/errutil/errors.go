package errutil

import (
	"errors"
	"fmt"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog/log"
	"net/http"
)

var NotFound = eris.New("not found")
var Unauthorized = eris.New("unauthorized")
var BadRequest = eris.New("bad request")

type FieldError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type ValidationError struct {
	FieldErrors []FieldError `json:"fieldErrors,omitempty"`
}

func New(msg string) error {
	return eris.New(msg)
}

func Errorf(format string, args ...interface{}) error {
	return eris.Errorf(format, args...)
}

func Wrap(err error, msg string) error {
	return eris.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return eris.Wrapf(err, format, args...)
}

func Is(err, target error) bool {
	return eris.Is(err, target)
}

func Cause(err error) error {
	return eris.Cause(err)
}

func NewUnauthorized(msg string) error {
	return eris.Wrap(Unauthorized, msg)
}
func NewBadRequest(msg string) error {
	return eris.Wrap(BadRequest, msg)
}

func NewFieldError(fieldErrors ...FieldError) error {
	return eris.Wrap(BadRequest, fmt.Sprintf("%v", fieldErrors))
}

func RenderError(w http.ResponseWriter, r *http.Request, err error) {

	if e, ok := err.(validation.Errors); ok {
		var result []FieldError
		for k, v := range e {
			result = append(result, FieldError{Field: k, Message: v.Error()})
		}
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, result)
		return
	}

	if errors.Is(err, Unauthorized) {
		render.Status(r, http.StatusUnauthorized)
		render.PlainText(w, r, err.Error())
		return
	}

	log.Error().Err(err).Send()
	render.Status(r, 500)
	render.PlainText(w, r, err.Error())
	return
}
