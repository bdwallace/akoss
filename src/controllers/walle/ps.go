package wallecontrollers

import (
	"library/components"
	"models"
)

// @Title 获取 service所有host上的状态
// @Description 根据 Service.Id 获取service所有的host上的状态
// @Param   service_id      query     int   		true         "service id"
// @Success 0 {list} {["tag", ...]}
// @Failure 1 获取 失败
// @Failure 2 User not found
// @router /walle/dockerps [get]
func (c *WalleController) DockerPs() {

	serviceId, _ := c.GetInt("service_id")
	lineData := c.GetString("line_data", "")

	service, err := models.GetServiceById(serviceId)
	if err != nil {
		c.SetJson(1, err, "获取 service host conf by id失败")
		return
	}

	components.GetServiceNacos(c.Service.Project,service)

	_, err = models.GetServiceAllRelated(service)
	if err != nil {
		c.SetJson(1, err, "获取 service host conf by id失败")
		return
	}

	s := components.BaseComponents{}

	s.SetProject(service.Project)
	s.SetService(service)
	s.SetUser(s.User)
	s.SetTask(s.Task)

	if service.Hosts == nil || len(service.Hosts) == 0 {
		c.SetJson(0, nil, "")
		return
	}
	s.SetHost(service.Hosts)

	d := components.BaseDocker{}
	d.SetBaseComponents(s)

	res, _ := d.DockerPs(lineData)

	c.SetJson(0, res, "")
	return
}
