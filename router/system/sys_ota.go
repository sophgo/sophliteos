package system

import (
	v1 "sophliteos/api/v1"
	"sophliteos/global"
	"sophliteos/middleware"

	"github.com/gin-gonic/gin"
)

type OtaRouter struct{}

func (s *OtaRouter) InitOtaRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	otaRouter := Router.Group("api/device/ota", middleware.TimeoutMiddleware(global.OtaTimeOut))

	api := v1.ApiGroupApp.SystemApiGroup.OtaApi
	{
		otaRouter.POST("upgrade", api.OtaUpgrate)
		otaRouter.GET("upgrade", api.OtaUpgradeList)
	}

	return otaRouter
}
