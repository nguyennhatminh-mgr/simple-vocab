package models

type User struct {
	Id    int   `json:"id"`
	FirstName  string `json:"first_name"`
	LastName int `json:"last_name"`
	Email int `json:"email"`
	Password string `json:"password"`
}