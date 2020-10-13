package model

type Channel struct {
	UUID string `db:"uuid" json:"uuid"`
	Name string `db:"name" json:"name"`
}