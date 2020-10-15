package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"k8s.io/klog/v2"
	"time"
)

func (handler *Handler) write(data gin.H) {
	handler.conn.SetWriteDeadline(time.Now().Add(time.Second*10))
	err := handler.conn.WriteJSON(data)
	if err != nil {
		klog.Error(err)
		handler.conn.Close()
	}
}

func (handler *Handler) writePings() {
	ticker := time.NewTicker(time.Second * 50)
	defer func() {
		ticker.Stop()
		handler.conn.Close()
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
