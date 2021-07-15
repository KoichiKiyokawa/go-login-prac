package router

import (
	"github.com/go-login-prac/src/controller"
	"github.com/go-login-prac/src/services"
	"github.com/gorilla/mux"
)

func IndexRouter(r *mux.Router) {
	appService := services.AppService{}
	appController := controller.AppController{AppService: appService}
	r.HandleFunc("/", appController.Index)
}
