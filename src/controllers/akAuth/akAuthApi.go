package akauthcontroller

import (
	"controllers"
	"encoding/json"
	"fmt"
	"library/components"
)

type AkAuthController struct {
	controllers.BaseController
}

// @Title akAuth -> akoss 添加用户
// @Description 当akAuth新增用户后， http请求到akoss，添加用户并更新token, 返回akAuth response
// 				当akAuth更新用户权限时， http请求到akoss，更新token/authKey, 返回akAuth response
// @router /akAuth [post]
func (c *AkAuthController)AkAuth() {

	reqObj := new(components.AkossRequest)
	var akossResp components.AkossResponse
	err := json.Unmarshal(c.Ctx.Input.RequestBody, reqObj)
	if err != nil {
		fmt.Println("error: AkAuth json.Unmarshal  ",err)
		c.SetJson(1,nil, "json解析失败")
		return
	}

	switch reqObj.Actions {
	case 101:
		// akAuth -> akoss注册用户
		akossResp, err = reqObj.AkossRigsterUser()
	case 102:
		// akAuth -> akoss 更新权限
		akossResp, err = reqObj.AkossUpdatePermiss()
	}
	if err != nil{
		fmt.Println("error: AkAuth failed  ",err)
		c.SetJson(1,nil, "json解析失败")
	}
	//
	//jsonResp, err := json.Marshal(akossResp)
	//if err != nil {
	//	fmt.Println("error: AkAuth json.Marshal  ",err)
	//	c.Controller.ServeJSON()
	//	return
	//}
	//
	c.SetJson(0,akossResp,"seccess")
	//c.Controller.ServeJSON()

}

