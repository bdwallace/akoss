package inspectcontrollers

import (
	"controllers"
	"library/inspect"
	"models"

	"encoding/json"
	"regexp"
)

type GrafanaController struct {
	controllers.BaseController
}

// @Title 添加
// @Description add the inspectgrafana
// @Param   name 			 query     string   	true         "inspectgrafana name"
// @Param   private_ip 		 query     string   	true         "private ip"
// @Param   public_ip 		 query     string   	true         "public ip"
// @Param   use_public 		 query     int		   	true         "use public IP"
// @Param   project_id       query     int   		true         "project id"
// @Success 0 {id} int64
// @Failure 1 添加 inspectgrafana 失败
// @Failure 2 User not found
// @router /inspectgrafana [post]
func (c *GrafanaController) PostInspectGrafana() {

	//beego.Info(string(c.Ctx.Input.RequestBody))
	var inspectgrafana models.InspectGrafana
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &inspectgrafana)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	reg := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	inspectgrafana.Domain = reg.FindString(inspectgrafana.Url)
	if inspectgrafana.Domain == "" {
		reg = regexp.MustCompile(`(?i)(?:\w+\.)*?(\w*\.(?:com\.cn|cn|com|net))[\\\/]*`)
		inspectgrafana.Domain = reg.FindString(inspectgrafana.Url)
	}

	if inspectgrafana.Id != 0 {
		_, err = models.UpdateInspectGrafanaById(&inspectgrafana)
	} else {
		_, err = models.AddInspectGrafana(&inspectgrafana)
	}
	if err != nil {
		if err.Error() == "<QuerySeter> return rows" {
			c.SetJson(1, nil, "命名有重复")
		}
		c.SetJson(1, nil, "数据库更新错误:"+err.Error())
		return
	}
	c.SetJson(0, inspectgrafana, "修改成功")
	return
}

// @Title 获取 inspectgrafana by id
// @Description get the inspectgrafana by id
// @Param   id      query     int   	true     "inspectgrafana id"
// @Success 0 {object} models.InspectGrafana
// @Failure 1 获取 inspectgrafana by id 失败
// @Failure 2 User not found
// @router /inspectgrafana/id/ [get]
func (c *GrafanaController) GetInspectGrafanaById() {

	id, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	resInspectGrafana, err := models.GetInspectGrafanaById(id)
	if err != nil {
		c.SetJson(1, err, "获取 inspectgrafana by id失败")
		return
	}

	c.SetJson(0, resInspectGrafana, "")
	return
}

// @Title 删除 inspectgrafana inspectgrafana by id
// @Description update the inspectgrafana inspectgrafana by id
// @Param   id      query     int   		true         "inspectgrafana id"
// @Success 0 {ok} bool
// @Failure 1 删除 inspectgrafana inspectgrafana by id 失败
// @Failure 2 User not found
// @router /inspectgrafana/id/ [delete]
func (c *GrafanaController) DeleteInspectGrafanaById() {

	id, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	_, err = models.DeleteInspectGrafanaById(id)
	if err != nil {
		c.SetJson(1, err, "删除 inspectgrafana 失败")
		return
	}

	c.SetJson(0, nil, "")
	return

}

// @Title 获取所有 inspectgrafana
// @Description get all the inspectgrafana
// @Success 0 {object} models.InspectGrafana
// @Failure 1 获取所有 inspectgrafana 失败
// @Failure 2 User not found
// @router /inspectgrafana/list/ [get]
func (c *GrafanaController) GetAllInspectGrafana() {

	page, _ := c.GetInt("page", 0)
	start := 0
	length, _ := c.GetInt("length", 20)
	if page > 0 {
		start = (page - 1) * length
	}

	count, resInspectGrafanas, err := models.GetAllInspectGrafanaPage(start, length)
	if err != nil {
		c.SetJson(1, err, "获取所有 InspectGrafana 失败")
		return
	}

	c.SetJson(0, map[string]interface{}{"total": count, "currentPage": page, "table_data": resInspectGrafanas}, "")
	return
}

// @Title 获取所有 inspectgrafana
// @Description get all the inspectgrafana
// @Success 0 {object} models.InspectGrafana
// @Failure 1 获取所有 inspectgrafana 失败
// @Failure 2 User not found
// @router /inspectgrafana/exec/ [get]
func (c *GrafanaController) ExecAllInspectGrafana() {

	resInspectGrafanas, err := models.GetAllInspectGrafana()
	if err != nil {
		c.SetJson(1, err, "获取所有 InspectGrafana 失败")
		return
	}

	for _, v := range resInspectGrafanas {
		inspect.Grafana(v)
	}

	c.SetJson(0, nil, "")
	return
}
