package system

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"sophliteos/client/ssm"
	"sophliteos/global"
	"sophliteos/logger"
	mvc "sophliteos/mvc/core"
	error2 "sophliteos/mvc/error"
	"sophliteos/mvc/i18n"
	services "sophliteos/mvc/services/opt"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type OtaApi struct{}

const (
	Ctrl = "ctrl"
	Core = "core"
)

func init() {
	i18n.SetString(i18n.Zh, Ctrl, "控制板")
	i18n.SetString(i18n.En, Ctrl, "Ctrl")
	i18n.SetString(i18n.Zh, Core, "核心板")
	i18n.SetString(i18n.En, Core, "Core")
}

func (b *OtaApi) OtaUpgrate(c *gin.Context) {

	// 参数判断
	module := c.Request.FormValue("module")
	if module != Ctrl && module != Core {
		c.JSON(http.StatusOK, mvc.Fail(error2.UpgradeParamErr, "param error"))
		return
	}

	cmd := exec.Command("rm", "-rf", "/data/ota/*")
	cmd.Dir = "/data/ota"

	// 执行命令
	err := cmd.Run()
	if err != nil {
		logger.Error("rm ota failed", err)
	}

	otaName, err := saveFile(c.Request, "/data/ota/")
	if err != nil {
		logger.Error("update failed", err)
		c.JSON(http.StatusOK, mvc.FailWithMsg(error2.UpgradeErr, "文件上传失败"))
		return
	}

	cmdFlag := "--target="
	ip := c.Request.FormValue("ip")
	if len(ip) > 0 {
		cmdFlag = cmdFlag + ip
	} else {
		cmdFlag = cmdFlag + "all"
	}

	if module == Ctrl {
		cmdFlag = "/data/ota/local_update.sh md5.txt 1"
	}

	deviceType := global.DeviceType

	otaInfo := ssm.OtaVersion{
		Name:       strings.ToLower(deviceType) + "_" + module + "_upgrade_" + time.Now().Format("20060102150405"),
		Product:    strings.ToUpper(deviceType),
		FileName:   otaName,
		ModuleName: module,
		CmdFlag:    cmdFlag,
	}
	logger.Info("OTA info is :%v", otaInfo)

	_, err = ssm.OtaUpgrade(otaInfo)
	mvc.HandleError(err, error2.UpgradeErr)

	services.SaveOptLog(c.Request, "%sOTA升级", i18n.GetString(mvc.GetLang(c.Request), module))

	c.JSON(http.StatusOK, mvc.OkWithMsg("操作成功"))
}

func (b *OtaApi) OtaUpgradeList(c *gin.Context) {

	result, err := ssm.OtaUpgradeList()
	mvc.HandleError(err)
	var otaTasks []ssm.OtaTask
	res, _ := json.Marshal(result.Result)
	_ = json.Unmarshal(res, &otaTasks)
	c.JSON(http.StatusOK, mvc.Success(otaTasks))

}

func (b *OtaApi) OtaRollback(c *gin.Context) {

	var task ssm.OtaTask
	data, _ := io.ReadAll(c.Request.Body)
	mapstructure.Decode(task, data)
	result, err := ssm.OtaUpgradeList()
	mvc.HandleError(err)
	var tasks []ssm.OtaTask
	mapstructure.Decode(tasks, result.Result)
	var version ssm.OtaVersion
	for _, t := range tasks {
		if t.WorkflowID == task.WorkflowID {
			version.Name = t.Name
			version.CmdFlag = t.CmdFlag
			version.FileName = t.FileName
			version.ModuleName = t.ModuleName
			version.Product = t.Product
			break
		}
	}
	if len(version.Name) <= 0 {
		c.JSON(http.StatusOK, mvc.Fail(error2.UpgradeTaskNotFound, "Upgrade Task Not Found"))
		return
	}
	result, err = ssm.OtaRollback(version)
	mvc.HandleError(err, error2.RollbackErr)
	services.SaveOptLog(c.Request, "OTA回滚")
	c.JSON(http.StatusOK, mvc.OkWithMsg("请求成功"))
}

// 文件上传控制
func handleFile(request *http.Request, module string) string {
	file, handler, err := request.FormFile("file")
	if err != nil {
		if module == Ctrl {
			panic(err)
		} else {
			return ""
		}
	}
	defer file.Close()

	dir := "/data/ota/"
	if module == Core {
		dir = "/recovery/tftp/"
	}

	f, err := os.OpenFile(dir+handler.Filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(f, file)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(handler.Filename)
		_ = request.MultipartForm.RemoveAll()
	}()
	return handler.Filename
}
