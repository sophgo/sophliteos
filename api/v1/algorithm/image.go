package algorithm

import (
	"algoliteos/global"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ImageApi struct{}

func (b *ImageApi) GetImage(c *gin.Context) {
	id := c.Query("cameraId")
	event := c.Query("type")
	fileName := c.Query("fileName")

	// 打开图片文件
	file, err := os.Open(global.PicDir + "/" + id + "/" + event + "/" + fileName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to open image")
		return
	}
	defer file.Close()

	// 设置响应头部信息
	c.Header("Content-Type", "image/jpeg")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to send image")
		return
	}
}
