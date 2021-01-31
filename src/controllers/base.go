package controllers

import (
	"fmt"
	"library/components"
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

	c.User = new(models.User)
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
					c.SetJson(1, "", "状态已过期，请重新登录")
					return
				}
			}
		}
	}
	c.BaseAkossAuth()
}

func (c *BaseController) BaseAkossAuth(){

	inputURL :=  c.Controller.Ctx.Input.URL()
	if CancelAuth(inputURL){
		return
	}
	reqAuth := c.InitAuthRequestParam()
	respAuth, err := c.AkossAuth(reqAuth)
	if err != nil{
		c.SetJson(1,nil,"鉴权已过期，请重新登录" )
		return
	}

	if !respAuth.AuthPassed {
		fmt.Println("=========================")
		fmt.Println("reqAuth.Uri: ",reqAuth.RequestUri)
		fmt.Println("respAuth.AuthPassed: ",respAuth.AuthPassed)
		fmt.Println("respAuth.Msg: ",respAuth.Msg)
		fmt.Println("=========================")

		if respAuth.Msg == "Couldn't handle this token:"{
			c.SetJson(1,nil,"登录状态已过期，请重新登录")
		}else {
			c.SetJson(1,nil,respAuth.Msg)
		}
	}
}



const CancelAuthUriProjectList = "/api/project/list"
const CancelAuthUriLogin = "/login"
const CancelAuthUriChangePwd = "/changePasswd"
const CancelAuthToAkoss= "/api/akAuth"
func CancelAuth(inputUrl string)bool{
	cancelAuthList := make([]string,0)
	cancelAuthList = append(cancelAuthList, CancelAuthUriProjectList)
	cancelAuthList = append(cancelAuthList, CancelAuthUriLogin)
	cancelAuthList = append(cancelAuthList, CancelAuthUriChangePwd)
	cancelAuthList = append(cancelAuthList, CancelAuthToAkoss)

	for _, v := range cancelAuthList{
		if v == inputUrl{
			return true
		}
	}
	return false
}


func (c *BaseController) AkossAuth(reqAuth *components.AuthRequest)(components.AuthRespones, error){

	authUrl := beego.AppConfig.String("AuthUrl")
	respAuth, err := reqAuth.HttpAuth(authUrl)

	return respAuth,err
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

func (c *BaseController) SetAkAuthJson(code int, users []components.AkossUser, errMsg string) {
	status := false
	if code == 0 {
		status = true
	}

	c.Data["json"] = map[string]interface{}{"status": status, "FailedUsers": users,"ErrMsg":errMsg}
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

/*
	获取http鉴权所需参数
 */
func (c *BaseController) InitAuthRequestParam()(authReq *components.AuthRequest){

	//if c.User.XToken != "" {
	//	c.Ctx.SetCookie("x-token",c.User.XToken)
	//}
	authReq = &components.AuthRequest{
		UserName:   c.User.Username,
		UserPwd:    c.User.PasswordHash,
		XToken:		c.User.XToken,
		RequestUri: c.Controller.Ctx.Input.URL(),
		Method:     c.Controller.Ctx.Input.Method(),
	}

	if authReq.UserName == "" {
		fmt.Println("authReq.UserName is empty")
	}
	if authReq.UserPwd == "" {
		fmt.Println("authReq.UserPwd is empty")
	}
	if authReq.RequestUri == "" {
		fmt.Println("authReq.RequestUrl is empty")
	}
	if authReq.Method == "" {
		fmt.Println("authReq.Method is empty")
	}

	return
}
