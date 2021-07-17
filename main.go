package main

import (
	"log"
	"net/http"

	"github.com/go-login-prac/middleware"
	"github.com/go-login-prac/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	router.SetupRoutes(r)
	r.Use(middleware.LoggingMiddleware, middleware.CorsMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
