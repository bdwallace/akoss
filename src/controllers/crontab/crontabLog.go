package crontabcontroller

import (
	"controllers"
	"models"
)


type CrontabLogController struct {
	controllers.BaseController
}

// @Title get crontab log by page
// @Description get crontab by id
// @Param   plage      query     int 		true       "crontab log page"
// @Success 0 {object} models.CrontabLog
// @Failure 1 获取 crontab by page 失败
// @Failure 2 User not found
// @router /crontabLog/page/ [get]
func (c *CrontabLogController) GetAllCrontabLog() {
	page, _ := c.GetInt("page", 0)
	start := 0
	length, _ := c.GetInt("length", 200000)
	if page > 0 {
		start = (page - 1) * length
	}

	count, data := models.GetCrontabLogAllPage(start, length)

	c.SetJson(0, map[string]interface{}{"total": count, "currentPage": page, "table_data": data}, "")
	// c.SetJson(0, data, "查询成功！")
	return
}

