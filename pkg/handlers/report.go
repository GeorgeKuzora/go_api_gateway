package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/GeorgeKuzora/go_api_gateway/pkg/api"
)

type Report struct {
	Client api.Client
}

func (rh Report) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		rh.Post(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (rh Report) Post(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != string(applicationJson) {
		http.Error(
			w,
			fmt.Sprint("%s Content-Type is not allowed", headerContentType),
		http.StatusUnsupportedMediaType,
		)
		return
	}
	var reportRequest api.ReportRequest
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reportRequest)
	if err != nil {
		if errors.As(err, *unmarshalErr) {
			http.Error(w, "Bad Request. Wrong type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			http.Error(w, "Bad Request. "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	report, clientErr := rh.Client.Post(reportRequest)
	if err != nil {
		http.Error(w, "Request failed. Reason: "+clientErr.Error(), clientErr.StatusCode())
		return
	}
	out, err := json.Marshal(report)
	if err != nil {
		http.Error(w, "Request failed. Reason: "+err.Error(), http.StatusServiceUnavailable)
	}
	w.Header().Set(string(contentType), string(applicationJson))
	w.Write(out)
}
