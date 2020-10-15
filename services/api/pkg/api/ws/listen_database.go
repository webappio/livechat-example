package ws

import (
	"fmt"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"github.com/lib/pq"
	"k8s.io/klog/v2"
	"time"
)

func (handler *Handler) handleChannelChange(uuid string) {
	handler.readDatabase()
}

func (handler *Handler) handleUserChange(uuid string) {
	handler.readDatabase()
}

func (handler *Handler) handleChannelMessage(channelUuid string, messageIndex uint64) {
	handler.readDatabase()
}

func (handler *Handler) handleDirectMessage(toUuid string, messageIndex uint64) {
	handler.readDatabase()
}

func (handler *Handler) listenDatabase() {
	listener := pq.NewListener(model.ConnString, time.Second, 10*time.Second, func(_ pq.ListenerEventType, err error) {
		if err != nil {
			klog.Error(err)
			handler.conn.Close()
		}
	})

	defer func() {
		listener.Close()
	}()

	if err := listener.Listen("channels"); err != nil {
		klog.Error(err)
		handler.conn.Close()
		return
	}
	if err := listener.Listen("users"); err != nil {
		klog.Error(err)
		handler.conn.Close()
		return
	}
	if err := listener.Listen("channel_messages"); err != nil {
		klog.Error(err)
		handler.conn.Close()
		return
	}
	if err := listener.Listen("direct_messages"); err != nil {
		klog.Error(err)
		handler.conn.Close()
		return
	}

	for {
		select {
		case event := <-listener.Notify:
			switch event.Channel {
			case "channels":
				handler.handleChannelChange(event.Extra)
			case "users":
				handler.handleUserChange(event.Extra)
			case "channel_messages":
				var channelUuid string
				var messageIndex uint64
				//_, err := fmt.Sscanf(event.Extra, "%s,%d", &channelUuid, &messageIndex)
				//if err != nil {
				//	klog.Error("invalid channel message data from trigger: ", event.Extra)
				//	continue
				//}
				handler.handleChannelMessage(channelUuid, messageIndex)
			case "direct_messages":
				var toUserUuid string
				var messageIndex uint64
				_, err := fmt.Sscanf(event.Extra, "%s,%d", &toUserUuid, &messageIndex)
				if err != nil {
					klog.Error("invalid channel message data from trigger: ", event.Extra)
					continue
				}
				handler.handleDirectMessage(toUserUuid, messageIndex)
			}
		case <-handler.done:
			handler.conn.Close()
			return
		}
	}
}
