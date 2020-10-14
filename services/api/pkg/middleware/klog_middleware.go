package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"time"
)

func Recover() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			klog.ErrorDepth(2, fmt.Sprintf("%+v", err))
		}()
		ginCtx.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		startTime := time.Now()
		path := ginCtx.Request.URL.Path
		ginCtx.Next()
		klog.Infof(
			"%v %v %v %v (%.2f)",
			ginCtx.Request.Method,
			path,
			ginCtx.Writer.Status(),
			ginCtx.ClientIP(),
			time.Now().Sub(startTime).Seconds(),
		)
	}
}