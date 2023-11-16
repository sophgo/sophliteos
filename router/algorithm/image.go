package algorithm

import (
	v1 "algoliteos/api/v1"
	"algoliteos/global"
	"algoliteos/middleware"

	"github.com/gin-gonic/gin"
)

type ImageRouter struct{}

func (s *ImageRouter) InitImageRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	router := Router.Group("algorithm", middleware.TimeoutMiddleware(global.TimeOut))
	api := v1.ApiGroupApp.AlgoGroup.ImageApi
	{
		router.GET("image", api.GetImage)
	}

	return router
}
