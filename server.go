package main

import (
	"net/http"
	"time"

	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/hackedu/backend/model"
	"github.com/hackedu/backend/route"
	"github.com/martini-contrib/binding"
	"github.com/zachlatta/cors"
)

func main() {
	m := martini.Classic()

	defer Dbm.Db.Close()

	m.Use(cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		MaxAge:          5 * time.Minute,
	}))
	m.MapTo(Dbm, (*gorp.SqlExecutor)(nil))

	m.Get("/schools", route.GetSchools)

	m.Group("/users", func(r martini.Router) {
		r.Post("", binding.Bind(model.User{}), route.AddUser)
		r.Get("/:id", route.GetUser)
	})

	// OPTIONS catchall for CORS.
	m.Options("/**", func() int {
		return http.StatusOK
	})

	m.Run()
}
