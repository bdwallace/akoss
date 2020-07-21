package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "Build",
            Router: `/walle/build`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "DockerPs",
            Router: `/walle/dockerps`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "Download",
            Router: `/walle/download/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "LineChange",
            Router: `/walle/linechange`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "LineGet",
            Router: `/walle/lineget`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "Reload",
            Router: `/walle/reload/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "Restart",
            Router: `/walle/restart/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "Stop",
            Router: `/walle/stop/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "GetTagList",
            Router: `/walle/taglist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/walle:WalleController"] = append(beego.GlobalControllerRouter["controllers/walle:WalleController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/walle/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
