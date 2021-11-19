package usercontrollers

import (
	"controllers"
	"golang.org/x/crypto/bcrypt"
	"library/common"
	"models"
	"time"

	"github.com/astaxie/beego/orm"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Get() {

	userId, _ := c.GetInt("id")
	if userId == 0 {
		o := orm.NewOrm()
		var users []orm.Params
		o.Raw("SELECT * FROM `user` WHERE is_del = 0 ").Values(&users)
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

	user, err := models.GetUserById(userId)
	if err != nil {
		c.SetJson(1, err, "获取 user 失败")
		return
	}
	user.IsDel = 1
	if err = models.UpdateUserById(user); err != nil {
		// if err := models.DeleteUser(userId); err != nil{
		c.SetJson(1, err, "删除用户失败")
		return
	}

	c.SetJson(0, nil, "删除用户成功")
	return
}

//重置用户密码，必须要admin用户才能重置
func (c *UserController) Put() {
	if c.User == nil || c.User.Id != 1 {
		c.SetJson(1, nil, "请联系admin重置密码")
		return
	}

	userId, _ := c.GetInt("id")

	user, err := models.GetUserById(userId)
	if err != nil {
		c.SetJson(1, err, "获取 user 失败")
		return
	}

	password := []byte("123456")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	userAuth := common.Md5String(user.Username + common.GetString(time.Now().Unix()))

	user.AuthKey = userAuth
	user.PasswordHash = string(hashedPassword)
	if err = models.UpdateUserById(user); err != nil {
		c.SetJson(1, err, "重置用户失败")
		return
	}

	c.SetJson(0, nil, "重置用户成功")
	return
}
