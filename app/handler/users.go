package handler

import (
	"encoding/json"
	"muxgoapi/app/model"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	responseUsers := model.ResponseUser{}
	user := model.User{}
	user.ID = "1"
	user.Name = "Ricardo"

	responseUsers.Users = append(responseUsers.Users, user)

	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(responseUsers)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(response))
}
