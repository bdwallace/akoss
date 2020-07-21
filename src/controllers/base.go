package controllers

import (
	"runtime"

	"models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


const (
	ACTIONUPLOAD   = "upload"
	ACTIONDOWNLOAD = "download"
	ACTIONBUILD    = "build"
)


//基类
type BaseController struct {
	beego.Controller
	Task    *models.Task
	User    *models.User
	Service *models.Service
	Project *models.Project

}

// Prepare implemented Prepare method for baseRouter.
func (c *BaseController) Prepare() {

	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			runtimec := runtime.Stack(buf, false)
			beego.Error("控制器错误:", panic_err, string(buf[0:runtimec]))
		}
	}()

	token := ""
	if ah := c.Ctx.Input.Header("Authorization"); ah != "" {
		if len(ah) > 5 && strings.ToUpper(ah[0:5]) == "TOKEN" {
			token = ah[6:]
			if token != "" {
				var users []models.User
				o := orm.NewOrm()
				s, err := o.Raw("SELECT * FROM `user` WHERE auth_key= ?", token).QueryRows(&users)
				if s > 0 && err == nil {
					c.User = &(users[0])
				}
				if s == 0 {
					c.SetJson(1, "", "帐号已在其它地方登录!")
					return
				}
			}
		}
	}
}

func (c *BaseController) SetJson(code int, data interface{}, Msg string) {
	if code == 0 {
		if Msg == "" {
			Msg = "success"
		}
	}

	c.Data["json"] = map[string]interface{}{"code": code, "msg": Msg, "data": data}
	c.ServeJSON()
}

func (c *BaseController) ErrExit(errType string, Msg string) {
	switch errType {
	case "add": Msg = "数据新增失败:" + Msg
	case "del": Msg = "数据删除失败:" + Msg
	case "up": Msg = "数据更新失败:" + Msg
	default: Msg = "操作失败" + Msg
	}
	c.Data["json"] = map[string]interface{}{"code": 1, "msg": Msg, "data": nil}
	c.ServeJSON()
}
