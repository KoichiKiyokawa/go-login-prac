package router

import (
	"github.com/go-login-prac/src/controller"
	"github.com/go-login-prac/src/repository"
	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	authController := controller.AuthController{AuthRepo: repository.AuthRepository{}}

	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", authController.AuthLogin).Methods("POST")
	s.HandleFunc("/", authController.AuthIndex).Methods("GET")
}
