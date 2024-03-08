package system

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"sophliteos/client/httpclient"
	"sophliteos/database"
	"sophliteos/global"
	"sophliteos/logger"
	mvc "sophliteos/mvc/core"
	error2 "sophliteos/mvc/error"
	services "sophliteos/mvc/services/opt"

	"sophliteos/mvc/types"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

// Login
func (b *BaseApi) Login(c *gin.Context) {
	req := types.LoginRequest{}
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &req)

	userName := req.UserName
	password := req.Password
	if userName == "" || password == "" {
		c.JSON(http.StatusOK, mvc.Fail(error2.InvalidUsernameOrPassword, "用户名或密码错误"))
		return
	}

	user, _ := database.QueryUserWithName(userName)
	var token string
	if user == nil || user.Password != password { // 验证密码
		c.JSON(http.StatusOK, mvc.Fail(error2.InvalidUsernameOrPassword, "用户名或密码错误"))
		return
	} else {
		now := time.Now()
		if now.After(user.ExpireTime) {
			token = strings.ReplaceAll(uuid.New().String(), "-", "")
			user.Token = token
			user.LoginTime = now
			user.ExpireTime = now.Add(time.Hour * 2)
			database.UpdateUser(user)
		} else {
			token = user.Token
		}
	}
	mvc.SetUser(token, user)

	services.SaveOptLog(c.Request, "登录")

	c.JSON(http.StatusOK, mvc.Success(types.LoginResponse{
		Token: token,
	}))
}

func (b *BaseApi) Logout(c *gin.Context) {
	req := types.LogoutRequest{}
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &req)

	if req.Token != "" {
		user, err := database.QueryUserWithToken(req.Token)
		if err == nil && user != nil {
			user.Token = ""
			user.ExpireTime = time.Now()
			database.UpdateUser(user)
			services.SaveOptLog(c.Request, "注销登录")
			mvc.RemoveUser(req.Token)
			c.JSON(http.StatusOK, mvc.Success(nil))

		} else {
			c.JSON(http.StatusOK, mvc.Fail(error2.InvalidUsernameOrPassword, "Invalid Token"))
			return
		}
	} else {
		c.JSON(http.StatusOK, mvc.Fail(error2.InvalidUsernameOrPassword, "Invalid Token"))
	}
}

func (b *BaseApi) AlarmListen(c *gin.Context) {
	var alarmRec database.AlarmRec
	data, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(data, &alarmRec)

	if alarmRec.DiskName == "/dev/mmcblk0p4" || alarmRec.DiskName == "/dev/mmcblk0p2" || alarmRec.DiskName == "/dev/mmcblk0p5" {
		c.JSON(http.StatusOK, mvc.Ok())
		return
	}

	logger.Debug("recive alarm:%s", string(data))

	if global.DeviceType == "" {
		GetArmResource(c)
	}

	var alarm database.Alarm
	switch global.DeviceType {
	case "SE5", "SE7", "SE9":
		alarm = database.Alarm{
			DeviceSn:      global.Resource.DeviceSn,
			DeviceIp:      "",
			CreatedAt:     time.Now(),
			ComponentType: getType(alarmRec.Code),
			Code:          alarmRec.Code,
			Msg:           alarmRec.Msg,
		}
		alarm.CoreUnitBoardSn = global.Resource.DeviceSn

	default:
		alarm = database.Alarm{
			DeviceSn:      alarmRec.DeviceSn,
			DeviceIp:      "",
			CreatedAt:     time.Now(),
			ComponentType: getType(alarmRec.Code),
			Code:          alarmRec.Code,
			Msg:           alarmRec.Msg,
		}

		alarm.CoreUnitBoardSn = alarmRec.BoardSn
		if alarm.CoreUnitBoardSn == "" {
			alarm.CoreUnitBoardSn = alarmRec.DeviceSn
		}
	}
	if alarm.ComponentType == "disk" && alarmRec.Code < 0 {
		alarm.Msg = "磁盘" + alarmRec.DiskName + ":  " + alarmRec.Msg
	}

	err := database.SaveAlarm(alarm)
	mvc.HandleError(err)
	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *BaseApi) AlgoRegister(c *gin.Context) {
	global.AlgoFlag = true
	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *BaseApi) AlgoExist(c *gin.Context) {
	c.JSON(http.StatusOK, mvc.Success(global.AlgoFlag))
}

func Register() {
	var req struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
	}
	logger.Info("尝试注册algoliteos服务")

	data, _ := httpclient.NewRequest("127.0.0.1:8081/algorithm/register", "GET", nil, nil)
	json.Unmarshal(data, &req)

	if req.Msg != "ok" {
		logger.Info("algoliteos未运行服务")
		return
	}
	logger.Info("注册到algoliteos服务成功")
	global.AlgoFlag = true
}

func getType(code int) string {
	if code < 0 {
		code = -code
	}
	code = code / 1000
	var res string

	switch code {
	case 101:
		res = "cpu"
	case 102:
		res = "memory"
	case 103:
		res = "disk"
	case 104:
		res = "netCard"
	case 201:
		res = "board"
	case 202:
		res = "chip"
	default:
		res = ""
	}
	return res
}
