package wallecontrollers

import (
	"encoding/json"
	"library/components"
	"models"
)

// @Title reload
// @Description reload
// @Param   service_id       query    	 int   	 true      "service id"
// @Param   host_id  	     query		 int  	 true	   "host id"
// @Success 0 {id} int64
// @Failure 1 reload 失败
// @Failure 2 User not found
// @router /walle/reload/ [post]
func (c *WalleController)Reload() {

	//beego.Info(string(c.Ctx.Input.RequestBody))
	operationRecord := new(models.OperationRecord)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, operationRecord)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	operationRecord.Class = "reload"

	s := components.BaseComponents{}
	s.SetProject(operationRecord.Service.Project)
	s.SetService(operationRecord.Service)
	s.SetUser(operationRecord.User)

	d := new(components.BaseDocker)
	d.SetBaseComponents(s)

	err = d.Reload(operationRecord)
	if err != nil {
		c.SetJson(1, "", "docker重载出错"+err.Error())
		return
	}

	c.SetJson(0, "", "")
	c.ServeJSON()
}


// @Title reload
// @Description 批量服务主机重载
// @Param   OperationRecord[]       query    	 obj   	 true      "OperationsRecord"
// @Success 0 {id} int64
// @Failure 1 reload 失败
// @Failure 2 User not found
// @router /walle/multiReload/ [post]
func (c *WalleController)MultiReload() {

	//beego.Info(string(c.Ctx.Input.RequestBody))

	type multiReload struct{
		User 	 		models.User
		Services 		[]models.Service
	}

	muReload := new(multiReload)
	muReload.Services = make([]models.Service,0)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, muReload)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	for _, service := range muReload.Services {
		operationRecord := new(models.OperationRecord)
		operationRecord.Service = &service
		if len(service.Hosts) == 0{
			continue
		}
		operationRecord.Host = service.Hosts[0]
		operationRecord.User = &muReload.User
		operationRecord.Class = "reload"
		s := components.BaseComponents{}
		s.SetProject(operationRecord.Service.Project)
		s.SetService(operationRecord.Service)
		s.SetUser(operationRecord.User)

		d := new(components.BaseDocker)
		d.SetBaseComponents(s)

		err = d.Reload(operationRecord)
		if err != nil {
			c.SetJson(1, "", "服务重载失败: "+s.Service.Name)
			return
		}

	}


	c.SetJson(0, "", "")
	c.ServeJSON()
}

