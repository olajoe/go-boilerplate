package errs

import (
	"go-boilerplate/domain"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

func (e ResponseError) Error() string {
	return e.Message
}

func NewResponseError(message string) error {
	return ResponseError{
		Message: message,
	}
}

func NewBadRequestError(message string) error {
	return NewResponseError(message)
}

func NewUnauthorizedError(message string) error {
	return NewResponseError(message)
}

func NewForbiddenError(message string) error {
	return NewResponseError(message)
}

func NewNotFoundError(message string) error {
	return NewResponseError(message)
}

func NewConflictError(message string) error {
	return NewResponseError(message)
}

func NewUnprocessableEntityError(message string) error {
	return NewResponseError(message)
}

func NewInternalServerError(message string) error {
	return NewResponseError(message)
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
