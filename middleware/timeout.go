package middleware

import (
	"context"
	"net/http"
	"sophliteos/logger"
	mvc "sophliteos/mvc/core"
	error2 "sophliteos/mvc/error"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeOut time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeOut)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan bool)
		go func() {
			c.Next()
			done <- true
		}()

		select {
		case <-done:

		case <-ctx.Done():
			// 请求超时，执行超时逻辑
			logger.Error("timeout on %s %s", c.Request.Method, c.Request.URL.Path)
			c.Abort()
			c.JSON(http.StatusGatewayTimeout, mvc.FailWithMsg(error2.UpgradeErr, "传输超时"))
		}
	}
}
