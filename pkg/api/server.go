package api

import (
	"net/http"
	"time"
)

type Server struct {
	AuthHandler Handler
	ReportHandler Handler
	TransactionHandler Handler
}

func (s *Server) Init() {
	// handler object
	transaction := http.NewServeMux()
	transaction.HandleFunc("/transaction", s.TransactionHandler.Handle)

	// router object
	mux := http.NewServeMux()
	mux.Handle("/trans", http.StripPrefix("/handle", transaction))

	// server object
	hs := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
}
