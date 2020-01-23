package errutil

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Error type that is OK to be displayed to user (Web, Mobile)
type clientError struct {
	FieldErrors []FieldError `json:"fieldErrors,omitempty"`
	Errors      []string     `json:"errors,omitempty"`
	Code        int
}

func (v clientError) Error() string {
	var s []string

	if len(v.FieldErrors) > 0 {
		for _, v := range v.FieldErrors {
			s = append(s, fmt.Sprintf("%s: %s", v.Field, v.Message))
		}
	}

	s = append(s, v.Errors...)

	return fmt.Sprintf("error:[ %s ]", strings.Join(s, ","))
}

func NewBadRequest(msg string) error {
	err := &clientError{Errors: []string{msg}, Code: http.StatusBadRequest}
	return errors.WithStack(err)
}

func NewUnauthorized(msg string) error {
	err := &clientError{
		Errors: []string{msg},
		Code:   http.StatusUnauthorized,
	}

	return errors.WithStack(err)
}

func NewFieldErrors(fieldErrors map[string]string) error {
	var result []FieldError
	for k, v := range fieldErrors {
		result = append(result, FieldError{Field: k, Message: v})
	}

	return errors.WithStack(&clientError{
		FieldErrors: result,
		Errors:      nil,
		Code:        http.StatusBadRequest,
	})
}

func NewFieldError(field, msg string) error {
	err := &clientError{
		FieldErrors: []FieldError{{
			Field:   field,
			Message: msg,
		}},
		Errors: []string{msg},
		Code:   http.StatusBadRequest,
	}

	return errors.WithStack(err)
}
