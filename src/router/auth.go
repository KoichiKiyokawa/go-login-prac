package router

import (
	"github.com/go-login-prac/src/controller"
	"github.com/go-login-prac/src/repository"
	"github.com/go-login-prac/src/service"
	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	// ここで依存を注入する
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", authController.AuthLogin).Methods("POST")
	s.HandleFunc("/", authController.AuthIndex).Methods("GET")
}
