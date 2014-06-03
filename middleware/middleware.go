package middleware

import "net/http"

type MiddlewareHandler interface {
	Handle(http.ResponseWriter, *http.Request) bool
}
