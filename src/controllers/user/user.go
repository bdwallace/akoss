package usercontrollers

import (
	"controllers"
	"github.com/astaxie/beego/orm"
	"models"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Get() {

	userId, _ := c.GetInt("id")
	if userId == 0 {
		o := orm.NewOrm()
		var users []orm.Params
		o.Raw("SELECT * FROM `user` ").Values(&users)
		c.SetJson(0, users, "")
		return
	} else {
		o := orm.NewOrm()
		var users []orm.Params
		var res orm.Params
		i, err := o.Raw("SELECT * FROM `user` where id = ? ", userId).Values(&users)
		if err == nil && i > 0 {
			res = users[0]
		}
		c.SetJson(0, res, "")
		return
	}

}


func (c *UserController) Delete() {

	userId, _ := c.GetInt("id")

	if err := models.DeleteUser(userId); err != nil{
		c.SetJson(1,err,"删除用户失败")
		return
	}

	c.SetJson(0, nil, "删除用户成功")
	return
}
