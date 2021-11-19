package robotcontroller

import (
	"controllers"
	"library/components"
	"models"
)

type RobotController struct {
	controllers.BaseController
}

// @Title restart
// @Description robot restart back-end service
// @Param   service_id       query    	 int   	 true      "service id"
// @Param   host_id 	     query		 int  	 true	   "host id"
// @Success 0 {id} int64
// @Failure 1 restart 失败
// @Failure 2 User not found
// @router /robot/restart/ [get]
func (c *RobotController) RobotRestart() {

	serviceId, err := c.GetInt("service_id")
	if err != nil {
		c.SetJson(1, err, "获取 serviceId 失败")
		return
	}
	hostId, err := c.GetInt("host_id")
	if err != nil {
		c.SetJson(1, err, "获取 hostId 失败")
		return
	}

	service, err := models.GetServiceById(serviceId)
	if err != nil {
		c.SetJson(1, err, "获取 service by service id 失败")
		return
	}

	host, err := models.GetHostById(hostId)
	if err != nil {
		c.SetJson(1, err, "获取 host by host id 失败")
		return
	}

	operationRecord := new(models.OperationRecord)
	operationRecord.Service = service
	operationRecord.Host = host
	operationRecord.User = c.User

	operationRecord.Class = "restart"

	s := components.BaseComponents{}
	s.SetProject(operationRecord.Service.Project)
	s.SetService(operationRecord.Service)
	s.SetUser(operationRecord.User)

	d := new(components.BaseDocker)
	d.SetBaseComponents(s)

	err = d.Restart(operationRecord)
	if err != nil {
		c.SetJson(1, "", "docker关闭出错"+err.Error())
		return
	}

	c.SetJson(0, "", "")
	c.ServeJSON()
}

// @Title reload
// @Description robot reload front-end service
// @Param   service_id       query    	 int   	 true      "service id"
// @Param   host_id     	 query		 int  	 true	   "host id"
// @Success 0 {id} int64
// @Failure 1 reload 失败
// @Failure 2 User not found
// @router /robot/reload/ [get]
func (c *RobotController) RobotReload() {

	serviceId, err := c.GetInt("service_id")
	if err != nil {
		c.SetJson(1, err, "获取 serviceId 失败")
		return
	}
	hostId, err := c.GetInt("host_id")
	if err != nil {
		c.SetJson(1, err, "获取 hostId 失败")
		return
	}

	service, err := models.GetServiceById(serviceId)
	if err != nil {
		c.SetJson(1, err, "获取 service by service id 失败")
		return
	}

	host, err := models.GetHostById(hostId)
	if err != nil {
		c.SetJson(1, err, "获取 host by host id 失败")
		return
	}

	operationRecord := new(models.OperationRecord)
	operationRecord.Service = service
	operationRecord.Host = host
	operationRecord.User = c.User

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
