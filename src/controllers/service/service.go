package servicecontroller

import (
	"controllers"
	"encoding/json"
	"fmt"
	"library/alertEmail"
	"library/blacklist"
	"library/components"
	"models"
)

type ServiceController struct {
	controllers.BaseController
}

// @Title 获取 default conf key
// @Description 获取 service default conf value json key
// @Param   project_id      query     int   		true         "project id"
// @Param   host_id	        query     int   		true         "host id"
// @Success 0 {object} models.Service
// @Failure 1 获取所有 default key 失败
// @Failure 2 User not found
// @router /service/defaultKey/ [get]
func (c *ServiceController) GetServiceDefaultConfKay() {

	projectId, err := c.GetInt("project_id")
	if err != nil {
		c.SetJson(1, err, "获取project_id 失败")
		return
	}

	templateConf := make(map[string]string, 0)
	templateConf["port"] = ""

	b, err := json.Marshal(templateConf)
	if err != nil {
		c.SetJson(1, err, "编码 valueMap 失败")
	}
	// 转 string
	valueStr := string(b)

	service := &models.Service{
		Value:   valueStr,
		Project: &models.Project{Id: projectId},
	}

	resId, err := models.AddService(service)
	if err != nil {
		c.SetJson(1, err, "添加 service value default key 失败")
		return
	}
	c.SetJson(0, resId, "")
	return

}

// @Title 添加 service
// @Description add the service
// @Param   name 			 query     string   	true         "service name"
// @Param   port 			 query     string   	true         "service port"
// @Param   image_path 		 query     string   	true         "service image path"
// @Param   value 			 query     string   	true         "service conf value"
// @Param   project_id      query     int   		true         "project id"
// @Success 0 {id} int64
// @Failure 1 添加 service 失败
// @Failure 2 User not found
// @router /service/ [post]
func (c *ServiceController) AddService() {

	//beego.Info(string(c.Ctx.Input.RequestBody))
	service := new(models.Service)
	var id int64
	err := json.Unmarshal(c.Ctx.Input.RequestBody, service)
	if err != nil {
		fmt.Println("error: AddService json.Unmarshal  ", err)
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	if 0 != service.Id {
		// update service
		resId, err := models.UpdateServiceById(service)
		if err != nil {
			fmt.Println("error: UpdateServiceById ", err)
			c.SetJson(1, err, "更新 service 失败")
			return
		}

		c.SetJson(0, resId, "")
		return
	} else {
		// add service
		id, err = models.AddService(service)
		if err != nil {
			fmt.Println("error: AddService   ", err)
			c.SetJson(1, err, "添加 service 失败")
			return
		}
	}

	c.SetJson(0, id, "")
	return

}

// @Title 获取所有 service
// @Description get all the service
// @Success 0 {object} models.Service
// @Failure 1 获取所有 service 失败
// @Failure 2 User not found
// @router /service/list/ [get]
func (c *ServiceController) GetAllServices() {

	projectId, err := c.GetInt("project_id")
	if err != nil {
		c.SetJson(1, err, "获取 project_id 失败")
		return
	}

	resServices, err := models.GetAllServices(projectId)
	if err != nil {
		c.SetJson(1, err, "获取所有 Services 失败")
		return
	}

	c.SetJson(0, resServices, "")
	return
}

// @Title 拷贝 service by id
// @Description copy the service conf by id
// @Param   id      query     int   		true         "service id"
// @Success 0 {ok} bool
// @Failure 1 拷贝 service conf by id 失败
// @Failure 2 User not found
// @router /service/copy [get]
func (c *ServiceController) CopyServiceById() {

	id, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	resService, err := models.GetServiceById(id)
	if err != nil {
		c.SetJson(1, err, "获取 service host conf by id失败")
		return
	}
	_, err = models.GetServiceAllRelated(resService)
	if err != nil {
		c.SetJson(1, err, "获取 service host conf by id失败")
		return
	}

	// copy service
	resService.Id = 0
	resService.Name = fmt.Sprintf("%s - copy", resService.Name)
	_, err = models.AddServiceAndRelated(resService)
	if err != nil {
		fmt.Println("error: AddServiceAndRelated  ", err)
		c.SetJson(1, err, "添加 service 失败")
		return
	}

	c.SetJson(0, resService, "")
	return

}

///////	多对多操作   service  <----->  conf

// @Title 获取 service by service id
// @Description 根据 service id 查询所有 service 信息
// @Param   service_id      query     int  		true         "service id"
// @Success 0 {object} models.Service
// @Failure 1 删除 service by service id 失败
// @Failure 2 User not found
// @router /service/id/ [get]
func (c *ServiceController) GetServiceRelatedById() {

	id, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	resService, err := models.GetServiceById(id)
	if err != nil {
		c.SetJson(1, err, "获取 service host conf by id失败")
		return
	}

	_, err = models.GetServiceAllRelated(resService)
	if err != nil {
		c.SetJson(1, err, "获取 service host conf by id失败")
		return
	}

	c.SetJson(0, resService, "")
	return
}

// @Title 获取 service by project id
// @Description 根据 project id 查询所有 service 信息
// @Param   project_id      query     int  		true         "project id"
// @Success 0 {object} models.Service
// @Failure 1 获取 service by project id 失败
// @Failure 2 User not found
// @router /service/backend/ [get]
func (c *ServiceController) GetServiceAllRelatedByProjectId() {

	searchText := c.GetString("search_text")
	projectId, err := c.GetInt("project_id")
	if err != nil {
		fmt.Println("error: GetProjectId  ", err)
		c.SetJson(1, err, "获取 project_id 失败")
		return
	}

	resServices := make([]*models.Service, 0)

	if searchText == "" {
		resServices, err = models.GetServiceClassOfJavaByProjectId(projectId)
		if err != nil {
			c.SetJson(1, err, "获取 services by project id and class 失败")
			return
		}

	} else {
		resServices, err = models.SearchBackendServices(searchText, projectId)
		if err != nil {
			fmt.Println("error: SearchService", err)
			c.SetJson(1, err, "搜索 Services 匹配内容 失败")
			return
		}
	}

	for _, service := range resServices {
		_, err = models.GetServiceAllRelated(service)
		if err != nil {
			c.SetJson(1, err, "获取 service host by id失败")
			return
		}
	}

	c.SetJson(0, resServices, "")
	return
}

// @Title 获取 service by project id
// @Description 根据 project id 查询所有 service 信息
// @Param   project_id      query     int  		true         "project id"
// @Success 0 {object} models.Service
// @Failure 1 获取 service by project id 失败
// @Failure 2 User not found
// @router /service/platformId/ [get]
func (c *ServiceController) GetServiceAllRelatedByPlatformId() {

	//searchText := c.GetString("search_text")

	platformId, err := c.GetInt("platform_id")
	if err != nil {
		fmt.Println("error: GetplatformId  ", err)
		c.SetJson(1, err, "获取 platform_id 失败")
		return
	}

	platform, err := models.GetPlatformById(platformId)
	if err != nil {
		c.SetJson(1, err, "获取 platform by id 失败")
		return
	}

	resServices := make([]*models.Service, 0)
	tmpPlatform := new(models.Platform)

	tmpPlatform, err = models.GetServiceByClassAndPlatform(platform)
	if err != nil {
		c.SetJson(1, err, "获取 services by platform id and class 失败")
		return
	}

	platform, err = models.GetPlatformAndDomainRelated(platform)
	if err != nil {
		c.SetJson(1, err, "获取 services by platform id and class 失败")
		return
	}

	for _, service := range tmpPlatform.Services {
		tmpService, err := models.GetServiceAllRelated(service)
		if err != nil {
			c.SetJson(1, err, "获取 services related 失败")
			return
		}
		resServices = append(resServices, tmpService)
	}

	/*
		添加平台 以及平台所关联的前端服务搜索
		resServices, err = models.SearchFrontendServices(searchText,projectId,platformId)
		if err != nil{
			fmt.Println("error: SearchService()",err)
			c.SetJson(1,err,"搜索 Services 匹配内容 失败")
			return
		}
	*/

	c.SetJson(0, resServices, "")
	return
}

// @Title  service by project_id
// @Description service by project_id for name
// @Param   project_id      query     int   		true         "project id"
// @Success 0 {ok} bool
// @Failure 1 查询每套环境里service的name
// @Failure 2 User not found
// @router /service/name [get]
func (c *ServiceController) GetServiceByProjectForName() {

	projectId, err := c.GetInt("project_id")
	if err != nil {
		c.SetJson(1, err, "获取project_id 失败")
		return
	}

	data, err := models.GetServiceByProjectNotJavaForName(projectId)
	if err != nil {
		c.SetJson(1, err, "获取 domain by project id失败")
		return
	}

	c.SetJson(0, data, "")
	return
}

// @Title 添加 service 以及多对多关联表数据
// @Description 添加 service 以及多对多关联表数据
// @Param   name 			 query     string   	true         "service name"
// @Param   port 			 query     string   	true         "service port"
// @Param   image_path 		 query     string   	true         "service image path"
// @Param   value 			 query     string   	false        "service conf value"
// @Param   project_id       query     int  		true         "project id"
// @Param   host	         query     object  		false        "host id"
// @Param   service	         query     object  		false        "host id"
// @Success 0 {id} int
// @Failure 1 添加 service 以及多对多关联表数据 失败
// @Failure 2 User not found
// @router /service/related/ [post]
func (c *ServiceController) AddServiceAllRelatedByService() {

	// 1. 获取数据
	//beego.Info(string(c.Ctx.Input.RequestBody))
	//var service models.Service
	service := new(models.Service)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, service)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	// 2. 校验数据
	if service.Hosts == nil {
		c.SetJson(1, nil, "host 为 nil,  获取 service.Hosts 失败")
		return
	}
	if service.Confs == nil {
		c.SetJson(1, nil, "conf 为 nil,  获取 service.Confs 失败")
		return
	}

	// 3. 处理更新或添加
	if 0 != service.Id {
		// update
		err := models.UpdateServiceAndRelated(service)
		if err != nil {
			fmt.Println("error: UpdateServiceAndRelated  ", err)
			c.SetJson(1, nil, "更新 service 多对多关系 失败")
			return
		}

	} else {
		// add
		_, err = models.AddServiceAndRelated(service)
		if err != nil {
			fmt.Println("error: AddServiceAndRelated ", err)
			c.SetJson(1, nil, "添加 service 多对多关系 失败")
			return
		}
	}

	blackList := false
	denyUserAgent := false
	if service.BlackList != "" {
		blackList = true
	}
	if service.DenyUserAgent != "" {
		denyUserAgent = true
	}

	if blackList || denyUserAgent {
		if cmdErr := blacklist.BlackListAndDenyUserAgentCP(service, blackList, denyUserAgent); cmdErr != nil {
			fmt.Println("error: service BlackListAndDenyUserAgentCP ", cmdErr)
			c.SetJson(1, nil, "更新 service 黑名单 || UA 失败")
			return
		}
	}

	// 4. 返回数据
	c.SetJson(0, service, "")
	return
}

// @Title 删除 service by id
// @Description delete the service conf by id
// @Param   id      query     int   		true         "service id"
// @Success 0 {ok} bool
// @Failure 1 删除 service conf by id 失败
// @Failure 2 User not found
// @router /service/id/ [delete]
func (c *ServiceController) DeleteServiceAndAllRelatedByServiceId() {

	serviceId, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	service, err := models.GetServiceById(serviceId)
	if err != nil {
		c.SetJson(1, err, "根据 serviceId 获取 service 失败")
		return
	}

	service, err = models.GetServiceAllRelated(service)
	if err != nil {
		c.SetJson(1, err, "根据 service 获取 service related 失败")
		return
	}

	err = models.DeleteServiceAndAllRelatedByService(service)
	if err != nil {
		c.SetJson(1, err, "删除 service 以及关联关系 失败")
		return
	}

	service.IsDel = 1
	_, err = models.UpdateServiceById(service)
	if err != nil {
		c.SetJson(1, err, "删除 service 失败")
		return
	}

	c.SetJson(0, nil, "")
	return
}


// @Title  service by project_id
// @Description service by project_id for name
// @Param   project_id      query     int   		true         "project id"
// @Success 0 {ok} bool
// @Failure 1 查询每套环境里service的name
// @Failure 2 User not found
// @router /service/backendHealth [get]
func (c *ServiceController) GetBackendServiceHealth() {

	projectId, err := c.GetInt("project_id")
	if err != nil {
		c.SetJson(1, err, "获取project_id 失败")
		return
	}

	// get backend service by project id
	services, err := models.GetServiceClassOfJavaByProjectId(projectId)
	if err != nil{
		c.SetJson(1, err, "获取 service 失败")
		return
	}

	unhealthyServices := make([]map[string]string,0)
	// check service health
	for _, s := range services {
		service, err := models.GetServiceById(s.Id)
		if err != nil {
			c.SetJson(1, err, "获取 service host conf by id失败")
			return
		}

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
			continue
		}
		s.SetHost(service.Hosts)
		d := components.BaseDocker{}
		d.SetBaseComponents(s)

		res, _ := d.DockerPs("")
		for _, item := range res{
			if item["health"] != "200"{
				unhealthyServices = append(unhealthyServices, item)
			}
		}
	}

	// if the service is unhealthy, email to the op group
	emailInfo := alertEmail.EmailInfoOfBackendService(unhealthyServices)
	err = alertEmail.SendEmail(emailInfo)
	if err != nil{
		c.SetJson(1, err, "服务告警邮件发送失败")
		return
	}

	c.SetJson(0, nil, "")
	return
}