package router

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
)

func ShowAllRoutes(r *mux.Router) {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
