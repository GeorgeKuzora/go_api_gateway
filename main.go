package main

import (
	"github.com/GeorgeKuzora/go_api_gateway/pkg/api"
	"github.com/GeorgeKuzora/go_api_gateway/pkg/clients"
	handlers "github.com/GeorgeKuzora/go_api_gateway/pkg/handlers"
)

func main() {
	ah := handlers.Auth{
		Client: clients.Auth{},
	}
	rh := handlers.Report{
		Client: clients.Report{},
	}
	th := handlers.Transaction{
		Client: clients.Transaction{},
	}
	s := api.Server{
		AuthHandler: ah,
		ReportHandler: rh,
		TransactionHandler: th,
	}
	s.Init()
}
