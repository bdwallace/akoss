package othercontrollers

import (
	"controllers"
	"fmt"
	"github.com/astaxie/beego"
)

type OtherController struct {
	controllers.BaseController
}

// @Title 获取所有一些配置文件的参数
// @Description get other the config
// @Success 0 {object} models.Project
// @Failure 1 获取所有 project 失败
// @Failure 2 User not found
// @router /other/appconfig [get]
func (c *OtherController) AppConfig() {
	var app = map[string]string{}
	app["RunMode"] = beego.BConfig.RunMode
	app["dockerUrl"] = beego.AppConfig.String("dockerUrl")
	app["dockerUrlPub"] = beego.AppConfig.String("dockerUrlPub")
	app["dockerRepo"] = beego.AppConfig.String("dockerRepo")
	app["dockerRepoPub"] = beego.AppConfig.String("dockerRepoPub")
	app["dockerUser"] = beego.AppConfig.String("dockerUser")
	app["dockerPass"] = beego.AppConfig.String("dockerPass")
	app["dockerManager"] = beego.AppConfig.String("dockerManager")

	app["mysqluser"] = beego.AppConfig.String("mysqluser")
	app["mysqlpass"] = beego.AppConfig.String("mysqlpass")
	app["mysqlhost"] = beego.AppConfig.String("mysqlhost")
	app["mysqlport"] = beego.AppConfig.String("mysqlport")
	app["mysqldb"] = beego.AppConfig.String("mysqldb")

	app["Alias"] = beego.AppConfig.String("Alias")
	fmt.Println("------", app)
	c.SetJson(0, app, "")
	return
}
