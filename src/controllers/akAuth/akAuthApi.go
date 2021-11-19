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
func (c *AkAuthController) AkAuth() {

	reqObj := new(components.AkossRequest)
	var akossResp components.AkossResponse
	err := json.Unmarshal(c.Ctx.Input.RequestBody, reqObj)
	if err != nil {
		fmt.Println("error: AkAuth json.Unmarshal  ", err)
		c.SetAkAuthJson(1, nil, "json解析失败")
		return
	}
	var errRegister, errDelete, errUpdate error
	switch reqObj.Actions {
	case 101:
		// akAuth -> akoss注册用户
		akossResp, errRegister = reqObj.AkossRigsterUser()
	case 102:
		// akAuth -> akoss删除用户
		akossResp, errDelete = reqObj.AkossDeleteUser()
	case 103:
		// akAuth -> akoss 更新权限
		akossResp, errUpdate = reqObj.AkossUpdatePermiss()
		//akossResp, err = reqObj.akossu()
	}
	if errRegister != nil {
		fmt.Println("error: AkAuth register failed userName: ", akossResp.ErrMsg)
		c.SetAkAuthJson(1, akossResp.FailedUsers, akossResp.ErrMsg)
		return
	}
	if errDelete != nil && errDelete.Error() != "<QuerySeter> no row found" {
		fmt.Println("error: AkAuth delete failed  ", akossResp.ErrMsg)
		c.SetAkAuthJson(1, akossResp.FailedUsers, akossResp.ErrMsg)
		return
	}
	if errUpdate != nil && errUpdate.Error() != "<QuerySeter> no row found" {
		fmt.Println("error: AkAuth update failed  ", err)
		c.SetAkAuthJson(1, akossResp.FailedUsers, akossResp.ErrMsg)
		return
	}

	if !akossResp.Status {
		fmt.Println("error: AkAuth status failed  ", akossResp.ErrMsg)
		c.SetAkAuthJson(1, akossResp.FailedUsers, akossResp.ErrMsg)
		return
	}

	c.SetAkAuthJson(0, nil, "akoss操作成功")
	return

}
