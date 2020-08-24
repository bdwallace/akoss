package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/monitor:MonitorController"] = append(beego.GlobalControllerRouter["controllers/monitor:MonitorController"],
        beego.ControllerComments{
            Method: "Monitor",
            Router: `/monitor/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
