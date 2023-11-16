package algorithm

import (
	"algoliteos/database"
	"algoliteos/global"
	"algoliteos/logger"
	"algoliteos/mvc"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ReceiveAlarmApi struct{}

var Rand *rand.Rand

func init() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	Rand = rand.New(source)
}

func (b *ReceiveAlarmApi) AlarmRev(c *gin.Context) {
	var alarmDate mvc.AlarmDate
	body, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(body, &alarmDate)

	dir := global.PicDir + "/" + alarmDate.CameraId
	if !fileIsExisted(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}
	dir = dir + "/" + alarmDate.AlarmType
	if !fileIsExisted(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}

	name := fmt.Sprintf("%c%c%d", Rand.Intn(26)+'A', Rand.Intn(26)+'A', alarmDate.Ts) + ".jpg"
	if err := jpegSave(&alarmDate.Scene, dir+"/"+name); err != nil {
		logger.Error("图片保存错误:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(1, "picture save error"))
		return
	}

	alarmInfo := mvc.AlarmInfo{
		Boxes: alarmDate.Boxes,
		Extra: alarmDate.Extra,
	}
	alarmInfoJson, err := json.Marshal(alarmInfo)
	if err != nil {
		logger.Error("序列化 JSON 错误:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(1, "error"))
		return
	}

	record := mvc.Record{
		CameraId: alarmDate.CameraId,
		Type:     alarmDate.AlarmType,
		Date:     alarmDate.Ts,
		Filename: name,
		JsonDate: string(alarmInfoJson),
	}
	if err = SaveRecord(record); err != nil {
		logger.Error("写数据库错误:%v", err)
		c.JSON(http.StatusOK, mvc.Fail(1, "error"))
		return
	}

	c.JSON(http.StatusOK, mvc.Ok())
}

func fileIsExisted(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		return false
	}
	return true
}

func jpegSave(base64ImageData *string, name string) error {
	// 将Base64数据解码成字节数组
	imageData, err := base64.StdEncoding.DecodeString(*base64ImageData)
	if err != nil {
		logger.Error("解码Base64数据失败:%v", err)
		return err
	}

	img, _, err := image.Decode(strings.NewReader(string(imageData)))
	if err != nil {
		logger.Error("解码图片失败:%v", err)
		return err
	}

	// 设置 JPEG 编码器的选项，包括图像质量（1-100，100表示最高质量）,数值越高，图片越清晰，磁盘占用也越高
	options := jpeg.Options{Quality: 80}

	// 图片保存
	outputFile, err := os.Create(name)
	if err != nil {
		logger.Error("创建输出文件失败:", err)
		return err
	}
	defer outputFile.Close()

	// 保存图片为JPEG格式
	err = jpeg.Encode(outputFile, img, &options)
	if err != nil {
		logger.Error("保存图片失败:", err)
		return err
	}
	return nil
}

// 保存告警
func SaveRecord(record mvc.Record) error {
	db := database.DB.Create(&record)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}
