package auth

import "github.com/GeorgeKuzora/go_api_gateway/pkg/api"



type Client struct {
}

func (c Client) Register(uc api.UserCredentials, url string) (api.Token, api.ClientError) {
	return nil, nil
} 