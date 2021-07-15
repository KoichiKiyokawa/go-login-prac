package router

import (
	"github.com/go-login-prac/src/controller"
	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", controller.AuthLogin).Methods("POST")
	s.HandleFunc("/", controller.AuthIndex).Methods("GET")
}
