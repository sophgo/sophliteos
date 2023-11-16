package algorithm

import (
	httpclient "algoliteos/client"
	"algoliteos/config"
	"algoliteos/database"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
)

// 算法应用api
const (
	taskFind   = "/dynamic/api/v1/find"
	taskCancle = "/dynamic/api/v1/cancel"
	taskSetup  = "/dynamic/api/v1/setup"
)

var algorithmHost string

type TaskApi struct{}

func (b *TaskApi) AddTask(c *gin.Context) {
	algoTask := mvc.AlgoTask{}
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &algoTask)

	if algoTask.DeviceId == "" || algoTask.DeviceName == "" || algoTask.Url == "" {
		logger.Error("添加任务参数错误")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "参数错误，不能为空"))
		return
	}

	// 通过 TaskName 查找数据
	var task mvc.AlgoTaskSql
	result := database.DB.Where("task_name = ?", algoTask.TaskName).First(&task)

	if result.Error != gorm.ErrRecordNotFound {
		logger.Error("任务已存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务已存在"))
		return
	}

	if len(algoTask.Abilities) == 0 {
		logger.Error("算法能力为空")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "算法能力为空"))
		return
	}

	jsonDate, err := json.Marshal(algoTask.Abilities)
	if err != nil {
		logger.Error("json Abilities marshal error")
	}

	algoTaskSql := mvc.AlgoTaskSql{
		TaskName:   algoTask.TaskName,
		Status:     0,
		DeviceName: algoTask.DeviceName,
		DeviceId:   algoTask.DeviceId,
		Url:        algoTask.Url,
		Abilities:  string(jsonDate),
	}

	// 保存到数据库
	_ = database.DB.Create(&algoTaskSql)
	logger.Info("任务创建成功%v", algoTaskSql)

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *TaskApi) ModTask(c *gin.Context) {
	algoTask := mvc.AlgoTask{}
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &algoTask)

	// 通过 TaskName 查找数据
	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_name = ?", algoTask.TaskName).First(&task)

	if db.Error == gorm.ErrRecordNotFound {
		logger.Error("任务不存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务不存在"))
		return
	}

	jsonDate, err := json.Marshal(algoTask.Abilities)
	if err != nil {
		logger.Error("json Abilities marshal error")
	}

	if task.DeviceId != algoTask.DeviceId && algoTask.DeviceId != "" {
		deviceMpa := getMediaDevices()
		device := deviceMpa[algoTask.DeviceId]
		task.DeviceId = device.DeviceId
		task.DeviceName = device.DeviceName
		task.Url = device.Url
	}

	task.Abilities = string(jsonDate)

	// 保存到数据库
	_ = database.DB.Save(&task)
	logger.Info("任务修改成功%v", task)

	stopAlgoTask(task.DeviceName)

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *TaskApi) DeleteTask(c *gin.Context) {
	var taskName struct {
		TaskName string `json:"taskName"`
	}
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &taskName)

	// 通过 TaskName 查找数据
	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_name = ?", taskName.TaskName).First(&task)

	if task.Status == 1 {
		stopAlgoTask(task.DeviceName)
	}

	// 删除找到的数据
	db = database.DB.Delete(&task)
	if db.Error != nil {
		logger.Error("删除数据时出错:%v", db.Error)
		return
	}
	logger.Info("任务删除成功%v", task)

	c.JSON(http.StatusOK, mvc.Ok())
}

func (b *TaskApi) StartTask(c *gin.Context) {
	var taskName struct {
		TaskName string `json:"taskName"`
	}
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &taskName)

	// 通过 TaskName 查找数据
	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_name = ?", taskName.TaskName).First(&task)
	if db.Error == gorm.ErrRecordNotFound {
		logger.Error("任务不存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务不存在"))
		return
	}

	if task.Status == 1 {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务正在运行"))
		return
	}

	if mediaPullMap[task.DeviceId].MediaPull == 0 {
		logger.Error("视频源未巡检")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "视频源未巡检或拉流失败"))
		return
	}

	// 解析Abilities字符串到数组
	var abilities []string
	err := json.Unmarshal([]byte(task.Abilities), &abilities)
	if err != nil {
		logger.Error("解析JSON失败:%v", err)
		return
	}
	//构造发送请求body
	upload := getConfigString("upload")
	requestBody := getTaskString("task.base")
	replaceMap := map[string]string{
		"{{taskId}}":    task.DeviceName,
		"{{streamUrl}}": task.Url,
		"{{uploadUrl}}": "http://" + upload + "/algorithm/upload",
	}
	// 执行替换
	for key, value := range replaceMap {
		requestBody = strings.ReplaceAll(requestBody, key, value)
	}

	var eventStrings string
	for _, ability := range abilities {
		temp := getTaskString("task.eventMap." + ability + ".template")
		temp = strings.ReplaceAll(temp, "\"{{hotRegion}}\"", "[]")

		value := getJsonString(ability + ".value").(map[string]interface{})
		jsonData, err := json.Marshal(value)
		if err != nil {
			logger.Error("JSON编码出错:%v", err)
		}

		var JsonValue mvc.AbilityValue
		json.Unmarshal(jsonData, &JsonValue)

		temp = strings.ReplaceAll(temp, "\"alarmInterval\":10", fmt.Sprintf(`"alarmInterval": %d`, int(JsonValue.AlarmInterval)))
		temp = strings.ReplaceAll(temp, "\"threshold\":0.4", fmt.Sprintf(`"threshold": %.2f`, JsonValue.Threshold))

		// 将结构体转为 JSON 字符串
		jsonString, err := json.Marshal(JsonValue.MinBox)
		if err != nil {
			logger.Error("JSON编码出错:%v", err)
			continue
		}

		// logger.Info(string(jsonString))
		temp = strings.ReplaceAll(temp, "{\"width\":50,\"height\":50}", string(jsonString))

		eventStrings += temp
		eventStrings += ","
	}
	// 删除最后一个逗号结尾字符
	if strings.HasSuffix(eventStrings, ",") {
		eventStrings = eventStrings[:len(eventStrings)-1]
	}
	requestBody = strings.ReplaceAll(requestBody, "{{abilities}}", eventStrings)
	logger.Info("start task body:%s", requestBody)

	getAlgorithmHost()
	var req mvc.AlgoReq
	data := NewRequestWithHeaders(algorithmHost+taskSetup, "POST", nil, []byte(requestBody))
	json.Unmarshal(data, &req)

	if req.Status != 200 {
		if req.Message == "" {
			c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务启动失败"))
			return
		}
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, req.Message))
		return
	}
	task.Status = 1
	_ = database.DB.Save(&task)
	c.JSON(http.StatusOK, mvc.Ok())

}

func (b *TaskApi) StopTask(c *gin.Context) {
	var taskName struct {
		TaskName string `json:"taskName"`
	}
	reqBody, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(reqBody, &taskName)

	var task mvc.AlgoTaskSql
	db := database.DB.Where("task_name = ?", taskName.TaskName).First(&task)
	if db.Error == gorm.ErrRecordNotFound {
		logger.Error("任务不存在")
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "任务不存在"))
		return
	}

	if !stopAlgoTask(task.DeviceName) {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "error"))
		return
	}

	task.Status = 0
	_ = database.DB.Save(&task)

	c.JSON(http.StatusOK, mvc.Ok())
}

func stopAlgoTask(deviceName string) bool {

	body := struct {
		CameraId string `json:"cameraId"`
	}{CameraId: deviceName}
	data, _ := json.Marshal(body)

	getAlgorithmHost()

	var req mvc.AlgoReq
	res := NewRequestWithHeaders(algorithmHost+taskCancle, "POST", nil, data)
	json.Unmarshal(res, &req)

	if req.Status != 200 {
		return false
	}
	return true
}

func (b *TaskApi) List(c *gin.Context) {
	var page struct {
		PageNo   int `json:"pageNo"`
		PageSize int `json:"pageSize"`
	}
	reqBody, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(reqBody, &page)

	var algoTasks []mvc.AlgoTaskSql
	database.DB.Find(&algoTasks)

	var cameraData CameraData
	var cameraId []string

	getAlgorithmHost()
	res := NewRequestWithHeaders(algorithmHost+taskFind, "GET", nil, nil)
	err := json.Unmarshal(res, &cameraData)
	if err != nil {
		logger.Error("解析JSON失败:%v", err)
	}
	// 遍历 CameraID 切片并输出 cameraId 值
	for _, cameraID := range cameraData.Data {
		cameraId = append(cameraId, cameraID.CameraID)
	}

	var taskList TaskList
	var items []Item
	taskList.Total = len(algoTasks)
	taskList.PageCount = taskList.Total / page.PageSize
	taskList.PageSize = page.PageSize

	for i := range algoTasks {
		if containsString(cameraId, algoTasks[i].DeviceName) {
			algoTasks[i].Status = 1
		} else {
			algoTasks[i].Status = 0
		}
		database.DB.Save(&algoTasks[i])

		var abilities []string
		err := json.Unmarshal([]byte(algoTasks[i].Abilities), &abilities)
		if err != nil {
			logger.Error("解析JSON失败:%v", err)
			continue
		}
		item := Item{
			TaskName:    algoTasks[i].TaskName,
			DeviceName:  algoTasks[i].DeviceName,
			Status:      algoTasks[i].Status,
			ErrorReason: "",
			Abilities:   abilities,
		}
		items = append(items, item)

	}
	taskList.Items = items
	c.JSON(http.StatusOK, mvc.Success(taskList))
}

func (b *TaskApi) GetTaskConfig(c *gin.Context) {

	var algoConfigs []mvc.AlgoConfig
	for _, ability := range AlgoAllAbilities {
		value := getJsonString(ability + ".value").(map[string]interface{})
		jsonData, err := json.Marshal(value)
		if err != nil {
			logger.Error("JSON编码出错:%v", err)
		}

		var JsonValue mvc.AbilityValue
		json.Unmarshal(jsonData, &JsonValue)

		algoConfig := mvc.AlgoConfig{
			Ability:   ability,
			Interval:  JsonValue.AlarmInterval,
			MinBox:    JsonValue.MinBox,
			Threshold: JsonValue.Threshold,
		}
		algoConfigs = append(algoConfigs, algoConfig)
	}
	c.JSON(http.StatusOK, mvc.Success(algoConfigs))
}

func (b *TaskApi) ModTaskConfig(c *gin.Context) {
	var algoConfig mvc.AlgoConfig
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &algoConfig)
	conf := &config.JsonConf
	conf.Lock()
	defer conf.Unlock()

	v := conf.GetViper()
	v.Set(algoConfig.Ability+".value.alarmInterval", algoConfig.Interval)
	v.Set(algoConfig.Ability+".value.threshold", algoConfig.Threshold)
	var data map[string]interface{}
	mapstructure.Decode(algoConfig.MinBox, &data)
	v.Set(algoConfig.Ability+".value.minBox", data)
	err := v.WriteConfig()
	if err != nil {
		c.JSON(http.StatusOK, mvc.FailWithMsg(-1, "failed"))
	}
	v.ReadInConfig() // 刷新Viper的配置

	c.JSON(http.StatusOK, mvc.Ok())
}

func NewRequestWithHeaders(url string, method string, header map[string]string, bytes []byte) []byte {
	data, err := httpclient.NewRequest(url, method, header, bytes)
	if err != nil {
		return nil
	}
	return data
}

func containsString(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true
		}
	}
	return false
}

func getTaskString(path string) string {
	conf := &config.Event
	conf.Lock()
	v := conf.GetViper()
	res := v.GetString(path)
	conf.Unlock()
	return res
}

func getConfigString(path string) string {
	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	res := v.GetString(path)
	conf.Unlock()
	return res
}

func getJsonString(path string) interface{} {
	conf := &config.JsonConf
	conf.Lock()
	v := conf.GetViper()
	res := v.Get(path)
	conf.Unlock()
	return res
}

func getAlgorithmHost() {
	conf := &config.Conf
	conf.Lock()
	v := conf.GetViper()
	algorithmHost = v.GetString("algorithm.host")
	conf.Unlock()
}

var AlgoAllAbilities = []string{
	"Smoking", "HumanBreakIn", "PeopleWithoutMask",
	"NoneMotorVehicleParking", "UnexpectedEvents",
	"RoadOccupatedInOperation", "ElectricCarEntersElevator",
	"OccupationOfFireAccess", "GarbageExposed",
	"GarbageOverflow", "SmokeDetected", "FireDetected",
	"LoessExposed", "RoadPonding", "MotorVehicleParking",
	"OffStoreOperation", "IllegalOutdoorAdvertising",
	"BannerOrSlogansHungging", "HangDownTheStreet",
	"ManholeCoverDamaged", "ConstructionTruck",
	"MotorVehicleBreakIn", "WithoutHelmetOnSite",
	"EngineeringVehicle", "RoadOccupatedInConstruction",
	"RoadDamaged", "HeapOfMaterial",
	"OccupationOfBarrierFreeAccess", "PlayPhone",
	"Fishing", "OffDutyAlarm", "WithoutSafeHelmet",
	"NoneMotorVehicleBreakIn", "GroundDust", "Climbing",
	"PersonNumber", "FollowIntoHousePerson",
	"WithoutChefHat", "WearDetection",
	"PersonnelFalls", "HumanCrossTheBorder",
	"PassengerFlow", "Retrograde",
	"SmokyEngineeringVehicle", "HumanHover",
}

type Item struct {
	TaskName    string   `json:"taskName"`
	DeviceName  string   `json:"deviceName"`
	Status      int      `json:"status"`
	ErrorReason string   `json:"errorReason"`
	Abilities   []string `json:"abilities"`
}

type TaskList struct {
	Total     int    `json:"total"`
	PageSize  int    `json:"pageSize"`
	PageCount int    `json:"pageCount"`
	PageNo    int    `json:"pageNo"`
	Items     []Item `json:"items"`
}

type CameraData struct {
	Data []struct {
		CameraID string `json:"cameraId"`
	} `json:"data"`
}
