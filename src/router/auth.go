package router

import (
	"github.com/go-login-prac/src/controller"
	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	r.HandleFunc("/auth/login", controller.AuthLogin).Methods("POST")
	r.HandleFunc("/auth", controller.AuthIndex).Methods("GET")
}
