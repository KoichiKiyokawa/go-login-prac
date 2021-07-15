package controller

import (
	"encoding/json"
	"errors"
	"net/http"
)

type loginBody struct {
	email    string
	password string
}
type loginResponse struct {
	message string
}

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var body loginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if body.email == "hoge@example.com" && body.password == "hogehoge" {
		respondJson(w, loginResponse{message: "ok"})
	}
	http.Error(w, errors.New("wrong").Error(), http.StatusUnauthorized)
}

func AuthIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
}
