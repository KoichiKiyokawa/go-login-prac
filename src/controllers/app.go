package controllers

import (
	"net/http"

	"github.com/go-login-prac/src/services"
)

type AppController struct {
	AppService services.IAppService
}

func (c AppController) Index(w http.ResponseWriter, r *http.Request) {
	respondJson(w, map[string]string{"message": c.AppService.GetHello()})
}
