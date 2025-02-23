package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/GeorgeKuzora/go_api_gateway/pkg/api"
)

type Transaction struct {
	Client api.Client
}

func (t Transaction) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		t.Post(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (t Transaction) Post(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != string(applicationJson) {
		http.Error(
			w,
			fmt.Sprint("%s Content-Type is not allowed", headerContentType),
		http.StatusUnsupportedMediaType,
		)
		return
	}
	var transactionData api.Transaction
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&transactionData)
	if err != nil {
		if errors.As(err, *unmarshalErr) {
			http.Error(w, "Bad Request. Wrong type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			http.Error(w, "Bad Request. "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	_, clientErr := t.Client.Post(transactionData)
	if err != nil {
		http.Error(w, "Request failed. Reason: "+clientErr.Error(), clientErr.StatusCode())
		return
	}
}
