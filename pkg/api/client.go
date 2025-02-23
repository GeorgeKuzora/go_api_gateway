package api

type Client interface {
	Post(Request) (Response, ClientError)
}

type ClientError interface {
	Error() string
	StatusCode() int
}
