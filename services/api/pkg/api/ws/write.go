package ws

import (
	"github.com/gorilla/websocket"
	"time"
)

func (handler *Handler) write() {
	ticker := time.NewTicker(time.Second * 50)
	defer func() {
		ticker.Stop()
		handler.conn.Close()
		close(handler.done)
	}()
	for {
		select {
		case <-ticker.C:
			if err := handler.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(10*time.Second)); err != nil {
				return
			}
		case <-handler.done:
			return
		}
	}
}
