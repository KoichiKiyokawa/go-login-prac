package router

import (
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
	s.HandleFunc("/", userController.GetAllUsers).Methods("GET")
	s.HandleFunc("/{id}", userController.GetUserById).Methods("GET")
	s.HandleFunc("/", userController.CreateUser).Methods("POST")
}
