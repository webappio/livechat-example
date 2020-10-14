package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"k8s.io/klog/v2"
	"time"
)

func (handler *Handler) handleContentMessage(message []byte) {
	var contentMessage struct {
		Type string `json:"type"`
		Contents string `json:"contents"`
		ChannelUUID string `json:"channel_uuid"`
	}
	err := json.Unmarshal(message, &contentMessage)
	if err != nil {
		klog.Error("Invalid content message: ", string(message))
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
	}
}

func (handler *Handler) read() {
	defer func() {
		handler.conn.Close()
		close(handler.done)
	}()
	handler.conn.SetReadLimit(4096)
	handler.conn.SetReadDeadline(time.Now().Add(time.Second*60))
	handler.conn.SetPongHandler(func(string) error {
		return handler.conn.SetReadDeadline(time.Now().Add(time.Second*60))
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