package api

type Client interface {
	Request(Request) (Response, error)
}
