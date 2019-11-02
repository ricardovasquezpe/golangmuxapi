package app

import (
	"context"
	"golangmuxapi/app/handler"
	"golangmuxapi/app/middle"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type App struct {
	Router  *mux.Router
	MClient *mongo.Client
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.MClient = GetClient()
	a.setVersionApi("v1")
	a.setRouters()
}

func (a *App) setVersionApi(v string) {
	a.Router = a.Router.PathPrefix("/api/" + v).Subrouter()
	a.Router.Use(middle.MiddlewareOne)
}

func (a *App) setRouters() {
	a.call("/users", "GET", a.handleRequest(handler.GetAllUsers))
	a.call("/createuser", "POST", a.handleRequest(handler.CreateUser))
	/*a.Router.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	a.Router.HandleFunc("/createuser", handler.CreateUser).Methods("POST")
	a.Router.HandleFunc("/getuser/{name}", handler.GetUser).Methods("GET")
	a.Router.HandleFunc("/searchuser", handler.SearchUser).Methods("GET")*/
}

func (a *App) call(path string, method string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(method)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *mongo.Client, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.MClient, w, r)
	}
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err_ping := client.Ping(context.Background(), readpref.Primary())
	if err_ping != nil {
		log.Fatal("Couldn't connect to the database", err_ping)
	} else {
		log.Println("Connected!")
	}

	return client
}
