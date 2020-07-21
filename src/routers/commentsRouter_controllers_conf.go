package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/conf:ConfController"] = append(beego.GlobalControllerRouter["controllers/conf:ConfController"],
        beego.ControllerComments{
            Method: "AddConf",
            Router: `/conf/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/conf:ConfController"] = append(beego.GlobalControllerRouter["controllers/conf:ConfController"],
        beego.ControllerComments{
            Method: "CopyConfById",
            Router: `/conf/copy`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/conf:ConfController"] = append(beego.GlobalControllerRouter["controllers/conf:ConfController"],
        beego.ControllerComments{
            Method: "GetConfById",
            Router: `/conf/id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/conf:ConfController"] = append(beego.GlobalControllerRouter["controllers/conf:ConfController"],
        beego.ControllerComments{
            Method: "DeleteConfById",
            Router: `/conf/id/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/conf:ConfController"] = append(beego.GlobalControllerRouter["controllers/conf:ConfController"],
        beego.ControllerComments{
            Method: "GetAllConfs",
            Router: `/conf/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/conf:ConfController"] = append(beego.GlobalControllerRouter["controllers/conf:ConfController"],
        beego.ControllerComments{
            Method: "GetConfByProjectId",
            Router: `/conf/projectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
