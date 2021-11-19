package operationRecord

import (
	"controllers"
	"models"

	"github.com/astaxie/beego/orm"
)

type OperationRecordController struct {
	controllers.BaseController
}

// @Title 获取 operation_record count by service_id, class
// @Description 根据 service_id, class 查询所有 count 信息
// @Param   service_id      query     int  		true         "service id"
// @Param   class 			query     string  	true         "class"
// @Success 0 [3, 2, 1]
// @Failure 1 获取 operation_record by service_id, class, count 失败
// @Failure 2 User not found
// @router /recordoperation/count [get]
func (c *OperationRecordController) RecordCount() {
	serviceId, _ := c.GetInt("serviceId", 0)
	class := c.GetString("class")
	recordCountList := new(orm.ParamsList)
	recordCountList, _ = models.GetOperationRecordCount(serviceId, class)

	c.SetJson(0, recordCountList, "")
	return

}

// @Title 获取 operation_record count by service_id, class, count
// @Description 根据 service_id, class, count 查询所有 operation_record 信息
// @Param   id      query     int  		true         "deploy id"
// @Param   class 			query     string  	true         "class"
// @Param   count			query     int  		true         "count"
// @Success 0 [3, 2, 1]
// @Failure 1 获取 service by project id 失败
// @Failure 2 User not found
// @router /recordoperation/listcount [get]
func (c *OperationRecordController) RecordListCount() {
	serviceId, _ := c.GetInt("serviceId", 0)
	class := c.GetString("class")
	count, _ := c.GetInt("count", 0)
	recordList, _ := models.GetOperationRecordListCount(serviceId, class, count)

	c.SetJson(0, recordList, "")
	return

}
