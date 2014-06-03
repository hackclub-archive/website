package middleware

import "net/http"

type MiddlewareProcessor struct {
	Middleware []MiddlewareHandler
}

func (mp *MiddlewareProcessor) Register(m MiddlewareHandler) {
	mp.Middleware = append(mp.Middleware, m)
}

func (mp *MiddlewareProcessor) Process(w http.ResponseWriter,
	r *http.Request) bool {
	for _, middleware := range mp.Middleware {
		if !middleware.Handle(w, r) {
			return false
		}
	}
	return true
}
