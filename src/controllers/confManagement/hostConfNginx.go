package confManagement

import (
	"controllers"
	"fmt"
	"library/confManagement"
	"models"
	"models/request"
	"models/response"
)

type HostConfNGController struct {
	controllers.BaseController
}


// @Title get nginx confs
// @Description get all of the nginx confs
// @Param   projectId 	      query     int   	  true     "project id"
// @Success 0 {obj} response
// @Failure 1 获取所有 host nginx confs 失败
// @Failure 2 User not found
// @router /hostConf/nginxConfs [get]
func (c *HostConfNGController) GetAllNginxConf(){

	projectId, err := c.GetInt("project_id")
	if err != nil {
		fmt.Println("error: GetInt(\"project_id\")",err)
		c.SetJson(1,err,"获取 project_id 失败")
		return
	}

	p, err := models.GetProjectById(projectId)
	if err != nil{
		fmt.Println("error: GetProjectById",err)
		c.SetJson(1,err,"获取 project 失败")
		return
	}

	producer := new(confManagement.Producer)
	producer.Project = new(request.ProjectConfNginx)
	producer.Project.ProjectId = p.Id
	producer.Project.ProjectAlias = p.Alias
	if err = producer.GetAllFrontServicesNginxConfToLocal(); err != nil{
		fmt.Println("error:  GetAllFrontServicesNginxConfToLocal is failed  ::  ",err)
		c.SetJson(1,err,"copy conf to local 失败")
		return
	}


	c.SetJson(0,nil,"")
	return
}



// @Title search nginx confs
// @Description search of the nginx confs
// @Param   projectId 	      query     int   	  true     "project id"
// @Param   platformId 	      query     int   	  true     "platform id"
// @Param   serviceClass      query     string    true     "service class"
// @Param   hostId            query     int       true     "host id"
// @Success 0 {obj} response
// @Failure 1 search host nginx confs 失败
// @Failure 2 User not found
// @router /hostConf/searchNginx [get]
func (c *HostConfNGController) SearchNginxConf(){


	projectId, _ := c.GetInt("project_id")
	platformId, _ := c.GetInt("platform_id")
	serviceClass := c.GetString("service_class")
	hostId, _:= c.GetInt("host_id")


	p, err := models.GetProjectById(projectId)
	if err != nil{
		fmt.Println("error: GetProjectById",err)
		c.SetJson(1,err,"获取 project 失败")
		return
	}
	producer := new(confManagement.Producer)
	producer.Project = new(request.ProjectConfNginx)
	producer.Project.ProjectId = p.Id
	producer.Project.ProjectAlias = p.Alias

	if err = producer.GetAllFrontServicesNginxConfToLocal(); err != nil{
		fmt.Println("error:  GetAllFrontServicesNginxConfToLocal is failed  ::  ",err)
		c.SetJson(1,err,"copy conf to local 失败")
		return
	}

	consumer := new(confManagement.Consumer)

	hostConfs, err := consumer.SearchNGConf(platformId,serviceClass,hostId)
	if len(hostConfs) == 0 && err != nil && err.Error() == "no search options"{
		fmt.Println("error:  SearchNGConf(platformId,serviceClass,hostId)  ::  ",err)
		c.SetJson(1,err,"请选择查询项:平台、服务类型、主机")
		return
	}
	if err != nil {
		fmt.Println("error:  SearchNGConf(platformId,serviceClass,hostId)  ::  ",err)
		c.SetJson(1,err,"查询失败，请联系管理员")
		return
	}
	if len(hostConfs) == 0 {
		fmt.Println("error:  SearchNGConf(platformId,serviceClass,hostId)  ::  ",err)
		c.SetJson(1,err,"很遗憾,没有相关主机配置")
		return
	}

	resp := make([]*response.ConfFile, 0)
	for _, conf := range hostConfs{
		//fmt.Printf("name: %s   class: %s    path: %s \n",c.ConfName, c.Class, c.LocalPath)
		confFiles, err := consumer.GetNginxConf(conf)
		if err != nil{
			fmt.Println("error:  GetNginxConf(conf)  ::  ",err)
			c.SetJson(1,err,"获取配饰文件失败，请联系管理员")
			return
		}
		resp = append(resp, confFiles...)
	}


	c.SetJson(0,resp,"")
	return
}



// @Title search nginx confs
// @Description search of the nginx confs
// @Param   platformId 	      query     int   	  true     "platform id"
// @Param   serviceClass      query     string    true     "service class"
// @Param   hostId            query     int       true     "host id"
// @Success 0 {obj} platform
// @Failure 1 search host nginx confs 失败
// @Failure 2 User not found
// @router /hostConf/nginx [post]
func (c *HostConfNGController) ModifyNginxConf(){
	//  TODO


	c.SetJson(0,nil,"")
	return
}


