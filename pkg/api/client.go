package api

type Client interface {}

type ClientError interface {
	Error() string
	StatusCode() int
}
