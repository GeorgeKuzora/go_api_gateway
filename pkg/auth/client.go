package auth

import "github.com/GeorgeKuzora/go_api_gateway/pkg/api"



type Client struct {
}

func (c Client) Register(uc api.UserCredentials, url string) (api.Token, api.ClientError) {
	return nil, nil
} 
func (c Client) Login(uc api.UserCredentials, token api.Token, url string) (api.Token, api.ClientError) {
	return nil, nil
} 