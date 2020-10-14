package ws

import (
	"fmt"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"github.com/lib/pq"
	"k8s.io/klog/v2"
	"time"
)

func handleChannelChange(uuid string) {

}

func handleUserChange(uuid string) {

}

func handleChannelMessage(channelUuid string, messageIndex uint64) {

}

func handleDirectMessage(toUuid string, messageIndex uint64) {

}

func (handler *Handler) Listen() error {
	listener := pq.NewListener(model.ConnString, time.Second, 10*time.Second, func(_ pq.ListenerEventType, err error) {
		if err != nil {
			klog.Error(err)
			handler.conn.Close()
		}
	})

	if err := listener.Listen("channels"); err != nil {
		return err
	}
	if err := listener.Listen("users"); err != nil {
		return err
	}
	if err := listener.Listen("channel_messages"); err != nil {
		return err
	}
	if err := listener.Listen("direct_messages"); err != nil {
		return err
	}

	for {
		select {
		case event := <-listener.Notify:
			switch event.Channel {
			case "channels":
				handleChannelChange(event.Extra)
			case "users":
				handleUserChange(event.Extra)
			case "channel_messages":
				var channelUuid string
				var messageIndex uint64
				_, err := fmt.Sscanf(event.Extra, "%s,%d", &channelUuid, &messageIndex)
				if err != nil {
					klog.Error("invalid channel message data from trigger: ", event.Extra)
					continue
				}
				handleChannelMessage(channelUuid, messageIndex)
			case "direct_messages":
				var toUserUuid string
				var messageIndex uint64
				_, err := fmt.Sscanf(event.Extra, "%s,%d", &toUserUuid, &messageIndex)
				if err != nil {
					klog.Error("invalid channel message data from trigger: ", event.Extra)
					continue
				}
				handleDirectMessage(toUserUuid, messageIndex)
			}
		case <-handler.done:
			return listener.Close()
		}
	}
}
