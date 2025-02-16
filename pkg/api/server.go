package api

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// handler object
	transaction := http.NewServeMux()
	transaction.HandleFunc("/trans", processTransaction)

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
}
