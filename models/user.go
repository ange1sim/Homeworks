package models

type User struct {
	Username string
	Email    string
}

var Database = make(map[string]User)