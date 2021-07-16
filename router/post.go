package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func PostRouter(r *mux.Router) {
	s := r.PathPrefix("/post").Subrouter()
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("/")
	})
}
