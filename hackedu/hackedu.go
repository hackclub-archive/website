package hackedu

import (
	"io"
	"net/http"
	"github.com/crhym3/go-endpoints/endpoints"

	"appengine"
)

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Internal Server Error")
	c.Errorf("%v", err)
}

func init() {
	schoolService := &SchoolService{}
	api, err := endpoints.RegisterService(
		schoolService,
		"school",
		"v1",
		"Schools API",
		true,
	)
	if err != nil {
		panic(err)
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "schools.list", "GET", "schools", "List schools."

	http.HandleFunc("/v1/apply", applyHandler)

	endpoints.HandleHttp()
}

func schoolsHandler(w http.ResponseWriter, r *http.Request) {
	middleware(w, r)
}

func applyHandler(w http.ResponseWriter, r *http.Request) {
	middleware(w, r)
	if r.Method == "POST" {
		Apply(w, r)
	}
}
