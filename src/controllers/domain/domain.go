package domaincontroller

import (
	"controllers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"library/common"
	"library/components"
	"models"
	"strings"
)

type DomainController struct {
	controllers.BaseController
}


// @Title 添加 domain
// @Description add the domain
// @Success 0 {id} int64
// @Failure 1 添加 domain 失败
// @Failure 2 User not found
// @router /domain [post]
func (c *DomainController)AddDomain(){

	beego.Info(string(c.Ctx.Input.RequestBody))
	// var domain *models.Domain
	domain := new(models.Domain)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, domain)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	allDomain := domain.Domain
	domainSplit := strings.Split(allDomain,"\n")

	// update     UpdateDomainRelated
	if 0 != domain.Id  {
		if len(domainSplit) > 1 {
			c.SetJson(1,nil,"不允许同时修改多个域名")
			return
		}
		domain.Domain = strings.Trim(domain.Domain," ")
		domain.Domain = strings.Trim(domain.Domain,",")
		err := models.UpdateDomainAndRelated(domain)
		if err != nil{
			fmt.Println("error: ",err)
			resErr := fmt.Sprintf("%s", err)
			c.SetJson(1,nil,resErr)
			return
		}
	}

	// add
	if 0 == domain.Id && len(domainSplit) > 0 {
		domains := make([]*models.Domain,0)
		for i := 0; i < len(domainSplit); i++ {
			if domainSplit[i] == "" {
				continue
			}
			tempDomain := new(models.Domain)
			*tempDomain = *domain
			tempDomain.Domain = strings.Trim(domainSplit[i]," ")
			tempDomain.Domain = strings.Trim(tempDomain.Domain,",")
			tempDomain.Conf = c.conf(tempDomain.Domain)
			tempDomain.Conf80 = c.conf80(tempDomain.Domain)
			domains = append(domains,tempDomain)
		}
		id, err := models.AddDomain(domains)
		if err != nil{
			fmt.Println("err:  ",err)
			c.SetJson(1,nil,"新建 domain 失败")
			return
		}

		domain.Id = int(id)
		err = models.AddDomainAndRelated(domain)
		if err != nil{
			fmt.Println("err:  ",err)
			c.SetJson(1,nil,"添加 domain 多对多关系 失败")
			return
		}
	}

	// 4. 返回数据
	c.SetJson(0, domain, "")
	return
}



// @Title 获取所有 domain
// @Description get all the domain
// @Success 0 {object} models.Domain
// @Failure 1 获取所有 domain 失败
// @Failure 2 User not found
// @router /domain/list/ [get]
func (c *DomainController)GetDomainList() {

	class := c.GetString("class")
	platforms := c.GetString("platforms")
	services := c.GetString("services")
	searchText := c.GetString("search_text")

	projectId, err := c.GetInt("project_id")
	if err != nil{
		fmt.Println("error: GetInt(\"project_id\")",err)
		c.SetJson(1,err,"获取 project_id 失败")
		return
	}

	page, _ := c.GetInt("page", 0)
	start := 0
	length, _ := c.GetInt("length", 20)

	if page > 0 {
		start = (page - 1) * length
	}

	count, resDomains, err := models.SearchDomain(projectId, start, length, class, platforms, services, searchText)
	if err != nil{
		fmt.Println("error: SearchDomain()",err)
		c.SetJson(1,err,"搜索 Domain匹配内容 失败")
		return
	}

	for _, s := range resDomains{
		_, err := models.GetDomainRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 conf services by id失败")
			return
		}
	}

	c.SetJson(0, map[string]interface{}{"total": count, "currentPage": page, "table_data": resDomains}, "")
	return
}



// @Title 获取所有 domain
// @Description get all the domain
// @Success 0 {object} models.Domain
// @Failure 1 获取所有 domain 失败
// @Failure 2 User not found
// @router /domain/platformAndClass/ [get]
func (c *DomainController)GetDomainByPlatformAndClass() {

	platformId, err := c.GetInt("platform_id")
	if err != nil{
		c.SetJson(1, err, "获取 platformId 失败")
		return
	}
	class := c.GetString("class")

	platform, err := models.GetPlatformById(platformId)
	if err != nil{
		c.SetJson(1, err, "获取 platform 失败")
		return
	}

	resPlatform, err := models.GetPlatformAndDomainRelated(platform)
	if err != nil{
		fmt.Println("error: GetAllDomain()",err)
		c.SetJson(1,err,"获取所有 Domain 失败")
		return
	}

	resDomains := make([]*models.Domain,0)
	for _, domain := range resPlatform.Domains{
		if domain.Class == class {
			resDomains = append(resDomains,domain)
		}
	}

	c.SetJson(0,resDomains,"")
	return
}



// @Title 获取 domain by id
// @Description get the domain by id
// @Param   id      query     int   	true     "domain id"
// @Success 0 {object} models.Domain
// @Failure 1 获取 domain by id 失败
// @Failure 2 User not found
// @router /domain/id/ [get]
func (c *DomainController)GetDomainById() {

	id, err := c.GetInt("id")
	if err != nil{
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	resDomain, err := models.GetDomainById(id)
	if err != nil{
		c.SetJson(1, err,"获取 domain by id失败")
		return
	}

	_, err = models.GetDomainRelated(resDomain)
	if err != nil {
		c.SetJson(1, err, "获取 domain related 失败")
		return
	}

	c.SetJson(0,resDomain,"")
	return
}


// @Title 获取 domain by projects id
// @Description get the domain by projects id
// @Param   projects_id      query     int   		true         "projects id"
// @Success 0 {object} models.Domain
// @Failure 1 获取 domain by projects id 失败
// @Failure 2 User not found
// @router /domain/projectId/ [get]
func (c *DomainController)GetDomainByProjectsId() {

	projectId, err :=  c.GetInt("project_id")
	if err != nil{
		c.SetJson(1, err, "获取project_id 失败")
		return
	}

	resDomain, err := models.GetDomainByProjectId(projectId)
	if err != nil{
		c.SetJson(1, err,"获取 domain by project id失败")
		return
	}

	for _, s := range resDomain{
		_, err := models.GetDomainRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 service host conf by id失败")
			return
		}
	}

	c.SetJson(0,resDomain,"")
	return
}


// @Title 获取 domain by projects id
// @Description get the domain by projects id
// @Param   service_id      query     int   		true         "service id"
// @Success 0 {object} models.Domain
// @Failure 1 获取 domain by projects id 失败
// @Failure 2 User not found
// @router /domain/serviceId/ [get]
func (c *DomainController)GetDomainByServiceId() {

	serviceId, err :=  c.GetInt("service_id")
	if err != nil{
		c.SetJson(1, err, "获取 service_id 失败")
		return
	}

	resDomain, err := models.GetDomainByServiceId(serviceId)
	if err != nil{
		c.SetJson(1, err,"获取 domain by service id失败")
		return
	}

	c.SetJson(0,resDomain,"")
	return
}




// @Title 删除 domain by id
// @Description delete the domain by id
// @Param   id      query     int   		true         "domain id"
// @Success 0 {ok} bool
// @Failure 1 删除 domain host by id 失败
// @Failure 2 User not found
// @router /domain/id/ [delete]
func (c *DomainController)DeleteDomainById(){


	id, err := c.GetInt("id")
	if err != nil{
		c.SetJson(1,err,"获取 id 失败")
		return
	}

	num, err := models.DeleteDomainById(id)
	if err != nil {
		c.SetJson(1,err,"删除 host 失败")
		return
	}

	c.SetJson(0,num,"")
	return

}



// @Title 删除 domain by id
// @Description delete the domain by id
// @Param   id      query     int   		true         "domain id"
// @Success 0 {ok} bool
// @Failure 1 删除 domain host by id 失败
// @Failure 2 User not found
// @router /domain/monitor/ [get]
func (c *DomainController)GetAllMonitor(){

	type Domain struct {
		Class  string
		Domain []string
	}

	var dataNew [4]Domain
	dataNew[0].Class = "download"
	dataNew[1].Class = "h5"
	dataNew[2].Class = "gateway"
	dataNew[3].Class = "mqtt"
	datas, _ := models.GetDomainMonitorAll()
	for _, data := range datas {
		switch data.Class {
		case "download":
			dataNew[0].Domain = append(dataNew[0].Domain, data.Domain)
		case "h5":
			dataNew[1].Domain = append(dataNew[1].Domain, data.Domain)
		case "gateway":
			dataNew[2].Domain = append(dataNew[2].Domain, data.Domain)
		case "mqtt":
			dataNew[3].Domain = append(dataNew[3].Domain, fmt.Sprintf("%s:%s", data.Domain, data.Port))
		}
	}

	c.SetJson(0, dataNew, "")
	return
}




// @Title  domain by project_id
// @Description doamin by project_id for class
// @Param   project_id      query     int   		true         "project id"
// @Success 0 {ok} bool
// @Failure 1 查询每套环境里域名的所有class
// @Failure 2 User not found
// @router /domain/class/ [get]
func (c *DomainController)GetDomainByProjectForClass(){

	projectId, err :=  c.GetInt("project_id")
	if err != nil{
		c.SetJson(1, err, "获取project_id 失败")
		return
	}

	data, err := models.GetDomainByProjectForClass(projectId)
	if err != nil{
		c.SetJson(1, err,"获取 domain by project id失败")
		return
	}

	c.SetJson(0, data, "")
	return
}


///////   util
func (c *DomainController) conf(domain string) (nginxconf string) {
	def := `
server {
	listen 443 ssl http2;
	server_name __ www.__;
	ssl_certificate      /etc/pki/tls/__.crt;
	ssl_certificate_key  /etc/pki/tls/__.key;
	include /etc/nginx/443.conf;
}
server {
	listen 80;
	server_name __ www.__;
	return 301   https://$host$request_uri;
}

`
	nginxconf = strings.Replace(def, "__", domain, -1)
	return
}

func (c *DomainController) conf80(domain string) (nginxconf string) {
	def := `
server {
	listen 80;
	server_name __ www.__;
	include /etc/nginx/443.conf;
}

`
	nginxconf = strings.Replace(def, "__", domain, -1)
	return
}



// @Title check domain health
// @Description check domain health
// @Param   domain      query     string 	  true       "domain"
// @Success 0 {ok}
// @Failure 1
// @Failure 2 User not found
// @router /domain/health/ [get]
func (c *DomainController)DomainHealthCheck() {

	domainId, err := c.GetInt("id")
	if err != nil{
		c.SetJson(1, err, "获取 domain_id 失败")
		return
	}

	domain, err := models.GetDomainById(domainId)
	if err != nil{
		c.SetJson(1, err, "根据 domain id 获取 domain 失败")
		return
	}

	url := ""
	health := ""
	if domain.Class == "mqtt" {
		health, err = components.GetMqttHealth(domain)
	}else {
		url = fmt.Sprintf("https://%s/", domain.Domain)
		response, err := common.HttpGet(url, nil)
		if err != nil {
			c.SetJson(1, -1,"获取 domain health 失败")
			return
		}
		health = common.IntToStr(response.StatusCode)
	}

	c.SetJson(0, health, "")
	return
}


