package api

import "net/http"


type Handler interface {
	Handle(http.ResponseWriter, *http.Request)
}
