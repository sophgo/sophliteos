package algorithm

import (
	"algoliteos/config"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type HostApi struct{}

var registerEntryID cron.EntryID
var cornRegister *cron.Cron

func init() {
	cornRegister = cron.New(cron.WithSeconds())
	var err error
	registerEntryID, err = cornRegister.AddFunc("0/5 * * * * ?", func() { //每5秒 执行任务
		register()
	})
	if err != nil {
		fmt.Println("err:", err)
	}

	cornRegister.Start()
}

func register() {
	var req struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}
	logger.Info("尝试注册sophliteos服务")

	data := NewRequestWithHeaders("127.0.0.1:8080/api/register", "GET", nil, nil)
	json.Unmarshal(data, &req)

	if req.Msg != "ok" {
		return
	}
	logger.Info("注册到sophliteos服务成功")
	cornRegister.Remove(registerEntryID)
}

func (b *HostApi) AddServerHost(c *gin.Context) {
	var serverHost MediaHost
	body, _ := io.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &serverHost)
	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "JSON解析失败"))
		return
	}

	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	v.Set("algorithm.host", serverHost.Ip+":"+strconv.Itoa(serverHost.Port))
	// 将更改保存到配置文件
	err = v.WriteConfig()
	conf.Unlock()

	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "set error"))
		return
	}
	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *HostApi) GetServerHost(c *gin.Context) {
	res := getHost()
	parts := strings.Split(res, ":")
	port, _ := strconv.Atoi(parts[1])
	mediaHost := MediaHost{
		Ip:   parts[0],
		Port: port,
	}
	c.JSON(http.StatusOK, mvc.Success(mediaHost))
}

func (b *HostApi) GetRegisterHost(c *gin.Context) {
	c.JSON(http.StatusOK, mvc.Ok())
}

func getHost() string {
	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	res := v.GetString("algorithm.host")
	conf.Unlock()

	return res

}
