package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-login-prac/entity"
	"github.com/go-login-prac/service"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) UserController {
	return UserController{userService: userService}
}

func (c UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idInt, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusInternalServerError)
		return
	}
	user, err := c.userService.GetUser(idInt)
	if err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusNotFound)
		return
	}

	respondJson(w, user)
}

func (c UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusInternalServerError)
		return
	}
	respondJson(w, users)
}

func (c UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body entity.User
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusInternalServerError)
		return
	}
	createdUser, err := c.userService.CreateUser(body)
	if err != nil {
		http.Error(w, errors.WithStack(err).Error(), http.StatusInternalServerError)
		return
	}
	respondJson(w, createdUser)
}
