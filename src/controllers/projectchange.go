package controllers

import (
	"models"
)

type ProjectChangeController struct {
	BaseController
}

func (c *ProjectChangeController) Get() {
	if c.User == nil || c.User.Id == 0 {
		c.SetJson(2, nil, "not login")
		return
	}

	id, _ := c.GetInt("id", 0)
	name := c.GetString("name")
	if id != 0 && name != "" {
		c.User.ProjectId = id
		c.User.ProjectName = name

		err := models.UpdateUserById(c.User)
		if err != nil {
			c.SetJson(1, c.User, "修改用户环境失败!")
			return
		}
	}

	c.SetJson(0, c.User, "修改成功")
	return
}
