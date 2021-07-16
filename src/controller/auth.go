package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-login-prac/src/service"
)

type AuthController struct {
	AuthService service.IAuthService
}

type loginBody struct {
	Email    string `json:"email"` // ここを emailと小文字にすると、decode/encodeがうまく行かない。違うパッケージ内で処理されるため
	Password string `json:"password"`
}

func (c AuthController) AuthLogin(w http.ResponseWriter, r *http.Request) {
	var body loginBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := c.AuthService.ValidateUser(body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// TODO: セッションに書き込む処理

	respondJson(w, user)
}

func (AuthController) AuthIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
}
