package models

import (
	"encoding/json"
	"os"
)

const UsersFileName = "users.json"

// changed Id type to string
type User struct {
	Id     string `json:"id"`
	Name   string `json:"name" `
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func LoadUsers() map[string]User {
	m := make(map[string]User)
	openedFile, err := os.Open(UsersFileName)
	defer openedFile.Close()
	if err != nil {
		openedFile, err = os.Create(UsersFileName)
		if err != nil {
			panic(err)
		}
	} else {
		json.NewDecoder(openedFile).Decode(&m)
	}
	return m
}
