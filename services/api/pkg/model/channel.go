package model

type Channel struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Topic string `json:"topic"`
	Description string `json:"description"`
}
