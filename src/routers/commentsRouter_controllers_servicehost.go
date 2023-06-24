package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers/servicehost:ServiceHostController"] = append(beego.GlobalControllerRouter["controllers/servicehost:ServiceHostController"],
		beego.ControllerComments{
			Method:           "GetHostByServiceId",
			Router:           `/serviceHost/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
