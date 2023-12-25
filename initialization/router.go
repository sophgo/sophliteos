package initialization

import (
	"net/http/httputil"
	"net/url"
	"sophliteos/config"
	"sophliteos/logger"
	"sophliteos/middleware"
	"sophliteos/router"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // 设置Gin的模式为release
	Router := gin.New()
	Router.Use(gin.Recovery())

	// Router := gin.Default()

	Router.MaxMultipartMemory = 64 << 20

	// 创建一个反向代理到算法业务
	algoURL, _ := url.Parse("http://localhost:8081")
	proxy := httputil.NewSingleHostReverseProxy(algoURL)

	algorithmGroup := Router.Group("/algorithm")
	// 添加反向代理处理器
	algorithmGroup.Any("/*path", func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	systemRouter := router.RouterGroupApp.System

	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	path := v.GetString("server.www")
	conf.Unlock()

	Router.StaticFile("/_app.config.js", path+"/_app.config.js")
	Router.StaticFile("/", path+"/index.html") // 前端网页入口页面
	Router.Static("/assets", path+"/assets")   // dist里面的静态资源

	Router.Use(middleware.BlockerMiddleware())

	PublicGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitDownRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.AuthMiddleware())
	{

		systemRouter.InitSsmUpgradeRouter(PrivateGroup)
		systemRouter.InitUpgradeRouter(PrivateGroup)
		systemRouter.InitVersionRouter(PrivateGroup)
		systemRouter.InitBasicRouter(PrivateGroup)
		systemRouter.InitResourceRouter(PrivateGroup)
		systemRouter.InitPasswordRouter(PrivateGroup)
		systemRouter.InitIpRouter(PrivateGroup)
		systemRouter.InitAlarmRouter(PrivateGroup)
		systemRouter.InitLogRouter(PrivateGroup)
		systemRouter.InitOtaRouter(PrivateGroup)

	}
	logger.Info("Router Init Ok")
	return Router
}

// NewProxy 创建一个反向代理
func NewProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}
