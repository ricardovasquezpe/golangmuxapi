package app

import (
	"golangmuxapi/app/handler"
	"golangmuxapi/app/middle"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRoutersv1()
}

func (a *App) setVersionApi(v string) {
	a.Router.PathPrefix("/api/" + v).Subrouter()
}

func (a *App) setRoutersv1() {
	apiv1 := a.Router.PathPrefix("/api/v1").Subrouter()
	apiv1.Use(middle.MiddlewareOne)
	apiv1.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	apiv1.HandleFunc("/createuser", handler.CreateUser).Methods("POST")
	apiv1.HandleFunc("/getuser/{name}", handler.GetUser).Methods("GET")
	apiv1.HandleFunc("/searchuser", handler.SearchUser).Methods("GET")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
