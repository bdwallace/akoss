package backupcontroller

import (
	"controllers"
	"library/backup"
)

type BackupController struct {
	controllers.BaseController
}


// @Title get backup path
// @Description get backup path
// @Success 0 back path string
// @Failure 1 删除 deploy by deploy id 失败
// @Failure 2 User not found
// @router /backup/sql [get]
func (c *BackupController) BackupAkgoSql() {
	backPath, err := backup.BackupSql()
	if err != nil {
		c.SetJson(1, nil, "")
		return
	}
	c.SetJson(0, backPath, "")
	return
}


// @Title get backup path
// @Description get backup path
// @Param   id      query     int 		true       "deploy id"
// @Success 0 back path string
// @Failure 1 删除 deploy by deploy id 失败
// @Failure 2 User not found
// @router /backup/service [get]
func (c *BackupController) GetService() {
	class := c.GetString("class")

	_, backPath, err := backup.BackupService(class)
	if err != nil {
		c.SetJson(1, nil, "")
		return
	}
	c.SetJson(0, backPath, "")
	// c.SetJson(0, data, backPath)
	return
}

