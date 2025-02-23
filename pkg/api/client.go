package api

type Client interface {
	Post(Request) ClientError
}

type ClientError interface {
	Error() string
	StatusCode() int
}
