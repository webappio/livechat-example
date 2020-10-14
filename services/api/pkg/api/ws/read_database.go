package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"k8s.io/klog/v2"
	"time"
)

func (handler *Handler) readDatabase() {
	channels := []*model.Channel{}
	err := model.Select(&channels, "SELECT * FROM channels")
	if err != nil {
		klog.Error(err)
		handler.conn.Close()
		return
	}

	if len(channels) == 0 {
		defaultChannel := model.Channel{
			Name: "default",
			Description: "workspace chatter",
		}
		_ = model.Get(&defaultChannel.UUID, "INSERT INTO channels(name, topic, description) VALUES ($1, $2, $3) RETURNING uuid", defaultChannel.Name, defaultChannel.Topic, defaultChannel.Description)
		channels = append(channels, &defaultChannel)
	}

	handler.conn.SetWriteDeadline(time.Now().Add(time.Second*10))
	handler.conn.WriteJSON(gin.H{"type": "channels", "channels": channels})
}