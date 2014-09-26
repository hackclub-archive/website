package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/hackedu/backend/database"
	"github.com/hackedu/backend/handler/v1"
	"github.com/hackedu/backend/middleware"
)

func httpLog(handler http.Handler,
	middleware middleware.MiddlewareProcessor) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		if middleware.Process(w, r) {
			handler.ServeHTTP(w, r)
		}
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

	m := middleware.MiddlewareProcessor{}

	m.Register(&middleware.CORS{})

	r := mux.NewRouter()

	r.Handle("/users", v1.Handler(v1.CreateUser)).Methods("POST")
	r.Handle("/users/authenticate",
		v1.Handler(v1.Authenticate)).Methods("POST")
	r.Handle("/users/{id}", v1.Handler(v1.GetUser)).Methods("GET")
	r.Handle("/users/{id}/clubs",
		v1.Handler(v1.GetAllClubsForUser)).Methods("GET")

	r.Handle("/schools",
		v1.Handler(v1.CreateSchool)).Methods("POST")
	r.Handle("/schools", v1.Handler(v1.GetSchools)).Methods("GET")
	r.Handle("/schools/{id}", v1.Handler(v1.GetSchool)).Methods("GET")

	r.Handle("/clubs", v1.Handler(v1.CreateClub)).Methods("POST")
	r.Handle("/clubs", v1.Handler(v1.GetAllClubs)).Methods("GET")
	r.Handle("/clubs/{id}", v1.Handler(v1.GetClub)).Methods("GET")
	r.Handle("/clubs/{id}/members",
		v1.Handler(v1.CreateClubMember)).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, httpLog(http.DefaultServeMux, m)))
}
