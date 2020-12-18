package wallecontrollers

import (
	"encoding/json"
	"library/components"
	"models"
)



// @Title restart
// @Description restart
// @Param   id       query    	 int   	 true      "service id"
// @Param   Host     query		*Task  	 true	   "host id"
// @Success 0 {id} int64
// @Failure 1 restart 失败
// @Failure 2 User not found
// @router /walle/stop/ [post]
func (c *WalleController)Stop() {


	//beego.Info(string(c.Ctx.Input.RequestBody))
	operationRecord := new(models.OperationRecord)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, operationRecord)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	operationRecord.Class = "stop"

	s := components.BaseComponents{}
	s.SetProject(operationRecord.Service.Project)
	s.SetService(operationRecord.Service)
	s.SetUser(operationRecord.User)

	d := new(components.BaseDocker)
	d.SetBaseComponents(s)

	err = d.Stop(operationRecord)
	if err != nil {
		c.SetJson(1, "", "docker关闭出错"+err.Error())
		return
	}

	c.SetJson(0, "", "")
	c.ServeJSON()
}
