package middleware

import (
	"algoliteos/global"
	"algoliteos/mvc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlockerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.BlockAllRequests {
			c.JSON(http.StatusServiceUnavailable, mvc.FailWithMsg(mvc.Upgradeing, "服务器升级中，暂不可用"))
			// c.File("/var/lib/algoliteos/dist/updating.html")
			c.Abort()
		}
	}
}
