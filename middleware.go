package main

import "net/http"

func allowCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept")
}
