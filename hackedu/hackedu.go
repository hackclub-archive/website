package hackedu

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/hello_world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
}
