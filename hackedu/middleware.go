package hackedu

import "net/http"

func middleware(w http.ResponseWriter, r *http.Request) {
	allowCORS(w, r)
}

func allowCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
