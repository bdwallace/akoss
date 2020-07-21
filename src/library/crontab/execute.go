package crontab


import (
	gopubssh "library/ssh"
	"models"

	"time"

	"github.com/astaxie/beego"
)

func Execute(id int, auto int) (err error) {
	crontab, _ := models.GetCrontabById(id)
	beego.Info("------正在执行任务：-----", crontab.Id)

	s, err := gopubssh.CommandLocal(crontab.Script, 360)
	switch {
	case err != nil:
		crontab.Result = err.Error()
		crontab.Status = 1
	case s.Error != nil:
		crontab.Result = s.Error.Error()
		crontab.Status = 1
		err = s.Error
	default:
		crontab.Result = s.Result
		crontab.Status = 0
	}

	crontab.UpdatedAt = time.Now()
	models.UpCrontab(crontab)

	var crontabLog models.CrontabLog
	crontabLog.CrontabName = crontab.Name
	crontabLog.Crontab = &models.Crontab{Id:crontab.Id}
	crontabLog.CreatedAt = crontab.UpdatedAt
	crontabLog.Status = crontab.Status
	crontabLog.Result = crontab.Result
	crontabLog.Auto = auto
	models.AddCrontabLog(&crontabLog)

	return
}
