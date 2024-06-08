package errs

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) error {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorizedError(message string) error {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewNotFoundError(message string) error {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewConflictError(message string) error {
	return &AppError{
		Code:    http.StatusConflict,
		Message: message,
	}
}
