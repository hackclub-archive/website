package hackedu

import (
	"io"
	"net/http"

	"appengine"
)

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Internal Server Error")
	c.Errorf("%v", err)
}

func init() {
	http.HandleFunc("/v1/schools", schoolsHandler)
	http.HandleFunc("/v1/apply", applyHandler)
}

func schoolsHandler(w http.ResponseWriter, r *http.Request) {
	middleware(w, r)
	Schools(w, r)
}

func applyHandler(w http.ResponseWriter, r *http.Request) {
	middleware(w, r)
	if r.Method == "POST" {
		Apply(w, r)
	}
}
