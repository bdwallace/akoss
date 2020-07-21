package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
	// BaseController
}

func (c *LogoutController) Post() {
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "sucess"}
	c.ServeJSON()
	return
}
