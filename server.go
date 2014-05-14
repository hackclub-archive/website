package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := mux.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
}
