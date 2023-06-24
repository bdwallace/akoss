package crontabcontroller

import (
	"controllers"
	"encoding/json"
	"fmt"
	"library/crontab"
	"models"
)

type CrontabController struct {
	controllers.BaseController
}

// @Title get crontab by id
// @Description get crontab by id
// @Param   id      query     int 		true       "crontab id"
// @Success 0 {object} models.Crontab
// @Failure 1 获取 crontab by id 失败
// @Failure 2 User not found
// @router /crontab/id/ [get]
func (c *CrontabController) GetCrontab() {
	crontabId, _ := c.GetInt("crontabId", 0)
	crontab, _ := models.GetCrontabById(crontabId)

	c.SetJson(0, crontab, "")
	return
}

// @Title get all crontab
// @Description get all crontab
// @Success 0 {object} models.Crontab
// @Failure 1 获取 get all crontab 失败
// @Failure 2 User not found
// @router /crontab/list/ [get]
func (c *CrontabController) GetAllCrontab() {
	data, _ := models.GetCrontabAll()

	c.SetJson(0, data, "修改成功")
	return
}

// @Title add crontab by new crontab
// @Description add crontab by new crontab
// @Success 0 {object} models.Crontab
// @Failure 1 添加 add crontab by new crontab 失败
// @Failure 2 User not found
// @router /crontab/ [post]
func (c *CrontabController) AddCrontab() {
	//projectId,_:=c.GetInt("projectId",0)
	//beego.Info(string(c.Ctx.Input.RequestBody))
	var crontab models.Crontab
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &crontab)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}

	if crontab.Id != 0 {
		err = models.UpCrontab(&crontab)
	} else {
		_, err = models.AddCrontab(&crontab)
	}
	if err != nil {
		c.SetJson(1, nil, "数据库更新错误")
		return
	}
	c.SetJson(0, crontab, "修改成功")
	return
}

// @Title copy crontab by crontab id
// @Description copy crontab by crontab id
// @Param   id      query     int 		true       "crontab id"
// @Success 0 {object} models.Crontab
// @Failure 1 复制 crontab by crontab id 失败
// @Failure 2 User not found
// @router /crontab/copy/ [get]
func (c *CrontabController) CopyCrontab() {
	if c.User == nil || c.User.Id == 0 {
		c.SetJson(2, nil, "not login")
		return
	}

	crontabId, _ := c.GetInt("crontabId", 0)

	Crontab, _ := models.GetCrontabById(crontabId)
	Crontab.Name = Crontab.Name + " - copy"
	Crontab.Id = 0
	_, err := models.AddCrontab(Crontab)
	if err != nil {
		c.SetJson(1, nil, "复制失败")
	}
	c.SetJson(0, Crontab, "")
	return

}

// @Title delete crontab by crontab id
// @Description delete crontab by crontab id
// @Param   id      query     int 		true       "crontab id"
// @Success 0 {object} models.Crontab
// @Failure 1 删除 crontab by crontab id 失败
// @Failure 2 User not found
// @router /crontab/id/ [delete]
func (c *CrontabController) DeleteCrontab() {
	crontabId, _ := c.GetInt("id")

	err := models.DeleteCrontab(crontabId)
	if err != nil {
		c.SetJson(1, nil, err.Error())
		return
	}

	c.SetJson(0, nil, "删除成功")
	return
}

// @Title exec crontab by crontab id
// @Description exec crontab by crontab id
// @Param   id      query     int 		true       "crontab id"
// @Success 0 {object} models.Crontab
// @Failure 1 执行 crontab by crontab id 失败
// @Failure 2 User not found
// @router /crontab/exec/ [get]
func (c *CrontabController) ExecCrontab() {
	crontabId, _ := c.GetInt("crontabId", 0)
	auto, _ := c.GetInt("auto", 0)
	err := crontab.Execute(crontabId, auto)
	if err != nil {
		c.SetJson(1, nil, fmt.Sprintf("执行计划任务失败, %s", err))
		return
	}

	c.SetJson(0, nil, "")
	return
}

// @Title reload crontab
// @Description reload crontab
// @Success 0 {object} models.Crontab
// @Failure 1 重启 reload crontab 失败
// @Failure 2 User not found
// @router /crontab/reload/ [get]
func (c *CrontabController) ReloadCrontab() {
	err := crontab.Crontab()
	if err != nil {
		c.SetJson(1, nil, "重新加载计划任务失败！")
		return
	}

	c.SetJson(0, nil, "")
	return
}
