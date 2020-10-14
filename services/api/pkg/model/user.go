package model

type User struct {
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	PasswordHash string `json:"-"`
}

