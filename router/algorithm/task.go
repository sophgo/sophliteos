package algorithm

import (
	v1 "algoliteos/api/v1"
	"algoliteos/global"
	"algoliteos/middleware"

	"github.com/gin-gonic/gin"
)

type TaskRouter struct{}

func (s *ReceiveAlarmRouter) InitKaolaRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm/task", middleware.TimeoutMiddleware(global.TimeOut))
	api := v1.ApiGroupApp.AlgoGroup.TaskApi
	{
		router.POST("add", api.AddTask)
		router.POST("modify", api.ModTask)
		router.POST("delete", api.DeleteTask)
		router.POST("start", api.StartTask)
		router.POST("stop", api.StopTask)
		router.POST("list", api.List)

	}

	return router
}

func (s *ReceiveAlarmRouter) InitConfigRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm/config", middleware.TimeoutMiddleware(global.TimeOut))
	api := v1.ApiGroupApp.AlgoGroup.TaskApi
	{
		router.GET("get", api.GetTaskConfig)
		router.POST("mod", api.ModTaskConfig)

	}

	return router
}
