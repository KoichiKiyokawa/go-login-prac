package routers

import (
	"github.com/go-login-prac/src/controllers"
	"github.com/go-login-prac/src/repositories"
	"github.com/go-login-prac/src/services"
	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	// ここで依存を注入する
	authRepository := repositories.AuthRepository{}
	authService := services.AuthService{AuthRepository: authRepository}
	authController := controllers.AuthController{AuthService: authService}

	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", authController.AuthLogin).Methods("POST")
	s.HandleFunc("/", authController.AuthIndex).Methods("GET")
}
