package confcontrollers

import (
	"controllers"
	"library/rsa"
	"models"

	"encoding/json"
	"fmt"
	"strconv"
)

type ConfController struct {
	controllers.BaseController
}

/////  project_conf

// @Title 添加 project conf
// @Description add the project conf
// @Param   name     		query      string 	    true        "project conf name"
// @Param   value    		query      string  		true        "project conf value is json type"
// @Param  	project_id      query      int   		true        "conf id"
// @Success 0 {id} int64
// @Failure 1 添加 project conf 失败
// @Failure 2 User not found
// @router /conf/ [post]
func (c *ConfController) AddConf() {

	//beego.Info(string(c.Ctx.Input.RequestBody))
	var conf *models.Conf
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &conf)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	if 0 != conf.Id {
		// update conf
		_, err := models.UpdateConf(conf)
		if err != nil {
			c.SetJson(1, err, "更新 project conf by id失败")
			return
		}

		err = models.DeleteConfRelatedServices(conf)
		if err != nil {
			c.SetJson(1, nil, "删除 link for project 错误")
			return
		}

	} else {
		// add conf
		_, err = models.AddConf(conf)
		if err != nil {
			c.SetJson(1, err, "添加 project conf 失败")
			return
		}
	}

	err = models.UpdateConfRelatedServices(conf)
	if err != nil {
		c.SetJson(1, nil, "删除 link for project 错误")
		return
	}

	c.SetJson(0, conf, "")
	return
}

// @Title 获取所有 project conf
// @Description get all the project conf
// @Success 0 {object} models.Conf
// @Failure 1 获取所有 project conf 失败
// @Failure 2 User not found
// @router /conf/list [get]
func (c *ConfController) GetAllConfs() {

	searchText := c.GetString("search_text")
	projectId, err := c.GetInt("project_id")
	if err != nil {
		fmt.Println("error: GetInt(\"project_id\")", err)
		c.SetJson(1, err, "获取 project_id 失败")
		return
	}

	resConfs := make([]*models.Conf, 0)
	if searchText == "" {
		resConfs, err = models.GetAllConfs(projectId)
		if err != nil {
			fmt.Println("error: GetAllConf()", err)
			c.SetJson(1, err, "获取所有 Host 失败")
			return
		}
	} else {
		resConfs, err = models.SearchConfs(searchText, projectId)
		if err != nil {
			fmt.Println("error: SearchConfs()", err)
			c.SetJson(1, err, "搜索 host匹配内容 失败")
			return
		}
	}

	for _, s := range resConfs {
		_, err := models.GetConfRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 conf services by id失败")
			return
		}
	}
	c.SetJson(0, resConfs, "")
	return
}

// @Title 获取 project conf by id
// @Description get the project conf by id
// @Param   id     	 query     int   	true        "project conf id"
// @Success 0 {object} models.Conf
// @Failure 1 获取 project conf by id失败
// @Failure 2 User not found
// @router /conf/id [get]
func (c *ConfController) GetConfById() {

	id := c.Input().Get("id")
	idInt, err := strconv.Atoi(id)
	resProConf, err := models.GetConfById(idInt)
	if err != nil {
		c.SetJson(1, err, "获取 project conf by id失败")
		return
	}

	_, err = models.GetConfRelated(resProConf)
	if err != nil {
		c.SetJson(1, err, "获取 conf 关联的 services 失败")
		return
	}

	c.SetJson(0, resProConf, "")
	return
}

// @Title 获取 project conf by project id
// @Description get the project conf by project id
// @Param   project_id     	 query     int   	true        "project id"
// @Success 0 {object} models.Conf
// @Failure 1 获取 project conf by project id失败
// @Failure 2 User not found
// @router /conf/projectId [get]
func (c *ConfController) GetConfByProjectId() {

	id := c.Input().Get("project_id")
	idInt, err := strconv.Atoi(id)

	resProConf, err := models.GetConfByProjectId(idInt)
	if err != nil {
		c.SetJson(1, err, "获取 project conf by project id失败")
		return
	}

	for _, s := range resProConf {
		_, err := models.GetConfRelated(s)
		if err != nil {
			c.SetJson(1, err, "获取 conf services by id失败")
			return
		}
	}

	c.SetJson(0, resProConf, "")
	return
}

// @Title 删除 project conf by project id
// @Param   id     	 		query     int   		true         "project conf id"
// @Description update the project conf by id
// @Success 0 {ok} bool
// @Failure 1 删除 project conf by id失败
// @Failure 2 User not found
// @router /conf/id/ [delete]
func (c *ConfController) DeleteConfById() {

	id, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 conf id 失败")
	}
	is, err := models.DeleteConfById(id)
	if err != nil && is == false {
		c.SetJson(1, err, "删除 project conf by id失败")
		return
	}

	c.SetJson(0, is, "")
	return
}

// @Title 拷贝 conf by id
// @Description copy the conf conf by id
// @Param   id      query     int   		true         "conf id"
// @Success 0 {ok} bool
// @Failure 1 拷贝 conf conf by id 失败
// @Failure 2 User not found
// @router /conf/copy [get]
func (c *ConfController) CopyConfById() {

	id, err := c.GetInt("id")
	if err != nil {
		c.SetJson(1, err, "获取 id 失败")
		return
	}

	resConf, err := models.GetConfById(id)
	if err != nil {
		c.SetJson(1, err, "获取 conf host conf by id失败")
		return
	}

	// copy conf
	resConf.Id = 0
	resConf.Name = fmt.Sprintf("%s - copy", resConf.Name)
	_, err = models.AddConf(resConf)
	if err != nil {
		fmt.Println("error:  ", err)
		c.SetJson(1, err, "添加 conf 失败")
		return
	}

	c.SetJson(0, resConf, "")
	return
}

// @Title get rsa
// @Description get rsa
// @Success 0 {ok} bool
// @Failure 1 拷贝 get rsa 失败
// @Failure 2 User not found
// @router /conf/rsa [get]
func (c *ConfController) GetRSA() {

	rsa, err := rsa.GenRsaKey(2048)
	if err != nil {
		c.SetJson(1, err, "获取 RSA 失败")
		return
	}

	c.SetJson(0, rsa, "")
	return

}
