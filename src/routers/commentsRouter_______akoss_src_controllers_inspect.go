package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/inspect:GrafanaController"] = append(beego.GlobalControllerRouter["controllers/inspect:GrafanaController"],
        beego.ControllerComments{
            Method: "PostInspectGrafana",
            Router: `/inspectgrafana`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/inspect:GrafanaController"] = append(beego.GlobalControllerRouter["controllers/inspect:GrafanaController"],
        beego.ControllerComments{
            Method: "ExecAllInspectGrafana",
            Router: `/inspectgrafana/exec/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/inspect:GrafanaController"] = append(beego.GlobalControllerRouter["controllers/inspect:GrafanaController"],
        beego.ControllerComments{
            Method: "GetInspectGrafanaById",
            Router: `/inspectgrafana/id/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/inspect:GrafanaController"] = append(beego.GlobalControllerRouter["controllers/inspect:GrafanaController"],
        beego.ControllerComments{
            Method: "DeleteInspectGrafanaById",
            Router: `/inspectgrafana/id/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/inspect:GrafanaController"] = append(beego.GlobalControllerRouter["controllers/inspect:GrafanaController"],
        beego.ControllerComments{
            Method: "GetAllInspectGrafana",
            Router: `/inspectgrafana/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
