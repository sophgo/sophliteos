package algorithm

import (
	v1 "algoliteos/api/v1"
	"algoliteos/global"
	"algoliteos/middleware"

	"github.com/gin-gonic/gin"
)

type ReceiveAlarmRouter struct{}

func (s *ReceiveAlarmRouter) InitReceiveAlarmRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm", middleware.TimeoutMiddleware(global.TimeOut))
	api := v1.ApiGroupApp.AlgoGroup.ReceiveAlarmApi
	{
		router.POST("upload", api.AlarmRev)
	}

	return router
}
