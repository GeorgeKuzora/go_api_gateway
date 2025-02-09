package api

import (
	"net/http"
	"time"
)

func main() {
	// handler object
	transaction := http.NewServeMux()
	transaction.HandleFunc("/trans", proccessTransaction)

	// router object
	mux := http.NewServeMux()
	mux.Handle("/trans", http.StripPrefix("/proc", transaction))

	// server object
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	// hendler methods

	// handler object
	// hendler methods
}
func proccessTransaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "applicaiton/json" {
		errorResponse
	}
}
