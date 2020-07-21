package wallecontrollers

import (
	"encoding/json"
	"library/components"
	"models"

	"github.com/astaxie/beego"
)

// @Title 下载站 download file
// @Description 下载站 download file
// @Param   id       query    	 int   	 true      "service id"
// @Param   Host     query		*Task  	 true	   "host id"
// @Success 0 {id} int64
// @Failure 1 download file 失败
// @Failure 2 User not found
// @router /walle/download/ [post]
func (c *WalleController)Download() {


	beego.Info(string(c.Ctx.Input.RequestBody))
	operationRecord := new(models.OperationRecord)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, operationRecord)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	operationRecord.Class = "download"

	s := components.BaseComponents{}
	s.SetProject(operationRecord.Service.Project)
	s.SetService(operationRecord.Service)
	s.SetUser(operationRecord.User)

	d := new(components.BaseDocker)
	d.SetBaseComponents(s)

	pathZip, err := d.DownloadZip(operationRecord)
	if err != nil {
		c.SetJson(1, "", "docker_cp_ZIP fail:"+err.Error())
		return
	}

	c.SetJson(0, pathZip, "")
	c.ServeJSON()
}

