package routers

import (
	"github.com/go-login-prac/src/controllers"
	"github.com/go-login-prac/src/services"
	"github.com/gorilla/mux"
)

func IndexRouter(r *mux.Router) {
	appService := services.AppService{}
	appController := controllers.AppController{AppService: appService}
	r.HandleFunc("/", appController.Index)
}
