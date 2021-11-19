package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"testing"
)

func initDB() {

	// 链接 mysql
	MaxIdleConn := 30
	MaxOpenConn := 30

	DbLink := "root:123456@tcp(192.168.75.11:3306)/ak_go?charset=utf8&loc=Asia%2FShanghai"

	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		fmt.Println("error: test initDB RegisterDriver: ", err)
	}

	if err := orm.RegisterDataBase("default", "mysql", DbLink, MaxIdleConn, MaxOpenConn); err != nil {
		fmt.Println("error: test initDB RegisterDataBase: ", err)
	}

}

/////  project
func TestGetAllProject(t *testing.T) {

	initDB()

	project, err := GetAllProject()
	if err != nil {
		fmt.Println("error: GetAllProject() ", err)
		return
	}

	for i, p := range project {
		fmt.Println(i, ": ", p)
	}

}

func TestGetProjectById(t *testing.T) {

	initDB()

	pro := new(Project)
	pro.Id = 2

	proRes, err := GetProjectById(pro.Id)
	if err != nil {
		fmt.Println("error: GetProjectById(pro.Id) ", err)
		return
	}
	fmt.Println("GetProjectById:")
	fmt.Println(proRes)

}

func TestGetAllProjectByAlias(t *testing.T) {

	initDB()
	pro := new(Project)
	pro.Alias = "bbb"

	proRes, err := GetProjectByAlias(pro.Alias)
	if err != nil {
		fmt.Println("error: GetProjectByAlias() ", err)
		return
	}

	fmt.Println("GetProjectByAlias:")
	fmt.Println(proRes)

}

func TestAddProject(t *testing.T) {

	initDB()

	pro := new(Project)

	pro.Name = "project_aaa"
	pro.Alias = "aaa"

	id, err := AddProject(pro)
	if err != nil {
		fmt.Println("error: AddProject(pro)  ", err)
		return
	}

	fmt.Println("AddProject success res: ", id)

	/*
		创建project点击next 发送请求,
		初始化projectConf到mysql并将数据返回前端
	*/
	if err := InitProjectConfKey(int(id)); err != nil {
		fmt.Println("error: InitProjectConfKey ", err)
		return
	}

	fmt.Println("InitProjectConfKey success res")

}

func TestUpdateProject(t *testing.T) {

	initDB()

	pro := new(Project)
	pro.Id = 2
	pro.Alias = "xxxxx"
	num, err := UpdateProject(pro)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}

func TestDeleteProjectById(t *testing.T) {

	initDB()
	ok, err := DeleteProjectById(5)
	if err != nil && ok == false {
		fmt.Println("error: project_id is not exists  ", err)
		return
	}

	fmt.Println("delete project_id is success")
}

/////  project_conf
func TestAddProjectConf(t *testing.T) {

	mysqlConf := make(map[string]string, 0)
	mysqlConf["spring.datasource.url"] = "aaaaaaaaaa"
	mysqlConf["spring.datasource.username"] = "bbbbbbbb"
	mysqlConf["spring.datasource.password"] = "cccccccccc"

	// json 编码
	b, err := json.Marshal(mysqlConf)
	if err != nil {
		fmt.Println("error: json.Marshal ", err)
		return
	}
	// 转 string
	bStr := string(b)

	// init project_conf obj
	proConf := &ProjectConf{
		ConfName:  "mysqlConfxxxx",
		ConfValue: bStr,
		Project:   &Project{Id: 1},
	}

	initDB()

	res, err := AddProjectConf(proConf)

	if err != nil {
		fmt.Println("error: AddProjectConf(proConf)  ", err)
		return
	}

	fmt.Println("AddProjectConf success res: ", res)

}

func TestGetAllProjectConf(t *testing.T) {

	initDB()

	p, err := GetAllProjectConf()
	if err != nil {
		fmt.Println("error: GetAllProjectConf()", err)
		return
	}

	for _, v := range p {
		fmt.Printf("id: %d\n", v.Id)
		fmt.Printf("name: %s\n", v.ConfName)
		fmt.Printf("value: %s\n", v.ConfValue)
		fmt.Printf("time: %s\n", v.CreatTime.String())
		fmt.Printf("project_id: %d\n", v.Project.Id)

	}
}

func TestGetProjectConfById(t *testing.T) {
	initDB()

	v, err := GetProjectConfById(7)
	if err != nil {
		fmt.Println("error: GetProjectConfById  ", err)
		return
	}

	fmt.Printf("id: %d\n", v.Id)
	fmt.Printf("name: %s\n", v.ConfName)
	fmt.Printf("value: %s\n", v.ConfValue)
	fmt.Printf("time: %s\n", v.CreatTime.String())
	fmt.Printf("project_id: %d\n", v.Project.Id)

}

func TestGetProjectConfByProjectId(t *testing.T) {

	initDB()
	p, err := GetProjectConfByProjectId(1)
	if err != nil {
		fmt.Println("error: GetProjectConfByProjectId   ", err)
		return
	}

	for _, v := range p {
		fmt.Printf("id: %d\n", v.Id)
		fmt.Printf("name: %s\n", v.ConfName)
		fmt.Printf("value: %s\n", v.ConfValue)
		fmt.Printf("time: %s\n", v.CreatTime.String())
		fmt.Printf("project_id: %d\n", v.Project.Id)
		fmt.Println()
	}

}

func TestUpdateProjectConfById(t *testing.T) {

	mysqlConf := make(map[string]string, 0)
	mysqlConf["spring.datasource.url"] = "000000"
	mysqlConf["spring.datasource.username"] = "9999999"
	mysqlConf["spring.datasource.password"] = "88888888"

	// json 编码
	b, err := json.Marshal(mysqlConf)
	if err != nil {
		fmt.Println("error: json.Marshal(redisConf) ", err)
		return
	}
	// 转 string
	bStr := string(b)

	// init project_conf obj
	proConf := &ProjectConf{
		Id:        24,
		ConfName:  "zzzzzz",
		ConfValue: bStr,
		Project:   &Project{Id: 1},
	}

	initDB()

	res, err := UpdateProjectConf(proConf)

	if err != nil {
		fmt.Println("error: UpdateProjectConf(proConf)  ", err)
		return
	}

	fmt.Println("UpdateProjectConf success res: ", res)

}

func TestDeleteProjectConfById(t *testing.T) {
	initDB()
	is, err := DeleteProjectConfById(14)
	if err != nil && is == false {
		fmt.Println("error:  DeleteProjectConfById  ", err)
	}

}

////  service
func TestGetAllService(t *testing.T) {

	initDB()

	service, err := GetAllService()
	if err != nil {
		fmt.Println("error: GetAllService() ", err)
		return
	}

	for i, s := range service {
		fmt.Printf("i = %d  %d  %s  %s  %d \n", i, s.Id, s.Name, s.CreatTime.String(), s.Project.Id)
	}
}

func TestAddService(t *testing.T) {

	initDB()

	srv := &Service{
		Name:    "service_ddd",
		Project: &Project{Id: 1},
	}

	id, err := AddService(srv)
	if err != nil {
		fmt.Println("error: AddService(srv)  ", err)
		return
	}

	fmt.Println("AddService(srv) success res: ", id)

	/*
		创建project点击next 发送请求,
		初始化projectConf到mysql并将数据返回前端
	*/

	//if err := InitProjectConfKey(int(id));err != nil{
	//	fmt.Println("error: InitProjectConfKey ",err)
	//	return
	//}
	//
	//fmt.Println("InitProjectConfKey success res")

}
