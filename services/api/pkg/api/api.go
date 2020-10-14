package api

import (
	"github.com/gin-gonic/gin"
	"github.com/layer-devops/livechat-example/services/api/pkg/api/ws"
)

func AddRoutes(group *gin.RouterGroup) {
	apiGroup := group.Group("/api")
	apiGroup.GET("/ws", ws.NewForContext)
	AddLoginRoutes(apiGroup)
}