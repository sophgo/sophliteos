package initialization

import (
	"algoliteos/logger"
	"algoliteos/middleware"
	"algoliteos/router"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // 设置Gin的模式为release
	Router := gin.New()
	Router.Use(gin.Recovery())

	// Router := gin.Default()

	algoRouter := router.RouterGroupApp.Algorithm

	Router.Use(middleware.BlockerMiddleware())

	PublicGroup := Router.Group("")
	{

		algoRouter.InitReceiveAlarmRouter(PublicGroup)
		algoRouter.InitImageRouter(PublicGroup)
		algoRouter.InitKaolaRouter(PublicGroup)
		algoRouter.InitMediaRouter(PublicGroup)
		algoRouter.InitQueryRouter(PublicGroup)
		algoRouter.InitHostRouter(PublicGroup)
		algoRouter.InitConfigRouter(PublicGroup)
		algoRouter.InitRegisterRouter(PublicGroup)

	}
	logger.Info("Router Init Ok")
	return Router
}
