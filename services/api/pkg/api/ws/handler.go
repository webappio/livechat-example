package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/layer-devops/livechat-example/services/api/pkg/middleware"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"k8s.io/klog/v2"
	"time"
)

type Handler struct {
	conn *websocket.Conn
	user *model.User
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

	handler.user, err = middleware.GetUser(ginCtx)
	if err != nil {
		klog.Info("User is not logged in.")
		handler.conn.SetWriteDeadline(time.Now().Add(time.Second*20))
		handler.conn.WriteJSON(gin.H{"type": "redirect-to-login"})
		return
	}

	go handler.read()
	go handler.write()
	return
}
