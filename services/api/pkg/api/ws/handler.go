package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/layer-devops/livechat-example/services/api/pkg/middleware"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"k8s.io/klog/v2"
	"net/http"
	"net/url"
	"strings"
)

type Handler struct {
	conn *websocket.Conn
	user *model.User
	done chan interface{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header["Origin"]
		if len(origin) == 0 {
			return true
		}
		parsedOrigin, err := url.Parse(origin[0])
		if err != nil {
			return false
		}
		if strings.EqualFold(r.Host, parsedOrigin.Host) {
			return true
		}
		if strings.EqualFold(strings.Split(r.Host, ":")[0], parsedOrigin.Hostname()) {
			return true
		}
		klog.Info(r.Host, parsedOrigin.Hostname())
		return false
	},
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
		handler.write(gin.H{"type": "redirect-to-login"})
		return
	} else {
		handler.write(gin.H{"type": "user-info", "uuid": handler.user.UUID})
	}

	go handler.read()
	go handler.writePings()
	go handler.listenDatabase()
	go handler.readDatabase()
	ginCtx.Abort()
	return
}
