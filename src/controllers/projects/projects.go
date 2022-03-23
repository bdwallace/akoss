package projectcontrollers

import (
	"controllers"
	"encoding/json"
	"fmt"
	"models"
	"models/response"
	"strconv"
)

type ProjectController struct {
	controllers.BaseController
}

func (c *ProjectController) URLMapping() {
	/*
		c.Mapping("list",c.GetAllProject)
		c.Mapping("GetProjectById",c.GetProjectById)
		c.Mapping("GetProjectByAlias",c.GetProjectByAlias)
		c.Mapping("AddProject",c.AddProject)
		c.Mapping("UpdateProject",c.UpdateProject)
		c.Mapping("DeleteProjectById",c.DeleteProjectById)
	*/
}

/////  project

// @Title 添加project
// @Description add the project
// @Param   name     query     string   true        "project name"
// @Param   alias    query     string  	true        "project alias"
// @Success 0 {id} int64
// @Failure 1 添加 project 失败
// @Failure 2 User not found
// @router /project/ [post]
func (c *ProjectController) AddProject() {

	//beego.Info(string(c.Ctx.Input.RequestBody))
	var project models.Project
	var resId int64
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &project)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	if 0 != project.Id {
		// update project
		resId, err := models.UpdateProject(&project)
		if err != nil {
			fmt.Println("error: AddProject  ", err)
			c.SetJson(1, err, "更新 project 失败")
			return
		}
		c.SetJson(0, resId, "")
		return
	} else {

		// add	project
		resId, err = models.AddProject(&project)
		if err != nil {
			fmt.Println("error: AddProject  ", err)
			c.SetJson(1, err, "添加 project 失败")
			return
		}
	}

	/*
		创建project点击next 发送请求,
		初始化projectConf到mysql并将数据返回前端
	*/
	if err := models.InitConfKey(int(resId)); err != nil {
		fmt.Println("error: InitConfKey ", err)
		c.SetJson(1, err, "创建 project 模板key失败")
		return
	}

	c.SetJson(0, resId, "")
	return

}

// @Title 获取所有project
// @Description get all the project
// @Success 0 {object} models.Project
// @Failure 1 获取所有 project 失败
// @Failure 2 User not found
// @router /project/list/ [get]
func (c *ProjectController) GetAllProject() {

	proRes, err := models.GetAllProject()
	if err != nil {
		fmt.Println("error: GetAllProject ", err)
		c.SetJson(1, nil, "获取所有 project 失败")
		return
	}
	projectListResp := new(response.ProjectListForLoginResponse)
	for _, item := range proRes {
		tempPro := new(response.Project)

		tempPro.Id = item.Id
		tempPro.Name = item.Name
		tempPro.Alias = item.Alias
		tempPro.Nacos1 = item.Nacos1
		tempPro.Nacos2 = item.Nacos2
		tempPro.Nacos3 = item.Nacos3
		tempPro.AwsKeyId = item.AwsKeyId
		tempPro.AwsKeySecret = item.AwsKeySecret
		tempPro.AwsRegion = item.AwsRegion
		tempPro.CreatedAt = item.CreatedAt
		tempPro.UpdatedAt = item.UpdatedAt

		projectListResp.ProjectList = append(projectListResp.ProjectList, tempPro)
		projectListResp.ProjectNameList = append(projectListResp.ProjectNameList, item.Name)
	}

	c.SetJson(0, projectListResp, "")
	return

}

// @Title 获取project by id
// @Description get the project by id
// @Param   id    query     int  	true        "project id"
// @Success 0 {object} models.Project
// @Failure 1 获取 project by id 失败
// @Failure 2 User not found
// @router /project/id/ [get]
func (c *ProjectController) GetProjectById() {

	id := c.Input().Get("id")
	intId, _ := strconv.Atoi(id)

	proRes, err := models.GetProjectById(intId)
	if err != nil {
		fmt.Println("error: GetProjectById ", err)
		c.SetJson(1, err, "获取 project by id 失败")
		return
	}

	c.SetJson(0, proRes, "")
	return

}

// @Title 获取project
// @Description get the project by alias
// @Param   alias    query     string  	true        "project alias"
// @Success 0 {object} models.Project
// @Failure 1 获取 project by alias 失败
// @Failure 2 User not found
// @router /project/alias/ [get]
func (c *ProjectController) GetProjectByAlias() {

	alias := c.GetString("alias")
	if alias == "" {
		c.SetJson(1, nil, "project alias is empty")
		return
	}

	proRes, err := models.GetProjectByAlias(alias)
	if err != nil {
		fmt.Println("error: GetProjectByAlias ", err)
		c.SetJson(1, err, "获取 project by alias 失败")
		return
	}

	c.SetJson(0, proRes, "")
	return
}

// @Title 删除project
// @Description delete the project
// @Param   id     	 query     int   	true        "project id"
// @Success 0 {ok} bool
// @Failure 1 删除 project 失败
// @Failure 2 User not found
// @router /project/id/ [delete]
func (c *ProjectController) DeleteProjectById() {

	id := c.Input().Get("id")
	intId, _ := strconv.Atoi(id)
	ok, err := models.DeleteProjectById(intId)
	if err != nil && ok != true {
		fmt.Println("error: delete project  ", err)
		c.SetJson(1, err, "删除 project 失败")
		return
	}

	c.SetJson(0, ok, "")
	return
}

// @Title 获取 nacos by projectId
// @Description get the nacos by projectId
// @Param   id    query     int  	true        "project id"
// @Success 0 {object} string
// @Failure 1 获取 nacos by projectId 失败
// @Failure 2 User not found
// @router /project/nacos/ [get]
func (c *ProjectController) GetNacosByProjectId() {

	id, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, nil, "project id is empty")
		return
	}

	nacos, err := models.GetNacosByProjectId(id)
	if err != nil {
		fmt.Println("error: GetNacosByProjectId ", err)
		c.SetJson(1, err, "获取 nacos by project id 失败")
		return
	}

	c.SetJson(0, nacos, "")
	return
}
