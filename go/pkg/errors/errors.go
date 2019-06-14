package errors

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Error struct {
	ID      uuid.UUID   `json:"id,omitempty"`
	Err     error       `json:"-"`
	Code    int         `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("ID:%v, Code:%d, Details:%v, Cause: %v", e.ID, e.Code, e.Details, e.Err)
}

var causeValidation = "Validation"

type errorDetails struct {
	Cause       string       `json:"cause,omitempty"`
	FieldErrors []FieldError `json:"fieldErrors,omitempty"`
}

type FieldError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func wrap(err error, msg string) error {
	if err != nil {
		return errors.Wrap(err, msg)
	} else {
		return errors.New(msg)
	}
}

func ToError(err error, msg string) Error {
	if ce, ok := err.(Error); ok {
		return ce
	}
	return NewInternal(err, msg)
}

func NewInternal(err error, msg string) Error {
	return Error{ID: uuid.New(), Err: wrap(err, msg), Code: http.StatusInternalServerError}
}

func NewBadRequest(details interface{}) Error {
	if reason, ok := details.(string); ok {
		return Error{ID: uuid.New(), Err: nil, Code: http.StatusBadRequest, Details: errorDetails{Cause: reason}}
	}
	return Error{ID: uuid.New(), Err: nil, Code: http.StatusBadRequest, Details: details}
}

func NewUnauthorized(details interface{}) Error {
	if reason, ok := details.(string); ok {
		return Error{ID: uuid.New(), Err: nil, Code: http.StatusUnauthorized, Details: errorDetails{Cause: reason}}
	}
	return Error{ID: uuid.New(), Err: nil, Code: http.StatusUnauthorized, Details: details}
}

func WithFieldErrors(fieldErrors []FieldError) Error {
	br := NewBadRequest(errorDetails{Cause: causeValidation, FieldErrors: fieldErrors})
	return br
}

func WithFieldError(field string, message string) Error {
	br := NewBadRequest(errorDetails{Cause: causeValidation, FieldErrors: []FieldError{{Field: field, Message: message}}})
	return br
}

func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	log.WithField("error", err).Errorf("Encountered error %v", err)
	var appErr Error
	if ce, ok := err.(Error); ok {
		appErr = ce
	} else if e, ok := err.(validation.InternalError); ok {
		appErr = NewInternal(e.InternalError(), "error during validation")
	} else if e, ok := err.(validation.Errors); ok {
		var result []FieldError
		for k, v := range e {
			result = append(result, FieldError{Field: k, Message: v.Error()})
		}
		appErr = WithFieldErrors(result)
	} else {
		appErr = NewInternal(err, "unknown internal error")
	}

	render.Status(r, appErr.Code)
	render.JSON(w, r, appErr)
	return
}
