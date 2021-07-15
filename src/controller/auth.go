package controller

import (
	"encoding/json"
	"errors"
	"net/http"
)

type loginBody struct {
	Email    string `json:"email"` // ここを emailと小文字にすると、decode/encodeがうまく行かない。違うパッケージ内で処理されるため
	Password string `json:"password"`
}
type loginResponse struct {
	Message string `json:"message"`
}

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var body loginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if body.Email == "hoge@example.com" && body.Password == "hogehoge" {
		respondJson(w, loginResponse{Message: "ok"})
		return
	}
	http.Error(w, errors.New("wrong").Error(), http.StatusUnauthorized)
}

func AuthIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
}
