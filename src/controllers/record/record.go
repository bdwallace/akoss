package recordcontrollers

import (
	"controllers"
	"models"

	"github.com/astaxie/beego/orm"
)

type RecordController struct {
	controllers.BaseController
}

// @Title 获取 service by project id
// @Description 根据 project id 查询所有 service 信息
// @Param   project_id      query     int  		true         "project id"
// @Success 0 {object} models.Service
// @Failure 1 获取 service by project id 失败
// @Failure 2 User not found
// @router /record/list [get]
func (c *RecordController) RecordList() {
	deployId := c.GetString("deployId")
	var records []orm.Params
	// if common.GetInt(deployId) <= 0 {
	// 	// timeNow := c.GetString("time")
	// 	o := orm.NewOrm()
	// 	o.Raw("SELECT * FROM `record` where deploy_id=? and created_at> ? ORDER BY `id` ASC ", deployId, timeNow).Values(&records)
	// } else {
	o := orm.NewOrm()
	o.Raw("SELECT * FROM `t_record` where deploy_id=? ORDER BY `id` ASC ", deployId).Values(&records)
	// }
	c.SetJson(0, records, "")
	return

}

// @Title 获取 record count by deploy id
// @Description 根据 project id 查询所有 service 信息
// @Param   id      query     int  		true         "deploy id"
// @Success 0 [3, 2, 1]
// @Failure 1 获取 service by project id 失败
// @Failure 2 User not found
// @router /record/count [get]
func (c *RecordController) RecordCount() {
	deployId, _ := c.GetInt("deployId", 0)
	recordCountList := new(orm.ParamsList)
	recordCountList, _ = models.GetRecordCount(deployId)

	c.SetJson(0, recordCountList, "")
	return

}

// @Title 获取 record count by deploy id
// @Description 根据 project id 查询所有 service 信息
// @Param   id      query     int  		true         "deploy id"
// @Success 0 [3, 2, 1]
// @Failure 1 获取 service by project id 失败
// @Failure 2 User not found
// @router /record/listcount [get]
func (c *RecordController) RecordListCount() {
	deployId, _ := c.GetInt("deployId", 0)
	count, _ := c.GetInt("count", 0)
	recordList, _ := models.GetRecordListCount(deployId, count)

	c.SetJson(0, recordList, "")
	return

}
