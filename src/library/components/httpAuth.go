package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"library/common"
	"models"
	"net/http"
	"time"
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

// 鉴权  akoss -> akauth
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






// 用户创建/删除/更改权限时更新token和authKay
type AkossRequest struct {
	Actions 	int			// 101注册用户， 102更新token/authKey
	UserName	string
	XToken 		string
}

type AkossResponse struct {
	Status		bool
	Msg 		string
}


// akAuth -> akoss 添加用户
func (req *AkossRequest) AkossRigsterUser() (akossResp AkossResponse, err error) {

	var u models.User
	o := orm.NewOrm()
	//先判断存在用户否
	err = o.Raw("SELECT * FROM `user` WHERE username= ?", req.UserName).QueryRow(&u)
	if err != nil{
		akossResp.Status = false
		akossResp.Msg = "akoss 查询用户信息失败"
		return
	}
	if u.Id != 0 {
		akossResp.Status = false
		akossResp.Msg = "akoss 该用户存在"
		return
	}
	userAuth := common.Md5String(req.UserName + common.GetString(time.Now().Unix()))
	user := new(models.User)
	user.Username = req.UserName
	user.XToken = req.XToken
	user.AuthKey = userAuth
	user.IsEmailVerified = 1
	user.Avatar = "default.jpg"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if _, err = models.AddUser(user); err != nil{
		akossResp.Status = false
		akossResp.Msg = "akoss 添加该用户失败"
		return
	}
	akossResp.Status = true

	return
}



func (req *AkossRequest) AkossUpdatePermiss() (akossResp AkossResponse, err error) {

	var u models.User
	o := orm.NewOrm()
	//先判断存在用户否
	err = o.Raw("SELECT * FROM `user` WHERE username= ?", req.UserName).QueryRow(&u)
	if err != nil{
		akossResp.Status = false
		akossResp.Msg = "akoss 查询用户信息失败"
		return
	}

	// 用户不存在
	if u.Id == 0{
		akossResp.Status = false
		akossResp.Msg = "akoss 用户不存在"
		return
	}

	userAuth := common.Md5String(req.UserName + common.GetString(time.Now().Unix()))
	user := new(models.User)
	user.Username = req.UserName
	user.XToken = req.XToken
	user.AuthKey = userAuth
	user.IsEmailVerified = 1
	user.Avatar = "default.jpg"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err = models.UpdateUserById(user); err != nil{
		akossResp.Status = false
		akossResp.Msg = "akoss 更新用户token authKey失败"
		return
	}

	akossResp.Status = true
	return
}
