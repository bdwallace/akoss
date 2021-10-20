package confManagement

import (
	"controllers"
	"encoding/json"
	"fmt"
	"library/confManagement"
	"models"
)

type HostConfNGController struct {
	controllers.BaseController
}

// @Title search nginx confs
// @Description search of the nginx confs
// @Param   projectId 	      query     int   	  true     "project id"
// @Param   class		      query     string    true     "service class"
// @Param   host              query     string    true     "host"
// @Success 0 {obj} response
// @Failure 1 search host nginx confs 失败
// @Failure 2 User not found
// @router /confManage/searchServiceHost [get]
func (c *HostConfNGController) SearchSericeHost(){


	projectId, _ := c.GetInt("project_id")
	class := c.GetString("service_class")
	host := c.GetString("host")

	p, err := models.GetProjectById(projectId)
	if err != nil{
		fmt.Println("error:  GetProjectById",err)
		c.SetJson(1,err,"获取 project 失败")
		return
	}

	confRelation, err := models.SearchServiceAndHost(host,class)
	if err != nil{
		fmt.Println("error:  SearchServiceAndHost",err)
		c.SetJson(1,err,"获取 host info 失败")
		return
	}


	for _, conf := range confRelation{
		service, err := models.GetServiceById(conf.ServiceId)
		if err != nil{
			break
		}
		docker := confManagement.RegisterBaseComm(p,service)
		if err = confManagement.GetConfAndCpToLocal(conf,docker); err != nil{
			break
		}
	}

	c.SetJson(0,confRelation,"")
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
// @router /confManage/nginx [post]
func (c *HostConfNGController) ModifyNginxConf(){

	confEdit := new(models.ConfEdit)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, confEdit); err != nil{
		fmt.Println("error:  Ctx.Input.RequestBody  ",err.Error())
		c.SetJson(1,err,"请求参数解析失败")
		return
	}

	if err := confManagement.WriteConfToLocal(confEdit); err != nil{
		fmt.Println("error:  WriteConfToLocal(confEdit)  ",err.Error())
		c.SetJson(1,err,"配置文件写入本地失败")
		return
	}
	if err := confManagement.DockerCPConfFileToContainer(confEdit); err != nil{
		fmt.Println("error:  DockerCPConfFileToContainer(confEdit)  ",err.Error())
		c.SetJson(1,err,"配置文件拷贝到容器失败")
		return
	}


	c.SetJson(0,nil,"")
	return
}

