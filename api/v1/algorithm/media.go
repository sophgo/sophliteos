package algorithm

import (
	"algoliteos/config"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type MediaApi struct{}

var host string
var mediaPullMap = make(map[string]mvc.DeviceInfo)
var mu sync.RWMutex

func init() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("0 0 0 * * ?", func() {
		autoDetect()
	})
	if err != nil {
		fmt.Println("err:", err)
	}

	c.Start()
}

func (b *MediaApi) AddMedia(c *gin.Context) {
	var mediaHost MediaHost
	body, _ := io.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &mediaHost)
	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "JSON解析失败"))
		return
	}

	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	v.Set("media.host", mediaHost.Ip+":"+strconv.Itoa(mediaHost.Port))
	// 将更改保存到配置文件
	err = v.WriteConfig()
	conf.Unlock()

	if err != nil {
		c.JSON(http.StatusOK, mvc.Fail(-1, "set error"))
		return
	}
	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *MediaApi) GetMedia(c *gin.Context) {
	getMediaHost()
	parts := strings.Split(host, ":")
	port, _ := strconv.Atoi(parts[1])
	mediaHost := MediaHost{
		Ip:   parts[0],
		Port: port,
	}
	c.JSON(http.StatusOK, mvc.Success(mediaHost))
}

func (b *MediaApi) AddDev(c *gin.Context) {
	var deviceInfo DeviceInfo
	body, _ := io.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &deviceInfo)
	if err != nil {
		logger.Error("JSON解析失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "JSON解析失败"))
		return
	}

	if containsSpecialCharacters(deviceInfo.Name) {
		logger.Error("设备名称不合法")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "设备名称不能包含符号，只能为中英文和数字"))
		return
	}

	deviceMap := getMediaDevices()
	for _, value := range deviceMap {
		if deviceInfo.Name == value.DeviceName {
			c.JSON(http.StatusOK, mvc.Fail(-1, "设备已存在"))
			return
		}
	}

	body, _ = json.Marshal(deviceInfo)

	var devices MediaRes
	res := NewRequestWithHeaders(host+"/addDevice", "POST", nil, body)
	err = json.Unmarshal(res, &devices)
	if err != nil {
		logger.Error("流媒体设备添加失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备添加失败"))
		return
	}
	if devices.Code != 0 {
		logger.Error("请求流媒体设备信息失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备添加失败"))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *MediaApi) ModDev(c *gin.Context) {
	var deviceMod DeviceMod
	body, _ := io.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &deviceMod)
	if err != nil {
		logger.Error("JSON解析失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "JSON解析失败"))
		return
	}

	deviceMap := getMediaDevices()
	if _, exists := deviceMap[deviceMod.DeviceId]; !exists {
		c.JSON(http.StatusOK, mvc.Fail(-1, "设备不存在"))
		return
	}

	body, _ = json.Marshal(deviceMod)

	var devices MediaRes
	res := NewRequestWithHeaders(host+"/modDevice", "POST", nil, body)
	err = json.Unmarshal(res, &devices)
	if err != nil {
		logger.Error("流媒体设备修改失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备修改失败"))
		return
	}
	if devices.Code != 0 {
		logger.Error("流媒体设备修改失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备修改失败"+devices.Msg))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *MediaApi) DelDev(c *gin.Context) {
	var deviceDel struct {
		Devices []string `json:"device"`
	}
	body, _ := io.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &deviceDel)
	if err != nil {
		logger.Error("JSON解析失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "JSON解析失败"))
		return
	}

	var devices MediaRes
	res := NewRequestWithHeaders(host+"/delDevice", "POST", nil, body)
	err = json.Unmarshal(res, &devices)
	if err != nil {
		logger.Error("流媒体设备删除失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备删除失败"))
		return
	}
	if devices.Code != 0 {
		logger.Error("流媒体设备删除失败:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(-1, "流媒体设备删除失败"+devices.Msg))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *MediaApi) GetLiveUrl(c *gin.Context) {
	deviceId := c.Query("deviceId")
	getMediaHost()

	body := fmt.Sprintf("{\"deviceId\":\"%s\"}", deviceId)
	res := NewRequestWithHeaders(host+"/getLiveUrl", "POST", nil, []byte(body))

	var liveUrl LiveUrl
	err := json.Unmarshal(res, &liveUrl)
	if err != nil {
		logger.Error("解析JSON失败:%v", err)
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "media server error"))
		return
	}

	c.JSON(http.StatusOK, mvc.Success(liveUrl.TsUrl))
}

func (b *MediaApi) GetDevices(c *gin.Context) {
	var page struct {
		PageNo   int `json:"pageNo"`
		PageSize int `json:"pageSize"`
	}
	reqBody, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(reqBody, &page)

	deviceMap := getMediaDevices()

	if deviceMap == nil {
		logger.Error("请求流媒体设备信息失败")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "请求流媒体设备信息失败"))
		return
	}

	devices := DevicesList{
		TotalSize: len(deviceMap),
		PageCount: int(math.Ceil(float64(len(deviceMap)) / float64(page.PageSize))),
		PageNo:    page.PageNo,
		PageSize:  page.PageSize,
	}

	keys := make([]string, 0, len(deviceMap))
	for key := range deviceMap {
		keys = append(keys, key)
	}
	// 对键进行排序
	sort.Strings(keys)

	for _, key := range keys {
		value := deviceMap[key]
		value.MediaPull = mediaPullMap[key].MediaPull
		devices.Devices = append(devices.Devices, value)
	}

	c.JSON(http.StatusOK, mvc.Success(devices))
}

func (b *MediaApi) DeviceDetectRev(c *gin.Context) {
	var detectInfo mvc.Detect
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &detectInfo)

	logger.Info("revice detect result:%v", detectInfo)
	mu.Lock()
	defer mu.Unlock()
	mediaPullMap[detectInfo.List[0].DeviceID] = detectInfo.List[0]
	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *MediaApi) DetectDev(c *gin.Context) {
	var devs struct {
		Devs []string `json:"deviceIds"`
	}
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &devs)
	logger.Info("巡检设备:%v", devs.Devs)
	detectDevice(devs.Devs)

	c.JSON(http.StatusOK, mvc.Ok())
}

func detectDevice(deviceId []string) bool {
	var deviceIds []DeviceId

	for _, dev := range deviceId {
		devID := DeviceId{
			DeviceID: dev,
		}
		deviceIds = append(deviceIds, devID)
	}
	upload := getConfigString("upload")
	logger.Info("上传地址:%s", upload)

	body := DeviceDetectBody{
		DetectReportURL: "http://" + upload + "/algorithm/media/detect",
		DevList:         deviceIds,
	}
	bodyByte, _ := json.Marshal(body)

	var req mvc.AlgoReq
	getMediaHost()
	res := NewRequestWithHeaders(host+"/manualDeviceDetect", "POST", nil, bodyByte)
	json.Unmarshal(res, &req)

	return req.Status == 0
}

func getMediaDevices() map[string]mvc.Device {
	getMediaHost()

	var devices DevicesList
	res := NewRequestWithHeaders(host+"/getDevListEx", "POST", nil, []byte("{}"))
	err := json.Unmarshal(res, &devices)
	if err != nil {
		logger.Error("请求流媒体设备信息失败:%v", err)
		return nil
	}
	if devices.Code != 0 {
		logger.Error("请求流媒体设备信息失败:%v", err)
		return nil
	}

	deviceMpa := make(map[string]mvc.Device)

	for i := range devices.Devices {
		devices.Devices[i].MediaServer = host

		device := devices.Devices[i]
		if device.Type != "camera" {
			continue
		}

		deviceMpa[device.DeviceId] = device
	}

	return deviceMpa
}

func autoDetect() {
	deviceMpa := getMediaDevices()

	var result []string
	for key := range deviceMpa {
		result = append(result, key)
	}
	logger.Info("巡检设备:%v", result)
	detectDevice(result)
}

func getMediaHost() {
	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	host = v.GetString("media.host")
	conf.Unlock()
}

func containsSpecialCharacters(input string) bool {
	specialCharacters := "&/?=#\\+-_:;*\" "

	for _, char := range specialCharacters {
		if strings.Contains(input, string(char)) {
			return true
		}
	}

	return false
}

type DeviceId struct {
	DeviceID string `json:"deviceId"`
}

type DeviceInfo struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
	Url      string `json:"url"`
	PtzType  int    `json:"ptzType"`
}

type DeviceMod struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
	Url      string `json:"url"`
	PtzType  int    `json:"ptzType"`
	DeviceId string `json:"deviceId"`
}

type DeviceDetectBody struct {
	DetectReportURL string     `json:"detectReportUrl"`
	DevList         []DeviceId `json:"devList"`
}

type DevicesList struct {
	Code      int          `json:"code"`
	Msg       string       `json:"msg"`
	TotalSize int          `json:"totalSize"`
	PageCount int          `json:"pageCount"`
	PageSize  int          `json:"pageSize"`
	PageNo    int          `json:"pageNo"`
	Devices   []mvc.Device `json:"device"`
}

type LiveUrl struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	TsUrl string `json:"tsUrl"`
}

type MediaRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type MediaHost struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}
