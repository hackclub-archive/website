package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type AppError struct {
	Error   error  `json:"-"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type AppHandler func(http.ResponseWriter, *http.Request) *AppError

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *AppError, not os.Error
		log.Println(e.Error)
		renderJSON(w, e, e.Code)
	}
}

func renderJSON(w http.ResponseWriter, data interface{}, code int) *AppError {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return &AppError{err, "error encoding json",
			http.StatusInternalServerError}
	}
	return nil
}
