package handler

import (
	"errors"
	"net/http"
)

func ErrCreatingModel(err error) *AppError {
	return &AppError{err, err.Error(), http.StatusBadRequest}
}

func ErrInvalidID(err error) *AppError {
	return &AppError{err, "invalid id", http.StatusBadRequest}
}

func ErrNotAuthorized() *AppError {
	err := errors.New("not authorized")
	return &AppError{err, err.Error(), http.StatusUnauthorized}
}

func ErrNotFound(err error) *AppError {
	return &AppError{err, "not found", http.StatusNotFound}
}

func ErrDatabase(err error) *AppError {
	return &AppError{err, "internal database error",
		http.StatusInternalServerError}
}

func ErrUnmarshalling(err error) *AppError {
	return &AppError{err, err.Error(), http.StatusBadRequest}
}
