package app

import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
}

func (app *App) Init() {
	app.Router = mux.NewRouter()
}

func (app *App) setRouters() {

}
