package mvc

// 接受告警信息
type AlarmDate struct {
	AlarmType       string  `json:"alarmType"`
	AlarmVideoStart int     `json:"alarmVideoStart"`
	AlarmVideoStop  int     `json:"alarmVideoStop"`
	Boxes           []Box   `json:"boxes"`
	CameraId        string  `json:"cameraId"`
	EventID         string  `json:"eventID"`
	Extra           Extra   `json:"extra"`
	Scene           string  `json:"scene"`
	SceneHeight     int     `json:"sceneHeight"`
	SceneWidth      int     `json:"sceneWidth"`
	Score           float64 `json:"score"`
	Ts              int64   `json:"ts"`
}

// 告警图片数据库存储
type Record struct {
	ID       uint `gorm:"primary_key"`
	CameraId string
	Type     string
	Date     int64
	Filename string
	JsonDate string `gorm:"type:text"`
}

type Box struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}

type ItemInBox struct {
	Confidence float64 `json:"confidence"`
	Type       string  `json:"type"`
	Content    struct {
		NotSatisfy int `json:"notSatisfy"`
	} `json:"content"`
	FaceId string `json:"faceId"`
}

type Extra struct {
	Cei        interface{} `json:"cei"`
	ItemsInBox []ItemInBox `json:"itemsInBox"`
}

type AlarmInfo struct {
	Boxes []Box `json:"boxes"`
	Extra Extra `json:"extra"`
}

// 算法任务接收结构体
type AlgoTask struct {
	TaskName   string   `json:"taskName"`
	DeviceName string   `json:"deviceName"`
	DeviceId   string   `json:"deviceId"`
	Url        string   `json:"url"`
	Abilities  []string `json:"abilities"`
}

// 算法任务数据库存储
type AlgoTaskSql struct {
	ID         uint   `gorm:"primary_key"`
	TaskName   string `json:"taskName"`
	Status     int    `json:"status"`
	DeviceName string `json:"deviceName"`
	DeviceId   string `json:"deviceId"`
	Url        string `json:"url"`
	Abilities  string `json:"abilities"`
}

// 算法应用响应
type AlgoReq struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// 告警信息查询请求
type AlgoQuery struct {
	BeginTime  int64    `json:"beginTime"`
	EndTime    int64    `json:"endTime"`
	PageNo     int      `json:"pageNo"`
	PageSize   int      `json:"pageSize"`
	DeviceName string   `json:"deviceName"`
	Alarms     []string `json:"alarms"`
}

// 告警信息查询响应
type AlarmReq struct {
	Total     int            `json:"total"`
	PageNo    int            `json:"pageNo"`
	PageSize  int            `json:"pageSize"`
	PageCount int            `json:"pageCount"`
	Items     []AlarmReqItem `json:"items"`
	UsedSize  string         `json:"usedSize"`
	MaxSize   string         `json:"maxSize"`
}

type AlarmReqItem struct {
	DeviceName string      `json:"deviceName"`
	Image      string      `json:"image"`
	Time       int64       `json:"time"`
	AlarmType  string      `json:"alarmType"`
	Boxes      []Box       `json:"boxes"`
	ItemsInBox []ItemInBox `json:"itemsInBox"`
}

// 接收巡检信息
type Detect struct {
	FaultSource int          `json:"faultSource"`
	List        []DeviceInfo `json:"list"`
	OperationID string       `json:"operationId"`
}

type DeviceInfo struct {
	DeviceID     string `json:"deviceId"`
	EncodeFormat string `json:"encodeFormat"`
	FaultType    int    `json:"faultType"`
	MediaPull    int    `json:"mediaPull"`
	OperationID  string `json:"operationId"`
	Resolution   string `json:"resolution"`
}

// 下发算法任务body
type AlgoStart struct {
	CameraID   string        `json:"cameraId"`
	URL        string        `json:"url"`
	ImageOut   string        `json:"imageOut"`
	InputType  string        `json:"inputType"`
	DecodeType string        `json:"decodeType"`
	NotifyURL  string        `json:"notifyUrl"`
	SkipFrame  int           `json:"skipFrame"`
	ROI        []interface{} `json:"roi"`
	AreaBoxes  []interface{} `json:"areaBoxes"`
	Abilities  []Ability     `json:"abilities"`
}

type Ability struct {
	Name  string       `json:"name"`
	Value AbilityValue `json:"value"`
}

// 算法能力结构体
type AbilityValue struct {
	Interval      float64 `json:"interval"`
	MinTarry      int     `json:"minTarry"`
	AlarmInterval float64 `json:"alarmInterval"`
	Threshold     float64 `json:"threshold"`
	AreaIsReverse bool    `json:"areaIsReverse"`
	PointType     int     `json:"pointType"`
	ZoomFactor    float64 `json:"zoomFactor"`
	ConfirmCount  int     `json:"confirmCount"`
	MinBox        MinRect `json:"minBox"`
	AreaBoxes     [][]struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"areaBoxes"`
}

type MinRect struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// 接受算法参数配置
type AlgoConfig struct {
	Ability   string  `json:"ability"`
	Interval  float64 `json:"interval"`
	MinBox    MinRect `json:"minBox"`
	Threshold float64 `json:"threshold"`
}

// 流媒体添加设备结构体
type Device struct {
	Codec       string `json:"codec"`
	DeviceId    string `json:"deviceId"  gorm:"primary_key"`
	DeviceName  string `json:"name"`
	Protocol    int    `json:"protocol"` //协议类型，0-未知，1-国标，2-RTSP，3-海康SDK，4-大华SDK
	PtzType     int    `json:"ptzType"`  //摄像机类型，0-未知，1-球机，2-半球，3-枪机，4-遥控枪机
	Resolution  string `json:"resolution"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	MediaServer string `json:"mediaServer"`
	MediaPull   int    `json:"mediaPull"`
}
