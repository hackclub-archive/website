package main

import (
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/hackedu/backend/v1"
	"github.com/hackedu/backend/v1/database"
	"github.com/hackedu/backend/v1/middleware"
)

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

	n := negroni.New()

	n.Use(negroni.HandlerFunc(middleware.CORS))

	n.UseHandler(r)
	n.Run(":" + port)
}
