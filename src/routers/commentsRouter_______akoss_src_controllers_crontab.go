package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/crontab:CrontabController"] = append(beego.GlobalControllerRouter["controllers/crontab:CrontabController"],
        beego.ControllerComments{
            Method: "AddCrontab",
            Router: `/crontab/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/crontab:CrontabController"] = append(beego.GlobalControllerRouter["controllers/crontab:CrontabController"],
        beego.ControllerComments{
            Method: "CopyCrontab",
            Router: `/crontab/copy/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/crontab:CrontabController"] = append(beego.GlobalControllerRouter["controllers/crontab:CrontabController"],
        beego.ControllerComments{
            Method: "ExecCrontab",
            Router: `/crontab/exec/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/crontab:CrontabController"] = append(beego.GlobalControllerRouter["controllers/crontab:CrontabController"],
        beego.ControllerComments{
            Method: "GetCrontab",
            Router: `/crontab/id/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/crontab:CrontabController"] = append(beego.GlobalControllerRouter["controllers/crontab:CrontabController"],
        beego.ControllerComments{
            Method: "DeleteCrontab",
            Router: `/crontab/id/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/crontab:CrontabController"] = append(beego.GlobalControllerRouter["controllers/crontab:CrontabController"],
        beego.ControllerComments{
            Method: "GetAllCrontab",
            Router: `/crontab/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/crontab:CrontabController"] = append(beego.GlobalControllerRouter["controllers/crontab:CrontabController"],
        beego.ControllerComments{
            Method: "ReloadCrontab",
            Router: `/crontab/reload/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/crontab:CrontabLogController"] = append(beego.GlobalControllerRouter["controllers/crontab:CrontabLogController"],
        beego.ControllerComments{
            Method: "GetAllCrontabLog",
            Router: `/crontabLog/page/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
