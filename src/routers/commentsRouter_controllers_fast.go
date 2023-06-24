package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers/fast:SgroupController"] = append(beego.GlobalControllerRouter["controllers/fast:SgroupController"],
		beego.ControllerComments{
			Method:           "DeleteSgroup",
			Router:           `/fast/deleteSgroup`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/fast:SgroupController"] = append(beego.GlobalControllerRouter["controllers/fast:SgroupController"],
		beego.ControllerComments{
			Method:           "GetSgroup",
			Router:           `/fast/getSgroup`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/fast:SgroupController"] = append(beego.GlobalControllerRouter["controllers/fast:SgroupController"],
		beego.ControllerComments{
			Method:           "SetSgroup",
			Router:           `/fast/setSgroup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
