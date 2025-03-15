package transaction

import (
	"net/http"

	"github.com/GeorgeKuzora/go_api_gateway/pkg/api"
)

type Client struct {
	Url string
	client *http.Client
}

func (c Client) Request(api.Transaction) error {

	return nil
}
