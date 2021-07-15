package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-login-prac/src/repository"
)

type AuthController struct {
	AuthRepo repository.IAuthRepository
}

type loginBody struct {
	Email    string `json:"email"` // ここを emailと小文字にすると、decode/encodeがうまく行かない。違うパッケージ内で処理されるため
	Password string `json:"password"`
}

func (c AuthController) AuthLogin(w http.ResponseWriter, r *http.Request) {
	var body loginBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := c.AuthRepo.FindByEmail(body.Email)
	if err != nil {
		http.Error(w, errors.New("something wrong").Error(), http.StatusUnauthorized)
		return
	}

	if user.Password == body.Password {
		respondJson(w, user)
		return
	}
	http.Error(w, errors.New("wrong").Error(), http.StatusUnauthorized)
}

func (AuthController) AuthIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
}
