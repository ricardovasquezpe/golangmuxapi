package handler

import (
	"encoding/json"
	"fmt"
	"golangmuxapi/app/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUsers(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/*responseUsers := model.ResponseUser{}
	user := model.User{}
	user.ID = "1"
	user.Name = "Ricardo"

	responseUsers.Users = append(responseUsers.Users, user)*/

	responseUsers := model.ResponseUser{}

	var users []model.User
	users = model.GetUsers(db, bson.M{})

	responseUsers.Users = users
	responseUsers.Status = 1

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(responseUsers)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(response))
}

func CreateUser(db *mongo.Client, w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		return
	}
	model.InsertUser(db, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Println(name)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	querys := r.URL.Query()
	param1 := querys.Get("param1")
	param2 := querys.Get("param2")
	fmt.Println(param1, param2)
}
