package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/GeorgeKuzora/go_api_gateway/pkg/api"
)

func processTransaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "applicaiton/json" {
		http.Error(
			w,
			fmt.Sprint("%s Content-Type is not allowed", headerContentType),
			http.StatusUnsupportedMediaType,
		)
		return
	}
	var t api.Transaction
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		if errors.As(err, *unmarshalErr) {
			http.Error(w, "Bad Request. Wrong type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			http.Error(w, "Bad Request. "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	
}
