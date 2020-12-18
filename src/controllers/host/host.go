package hostcontrollers

import (
	"controllers"
	"library/components"
	"models"

	"encoding/json"
	"fmt"
)

type HostController struct {
	controllers.BaseController
}


// @Title 添加 host
// @Description add the host
// @Param   name 			 query     string   	true         "host name"
// @Param   private_ip 		 query     string   	true         "private ip"
// @Param   public_ip 		 query     string   	true         "public ip"
// @Param   use_public 		 query     int		   	true         "use public IP"
// @Param   project_id       query     int   		true         "project id"
// @Success 0 {id} int64
// @Failure 1 添加 host 失败
// @Failure 2 User not found
// @router /host [post]
func (c *HostController)AddHost(){

	//beego.Info(string(c.Ctx.Input.RequestBody))
	var host models.Host
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &host)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}


	if host.Id != 0 {
		_, err = models.UpdateHostById(&host)
	} else {
		_, err = models.AddHost(&host)
	}
	if err != nil {
		if err.Error() == "<QuerySeter> return rows" {
			c.SetJson(1, nil, "命名有重复")
		}
		c.SetJson(1, nil, "数据库更新错误:"+err.Error())
		return
	}
	c.SetJson(0, host, "修改成功")
	return
}



// @Title 获取所有 host
// @Description get all the host
// @Success 0 {object} models.Host
// @Failure 1 获取所有 host 失败
// @Failure 2 User not found
// @router /host/list/ [get]
func (c *HostController)GetAllHost() {

	searchText := c.GetString("search_text")
	awsRegion := c.GetString("aws_region")
	projectId, err := c.GetInt("project_id")

	if err != nil{
		fmt.Println("error: GetInt(\"project_id\")",err)
		c.SetJson(1,err,"获取 project_id 失败")
		return
	}

	resHosts := make([]*models.Host, 0)

	// if searchText == ""{
	// 	resHosts, err = models.GetAllHost(projectId)
	// 	if err != nil{
	// 		fmt.Println("error: GetAllHost()",err)
	// 		c.SetJson(1,err,"获取所有 Host 失败")
	// 		return
	// 	}
	// }else {
		resHosts, err = models.SearchHosts(projectId, awsRegion, searchText)
		if err != nil{
			fmt.Println("error: SearchHosts ",err)
			c.SetJson(1,err,"搜索 host匹配内容 失败")
			return
		}
	// }

	for _, s := range resHosts{
		err = models.GetHostRelatedProject(s)
		if err != nil {
			c.SetJson(1, err, "获取 host 的 porject 关联失败")
			return
		}
		_, err := models.GetHostRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 host 的 services 关联失败")
			return
		}
	}

	c.SetJson(0,resHosts,"")
	return
}




/*
func (c *HostController)GetAllHost() {

	searchText := c.GetString("search_text")
	projectId, err := c.GetInt("project_id")
	if err != nil{
		fmt.Println("error: GetInt(\"project_id\")",err)
		c.SetJson(1,err,"获取 project_id 失败")
		return
	}

	resHosts := make([]*models.Host,0)
	if searchText == ""{
		resHosts, err = models.GetAllHost(projectId)
		if err != nil{
			fmt.Println("error: GetAllHost()",err)
			c.SetJson(1,err,"获取所有 Host 失败")
			return
		}
	}else {
		resHosts, err = models.SearchHosts(searchText,projectId)
		if err != nil{
			fmt.Println("error: SearchHosts()",err)
			c.SetJson(1,err,"搜索 host匹配内容 失败")
			return
		}
	}

	for _, s := range resHosts{
		_, err := models.GetHostRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 conf services by id失败")
			return
		}
	}

	c.SetJson(0,resHosts,"")
	return
}

*/


// @Title 获取 host by id
// @Description get the host by id
// @Param   id      query     int   	true     "host id"
// @Success 0 {object} models.Host
// @Failure 1 获取 host by id 失败
// @Failure 2 User not found
// @router /host/id/ [get]
func (c *HostController)GetHostById() {

	id, err := c.GetInt("id")
	if err != nil{
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	resHost, err := models.GetHostById(id)
	if err != nil{
		c.SetJson(1, err,"获取 host by id失败")
		return
	}

	c.SetJson(0,resHost,"")
	return
}





// @Title 获取 host by projects id
// @Description get the host by projects id
// @Param   projects_id      query     int   		true         "projects id"
// @Success 0 {object} models.Host
// @Failure 1 获取 host by projects id 失败
// @Failure 2 User not found
// @router /host/projectId/ [get]
func (c *HostController)GetHostByProjectsId() {

	projectId, err :=  c.GetInt("project_id")
	if err != nil{
		c.SetJson(1, err, "获取project_id 失败")
		return
	}

	resHost, err := models.GetHostByProjectId(projectId)
	if err != nil{
		c.SetJson(1, err,"获取 host by project id失败")
		return
	}

	for _, s := range resHost{
		_, err := models.GetHostRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 service host conf by id失败")
			return
		}
	}


	c.SetJson(0,resHost,"")
	return
}



// @Title 删除 host host by id
// @Description update the host host by id
// @Param   id      query     int   		true         "host id"
// @Success 0 {ok} bool
// @Failure 1 删除 host host by id 失败
// @Failure 2 User not found
// @router /host/id/ [delete]
func (c *HostController)DeleteHostById(){


	id, err := c.GetInt("id")
	if err != nil{
		c.SetJson(1,err,"获取 id 失败")
		return
	}

	ok, err := models.DeleteHostById(id)
	if err != nil && ok != true{
		c.SetJson(1,err,"删除 host 失败")
		return
	}

	c.SetJson(0,ok,"")
	return

}



// @Title 拷贝 host by id
// @Description copy the host host by id
// @Param   id      query     int   		true         "host id"
// @Success 0 {ok} bool
// @Failure 1 拷贝 host host by id 失败
// @Failure 2 User not found
// @router /host/copy [get]
func (c *HostController) CopyHostById() {

	id, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	resHost, err := models.GetHostById(id)
	if err != nil {
		c.SetJson(1, err, "获取 host host host by id失败")
		return
	}

	// copy host
	resHost.Id = 0
	resHost.Name = fmt.Sprintf("%s - copy", resHost.Name)
	_, err = models.AddHost(resHost)
	if err != nil {
		fmt.Println("error:  ",err)
		c.SetJson(1, err, "添加 host 失败")
		return
	}

	c.SetJson(0, resHost, "")
	return

}



// @Title 获取 host docker ps
// @Description get host docker ps
// @Param   use_ip      query     string   		true         "use_ip"
// @Failure 1 获取 获取 host docker ps 状态 失败
// @Failure 2 User not found
// @router /host/dockerPs/ [get]
func (c *HostController)GetHostDockerPs() {

	useIp :=  c.GetString("use_ip")

	d := components.BaseDocker{}
	res, err := d.GetStatusHost(useIp)
	if err != nil{
		c.SetJson(1, res, fmt.Sprintf("获取 %s docker ps 状态 失败", useIp))
		return
	}
	c.SetJson(0,res,"")
	return
}

