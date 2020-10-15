package model

import "time"

type ChannelMessage struct {
	ChannelUUID string `json:"channel_uuid"`
	UserUUID string `json:"user_uuid"`
	Index int `json:"index"`
	Text string `json:"text"`
	Time time.Time `json:"time"`
}