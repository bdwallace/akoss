package cloudcontrollers

import (
	"controllers"
	"library/cloud/aliyun"
	"library/cloud/baiduyun"
	"library/cloud/qiniu"
	"library/cloud/tencent"
	"library/cloud/tianyi"
	"library/common"
	"models"
	"strings"
	"time"

	"encoding/json"

	"github.com/astaxie/beego"
)

type CloudController struct {
	controllers.BaseController
}

// @Title get cloud by id
// @Description get cloud by id
// @Param   id      query     int 		true       "cloud id"
// @Success 0 {object} models.Cloud
// @Failure 1 获取 cloud by id 失败
// @Failure 2 User not found
// @router /cloud/id/ [get]
func (c *CloudController) GetCloudById() {
	cloudId, _ := c.GetInt("id")
	cloud, err := models.GetCloudById(cloudId)
	if err != nil {
		c.SetJson(1, err, "获取 cloud 失败")
		return
	}

	c.SetJson(0, cloud, "")
	return
}

// @Title get all cloud
// @Description get all cloud
// @Success 0 {object} models.Cloud
// @Failure 1 获取 get all cloud 失败
// @Failure 2 User not found
// @router /cloud/list/ [get]
func (c *CloudController) GetAllCloud() {
	data, err := models.GetCloudAll()
	if err != nil {
		c.SetJson(1, err, "获取 cloud 列表 失败")
		return
	}

	c.SetJson(0, data, "查询成功")
	return
}

// @Title add cloud by new cloud
// @Description add cloud by new cloud
// @Success 0 {object} models.Cloud
// @Failure 1 添加 add cloud by new cloud 失败
// @Failure 2 User not found
// @router /cloud/ [post]
func (c *CloudController) AddCloud() {
	//projectId,_:=c.GetInt("projectId",0)
	beego.Info(string(c.Ctx.Input.RequestBody))
	var cloud *models.Cloud
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cloud)
	if err != nil {
		c.SetJson(1, nil, "格式错误")
		return
	}

	o, _ := models.OrmBegin()

	if cloud.Id != 0 {
		if err = models.UpdateCloud(o, cloud); err != nil {
			models.OrmRollback(o)
			c.SetJson(1, nil, "修改失败: "+err.Error())
			return
		}

	} else {
		if _, err = models.AddCloud(o, cloud); err != nil {
			models.OrmRollback(o)
			c.SetJson(1, nil, "新增失败:"+err.Error())
			return
		}
	}

	if err := models.OrmCommit(o); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "操作失败: "+err.Error())
		return
	} else {
		c.SetJson(0, cloud, "操作成功")
		return
	}
}

// @Title copy cloud by cloud id
// @Description copy cloud by cloud id
// @Param   id      query     int 		true       "cloud id"
// @Success 0 {object} models.Cloud
// @Failure 1 复制 cloud by cloud id 失败
// @Failure 2 User not found
// @router /cloud/copy/ [get]
func (c *CloudController) CopyCloud() {
	// if c.User == nil || c.User.Id == 0 {
	// 	c.SetJson(2, nil, "not login")
	// 	return
	// }

	cloudId, _ := c.GetInt("id")

	cloud, err := models.GetCloudById(cloudId)
	if err != nil {
		c.SetJson(1, nil, "查询失败")
		return
	}

	cloud.Name = cloud.Name + " - copy"
	cloud.Id = 0

	o, _ := models.OrmBegin()
	if _, err := models.AddCloud(o, cloud); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "复制失败")
		return
	}

	if err = models.OrmCommit(o); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "操作失败: "+err.Error())
		return
	} else {
		c.SetJson(0, nil, "操作成功")
		return
	}
}

// @Title delete cloud by cloud id
// @Description delete cloud by cloud id
// @Param   id      query     int 		true       "cloud id"
// @Success 0 {object} models.Cloud
// @Failure 1 删除 cloud by cloud id 失败
// @Failure 2 User not found
// @router /cloud/id/ [delete]
func (c *CloudController) DeleteCloud() {
	cloudId, _ := c.GetInt("id")

	cloud, err := models.GetCloudById(cloudId)
	if err != nil {
		c.SetJson(1, nil, "查询异常: "+err.Error())
		return
	}

	o, _ := models.OrmBegin()

	if err = models.DeleteCloud(o, cloud); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "删除失败: "+err.Error())
		return
	}

	if err = models.OrmCommit(o); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "操作失败: "+err.Error())
		return
	} else {
		c.SetJson(0, nil, "操作成功")
		return
	}
}

// @Title sync balance cloud by cloud id
// @Description sync balance cloud by cloud id
// @Param   id      query     int 		true       "cloud id"
// @Success 0 {object} models.Cloud
// @Failure 1 同步 cloud by cloud id 失败
// @Failure 2 User not found
// @router /cloud/sync/ [get]
func (c *CloudController) SyncCloud() {
	cloudId, _ := c.GetInt("id")

	cloud, err := models.GetCloudById(cloudId)
	if err != nil {
		c.SetJson(1, nil, "查询异常: "+err.Error())
		return
	}

	if cloud.Class == "阿里云" {
		data, err := aliyun.Balance(cloud.AccessKeyId, cloud.AccessSecret)
		if err != nil {
			c.SetJson(1, nil, "请求阿里云失败: "+err.Error())
			return
		}

		cloud.Balance = common.StrToFloat(strings.Replace(data.Data.AvailableAmount, ",", "", -1))
		cloud.UpdatedAt = time.Now()
	} else if cloud.Class == "腾讯云" {
		data, err := tencent.Balance(cloud.AccessKeyId, cloud.AccessSecret)
		if err != nil {
			c.SetJson(1, nil, "请求腾讯云失败: "+err.Error())
			return
		}

		cloud.Balance = data
		cloud.UpdatedAt = time.Now()
	} else if cloud.Class == "百度云" {
		data, err := baiduyun.Balance(cloud.AccessKeyId, cloud.AccessSecret)
		if err != nil {
			c.SetJson(1, nil, "请求腾讯云失败: "+err.Error())
			return
		}

		cloud.Balance = data
		cloud.UpdatedAt = time.Now()
	} else if cloud.Class == "七牛云" {
		data, err := qiniu.Balance(cloud.AccessKeyId,cloud.AccessSecret)
		if err != nil {
			c.SetJson(1, nil, "请求七牛云失败: "+err.Error())
			return
		}
		cloud.Balance = *data
		cloud.UpdatedAt = time.Now()
	} else if cloud.Class == "天翼云" {
		data, err := tianyi.Balance(cloud.AccessKeyId, cloud.AccessSecret)
		if err != nil {
			c.SetJson(1, nil, "请求天翼云失败: "+err.Error())
			return
		}
		cloud.Balance = *data
		cloud.UpdatedAt = time.Now()
	} else {
		c.SetJson(1, nil, "云平台类型错误: " + err.Error())
		return
	}

	o, _ := models.OrmBegin()

	if err = models.UpdateCloud(o, cloud); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "修改失败: "+err.Error())
		return
	}

	if err = models.OrmCommit(o); err != nil {
		models.OrmRollback(o)
		c.SetJson(1, nil, "操作失败: "+err.Error())
		return
	} else {
		c.SetJson(0, cloud, "操作成功")
		return
	}
}
