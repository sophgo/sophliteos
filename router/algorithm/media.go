package algorithm

import (
	v1 "algoliteos/api/v1"
	"algoliteos/global"
	"algoliteos/middleware"

	"github.com/gin-gonic/gin"
)

type MediaRouter struct{}

func (s *MediaRouter) InitMediaRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm/media", middleware.TimeoutMiddleware(global.TimeOut))
	api := v1.ApiGroupApp.AlgoGroup.MediaApi
	{
		router.POST("add", api.AddMedia)
		router.GET("get", api.GetMedia)
		router.POST("list", api.GetDevices)
		router.GET("live", api.GetLiveUrl)
		router.POST("detect", api.DeviceDetectRev)
		router.POST("check", api.DetectDev)

		router.POST("dev/add", api.AddDev)
		router.POST("dev/mod", api.ModDev)
		router.POST("dev/del", api.DelDev)

	}

	return router
}
