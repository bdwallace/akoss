package resourcecontrollers

import (
	"controllers"
	"models"

	"fmt"
	"library/aws"
	"library/backup"

	"github.com/astaxie/beego"
)

type ResourceController struct {
	controllers.BaseController
}

// @Title AWS 资源监控
// @Description AWS 资源监控
// @Param   awsRegion      query     string   	true     "awsRegion"
// @Param   service        query     string   	true     "service class"
// @Success 0 {id} int64
// @Failure 1 AWS 资源监控
// @Failure 2 User not found
// @router /resource/aws [get]
func (c *ResourceController) Aws() {

	//  获取区域选择
	awsRegion := c.GetString("awsRegion")
	server := c.GetString("server")

	project, err := models.GetProjectById(c.User.ProjectId)
	if err != nil {
		fmt.Println("error: GetProjectById ", err)
		c.SetJson(1, err, "获取 project by id 失败")
		return
	}

	//  init
	myAws := new(aws.BaseAws)
	myAws.Region = awsRegion
	myAws.Alias = project.Alias
	myAws.SourceInfo = new(aws.ResourceInfo)
	_ = myAws.InitSession()

	switch server {
	case "ec2":
		err := myAws.GetEc2CollocationSubnet()
		if err != nil {
			beego.Error(err)
			c.SetJson(501, nil, "获取EC2Info失败")
		}
		resEc2InstanceInfos := make([]*aws.EC2InstanceInfo, 0)
		for _, v := range myAws.SourceInfo.EC2Info.EC2Reservations {
			resEc2InstanceInfos = append(resEc2InstanceInfos, v.Instances...)
		}

		c.SetJson(200, resEc2InstanceInfos, "")

	case "vpc":
		err := myAws.GetVPCAndAnalytical()
		if err != nil {
			beego.Error(err)
			c.SetJson(501, nil, "获取VPC失败")
		}
		c.SetJson(200, myAws.SourceInfo.VPCInfo.VPCs, "")

	case "subnet":
		err := myAws.GetSubnetAndAnalytical()
		if err != nil {
			beego.Error(err)
			c.SetJson(501, nil, "获取subnet失败")
		}
		c.SetJson(200, myAws.SourceInfo.Subnets.Subnets, "")

	case "rds":
		err := myAws.GetRDSAndAnalytical()
		if err != nil {
			beego.Error(err)
			c.SetJson(501, nil, "获取RDSInfo失败")
		}
		c.SetJson(200, myAws.SourceInfo.RDSCluster.RDSInstances, "")

	case "redis":
		err := myAws.GetElastiCacheAndAnalytical()
		if err != nil {
			beego.Error(err)
			c.SetJson(501, nil, "获取redisInfo失败")
		}
		c.SetJson(200, myAws.SourceInfo.ElastiCacheClusters.InstancesInfo, "")

	case "es":
		err := myAws.GetESAndAnalytical()
		if err != nil {
			beego.Error(err)
			c.SetJson(501, nil, "获取ESInfo失败")
		}
		c.SetJson(200, myAws.SourceInfo.ESCluster.DomainInfo, "")

	case "kafka":
		err := myAws.GetKafkaAndAnalytical()
		if err != nil {
			beego.Error(err)
			c.SetJson(501, nil, "获取KafkaInfo失败")
		}
		c.SetJson(200, myAws.SourceInfo.KafkaCluster.Cluster, "")
	}
	return
}

// @Title AWS 资源监控
// @Description AWS 资源监控
// @Param   awsRegion      query     string   	true     "awsRegion"
// @Param   service        query     string   	true     "service class"
// @Success 0 {id} int64
// @Failure 1 AWS 资源监控
// @Failure 2 User not found
// @router /resource/tag [get]
func (c *ResourceController) Tag() {
	class := c.GetString("class")

	//data, backPath, err := backup.BackupService(class)
	serviceObj, err := backup.GetServiceTag(class)
	if err != nil {
		c.SetJson(1, nil, "获取 服务tag列表 失败")
		return
	}
	c.SetJson(0, serviceObj, "")
	return
}
