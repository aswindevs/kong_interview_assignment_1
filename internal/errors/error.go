package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code     string
	Message  string
	HTTPCode int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:     http.StatusText(http.StatusNotFound),
		HTTPCode: http.StatusNotFound,
		Message:  message,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:     http.StatusText(http.StatusBadRequest),
		HTTPCode: http.StatusBadRequest,
		Message:  message,
	}
}

func NewInternalError(message string) *AppError {
	return &AppError{
		Code:     http.StatusText(http.StatusInternalServerError),
		HTTPCode: http.StatusInternalServerError,
		Message:  message,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Code:     http.StatusText(http.StatusUnauthorized),
		HTTPCode: http.StatusUnauthorized,
		Message:  message,
	}
}

func NewAlreadyExistsError(message string) *AppError {
	return &AppError{
		Code:     http.StatusText(http.StatusConflict),
		HTTPCode: http.StatusConflict,
		Message:  message,
	}
}