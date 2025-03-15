package report

import (
	"net/http"

	"github.com/GeorgeKuzora/go_api_gateway/pkg/api"
)

type Client struct {
	client *http.Client
}

func (c Client) Request(url string, request api.ReportRequest) (api.Report, error) {

	return nil, nil
}