package algorithm

import (
	v1 "algoliteos/api/v1"
	"algoliteos/global"
	"algoliteos/middleware"

	"github.com/gin-gonic/gin"
)

type ServerHostRouter struct{}

func (s *ServerHostRouter) InitHostRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm/server", middleware.TimeoutMiddleware(global.TimeOut))
	api := v1.ApiGroupApp.AlgoGroup.HostApi
	{
		router.POST("add", api.AddServerHost)
		router.GET("get", api.GetServerHost)

	}

	return router
}

func (s *ServerHostRouter) InitRegisterRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm", middleware.TimeoutMiddleware(global.TimeOut))
	api := v1.ApiGroupApp.AlgoGroup.HostApi
	{
		router.GET("register", api.GetRegisterHost)

	}

	return router
}
