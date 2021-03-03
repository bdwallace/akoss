package fastcontrollers

import (
	"controllers"
	"encoding/json"
	"fmt"
	"library/aws"
	"models"

	"github.com/astaxie/beego"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type SgroupController struct {
	controllers.BaseController
}


type GroupData struct {
	ProjectId int
	Data  ec2.SecurityGroup
}


// @Title 获取 sgroup
// @Description 获取 sgroup
// @Success 0 {id} int64
// @Failure 1 获取 sgroup 失败
// @Failure 2 User not found
// @router /fast/getSgroup [get]
func (c *SgroupController) GetSgroup() {
	var myAws aws.BaseAws

	projectId, err := c.GetInt("projectId")
	project, err := models.GetProjectById(projectId)
	if err != nil{
		fmt.Println("error: GetProjectById(pro.Id) ",err)
		c.SetJson(1, err, "获取 project by id 失败")
		return
	}

	myAws.Region = project.AwsRegion
	myAws.Alias = project.Alias

	_ = myAws.InitSession()
	myAws.SourceInstance.Ec2 = ec2.New(myAws.Session)

	data, err := myAws.GetAllSgroupForMerchant()
	if err != nil {
		beego.Error(err)
		c.SetJson(1, nil, "获取安全组失败")
		return
	}
	c.SetJson(0, data, "")
	return
}



// @Title 设置 sgroup
// @Description set sgroup
// @Success 0 {id} int64
// @Failure 1 设置 sgroup 失败
// @Failure 2 User not found
// @router /fast/setSgroup [post]
func (c *SgroupController) SetSgroup() {
	//beego.Info(string(c.Ctx.Input.RequestBody))
	var myAws aws.BaseAws
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &myAws)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误1")
		return
	}

	project	, _ := models.GetProjectById(myAws.Project)
	myAws.Region = project.AwsRegion
	myAws.Alias = project.Alias

	_ = myAws.InitSession()
	myAws.SourceInstance.Ec2 = ec2.New(myAws.Session)

	data, err := myAws.IngressSgroup()
	if err != nil {
		beego.Error(err)
		c.SetJson(1, nil, "保存安全组失败")
		return
	}
	c.SetJson(0, data, "")
	return
}




// @Title 删除 sgroup
// @Description delete sgroup
// @Success 0 {id} int64
// @Failure 1 删除 sgroup 失败
// @Failure 2 User not found
// @router /fast/deleteSgroup [delete]
func (c *SgroupController) DeleteSgroup() {
	//beego.Info(string(c.Ctx.Input.RequestBody))

	// var sgroup ec2.SecurityGroup
	var sgroup GroupData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &sgroup)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	var myAws aws.BaseAws
	project, _ := models.GetProjectById(sgroup.ProjectId)
	myAws.Region = project.AwsRegion
	myAws.Alias = project.Alias

	_ = myAws.InitSession()
	myAws.SourceInstance.Ec2 = ec2.New(myAws.Session)

	myAws.GroupIds = []string{*sgroup.Data.GroupId}
	iprangstr := sgroup.Data.IpPermissions

	data, err := myAws.DeleteSgroup(iprangstr)
	if err != nil {
		c.SetJson(1, nil, "删除安全组失败")
		return
	}
	c.SetJson(0, data, "")
	return
}
