package router

import (
	"net/http"

	"github.com/go-login-prac/controller"
	"github.com/go-login-prac/repository"
	"github.com/go-login-prac/service"
	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router) {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("/", userController.GetAllUsers).Methods(http.MethodGet, http.MethodOptions)
	s.HandleFunc("/{id}", userController.GetUserById).Methods(http.MethodGet, http.MethodOptions)
	s.HandleFunc("/", userController.CreateUser).Methods(http.MethodPost, http.MethodOptions)
}
