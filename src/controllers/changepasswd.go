package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"library/common"
	"library/components"
	"models"
	"time"

	"encoding/json"
)

type ChangePasswdController struct {
	BaseController
}

/*
func (c *ChangePasswdController) Post() {
	//哈希校验成功后 更新 auth_key
	//beego.Info(string(c.Ctx.Input.RequestBody))

	postData := map[string]string{"old_password": "", "newpassword": "", "repeat_newpassword": ""}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &postData)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	oldPassword := postData["old_password"]
	newPassword := postData["newpassword"]
	repeatNewpassword := postData["repeat_newpassword"]
	if oldPassword == "" || newPassword == "" || repeatNewpassword == "" {
		c.SetJson(1, nil, "请输入密码")
		return
	}
	var user models.User
	o := orm.NewOrm()
	err = o.Raw("SELECT * FROM `user` WHERE id= ?", c.User.Id).QueryRow(&user)
	//beego.Info(user)
	//验证旧密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword))
	if err != nil {
		c.SetJson(1, nil, "旧密码有误，请重新输入")
		return
	} else {
		if newPassword == repeatNewpassword {
			password := []byte(newPassword)
			hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}
			user.PasswordHash = string(hashedPassword)
			models.UpdateUserById(&user)
			c.Data["json"] = map[string]interface{}{"code": 0, "msg": "sucess"}
			c.ServeJSON()
			return
		} else {
			c.SetJson(1, nil, "两次密码输入不一致，请重新输入")
			return
		}
	}
}

*/

func (c *ChangePasswdController) Post() {
	postData := map[string]string{"old_password": "", "newpassword": "", "repeat_newpassword": ""}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &postData)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	oldPassword := postData["old_password"]
	newPassword := postData["newpassword"]
	repeatNewpassword := postData["repeat_newpassword"]
	if oldPassword == "" || newPassword == "" || repeatNewpassword == "" {
		c.SetJson(1, nil, "请输入密码")
		return
	}
	var user models.User
	o := orm.NewOrm()
	err = o.Raw("SELECT * FROM `user` WHERE id= ?", c.User.Id).QueryRow(&user)
	if err != nil {
		c.SetJson(1, nil, "获取用户信息失败")
		return
	}
	//beego.Info(user)
	//验证旧密码
	reqAuth := &components.AuthRequest{
		UserName:   user.Username,
		UserPwd:    oldPassword,
		NewPwd:     newPassword,
		XToken:     user.XToken,
		RequestUri: c.Controller.Ctx.Input.URL(),
		Method:     c.Controller.Ctx.Input.Method(),
	}
	if newPassword == repeatNewpassword {
		respAuth, err := c.AkossAuth(reqAuth)
		user.AuthKey = common.Md5String(user.Username + common.GetString(time.Now().Unix()))
		models.UpdateUserById(&user)
		if err != nil {
			c.SetJson(1, nil, respAuth.Msg)
			return
		}
		if !respAuth.AuthPassed {
			fmt.Println("*************************")
			fmt.Println("reqAuth.Uri: ", reqAuth.RequestUri)
			fmt.Println("respAuth.AuthPassed: ", respAuth.AuthPassed)
			fmt.Println("respAuth.Msg: ", respAuth.Msg)
			fmt.Println("*************************")

			if respAuth.Msg == "Couldn't handle this token:" {
				c.SetJson(1, nil, "登录状态已过期，请重新登录")
			} else {
				c.SetJson(1, nil, respAuth.Msg)
			}
		}
	} else {
		c.SetJson(1, nil, "两次密码输入不一致，请重新输入")
	}

	c.SetJson(0, nil, "密码更改成功请重新登录")
	return
}
