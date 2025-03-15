package main

import (
	"github.com/GeorgeKuzora/go_api_gateway/pkg/api"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	ah := handlers.Register{
		Client: clients.Auth{},
	}
	rh := handlers.Report{
		Client: clients.Report{},
	}
	th := handlers.Transaction{
		Client: clients.Transaction{},
	}
	s := api.Server{
		AuthHandler:        ah,
		ReportHandler:      rh,
		TransactionHandler: th,
	}
	s.Init()
}

