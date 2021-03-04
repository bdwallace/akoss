package platformcontroller

import (
	"controllers"
	"encoding/json"
	"fmt"
	"models"
)

type PlatformController struct {
	controllers.BaseController
}

// @Title 添加 platform
// @Description add the platform
// @Success 0 {id} int64
// @Failure 1 添加 platform 失败
// @Failure 2 User not found
// @router /platform [post]
func (c *PlatformController)AddPlatform(){

	//beego.Info(string(c.Ctx.Input.RequestBody))
	var platform models.Platform
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &platform)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}


	if platform.Id != 0 {
		_, err = models.UpdatePlatformById(&platform)
	} else {
		_, err = models.AddPlatform(&platform)
	}
	if err != nil {
		if err.Error() == "<QuerySeter> return rows" {
			c.SetJson(1, nil, "命名有重复")
		}
		c.SetJson(1, nil, "数据库更新错误:"+err.Error())
		return
	}
	c.SetJson(0, platform, "修改成功")
	return
}



// @Title 获取所有 platform
// @Description get all the platform
// @Success 0 {object} models.Platform
// @Failure 1 获取所有 platform 失败
// @Failure 2 User not found
// @router /platform/list/ [get]
func (c *PlatformController)GetAllPlatform() {

	searchText := c.GetString("search_text")
	projectId, err := c.GetInt("project_id")
	if err != nil{
		fmt.Println("error: GetInt(\"project_id\")",err)
		c.SetJson(1,err,"获取 project_id 失败")
		return
	}

	resPlatform := make([]*models.Platform,0)
	if searchText == ""{
		resPlatform, err = models.GetAllPlatform(projectId)
		if err != nil{
			fmt.Println("error: GetAllPlatform",err)
			c.SetJson(1,err,"获取所有 Platform 失败")
			return
		}
	}else {
		resPlatform, err = models.SearchPlatform(searchText,projectId)
		if err != nil{
			fmt.Println("error: GetAllPlatform",err)
			c.SetJson(1,err,"搜索 Platform 匹配内容 失败")
			return
		}
	}

	c.SetJson(0,resPlatform,"")
	return
}

// @Title 获取所有 platform
// @Description 获取全部平台以及平台所关联域名， 不需要参数
// @Success 0 {object} models.Platform
// @Failure 1 获取所有 platform 失败
// @Failure 2 User not found
// @router /platform/domains [get]
func (c *PlatformController)GetAllPlatformforMonitor() {

	projects, err := models.GetProjectAll()
	if err != nil{
		c.SetJson(1, err, "获取 all project 失败")
		return
	}



	resPlatform := make([]*models.PlatformDomains,0)
	for _, p := range projects{

		projectAllPlatform, err := models.GetPlatformByProjectId(p.Id)
		if err != nil{
			c.SetJson(1, err, "获取 platform 失败")
			return
		}

		for _, platform := range projectAllPlatform {

			platformDomains := new(models.PlatformDomains)
			platformDomains.Domains =  make([]string,0)
			platformRel, err := models.GetPlatformAndDomainRelated(platform)
			if err != nil{
				continue
			}
			if len(platformRel.Domains) == 0 {
				continue
			}
			platformDomains.PlatformName = platformRel.Name

			for _, domain := range platformRel.Domains{
				platformDomains.Domains = append(platformDomains.Domains,domain.Name)
			}
			resPlatform = append(resPlatform,platformDomains)
		}
	}

	c.SetJson(0,resPlatform,"")
	return
}



// @Title 获取 platform by id
// @Description get the platform by id
// @Param   id      query     int   	true     "platform id"
// @Success 0 {object} models.Platform
// @Failure 1 获取 platform by id 失败
// @Failure 2 User not found
// @router /platform/id/ [get]
func (c *PlatformController)GetPlatformById() {

	id, err := c.GetInt("id")
	if err != nil{
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	resPlatform, err := models.GetPlatformById(id)
	if err != nil{
		c.SetJson(1, err,"获取 platform by id失败")
		return
	}

	_, err = models.GetPlatformRelated(resPlatform)
	if err != nil {
		c.SetJson(1, err, "获取 service host conf by id失败")
		return
	}

	c.SetJson(0,resPlatform,"")
	return
}



// @Title 获取 platform by projects id
// @Description get the platform by projects id
// @Param   projects_id      query     int   		true         "projects id"
// @Success 0 {object} models.Platform
// @Failure 1 获取 platform by projects id 失败
// @Failure 2 User not found
// @router /platform/projectId/ [get]
func (c *PlatformController)GetPlatformByProjectsId() {

	projectId, err :=  c.GetInt("project_id")
	if err != nil{
		c.SetJson(1, err, "获取project_id 失败")
		return
	}

	resPlatform, err := models.GetPlatformByProjectId(projectId)
	if err != nil{
		c.SetJson(1, err,"获取 platform by project id失败")
		return
	}

	for _, s := range resPlatform {
		_, err := models.GetPlatformRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 service host conf by id失败")
			return
		}
	}

	c.SetJson(0,resPlatform,"")
	return
}


// @Title platform by project_id
// @Description platform by project_id for name
// @Param   project_id      query     int   		true         "project id"
// @Success 0 {ok} bool
// @Failure 1 查询每套环境里platform的name
// @Failure 2 User not found
// @router /platform/name [get]
func (c *PlatformController)GetPlatformByProjectForName(){

	projectId, err :=  c.GetInt("project_id")
	if err != nil{
		c.SetJson(1, err, "获取project_id 失败")
		return
	}

	data, err := models.GetPlatformByProjectForName(projectId)
	if err != nil{
		c.SetJson(1, err,"获取 domain by project id失败")
		return
	}

	c.SetJson(0, data, "")
	return
}



// @Title 删除 platform by id
// @Description delete the platform by id
// @Param   id      query     int   		true         "platform id"
// @Success 0 {ok} bool
// @Failure 1 删除 platform by id 失败
// @Failure 2 User not found
// @router /platform/relatedCheck/ [get]
func (c *PlatformController)DeletePlatformCheck(){

	id, err := c.GetInt("id")
	if err != nil{
		c.SetJson(1,err,"获取 id 失败")
		return
	}

	platform, err := models.GetPlatformById(id)
	if err != nil{
		c.SetJson(1, err, "根据 platformId 获取 platform 失败")
		return
	}

	platform, err = models.GetPlatformAndServiceRelated(platform)
	if err != nil{
		c.SetJson(1, err, "根据 platform 获取 platform 和 service related 失败")
		return
	}
	if len(platform.Services) != 0 {
		c.SetJson(1, platform.Services, "该平台关联服务未删除，请将关联服务删除或取消关联后再删除该平台")
		return
	}


	platform, err = models.GetPlatformAndDomainRelated(platform)
	if err != nil{
		c.SetJson(1, err, "根据 platform 获取 platform 和 domain related 失败")
		return
	}
	if len(platform.Domains) != 0 {
		c.SetJson(0, platform.Domains, "该平台关联以下域名，是否全部删除?")
		return
	}

	c.SetJson(0,nil,"")
	return
}




// @Title 删除 platform by id
// @Description delete the platform by id
// @Param   id      query     int   		true         "platform id"
// @Success 0 {ok} bool
// @Failure 1 删除 platform by id 失败
// @Failure 2 User not found
// @router /platform/id/ [delete]
func (c *PlatformController)DeletePlatformId(){

	id, err := c.GetInt("id")
	if err != nil{
		c.SetJson(1,err,"获取 id 失败")
		return
	}

	platform, err := models.GetPlatformById(id)
	if err != nil{
		c.SetJson(1, err, "根据 platformId 获取 platform 失败")
		return
	}

	platform, err = models.GetPlatformAndDomainRelated(platform)
	if err != nil{
		c.SetJson(1, err, "根据 platform 获取 platform 和 domain related 失败")
		return
	}
	if len(platform.Domains) != 0 {
		if err = models.DeletePlatformAndDomainRelatedByPlatformId(platform); err != nil{
			c.SetJson(1, err, "根据 platform 删除 platform 和 domain related 失败")
			return
		}

		for _, domain := range platform.Domains {
			if _, err = models.DeleteDomainById(domain.Id); err != nil{
				c.SetJson(1, err, "删除 domain 失败")
				return
			}
		}
	}

	num, err := models.DeletePlatformById(id)
	if err != nil {
		c.SetJson(1,err,"删除 platform 失败")
		return
	}

	c.SetJson(0,num,"")
	return
}





///////  多对多操作  platform <---> service

// @Title 获取 platform by platform id
// @Description 根据 platform id 查询所有 service 信息
// @Param   platform_id      query     int  		true         "platform id"
// @Success 0 {object} models.Service
// @Failure 1 获取 platform and service related 失败
// @Failure 2 User not found
// @router /platformAndService/platformId/ [get]
func (c *PlatformController) GetPlatformAndServiceRelatedByPlatformId(){

	id, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	resPlatform, err := models.GetPlatformById(id)
	if err != nil {
		c.SetJson(1, err, "获取 platform by id 失败")
		return
	}

	_, err = models.GetPlatformAndServiceRelated(resPlatform)
	if err != nil {
		c.SetJson(1, err, "获取 platform and service related by platform 失败")
		return
	}

	c.SetJson(0, resPlatform, "")
	return
}



// @Title 添加 platform and service 多对多关联表数据
// @Description 添加 platform and service 多对多关联表数据
// @Success 0 {id} int
// @Failure 1 添加 platform and service 多对多关联表数据 失败
// @Failure 2 User not found
// @router /platformAndService/related/ [post]
func (c *PlatformController) AddPlatformAndServiceRelated(){

	// 1. 获取数据
	//beego.Info(string(c.Ctx.Input.RequestBody))
	//var service models.Service
	platform := new(models.Platform)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, platform)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	// 2. 校验数据
	if len(platform.Services) == 0{
		c.SetJson(1, nil, "host 为 nil,  获取 service.Hosts 失败")
		return
	}

	// 3. 处理更新或添加
	if 0 != platform.Id {
		// update
		err := models.UpdatePlatformAndServicesRelated(platform)
		if err != nil{
			fmt.Println(" error: ",err)
			c.SetJson(1,nil,"更新 platform and service 多对多关系 失败")
			return
		}
	}else {
		// add
		_, err = models.AddPlatformAndServiceRelated(platform)
		if err != nil{
			fmt.Println("error: AddPlatformAndServiceRelated  ",err)
			c.SetJson(1,nil,"添加 platform and service 多对多关系 失败")
			return
		}
	}

	// 4. 返回数据
	c.SetJson(0, platform, "")
	return
}


// @Title 删除 platform and service related by platform id
// @Description delete the platform and service by related id
// @Param   id      query     int   		true         "platform id"
// @Success 0 {ok} bool
// @Failure 1 删除 platform and service related by id 失败
// @Failure 2 User not found
// @router /platformAndService/id/ [delete]
func (c *PlatformController) DeletePlatformAndServiceRelatedByPlatformId() {

	platformId, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	platform, err := models.GetPlatformById(platformId)
	if err != nil{
		c.SetJson(1, err, "根据 platformId 获取 platform 失败")
		return
	}

	err = models.DeletePlatformAndServiceRelatedByPlatformId(platform)
	if err != nil  {
		c.SetJson(1, err, "删除 platform and service 关联关系 失败")
		return
	}

	c.SetJson(0,nil , "")
	return
}
