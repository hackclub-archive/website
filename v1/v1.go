package v1

import (
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/hackedu/backend/v1/model"
	"github.com/hackedu/backend/v1/route"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/cors"
)

func Setup(m *martini.ClassicMartini) {
	// TODO: Only apply middleware on /v1/** routes
	m.Use(cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		MaxAge:          5 * time.Minute,
	}))
	m.MapTo(Dbm, (*gorp.SqlExecutor)(nil))

	m.Get("/v1/schools", route.GetSchools)

	m.Post("/v1/users", binding.Bind(model.User{}), route.AddUser)

	// OPTIONS catchall for CORS.
	m.Options("/**", func() int {
		return http.StatusOK
	})
}
