// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"controllers"
	backupcontrollers "controllers/backup"
	cloudcontrollers "controllers/cloud"
	confcontrollers "controllers/conf"
	crontabcontroller "controllers/crontab"
	deploycontrollers "controllers/deploy"
	domaincontroller "controllers/domain"
	fastcontrollers "controllers/fast"
	hostcontrollers "controllers/host"
	inspectcontrollers "controllers/inspect"
	linkcontrollers "controllers/link"
	monitorcontrollers "controllers/monitor"
	"controllers/operationRecord"
	othercontrollers "controllers/other"
	platformcontroller "controllers/platform"
	projectcontrolers "controllers/projects"
	recordcontrollers "controllers/record"
	resourcecontrollers "controllers/resource"
	robotcontroller "controllers/robot"
	servicecontroller "controllers/service"
	taskcontrollers "controllers/task"
	usercontrollers "controllers/user"
	wallecontrollers "controllers/walle"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "UserToken", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		MaxAge:          5 * time.Minute,
	}))

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/loginbydocke", &controllers.LoginByDockerController{})
	beego.Router("/changePasswd", &controllers.ChangePasswdController{})
	//beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/changeProject", &controllers.ProjectChangeController{})

	beego.Router("/api/get/user", &usercontrollers.UserController{})
	beego.Router("/api/delete/user", &usercontrollers.UserController{})
	//在使用Put方法时，页面传给后端的id始终为零，改成get方法就可以！记一下
	beego.Router("/api/user/reset", &usercontrollers.UserController{}, "get:Put")
	beego.Router("/", &controllers.MainController{})

	akNs := beego.NewNamespace("/api",
		beego.NSInclude(
			&projectcontrolers.ProjectController{},
		),
		beego.NSInclude(
			&confcontrollers.ConfController{},
		),
		beego.NSInclude(
			&servicecontroller.ServiceController{},
		),
		beego.NSInclude(
			&hostcontrollers.HostController{},
		),
		beego.NSInclude(
			&taskcontrollers.TaskController{},
		),
		beego.NSInclude(
			&deploycontrollers.DeployController{},
		),
		beego.NSInclude(
			&wallecontrollers.WalleController{},
		),
		beego.NSInclude(
			&recordcontrollers.RecordController{},
		),
		beego.NSInclude(
			&othercontrollers.OtherController{},
		),
		beego.NSInclude(
			&domaincontroller.DomainController{},
		),
		beego.NSInclude(
			&platformcontroller.PlatformController{},
		),
		beego.NSInclude(
			&operationRecord.OperationRecordController{},
		),
		beego.NSInclude(
			&recordcontrollers.RecordController{},
		),
		beego.NSInclude(
			&resourcecontrollers.ResourceController{},
		),
		beego.NSInclude(
			&crontabcontroller.CrontabController{},
		),
		beego.NSInclude(
			&crontabcontroller.CrontabLogController{},
		),
		beego.NSInclude(
			&hostcontrollers.HostCheckController{},
		),
		beego.NSInclude(
			&linkcontrollers.LinkController{},
		),
		beego.NSInclude(
			&fastcontrollers.SgroupController{},
		),
		beego.NSInclude(
			&backupcontrollers.BackupController{},
		),
		beego.NSInclude(
			&robotcontroller.RobotController{},
		),
		beego.NSInclude(
			&inspectcontrollers.GrafanaController{},
		),
		beego.NSInclude(
			&monitorcontrollers.MonitorController{},
		),
		beego.NSInclude(
			&cloudcontrollers.CloudController{},
		),
	)
	beego.AddNamespace(akNs)
}
