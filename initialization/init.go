package initialization

import (
	"algoliteos/config"
	"algoliteos/database"
	"algoliteos/global"
	"algoliteos/logger"
	"time"
)

func InitBase() {
	// 加载配置
	config.LoadConfig()
	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	logLevel := v.GetString("log.level")
	logPath := v.GetString("log.path")
	timeout := v.GetString("server.timeout")
	global.PicDir = v.GetString("dir.source")
	conf.Unlock()

	// 日志处理
	logger.InitLogging(logPath, "algo.log", logLevel)

	// 初始化sqlite
	database.InitDB()

	global.TimeOut, _ = time.ParseDuration("30s")
	global.OtaTimeOut, _ = time.ParseDuration(timeout)
	global.BlockAllRequests = false
}
