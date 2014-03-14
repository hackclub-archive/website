package v1

import (
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/hackedu/backend/v1/model"
	"github.com/hackedu/backend/v1/route"
	"github.com/martini-contrib/binding"
)

func Setup(m *martini.ClassicMartini) {
	// TODO: Only apply middleware on /v1/** routes
	m.Use(allowCORS)
	m.MapTo(Dbm, (*gorp.SqlExecutor)(nil))

	m.Get("/v1/schools", route.GetSchools)

	m.Post("/v1/users", binding.Bind(model.User{}), route.AddUser)
}
