package hostcontrollers

import (
	"controllers"
	"encoding/json"
	"fmt"
	"library/aws"
	"library/common"
	"library/components"
	"models"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type HostCheckController struct {
	controllers.BaseController
}

type ResAwsHost struct {
	Region          string
	EC2InstanceInfo *aws.EC2InstanceInfo
}

type ResAwsHosts struct {
	Region           string
	EC2InstanceInfoS []*aws.EC2InstanceInfo
}



// @Title get awsrsync by ip
// @Description get awsrsync by ip
// @Param   ip      query     int 		true       "host ip"
// @Success 0 {object} models.Host
// @Failure 1 获取 awsrsync by ip 失败
// @Failure 2 User not found
// @router /awsHost/awsrsync/ [post]
func (c *HostCheckController)AwsRsync() {
	// hostId, _ := c.GetInt("hostId", 0)
	// host, _ := models.GetHostById(hostId)
	// ip := c.GetString("ip")
	beego.Info(string(c.Ctx.Input.RequestBody))
	var host models.Host
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &host)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	project, err := models.GetProjectById(host.Project.Id)
	if err != nil {
		c.SetJson(1, nil, "查询环境参数失败")
		return
	}

	var myAws aws.BaseAws
	myAws.Alias = project.Alias

	regions, err := myAws.GetAllRegion()
	if err != nil {
		beego.Error(fmt.Sprintf("--同步aws到host失败！--"))
		c.SetJson(1, nil, "")
		return
	}

	resAwsHoatS := make([]*ResAwsHosts, 0)
	for _, region := range regions {
		myAws.Region = region
		myAws.SourceInfo = new(aws.ResourceInfo)
		_ = myAws.InitSession()

		err := myAws.GetEc2CollocationSubnet()
		if err != nil {
			beego.Error(err)
			continue
		}

		resEc2InstanceInfos := make([]*aws.EC2InstanceInfo, 0)
		for _, v := range myAws.SourceInfo.EC2Info.EC2Reservations {
			resEc2InstanceInfos = append(resEc2InstanceInfos, v.Instances...)
		}
		res := &ResAwsHosts{Region: region, EC2InstanceInfoS: resEc2InstanceInfos}
		resAwsHoatS = append(resAwsHoatS, res)
	}

	resAwsHost := make([]*ResAwsHost, 0)
	for _, v := range resAwsHoatS {
		for _, vv := range v.EC2InstanceInfoS {
			if host.PrivateIp == vv.PrivateIP {
				res := &ResAwsHost{Region: v.Region, EC2InstanceInfo: vv}
				resAwsHost = append(resAwsHost, res)
			}
		}
	}

	c.SetJson(0, resAwsHost, "")
	return
}


// @Title get awsRsyncUpdate by project id
// @Description get awsrsync by ip
// @Param   id      query     int 		true       "project id"
// @Success 0 {object} models.Host
// @Failure 1 获取 awsRsyncUpdate by project id 失败
// @Failure 2 User not found
// @router /awsHost/awsRsyncUpdate/ [get]
func (c *HostCheckController) AwsRsyncUpdate() {
	projectId, _ := c.GetInt("projectId", 0)
	project, err := models.GetProjectById(projectId)
	if err != nil {
		beego.Error(fmt.Sprintf("--查询第 %d 套主机失败！--", c.Project.Id))
		c.SetJson(1, nil, "")
		return
	}
	hosts, err := models.GetHostByProjectId(projectId)
	if err != nil {
		beego.Error(fmt.Sprintf("--查询第 %d 套主机失败！--", c.Project.Id))
		c.SetJson(1, nil, "")
		return
	}

	var myAws aws.BaseAws
	myAws.Alias = project.Alias

	regions, err := myAws.GetAllRegion()
	if err != nil {
		beego.Error(fmt.Sprintf("--同步aws到host失败！--"))
		c.SetJson(1, nil, "")
		return
	}

	resAwsHoatS := make([]*ResAwsHosts, 0)
	for _, region := range regions {
		myAws.Region = region
		myAws.SourceInfo = new(aws.ResourceInfo)
		_ = myAws.InitSession()

		err := myAws.GetEc2CollocationSubnet()
		if err != nil {
			beego.Error(err)
			continue
		}

		resEc2InstanceInfos := make([]*aws.EC2InstanceInfo, 0)
		for _, v := range myAws.SourceInfo.EC2Info.EC2Reservations {
			resEc2InstanceInfos = append(resEc2InstanceInfos, v.Instances...)
		}
		res := &ResAwsHosts{Region: region, EC2InstanceInfoS: resEc2InstanceInfos}
		resAwsHoatS = append(resAwsHoatS, res)
	}

	var resDataList []map[string]string
	for _, host := range hosts {
		for _, v := range resAwsHoatS {
			for _, vv := range v.EC2InstanceInfoS {
				if host.PrivateIp == vv.PrivateIP {
					if host.Region != v.Region || host.InstanceId != vv.InstanceID || host.InstanceType != vv.InstanceType || host.PublicIp != vv.PublicIP {
						resData := map[string]string{host.PrivateIp:""}
						if host.Region != v.Region {
							resData[host.PrivateIp] = fmt.Sprintf("%s,%s to %s", resData[host.PrivateIp], host.Region, v.Region)
							host.Region = v.Region
						}
						if host.InstanceId != vv.InstanceID {
							resData[host.PrivateIp] = fmt.Sprintf("%s,%s to %s", resData[host.PrivateIp], host.InstanceId, vv.InstanceID)
							host.InstanceId = vv.InstanceID
						}
						if host.InstanceType != vv.InstanceType {
							resData[host.PrivateIp] = fmt.Sprintf("%s,%s to %s", resData[host.PrivateIp], host.InstanceType, vv.InstanceType)
							host.InstanceType = vv.InstanceType
						}
						if host.PublicIp != vv.PublicIP {
							resData[host.PrivateIp] = fmt.Sprintf("%s,%s to %s", resData[host.PrivateIp], host.PublicIp, vv.PublicIP)
							host.PublicIp = vv.PublicIP
						}
						err := models.UpdateHostByHost(host)
						if err != nil {
							beego.Error(fmt.Sprintf("--更新id:%d 的host失败！--", host.Id))
							resData[host.PrivateIp] = fmt.Sprintf("%s%s", "failure", resData[host.PrivateIp])
						} else {
							resData[host.PrivateIp] = fmt.Sprintf("%s%s", "success", resData[host.PrivateIp])
						}

						resDataList = append(resDataList, resData)
						break
					}
				}
			}
		}

	}

	c.SetJson(0, resDataList, "")
	return
}



// @Title get AwsStatus by instanceId and region
// @Description get AwsStatus by instanceId and region
// @Param   id      query     int 		true       "instanceId"
// @Param   id      query     int 		true       "region"
// @Failure 1 获取 AwsStatus by instanceId and region 失败
// @Failure 2 User not found
// @router /awsHost/awsStatus/ [post]
func (c *HostCheckController) AwsStatus() {

	beego.Info(string(c.Ctx.Input.RequestBody))
	var host models.Host
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &host)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	var myAws aws.BaseAws
	myAws.Alias = host.Project.Alias

	myAws.Region = host.Region
	myAws.SourceInfo = new(aws.ResourceInfo)
	_ = myAws.InitSession()

	result, err := myAws.GetEC2Status(host.InstanceId)
	if err != nil {
		beego.Error(err)
		c.SetJson(0, nil, "")
		return
	}

	c.SetJson(0, result.Reservations[0].Instances[0].State.Name, "")
	return
}



// @Title get AwsStatus by instanceId and region
// @Description get AwsStatus by instanceId and region
// @Param   id      query     int 		true       "instanceId"
// @Param   id      query     int 		true       "region"
// @Failure 1 获取 AwsStatus by instanceId and region 失败
// @Failure 2 User not found
// @router /awsHost/awsStop/ [post]
func (c *HostCheckController) AwsStop() {

	beego.Info(string(c.Ctx.Input.RequestBody))
	var host models.Host
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &host)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	url := "http://" + host.Project.Nacos + "/nacos/v1/cs/configs"

	err = components.LineChangeHost(url, &host, "off")
	if err != nil {
		c.SetJson(1, err.Error(), "获取nacos接口数据返回失败")
		return
	}

	time.Sleep(time.Duration(10)*time.Second)

	var myAws aws.BaseAws
	myAws.Alias = host.Project.Alias
	myAws.Region = host.Region
	myAws.SourceInfo = new(aws.ResourceInfo)
	_ = myAws.InitSession()

	result, err := myAws.StopEC2Status(host.InstanceId)
	if err != nil {
		beego.Error(err)
		c.SetJson(1, nil, "")
		return
	}

	c.SetJson(0, result, "")
	return
}



// @Title get AwsStatus by instanceId and region
// @Description get AwsStatus by instanceId and region
// @Param   id     		query     string 		true       "instance id"
// @Param   region      query     string 		true       "region"
// @Param   hostId      query     string 		true       "host id"
// @Failure 1 获取 AwsStatus by instanceId and region 失败
// @Failure 2 User not found
// @router /awsHost/awsStart/ [post]
func (c *HostCheckController) AwsStart() {


	beego.Info(string(c.Ctx.Input.RequestBody))
	var host models.Host
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &host)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	var myAws aws.BaseAws
	myAws.Alias = host.Project.Alias
	myAws.Region = host.Region
	myAws.SourceInfo = new(aws.ResourceInfo)
	_ = myAws.InitSession()

	result, err := myAws.StartEC2Status(host.InstanceId)
	if err != nil {
		beego.Error(err)
		c.SetJson(1, nil, "")
		return
	}

	time.Sleep(time.Duration(10)*time.Second)

	url := "http://" + host.Project.Nacos + "/nacos/v1/cs/configs"

	err = components.LineChangeHost(url, &host, "on")
	if err != nil {
		c.SetJson(1, err.Error(), "获取nacos接口数据返回失败")
		return
	}


	c.SetJson(0, result, "")
	return
}




// @Title get AwsStatus by instanceId and region
// @Description get AwsStatus by instanceId and region
// @Param   id     		query     string 		true       "instance id"
// @Param   region      query     string 		true       "region"
// @Param   hostId      query     string 		true       "host id"
// @Failure 1 获取 AwsStatus by instanceId and region 失败
// @Failure 2 User not found
// @router /awsHost/awsCheckStop/ [get]
func (c *HostCheckController) AwsCheckStop() {

	projectId, _ := c.GetInt("projectId", 0)
	project, _ := models.GetProjectById(projectId)

	newTime := common.StrToInt(time.Now().Format("20060102150405")[8:12])
	var myAws aws.BaseAws

	myAws.Alias = project.Alias

	url := "http://" + project.Nacos + "/nacos/v1/cs/configs"

	hosts, err := models.GetHostStopCheck(project.Id)
	if err != nil {
		beego.Error(err)
		c.SetJson(1, nil, "")
		return
	}

	var result string
	for _, host := range hosts {
		myAws.Region = host.Region
		myAws.SourceInfo = new(aws.ResourceInfo)
		_ = myAws.InitSession()

		endTime := common.StrToInt(strings.Replace(host.EndTime, ":", "", -1))
		startTime := common.StrToInt(strings.Replace(host.StartTime, ":", "", -1))
		if startTime < newTime && newTime < endTime {
			err = components.LineChangeHost(url, host, "off")
			if err != nil {
				c.SetJson(1, err.Error(), "获取nacos接口数据返回失败")
				return
			}

			_, err := myAws.StopEC2Status(host.InstanceId)
			if err != nil {
				beego.Error(err)
				c.SetJson(1, nil, "")
				return
			}
			result = fmt.Sprintf("%s%s_%s_关闭#", result, host.UseIp, host.InstanceId)
		} else {
			_, err := myAws.StartEC2Status(host.InstanceId)
			if err != nil {
				beego.Error(err)
				c.SetJson(1, nil, "")
				return
			}

			result = fmt.Sprintf("%s%s_%s_启动#", result, host.UseIp, host.InstanceId)
			err = components.LineChangeHost(url, host, "on")
			if err != nil {
				c.SetJson(1, err.Error(), "获取nacos接口数据返回失败")
				return
			}
		}
	}

	c.SetJson(0, result, "")
	return
}




// @Title get AwsStatus by instanceId and region
// @Description get AwsStatus by instanceId and region
// @Param   id     		query     string 		true       "instance id"
// @Param   region      query     string 		true       "region"
// @Param   hostId      query     string 		true       "host id"
// @Failure 1 获取 AwsStatus by instanceId and region 失败
// @Failure 2 User not found
// @router /awsHost/awsSetStopTime/ [post]
func (c *HostCheckController) AwsSetStopTime() {
	beego.Info(string(c.Ctx.Input.RequestBody))
	var hosts []*models.Host
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &hosts)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	err = models.SetHostStopTime(hosts)
	if err != nil {
		c.SetJson(1, err, "批量更新主机出错")
		return
	}

	c.SetJson(0, nil, "修改成功")
	return
}



// @Title get AwsStatus by instanceId and region
// @Description get AwsStatus by instanceId and region
// @Param   id     		query     string 		true       "instance id"
// @Param   region      query     string 		true       "region"
// @Param   hostId      query     string 		true       "host id"
// @Failure 1 获取 AwsStatus by instanceId and region 失败
// @Failure 2 User not found
// @router /awsHost/awsStopOnlineHost/ [get]
func (c *HostCheckController) AwsStopOnlineHost() {
	hostId, _ := c.GetInt("hostId", 0)
	host, _ := models.GetHostById(hostId)

	level, _ := models.GetProjectById(host.Project.Id)
	url := "http://" + level.Nacos + "/nacos/v1/cs/configs"

	ipList, err := components.Line(url)
	if err != nil {
		c.SetJson(1, err.Error(), "访问nacos的下线列表失败")
		return
	}

	var newList string
	for _, ipPort := range strings.Split(ipList, ",") {
		if strings.Index(ipPort, host.UseIp) == -1 {
			if newList == "" {
				newList = ipPort
			} else {
				newList = fmt.Sprintf("%s,%s", newList, ipPort)
			}
		}
	}

	reStr := fmt.Sprintf("列表由：\n%s \n变成：\n%s", ipList, newList)

	c.SetJson(0, reStr, "")
	return

}


