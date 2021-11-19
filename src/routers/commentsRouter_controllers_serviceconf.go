package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers/serviceconf:ServiceConfController"] = append(beego.GlobalControllerRouter["controllers/serviceconf:ServiceConfController"],
		beego.ControllerComments{
			Method:           "GetConfByProjectId",
			Router:           `/serviceConf/projectId/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
