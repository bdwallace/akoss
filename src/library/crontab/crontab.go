package crontab

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"models"
)

func Crontab() (err error) {
	// CrontabList, err := models.GetCrontabByAuto(1)
	CrontabList, err := models.GetCrontabAll()
	if err != nil || len(CrontabList) == 0 {
		beego.Info("-----no crontab")
		return
	}

	for _, crontab := range CrontabList {
		if crontab.Auto == 1 {
			beego.Info("----加载到自动备份任务-", crontab.Name)
			id := crontab.Id
			cron := toolbox.NewTask(crontab.Name, crontab.Crontab, func() error { Execute(id, 1); return nil })
			toolbox.AddTask(crontab.Name, cron)
		} else {
			beego.Info("----删除自动备份任务-", crontab.Name)
			toolbox.DeleteTask(crontab.Name)
		}
	}

	toolbox.StartTask()

	return
}
