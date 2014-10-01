package v1

import (
	"errors"
	"net/http"

	"github.com/hackedu/hackedu/httputil"
)

func ErrCreatingModel(err error) *httputil.HTTPError {
	return &httputil.HTTPError{http.StatusBadRequest, err}
}

func ErrForbidden() *httputil.HTTPError {
	return &httputil.HTTPError{http.StatusForbidden, errors.New("forbidden")}
}

func ErrInvalidID() *httputil.HTTPError {
	return &httputil.HTTPError{http.StatusForbidden, errors.New("invalid id")}
}

func ErrNotAuthorized() *httputil.HTTPError {
	return &httputil.HTTPError{http.StatusUnauthorized,
		errors.New("not authorized")}
}

func ErrNotFound() *httputil.HTTPError {
	return &httputil.HTTPError{http.StatusNotFound, errors.New("not found")}
}
