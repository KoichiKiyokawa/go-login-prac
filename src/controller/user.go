package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-login-prac/src/entity"
	"github.com/go-login-prac/src/service"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) UserController {
	return UserController{userService: userService}
}

func (c UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body entity.User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdUser, err := c.userService.CreateUser(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondJson(w, createdUser)
}
