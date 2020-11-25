package linkcontrollers

import (
	"controllers"
	"models"

	"encoding/json"
)

type LinkController struct {
	controllers.BaseController
}


// @Title get link by id
// @Description get link by id
// @Param   id      query     int 		true       "link id"
// @Success 0 {object} models.Link
// @Failure 1 获取 link by id 失败
// @Failure 2 User not found
// @router /link/id/ [get]
func (c *LinkController) GetLinkById() {
	linkId, _ := c.GetInt("id")
	link, err := models.GetLinkById(linkId)
	if err != nil {
		c.SetJson(1, err, "获取 link 失败")
		return
	}

	err = models.GetLinkRelatedProjects(link)
	if err != nil {
		c.SetJson(1, err, "获取 service host by id失败")
		return
	}

	c.SetJson(0, link, "")
	return
}


// @Title get all link
// @Description get all link
// @Success 0 {object} models.Link
// @Failure 1 获取 get all link 失败
// @Failure 2 User not found
// @router /link/list/ [get]
func (c *LinkController) GetAllLink() {
	data, err := models.GetLinkAll()
	if err != nil {
		c.SetJson(1, err, "获取 link 列表 失败")
		return
	}

	for _, v := range data {
		err = models.GetLinkRelatedProjects(v)
		if err != nil {
			c.SetJson(1, err, "获取 service host by id失败")
			return
		}
	}

	c.SetJson(0, data, "查询成功")
	return
}


// @Title add link by new link
// @Description add link by new link
// @Success 0 {object} models.Link
// @Failure 1 添加 add link by new link 失败
// @Failure 2 User not found
// @router /link/ [post]
func (c *LinkController) AddLink() {
	//projectId,_:=c.GetInt("projectId",0)
	//beego.Info(string(c.Ctx.Input.RequestBody))
	var link *models.Link
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &link)
	if err != nil {
		c.SetJson(1, nil, "格式错误")
		return
	}

	o, _ := models.OrmBegin()

	if link.Id != 0 {
		if err = models.UpdateLink(o, link); err != nil {
			models.OrmRollback(o)
			c.SetJson(1, nil, "修改失败: " + err.Error())
			return
		}
		if err = models.DeleteLinkRelatedProjects(o, link); err != nil {
			models.OrmRollback(o)
			c.SetJson(1, nil, "删除关联失败: " + err.Error())
			return
		}

	} else {
		if _, err = models.AddLink(o, link); err != nil {
			models.OrmRollback(o)
			c.SetJson(1, nil, "新增失败:" + err.Error())
			return
		}
	}
	
	if err := models.UpdateLinkRelatedProjects(o, link); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "更新关联失败")
		return
	}

	if err := models.OrmCommit(o); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "操作失败: " + err.Error())
		return
	} else {
		c.SetJson(0, link, "操作成功")
		return
	}
}


// @Title copy link by link id
// @Description copy link by link id
// @Param   id      query     int 		true       "link id"
// @Success 0 {object} models.Link
// @Failure 1 复制 link by link id 失败
// @Failure 2 User not found
// @router /link/copy/ [get]
func (c *LinkController) CopyLink() {
	if c.User == nil || c.User.Id == 0 {
		c.SetJson(2, nil, "not login")
		return
	}

	linkId, _ := c.GetInt("id")

	link, err := models.GetLinkById(linkId)
	if err != nil {
		c.SetJson(1, nil, "查询失败")
		return
	}
	if err := models.GetLinkRelatedProjects(link); err != nil {
		c.SetJson(1, nil, "查询关联失败")
		return
	}

	link.Name = link.Name + " - copy"
	link.Id = 0


	o, _ := models.OrmBegin()
	if _, err := models.AddLink(o, link); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "复制失败")
		return
	}
	
	if err := models.UpdateLinkRelatedProjects(o, link); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "更新关联失败")
		return
	}

	if err = models.OrmCommit(o); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "操作失败: " + err.Error())
		return
	} else {
		c.SetJson(0, nil, "操作成功")
		return
	}
}


// @Title delete link by link id
// @Description delete link by link id
// @Param   id      query     int 		true       "link id"
// @Success 0 {object} models.Link
// @Failure 1 删除 link by link id 失败
// @Failure 2 User not found
// @router /link/id/ [delete]
func (c *LinkController) DeleteLink() {
	linkId, _ := c.GetInt("id")

	link, err := models.GetLinkById(linkId)
	if err != nil {
		c.SetJson(1, nil, "查询异常: " + err.Error())
		return
	}

	o, _ := models.OrmBegin()

	if err = models.DeleteLinkRelatedProjects(o, link); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "删除关联失败: " + err.Error())
		return
	}

	if err = models.DeleteLink(o, link); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "删除失败: " + err.Error())
		return
	}

	if err = models.OrmCommit(o); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "操作失败: " + err.Error())
		return
	} else {
		c.SetJson(0, nil, "操作成功")
		return
	}
}
