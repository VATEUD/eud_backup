package handlers

import "net/http"

// Stats returns the cached information about the last backup
func Stats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test"))
}
