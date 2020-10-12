package model

type Channel struct {
	UUID string `db:"uuid"`
	Name string `db:"name"`
}