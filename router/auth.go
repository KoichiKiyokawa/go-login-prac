package router

import (
	"net/http"

	"github.com/go-login-prac/controller"
	"github.com/go-login-prac/repository"
	"github.com/go-login-prac/service"
	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	// ここで依存を注入する
	userRepository := repository.NewUserRepository()
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", authController.AuthLogin).Methods(http.MethodPost, http.MethodOptions)
	s.HandleFunc("/check", authController.AuthCheck).Methods(http.MethodGet, http.MethodOptions)
	s.HandleFunc("/", authController.AuthIndex).Methods(http.MethodGet, http.MethodOptions)
}
