package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/layer-devops/livechat-example/services/api/pkg/api"
	"github.com/layer-devops/livechat-example/services/api/pkg/middleware"
	"github.com/layer-devops/livechat-example/services/api/pkg/model"
	"k8s.io/klog/v2"
)

func main()  {
	klog.InitFlags(nil)
	flag.Parse()

	err := model.Init(10)
	if err != nil {
		panic(err)
	}

	router := gin.New()
	router.Use(middleware.Recover(), middleware.Logger(), middleware.Session())
	api.AddRoutes(&router.RouterGroup)
	klog.Fatal(router.Run())
}