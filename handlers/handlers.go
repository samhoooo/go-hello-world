package handlers

import (
	"fmt"
	"net/http"
)

// API Handler for the GET /api/hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "hello world")
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
