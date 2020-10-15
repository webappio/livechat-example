package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"k8s.io/klog/v2"
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

	handler.write(gin.H{"type": "channels", "channels": channels})
	for _, channel := range channels {
		messages := []*model.ChannelMessage{}
		err = model.Select(&messages, "SELECT * FROM channel_messages WHERE channel_uuid=$1 ORDER BY index DESC LIMIT 100", channel.UUID)
		if err != nil {
			klog.Error(err)
			handler.conn.Close()
			return
		}
		handler.write(gin.H{"type": "channel_messages", "messages": messages})
	}

	users := []*model.User{}
	err = model.Select(&users, "SELECT * FROM users")
	if err != nil {
		klog.Error(err)
		handler.conn.Close()
		return
	}
	handler.write(gin.H{"type": "users", "users": users})
}