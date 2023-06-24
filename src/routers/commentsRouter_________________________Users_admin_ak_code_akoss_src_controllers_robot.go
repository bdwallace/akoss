package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/robot:RobotController"] = append(beego.GlobalControllerRouter["controllers/robot:RobotController"],
        beego.ControllerComments{
            Method: "RobotReload",
            Router: `/robot/reload/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/robot:RobotController"] = append(beego.GlobalControllerRouter["controllers/robot:RobotController"],
        beego.ControllerComments{
            Method: "RobotRestart",
            Router: `/robot/restart/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
