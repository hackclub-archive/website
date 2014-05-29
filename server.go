package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/hackedu/backend/database"
	"github.com/hackedu/backend/handler"
)

func httpLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
		log.Printf("Completed in %s", time.Now().Sub(start).String())
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := database.Init("postgres",
		os.ExpandEnv("postgres://docker:docker@$DB_1_PORT_5432_TCP_ADDR/docker"))
	if err != nil {
		panic(err)
	}
	defer database.Close()

	r := mux.NewRouter()

	r.Handle("/users", handler.AppHandler(handler.CreateUser)).Methods("POST")
	r.Handle("/users/authenticate",
		handler.AppHandler(handler.Authenticate)).Methods("POST")
	r.Handle("/users/me",
		handler.AppHandler(handler.GetCurrentUser)).Methods("GET")
	r.Handle("/users/{id}", handler.AppHandler(handler.GetUser)).Methods("GET")

	r.Handle("/schools",
		handler.AppHandler(handler.CreateSchool)).Methods("POST")
	r.Handle("/schools", handler.AppHandler(handler.GetSchools)).Methods("GET")
	r.Handle("/schools/{id}",
		handler.AppHandler(handler.GetSchool)).Methods("GET")

	r.Handle("/clubs", handler.AppHandler(handler.CreateClub)).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, httpLog(http.DefaultServeMux)))
}
