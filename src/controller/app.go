package controller

import (
	"net/http"

	"github.com/go-login-prac/src/service"
)

type AppController struct {
	AppService service.IAppService
}

func (c AppController) Index(w http.ResponseWriter, r *http.Request) {
	respondJson(w, map[string]string{"message": c.AppService.GetHello()})
}
