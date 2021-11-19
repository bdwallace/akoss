package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers/resource:ResourceController"] = append(beego.GlobalControllerRouter["controllers/resource:ResourceController"],
		beego.ControllerComments{
			Method:           "Aws",
			Router:           `/resource/aws`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/resource:ResourceController"] = append(beego.GlobalControllerRouter["controllers/resource:ResourceController"],
		beego.ControllerComments{
			Method:           "Tag",
			Router:           `/resource/tag`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
