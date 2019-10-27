package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResponseUser struct {
	Users  []User `json:"users"`
	Status int    `json:"status"`
}
