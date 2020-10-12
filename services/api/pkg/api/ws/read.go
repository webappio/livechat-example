package ws

import (
	"github.com/gorilla/websocket"
	"k8s.io/klog/v2"
	"time"
)

func (handler *Handler) handleMessage(message []byte) {

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