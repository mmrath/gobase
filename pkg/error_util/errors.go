package error_util

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type errorCode int

const ErrorCodeBadRequest errorCode = http.StatusBadRequest
const ErrorCodeInternal errorCode = http.StatusInternalServerError
const ErrorCodeUnauthorized errorCode = http.StatusUnauthorized

type Error struct {
	ID      uuid.UUID   `json:"id,omitempty"`
	Err     error       `json:"-"`
	Code    errorCode   `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("ID:%v, Code:%d, Details:%v, Cause: %v", e.ID, e.Code, e.Details, e.Err)
}

func (e Error) IsBadRequest() bool {
	return e.Code == ErrorCodeBadRequest
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
		return fmt.Errorf(msg, err)
	} else {
		return errors.New(msg)
	}
}
func GetErrorID(err error) string {
	if ce, ok := err.(Error); ok {
		return ce.ID.String()
	}
	return ""
}
func ToError(err error, msg string) Error {
	if ce, ok := err.(Error); ok {
		return ce
	}
	return NewInternal(err, msg)
}

func NewInternal(err error, msg string) Error {
	if ce, ok := err.(Error); ok {
		if ce.Code == ErrorCodeInternal {
			return ce
		}
	}
	return Error{ID: uuid.New(), Err: wrap(err, msg), Code: ErrorCodeInternal}
}

func NewBadRequest(details interface{}) Error {
	if reason, ok := details.(string); ok {
		return Error{ID: uuid.New(), Err: nil, Code: ErrorCodeBadRequest, Details: errorDetails{Cause: reason}}
	}
	return Error{ID: uuid.New(), Err: nil, Code: ErrorCodeBadRequest, Details: details}
}

func NewUnauthorized(details interface{}) Error {
	if reason, ok := details.(string); ok {
		return Error{ID: uuid.New(), Err: nil, Code: ErrorCodeUnauthorized, Details: errorDetails{Cause: reason}}
	}
	return Error{ID: uuid.New(), Err: nil, Code: ErrorCodeUnauthorized, Details: details}
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
	log.Error().Err(err).Send()
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

	render.Status(r, int(appErr.Code))
	render.JSON(w, r, appErr)
	return
}
