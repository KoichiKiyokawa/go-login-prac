package router

import (
	"github.com/go-login-prac/src/controller"
	"github.com/go-login-prac/src/service"
	"github.com/gorilla/mux"
)

func IndexRouter(r *mux.Router) {
	appService := service.AppService{}
	appController := controller.AppController{AppService: appService}
	r.HandleFunc("/", appController.Index)
}
