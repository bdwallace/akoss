package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"library/common"
	"library/components"
	"models"
	"models/request"
	"time"
)

type LoginController struct {
	BaseController
}

/*
func (c *LoginController) Post() {
	//哈希校验成功后 更新 auth_key
	//beego.Info(string(c.Ctx.Input.RequestBody))
	postData := map[string]string{"user_password": "", "user_name": "", "ProjectId": "", "ProjectName": ""}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &postData)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	password := postData["user_password"]
	userName := postData["user_name"]
	ProjectId := common.StrToInt(postData["ProjectId"])
	ProjectName := postData["ProjectName"]
	if userName == "" || password == "" {
		c.SetJson(1, nil, "用户名或密码不存在")
		return
	}
	var user models.User
	o := orm.NewOrm()
	err = o.Raw("SELECT * FROM `user` WHERE username= ?", userName).QueryRow(&user)
	//beego.Info(user)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		c.SetJson(1, nil, "用户名或密码错误")
		return
	} else if user.IsDel == 1 {
		c.SetJson(1, nil, "这是已删除用户，不能登录")
		return
	} else {
		// if user.AuthKey == "" {
		userAuth := common.Md5String(user.Username + common.GetString(time.Now().Unix()))
		user.AuthKey = userAuth
		user.ProjectId = ProjectId
		user.ProjectName = ProjectName
		models.UpdateUserById(&user)
		// }
		user.PasswordHash = ""
		c.SetJson(0, user, "")
		return
	}
}

*/
func (c *LoginController) Post() {
	//哈希校验成功后 更新 auth_key
	//beego.Info(string(c.Ctx.Input.RequestBody))
/*
	postData := map[string]string{"user_password": "", "user_name": "","Project": ""}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &postData)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	password := postData["user_password"]
	userName := postData["user_name"]
	projectIdStr := postData["Project"]
	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		c.SetJson(1, nil, "获取项目环境失败，请联系管理员")
		return
	}
*/
/*
	userName := c.GetString("user_name")
	password := c.GetString("user_password")
	project := new(models.Project)
	*/
	loginReq := new(request.LoginRequest)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &loginReq)
	if err != nil{
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	if loginReq.UserName == "" || loginReq.Password == "" {
		c.SetJson(1, nil, "请填写用户名和密码")
		return
	}

	var user models.User
	o := orm.NewOrm()
	err = o.Raw("SELECT * FROM `user` WHERE username= ?", loginReq.UserName).QueryRow(&user)
	if err != nil {
		c.SetJson(1, nil, "登录失败，没有该用户，请联系管理员")
		return
	}

	loginProject, err := models.GetProjectByName(loginReq.ProjectName)
	if err != nil{
		c.SetJson(1, nil, "获取项目环境失败，请联系管理员")
		return
	}
	reqAuth := &components.AuthRequest{
		UserName:   loginReq.UserName,
		UserPwd:    loginReq.Password,
		XToken:     user.XToken,
		RequestUri: c.Controller.Ctx.Input.URL(),
		Method:     c.Controller.Ctx.Input.Method(),
		ProjectId: loginProject.Id,
	}
	respAuth, err := c.AkossAuth(reqAuth)
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
		return
	}

	/*project,err := models.GetProjectById(project.Id)
	if err != nil {
		c.SetJson(1, nil, "获取项目环境失败，请联系管理员")
		return
	}*/

	userAuth := common.Md5String(user.Username + common.GetString(time.Now().Unix()))
	user.AuthKey = userAuth
	user.ProjectId = loginProject.Id
	user.ProjectName = loginProject.Name
	user.XToken = respAuth.XToken
	menusBase := respAuth.Menus
	components.FindDeployList(&menusBase, user)
	rootMenu := components.FindRootMenu(menusBase)
	menusTree := components.FindChild(&menusBase, &rootMenu)
	if menusTree == nil {
		c.SetJson(1, nil, "菜单权限校验失败,请联系管理员")
		return
	}
	menusJson, err := json.Marshal(menusTree.Child)
	if err != nil {
		c.SetJson(1, nil, "登录状态已过期，请重新登录")
		return
	}
	user.UserMenus = string(menusJson)
	if err := models.UpdateUserById(&user); err != nil {
		fmt.Println("error: models.UpdateUserById failed  ", err)
		c.SetJson(1, nil, "用户权限获取失败")
	}
	user.PasswordHash = ""
	if c.User.XToken != "" {
		c.Ctx.SetCookie("x-token", c.User.XToken)
	}
	resUserInfo := map[string]interface{}{"user": user, "login": true}
	userInfoJson, err := json.Marshal(resUserInfo)

	c.Ctx.SetCookie("gopub_userinfo", string(userInfoJson), 3600*24*2, "/")

	c.SetJson(0, user, "")
	return
}
