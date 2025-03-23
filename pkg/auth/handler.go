package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/GeorgeKuzora/go_api_gateway/pkg/api"
)

type Register struct {
	Client Client
	Url string
}

func (th Register) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		th.Post(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (th Register) Post(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		http.Error(
			w,
			fmt.Sprintf("%s Content-Type is not allowed", headerContentType),
			http.StatusUnsupportedMediaType,
		)
		return
	}
	var uc api.UserCredentials
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&uc)
	if err != nil {
		if errors.As(err, *unmarshalErr) {
			http.Error(w, "Bad Request. Wrong type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
			} else {
			http.Error(w, "Bad Request. "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	token, clientErr := th.Client.Register(uc, "register/")
	if clientErr != nil {
		http.Error(w, fmt.Sprintf("Bad Response. %s", clientErr.Error()), clientErr.StatusCode())
	}
	encodedToken, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "Bad Response. Encoding Error", http.StatusInternalServerError)
	}
	w.Write(encodedToken)
}

type Login struct {
	Client Client
	Url string
}

func (l Login) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
			l.Post(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
}

// request body
// {
	// UserCredentials,
	// token,
// }
func (l Login) Post(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		http.Error(
			w,
			fmt.Sprintf("%s Content-Type is not allowed", headerContentType),
			http.StatusUnsupportedMediaType,
		)
	}
	lr := struct {
		UserCredentials api.UserCredentials  `json:"userCredentials"`
		Token api.Token  `json:"token"`
	}{
		UserCredentials: api.UserCredentials{},
		Token: api.Token{},
	}
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&lr)
	if err != nil{
		if errors.As(err, *unmarshalErr) {
			http.Error(w , "Bad Request. Wrong type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			http.Error(w, "Bad Request. "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	token, clientErr := l.Client.Login(lr.UserCredentials, lr.Token, "/login")
	if clientErr != nil {
		http.Error(w, fmt.Sprintf("Bad Response. %s", clientErr.Error()), clientErr.StatusCode())
	}
	encodedToken, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "Bad Response. Encoding Error", http.StatusInternalServerError)
	}
	w.Write(encodedToken)
}

type Verify struct {
	Client Client
	Url string
}

func (v Verify) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
			v.Post(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
}

func (v Verify) Post(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "multipart/form-data" {
		http.Error(
			w,
			fmt.Sprintf("%s Content-Type is not allowed", headerContentType),
			http.StatusUnsupportedMediaType,
		)
	}
}
