package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"k8s.io/klog/v2"
	"time"
)

func (handler *Handler) handleContentMessage(message []byte) {
	var contentMessage struct {
		Type        string `json:"type"`
		Contents    string `json:"contents"`
		ChannelUUID string `json:"channel_uuid"`
	}

	err := json.Unmarshal(message, &contentMessage)
	if err != nil {
		klog.Error("Invalid content message: ", string(message))
		return
	}

	err = model.Exec(`
INSERT INTO channel_messages(channel_uuid, user_uuid, index, text)
VALUES($1, $2, (
	SELECT COALESCE(MAX(index)+1, 1)
	FROM channel_messages WHERE channel_uuid=$1
), $3)`, contentMessage.ChannelUUID, handler.user.UUID, contentMessage.Contents)
	if err != nil {
		klog.Error("Could not insert channel message: ", err)
		return
	}
}

func (handler *Handler) handleNewChannelMessage(message []byte) {
	var contentMessage struct {
		Type        string `json:"type"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	err := json.Unmarshal(message, &contentMessage)
	if err != nil {
		klog.Error("Invalid content message: ", string(message))
		return
	}

	err = model.Exec("INSERT INTO channels(name, topic, description) VALUES ($1, '', $2)", contentMessage.Name, contentMessage.Description)
	if err != nil {
		klog.Error(err)
		return
	}
}

func (handler *Handler) handleMessage(message []byte) {
	var typer struct {
		Type string `json:"type"`
	}
	err := json.Unmarshal(message, &typer)
	if err != nil {
		klog.Error(err)
		return
	}

	switch typer.Type {
	case "new_message":
		handler.handleContentMessage(message)
	case "new_channel":
		handler.handleNewChannelMessage(message)
	}
}

func (handler *Handler) read() {
	defer func() {
		handler.conn.Close()
		close(handler.done)
	}()
	handler.conn.SetReadLimit(4096)
	handler.conn.SetReadDeadline(time.Now().Add(time.Second * 60))
	handler.conn.SetPongHandler(func(string) error {
		return handler.conn.SetReadDeadline(time.Now().Add(time.Second * 60))
	})
	for {
		_, msg, err := handler.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				klog.Error(err)
			}
			return
		}
		handler.handleMessage(msg)
	}
}
