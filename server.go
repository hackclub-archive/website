package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/hackedu/backend/database"
	"github.com/hackedu/backend/handler"
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

	r.Handle("/users", handler.AppHandler(handler.CreateUser)).Methods("POST")
	r.Handle("/users/authenticate",
		handler.AppHandler(handler.Authenticate)).Methods("POST")
	r.Handle("/users/{id}", handler.AppHandler(handler.GetUser)).Methods("GET")
	r.Handle("/users/{id}/clubs",
		handler.AppHandler(handler.GetAllClubsForUser)).Methods("GET")

	r.Handle("/schools",
		handler.AppHandler(handler.CreateSchool)).Methods("POST")
	r.Handle("/schools", handler.AppHandler(handler.GetSchools)).Methods("GET")
	r.Handle("/schools/{id}",
		handler.AppHandler(handler.GetSchool)).Methods("GET")

	r.Handle("/clubs", handler.AppHandler(handler.CreateClub)).Methods("POST")
	r.Handle("/clubs", handler.AppHandler(handler.GetAllClubs)).Methods("GET")
	r.Handle("/clubs/{id}", handler.AppHandler(handler.GetClub)).Methods("GET")
	r.Handle("/clubs/{id}/members",
		handler.AppHandler(handler.CreateClubMember)).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, httpLog(http.DefaultServeMux, m)))
}
