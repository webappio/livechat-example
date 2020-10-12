package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"k8s.io/klog/v2"
)

type Handler struct {
	conn *websocket.Conn
	done chan interface{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewForContext(ginCtx *gin.Context) {
	var err error
	handler := &Handler{done: make(chan interface{})}
	handler.conn, err = upgrader.Upgrade(ginCtx.Writer, ginCtx.Request, nil)
	if err != nil {
		klog.Error(err)
		return
	}
	go handler.read()
	go handler.write()
	return
}
