package main

import (
	"log"
	"net/http"

	"github.com/go-login-prac/src/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	router.AuthRouter(r)
	router.PostRouter(r)

	router.ShowAllRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
