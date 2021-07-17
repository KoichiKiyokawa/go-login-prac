package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-login-prac/service"

	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) AuthController {
	return AuthController{authService: authService}
}

type loginBody struct {
	Email    string `json:"email"` // ここを emailと小文字にすると、decode/encodeがうまく行かない。違うパッケージ内で処理されるため
	Password string `json:"password"`
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func (c AuthController) AuthLogin(w http.ResponseWriter, r *http.Request) {
	var body loginBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusInternalServerError)
		return
	}

	user, err := c.authService.ValidateUser(body.Email, body.Password)
	if err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusUnauthorized)
		return
	}

	// write to session
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, errors.New("error in write session").Error(), http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	session.Values["currentUser"] = json
	if err := session.Save(r, w); err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusInternalServerError)
		return
	}

	respondJson(w, user)
}

func (c AuthController) AuthCheck(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusUnauthorized)
		return
	}

	if session.Values["currentUser"] == nil {
		http.Error(w, errors.New("not logged in").Error(), http.StatusUnauthorized)
		return
	}

	respondJson(w, true)
}

func (AuthController) AuthIndex(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "ok"}); err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusInternalServerError)
	}
}
