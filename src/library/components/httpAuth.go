package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthRequest struct {
	UserName	string
	UserPwd		string
	NewPwd 		string		// 用户修改密码使用
	XToken		string 		// 认证 token
	RequestUri	string 		// akoss 鉴权的api
	Method 		string 		// akoss api 鉴权的method
}

type AuthRespones struct {
	XToken		string 		// 用户登陆成功后，返回 xtoken
	AuthPassed 	bool 		// 是否通过api鉴权
	Msg 		string 		// 失败信息
}


func (authReq *AuthRequest) HttpAuth(url string) (respAuth AuthRespones, err error) {
	reqJson, err := json.Marshal(authReq)
	if err != nil{
		fmt.Println("error: HttpAuth json.Marshal  ",err)
		respAuth.Msg = err.Error()
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST",url, bytes.NewBuffer(reqJson))
	if err != nil{
		fmt.Println("error: HttpAuth http.NewRequest ",err)
		respAuth.Msg = err.Error()
		return
	}
	defer req.Body.Close()

	req.Header.Add("x-token",authReq.XToken)
	req.Header.Add("content-type","application/json")
	httpResp, err := client.Do(req)
	if err != nil{
		fmt.Println("error: HttpAuth client.Do  ",err)
		respAuth.Msg = err.Error()
		return
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		fmt.Println("error: HttpAuth ioutil.ReadAll  ",err)
		respAuth.Msg = err.Error()
		return
	}

	if err = json.Unmarshal(body,&respAuth); err != nil {
		fmt.Println("error: HttpAuth json.Unmarshal  ",err)
		respAuth.Msg = err.Error()
		return
	}

	return
}


