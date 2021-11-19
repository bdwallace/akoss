package deploycontrollers

import (
	"controllers"
	"encoding/json"
	"fmt"
	"models"
)

type DeployController struct {
	controllers.BaseController
	//Coment components.BaseComponents
	Task     *models.Task
	TaskUser *models.User
}

///////////
/*
	deploy  curd
*/

// @Title 获取 deploy by Id
// @Description 获取 deploy by Id
// @Param   id       query     int  	true       "deploy id"
// @Success 0 {object} models.Deploy
// @Failure 1 获取 deploy by Id 失败
// @Failure 2 User not found
// @router /deploy/id/ [get]
func (c *DeployController) GetDeployById() {

	deployId, _ := c.GetInt("id")
	deploy, err := models.GetDeployById(deployId)
	if err != nil {
		c.SetJson(1, nil, "获取 deploy by Id 失败")
		return
	}

	err = models.GetDeployRelated(deploy)
	if err != nil {
		c.SetJson(1, err, "获取 service by id 失败")
		return
	}

	c.SetJson(0, deploy, "")
	return
}

// @Title 获取 deploys by project id
// @Description get deploys by project id
// @Param   project_id      query     int  	  true      "project id"
// @Success 0 {object} models.Deploy
// @Failure 1 获取 deploy by Id 失败
// @Failure 2 User not found
// @router /deploy/projectId/userId [get]
func (c *DeployController) GetDeployByProjectIdUserId() {

	projectId, err := c.GetInt("project_id")
	if err != nil {
		c.SetJson(1, err, "获取 project id 失败")
		return
	}

	userId, err := c.GetInt("user_id", 0)
	if err != nil {
		c.SetJson(1, err, "获取 user id 失败")
		return
	}

	page, _ := c.GetInt("page", 0)
	start := 0
	length, _ := c.GetInt("length", 200000)
	if page > 0 {
		start = (page - 1) * length
	}

	count, resDeploy, err := models.GetDeployByProjectIdUserId(projectId, userId, start, length)
	if err != nil {
		c.SetJson(1, err, "获取 host by project id失败")
		return
	}

	for _, s := range resDeploy {
		err = models.GetDeployRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 service by id 失败")
			return
		}
	}

	c.SetJson(0, map[string]interface{}{"total": count, "currentPage": page, "table_data": resDeploy}, "")

	return
}

// @Title 获取 deploys by project id
// @Description get deploys by project id
// @Param   project_id      query     int  	  true      "project id"
// @Success 0 {object} models.Deploy
// @Failure 1 获取 deploy by Id 失败
// @Failure 2 User not found
// @router /deploy/projectId/ [get]
func (c *DeployController) GetDeployByProjectId() {

	projectId, err := c.GetInt("project_id")
	if err != nil {
		c.SetJson(1, err, "获取 project id 失败")
		return
	}

	resDeploy, err := models.GetDeployByProjectId(projectId)
	if err != nil {
		c.SetJson(1, err, "获取 host by project id失败")
		return
	}

	for _, s := range resDeploy {
		err = models.GetDeployRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 service by id 失败")
			return
		}
	}

	c.SetJson(0, resDeploy, "")
	return
}

// @Title 获取 deploys by service id
// @Description get deploys by service id
// @Param   service_id      query     int  	  true      "service id"
// @Success 0 {object} models.Deploy
// @Failure 1 获取 deploy by Id 失败
// @Failure 2 User not found
// @router /deploy/serviceId/ [get]
func (c *DeployController) GetDeployByServiceId() {

	serviceId, err := c.GetInt("service_id")
	if err != nil {
		c.SetJson(1, err, "获取 service id 失败")
		return
	}

	resDeploy, err := models.GetDeployByServiceId(serviceId)
	if err != nil {
		c.SetJson(1, err, "获取 host by service id失败")
		return
	}

	c.SetJson(0, resDeploy, "")
	return
}

// @Title 获取所有 deploy
// @Description get all the deploy
// @Success 0 {object} models.Deploy
// @Failure 1 获取所有 deploy 失败
// @Failure 2 User not found
// @router /deploy/list/ [get]
func (c *DeployController) GetAllDeploy() {

	resDeploy, err := models.GetAllDeploy()

	if err != nil {
		fmt.Println("error: GetAllDeploy()", err)
		c.SetJson(1, err, "获取所有 deploy 失败")
		return
	}

	c.SetJson(0, resDeploy, "")
	return

}

// @Title delete deploy by deploy id
// @Description delete deploy by deploy id
// @Param   id      query     int 		true       "deploy id"
// @Success 0 {object} models.Deploy
// @Failure 1 删除 deploy by deploy id 失败
// @Failure 2 User not found
// @router /deploy/id/ [delete]
func (c *DeployController) DeleteDeploy(id int) {

	resId, err := models.DeleteDeploy(id)
	if err != nil {
		c.SetJson(1, err, "删除 deploy by id 失败")
		return
	}

	c.SetJson(0, resId, "")
	return

}

// @Title 添加空白 deploy
// @Description 添加没有tag和host的deploy
// @Param servcies body []models.Service ture
// @Success 0 {id} int
// @Failure 1 添加 deploy 失败
// @Failure 2 User not found
// @router /deploy [post]
func (c *DeployController) AddDeploy() {

	//beego.Info(string(c.Ctx.Input.RequestBody))
	var services []models.Service
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &services)
	if err != nil {
		c.SetJson(1, nil, "数据格式解析错误")
		return
	}

	deploy := new(models.Deploy)

	deploy.User = c.User
	deploy.Class = services[0].Class
	deploy.Project = &models.Project{Id: c.User.ProjectId}

	deployId, err := models.AddDeploy(deploy)
	if err != nil {
		c.SetJson(1, nil, "新建 deploy 失败!")
		return
	}

	for _, s := range services {
		task := new(models.Task)
		task.Deploy = &models.Deploy{Id: int(deployId)}
		task.Service = &s
		task.Project = s.Project
		task.Hosts = s.Hosts
		task.User = deploy.User

		// if s.Class != deploy.Class {
		// 	c.SetJson(1, nil, "新建 tasks 失败! service class 不统一")
		// 	return
		// }

		_, err = models.AddTask(task)
		if err != nil {
			c.SetJson(1, nil, "创建 docker cmd 插入 deploy 相关 task cmd 失败")
			return
		}
	}

	c.SetJson(0, deploy.Id, "")
	return

}

// @Title 添加空白 deploy
// @Description 添加没有tag和host的deploy
// @Param servcies body []models.Service ture
// @Success 0 {id} int
// @Failure 1 添加 deploy 失败
// @Failure 2 User not found
// @router /deploy/build [post]
func (c *DeployController) AddDeployFromBuild() {

	//beego.Info(string(c.Ctx.Input.RequestBody))
	var service models.Service
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &service)
	if err != nil {
		c.SetJson(1, nil, "数据格式解析错误")
		return
	}

	deploy := new(models.Deploy)

	deploy.User = c.User
	deploy.Class = service.Class
	deploy.Project = &models.Project{Id: c.User.ProjectId}

	deployId, err := models.AddDeploy(deploy)
	if err != nil {
		c.SetJson(1, nil, "新建 deploy 失败!")
		return
	}

	task := new(models.Task)
	task.Deploy = &models.Deploy{Id: int(deployId)}
	task.Tag = service.Tag
	task.Service = &service
	task.Project = service.Project
	task.Hosts = service.Hosts
	task.User = deploy.User

	_, err = models.AddTask(task)
	if err != nil {
		c.SetJson(1, nil, "创建 docker cmd 插入 deploy 相关 task cmd 失败")
		return
	}

	//  创建 发布 命令
	deployPlatform := new(models.Platform)
	if deploy.Class != "java" {
		// 初始化 平台
		s, err := models.GetServiceAndPlatformRelated(&service)
		if err != nil {
			c.SetJson(1, nil, "获取 前端服务所属平台 失败")
			return
		}
		deployPlatform, err = models.GetPlatformById(s.Platforms[0].Id)
		if err != nil {
			c.SetJson(1, nil, "获取 前端服务所属平台 失败")
			return
		}
	}
	docker := c.initComponents(task)
	docker.BaseComponents.SetProject(task.Project)
	docker.BaseComponents.SetPlatform(deployPlatform)
	docker.BaseComponents.SetService(task.Service)

	err = docker.CreateDockerCmd(task, deploy.Count, deploy.Class)
	if err != nil {
		c.SetJson(1, nil, "创建 docker cmd 失败")
		return
	}
	task.Cmd = docker.Cmds
	_, err = models.UpdateTaskById(task)
	if err != nil {
		c.SetJson(1, nil, "创建 docker cmd 插入 deploy 相关 task cmd 失败")
		return
	}

	c.SetJson(0, deploy.Id, "")
	return

}

// @Title 添加tagcmd deploy
// @Description 添加有tag会生成docker命令,但没有确实要发布哪些主机
// @Param deploy body models.Deploy true
// @Success 0 {id} int
// @Failure 1 添加 deploy 失败
// @Failure 2 User not found
// @router /deploy/tagcmd [post]
func (c *DeployController) AddDeployTagCmd() {

	//beego.Info(string(c.Ctx.Input.RequestBody))
	var deploy models.Deploy
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &deploy)
	if err != nil {
		c.SetJson(1, nil, "数据格式解析错误")
		return
	}

	for _, task := range deploy.Tasks {
		task.Hosts = task.Service.Hosts
		docker := c.initComponents(task)
		if len(task.Service.Platforms) == 1 && task.Service.Class != "java" {
			docker.BaseComponents.SetPlatform(task.Service.Platforms[0])
		}

		err := docker.CreateDockerCmd(task, deploy.Count, task.Service.Class)
		if err != nil {
			c.SetJson(1, nil, "创建 docker cmd 失败")
			return
		}
		task.Cmd = docker.Cmds
		_, err = models.UpdateTaskById(task)
		if err != nil {
			c.SetJson(1, nil, "创建 docker cmd 插入 deploy 相关 task cmd 失败")
			return
		}

		// 更新服务最新tag
		if err := c.UpdateServiceTagAndLastTag(task); err != nil {
			c.SetJson(1, nil, "发布 service 更新 tag 失败")
			return
		}
	}

	c.SetJson(0, deploy, "")
	return

}

////// 以服务发布
// @Title 发布
// @Description 添加 deploy
// @Param deploy body models.Deploy true
// @Success 0 {id} int
// @Failure 1 添加 deploy 失败
// @Failure 2 User not found
// @router /deploy/service [post]
func (c *DeployController) DeployService() {

	//fmt.Println("deploy service...")
	var deploy *models.Deploy
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &deploy)
	if err != nil {
		c.SetJson(1, nil, "数据格式解析错误")
		return
	}
	c.DeployCheck(deploy)

	_, err = models.UpdateDeployById(deploy)
	if err != nil {
		c.SetJson(1, nil, "UpdateTaskMulit status errr")
		return
	}

	ch := make(chan int, 0)
	//  多个 task 需要并发
	for _, t := range deploy.Tasks {

		t.Count = deploy.Count
		if _, err := models.UpdateTaskById(t); err != nil {
			c.SetJson(1, nil, "修改 task count 失败")
			return
		}

		go c.releaseHandling(ch, t, deploy, t.Hosts)
		//  需要添加 goroutine 间同步机制 保证
	}
	var res int
	for {
		res = res + <-ch
		if res == len(deploy.Tasks) {
			break
		}
		continue
	}

	deploy.IsRun = 0
	_, err = models.UpdateDeployById(deploy)
	if err != nil {
		c.SetJson(1, nil, "UpdateTaskMulit status errr")
		return
	}

	c.SetJson(0, deploy.Count, "")
	return
}
