package controller

import (
	"net/http"

	"github.com/go-login-prac/service"
)

type AppController struct {
	appService service.IAppService
}

func NewAppController(appService service.IAppService) AppController {
	return AppController{appService: appService}
}

func (c AppController) Index(w http.ResponseWriter, r *http.Request) {
	respondJson(w, map[string]string{"message": c.appService.GetHello()})
}
