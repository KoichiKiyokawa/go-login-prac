package main

import (
	"log"
	"net/http"

	"github.com/go-login-prac/src/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	router.SetupRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
