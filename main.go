package main

import (
	"log"
	"net/http"

	"github.com/go-login-prac/src/routers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	routers.AuthRouter(r)
	routers.PostRouter(r)

	routers.ShowAllRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
