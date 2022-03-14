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

/*
	鉴权  akoss -> akauth
*/

type AuthRequest struct {
	UserName   string
	UserPwd    string
	NewPwd     string // 用户修改密码使用
	XToken     string // 认证 token
	RequestUri string // akoss 鉴权的api
	Method     string // akoss api 鉴权的method
	ProjectId	int
}

type AuthRespones struct {
	XToken     string // 用户登陆成功后，返回 xtoken
	AuthPassed bool   // 是否通过api鉴权
	Msg        string // 失败信息
	Menus      []MenusAuth
}

type MenusAuth struct {
	Id       int    `gorm:"id" json:"id"`
	ParentId int    `gorm:"parent_id" json:"parent_id"`
	Title    string `gorm:"title" json:"title"`
	//Name 		string	`gorm:"name" json:"name"`
	Path  string      `gorm:"path" json:"path"`
	Icon  string      `gorm:"icon" json:"icon"`
	Child []MenusAuth `json:"child"`
}

func (authReq *AuthRequest) HttpAuth(url string) (respAuth AuthRespones, err error) {
	reqJson, err := json.Marshal(authReq)
	if err != nil {
		fmt.Println("error: HttpAuth json.Marshal  ", err)
		respAuth.Msg = err.Error()
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	if err != nil {
		fmt.Println("error: HttpAuth http.NewRequest ", err)
		respAuth.Msg = err.Error()
		return
	}
	defer req.Body.Close()

	req.Header.Add("x-token", authReq.XToken)
	req.Header.Add("content-type", "application/json")
	httpResp, err := client.Do(req)
	if err != nil {
		fmt.Println("error: HttpAuth client.Do  ", err)
		respAuth.Msg = err.Error()
		return
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		fmt.Println("error: HttpAuth ioutil.ReadAll  ", err)
		respAuth.Msg = err.Error()
		return
	}

	respAuth.Menus = make([]MenusAuth, 0)
	if err = json.Unmarshal(body, &respAuth); err != nil {
		fmt.Println("error: HttpAuth json.Unmarshal  ", err)
		respAuth.Msg = err.Error()
		return
	}

	return
}

/*
	akAuth -> akoss
*/
// 用户创建/删除/更改权限时更新token和authKay
type AkossRequest struct {
	Actions int         // 101注册用户， 102更新token/authKey
	Users   []AkossUser // 用户注册时 len==1
}

type AkossResponse struct {
	Status      bool
	FailedUsers []AkossUser
	ErrMsg      string
}

type AkossUser struct {
	UserName string
	XToken   string
	Msg      string
}

// akAuth -> akoss 添加用户
func (req *AkossRequest) AkossRigsterUser() (akossResp AkossResponse, err error) {

	var u models.User
	o := orm.NewOrm()
	respUser := AkossUser{}
	akossResp.FailedUsers = make([]AkossUser, 0)
	//先判断存在用户否

	o.Raw("SELECT * FROM `user` WHERE username= ?", req.Users[0].UserName).QueryRow(&u)

	userAuth := common.Md5String(req.Users[0].UserName + common.GetString(time.Now().Unix()))
	registerUser := new(models.User)
	registerUser.Username = req.Users[0].UserName
	registerUser.Realname = req.Users[0].UserName
	registerUser.XToken = req.Users[0].XToken
	registerUser.AuthKey = userAuth
	registerUser.IsEmailVerified = 1
	registerUser.Status = 10
	registerUser.Role = 1
	registerUser.IsDel = 0
	registerUser.Email = req.Users[0].UserName + "@ak.com"
	registerUser.Avatar = "default.jpg"
	registerUser.CreatedAt = time.Now()
	registerUser.UpdatedAt = time.Now()

	fmt.Println("query user name: ", req.Users[0].UserName)
	fmt.Println("u.name: ", u.Username, "  u.AuthKey: ", u.AuthKey, "  u.xtoken: ", u.XToken)

	if u.Id != 0 {
		//akossResp.Status = false
		//respUser.UserName = req.Users[0].UserName
		//respUser.Msg = "akoss 该用户存在"
		//akossResp.FailedUsers = append(akossResp.FailedUsers,respUser)
		//err = fmt.Errorf("akoss 该用户存在")

		fmt.Println("用户已存在，更新用户信息")
		// 用户存在 更新用户信息
		if err := models.UpdateUserById(registerUser); err != nil {
			fmt.Println("更新用户失败")
			akossResp.Status = false
			respUser.UserName = req.Users[0].UserName
			respUser.Msg = "akoss用户已存在， 更新用户信息失败"
			akossResp.FailedUsers = append(akossResp.FailedUsers, respUser)
		}

	} else {
		fmt.Println("用户已存在，添加用户")
		// 用户不存在 添加用户
		if _, err = models.AddUser(registerUser); err != nil {
			fmt.Println("添加用户失败")
			akossResp.Status = false
			respUser.UserName = req.Users[0].UserName
			respUser.Msg = "akoss用户不存在， 添加用户失败"
			akossResp.FailedUsers = append(akossResp.FailedUsers, respUser)
			return
		}
	}
	if len(akossResp.FailedUsers) > 0 {
		akossResp.ErrMsg = akossResp.FailedUsers[0].Msg + " " + akossResp.FailedUsers[0].UserName
	} else {
		akossResp.Status = true
	}

	return
}

// akAuth -> akoss 添加用户
func (req *AkossRequest) AkossDeleteUser() (akossResp AkossResponse, err error) {

	respUser := AkossUser{}
	akossResp.FailedUsers = make([]AkossUser, 0)

	qureyUser, err := models.GetUserByName(req.Users[0].UserName)
	if err != nil {
		errMsg := fmt.Errorf("error: models.GetUserByName not found username %s \n", req.Users[0].UserName)
		fmt.Println(errMsg)
		respUser.Msg = "akoss 获取用户失败"
		akossResp.Status = false
		akossResp.FailedUsers = append(akossResp.FailedUsers, respUser)
		goto End
	}
	qureyUser.IsDel = 1
	if err = models.UpdateUserById(qureyUser); err != nil {
		errMsg := fmt.Errorf("error: models.DeleteUser failed  %s \n", req.Users[0].UserName)
		fmt.Println(errMsg)
		respUser.Msg = "akoss 删除用户失败"
		akossResp.Status = false
		akossResp.FailedUsers = append(akossResp.FailedUsers, respUser)
		goto End
	}

End:
	if len(akossResp.FailedUsers) > 0 {
		akossResp.ErrMsg = akossResp.FailedUsers[0].Msg + " " + akossResp.FailedUsers[0].UserName
		return
	} else {
		akossResp.Status = true
	}
	return
}

func (req *AkossRequest) AkossUpdatePermiss() (akossResp AkossResponse, err error) {

	akossResp = AkossResponse{}
	akossResp.FailedUsers = make([]AkossUser, 0)
	for _, reqUser := range req.Users {
		var quireUser models.User
		o := orm.NewOrm()
		respUser := AkossUser{
			UserName: reqUser.UserName,
		}
		//先判断存在用户否
		err = o.Raw("SELECT * FROM `user` WHERE username= ?", reqUser.UserName).QueryRow(&quireUser)
		if err != nil {
			respUser.Msg = "akoss 用户不存在"
			akossResp.FailedUsers = append(akossResp.FailedUsers, respUser)
			continue
		}

		userAuth := common.Md5String(reqUser.UserName + common.GetString(time.Now().Unix()))
		updateUser := new(models.User)
		updateUser.Id = quireUser.Id
		updateUser.Username = reqUser.UserName
		updateUser.XToken = reqUser.XToken
		updateUser.AuthKey = userAuth
		updateUser.Role = 1
		updateUser.Status = 10
		updateUser.Realname = reqUser.UserName
		updateUser.ProjectId = quireUser.ProjectId
		updateUser.ProjectName = quireUser.ProjectName
		updateUser.IsDel = quireUser.IsDel
		updateUser.Avatar = "default.jpg"
		updateUser.Email = reqUser.UserName + "@ak.com"
		updateUser.UpdatedAt = time.Now()

		if err = models.UpdateUserById(updateUser); err != nil {
			respUser.Msg = "akoss 更新用户token authKey失败"
			akossResp.FailedUsers = append(akossResp.FailedUsers, respUser)
			continue
		}
	}

	if len(akossResp.FailedUsers) > 0 {
		akossResp.Status = false
		for _, u := range akossResp.FailedUsers {
			akossResp.ErrMsg += u.Msg + " " + u.UserName + "\n"
		}
	} else {
		akossResp.Status = true
	}
	return
}
