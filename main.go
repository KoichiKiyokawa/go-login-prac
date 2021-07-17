package main

import (
	"log"
	"net/http"

	"github.com/go-login-prac/middleware"
	"github.com/go-login-prac/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	router.SetupRoutes(r)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	r.Use(middleware.LoggingMiddleware)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
