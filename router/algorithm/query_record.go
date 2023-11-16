package algorithm

import (
	v1 "algoliteos/api/v1"
	"algoliteos/global"
	"algoliteos/middleware"

	"github.com/gin-gonic/gin"
)

type QueryRouter struct{}

func (s *QueryRouter) InitQueryRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm/alarm", middleware.TimeoutMiddleware(global.TimeOut))
	api := v1.ApiGroupApp.AlgoGroup.QueryApi
	{
		router.POST("list", api.GetRecord)
		router.POST("modSize", api.ModSize)

	}

	return router
}
