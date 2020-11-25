package wallecontrollers

import (
	"encoding/json"
	"library/components"
	"models"
)

// @Title 下载站 download file
// @Description 下载站 download file
// @Param   id      query     int   	true     "service id"
// @Success 0 {id} int64
// @Failure 1 download file 失败
// @Failure 2 User not found
// @router /walle/build [post]
func (c *WalleController)Build() {


	//beego.Info(string(c.Ctx.Input.RequestBody))
	// var domain *models.Domain
	operationRecord := new(models.OperationRecord)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, operationRecord)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	operationRecord.Class = "build"

	b := components.BaseComponents{}
	b.SetProject(operationRecord.Service.Project)
	b.SetService(operationRecord.Service)
	b.SetUser(operationRecord.User)
	d := new(components.BaseDocker)
	d.SetBaseComponents(b)


	tag, err := d.Build(operationRecord)
	if err != nil {
		c.SetJson(1, nil, err.Error())
		return
	} else {
		s := new(models.Service)
		s = operationRecord.Service
		s.Tag = tag
		_, err := models.UpdateServiceById(s)
		if err != nil {
			c.SetJson(1, nil, err.Error())
			return
		}
		c.SetJson(0, tag, "")
	}

}

