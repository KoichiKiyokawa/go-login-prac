package router

import (
	"github.com/go-login-prac/src/controller"
	"github.com/go-login-prac/src/service"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	appService := service.NewAppService()
	appController := controller.NewAppController(appService)
	r.HandleFunc("/", appController.Index)

	AuthRouter(r)
	UserRouter(r)
	PostRouter(r)
	ShowAllRoutes(r)
}
