package auth

import (
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

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
			fmt.Sprint("%s Content-Type is not allowed", headerContentType),
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(uc.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Bad Request. Invalid password", http.StatusBadRequest)
	}
	
	uc.Password = string(hashedPassword)
	token, clientErr := th.Client.Register(uc, "register/")
	if clientErr != nil {
		http.Error(w, fmt.Sprintf("Bad Response. %s", err.Error()), clientErr.StatusCode())
	}
	encodedToken, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "Bad Response. Encoding Error", http.StatusInternalServerError)
	}
	w.Write(encodedToken)
}
