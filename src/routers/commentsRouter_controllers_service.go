package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "AddService",
			Router:           `/service/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "GetServiceAllRelatedByProjectId",
			Router:           `/service/backend/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "CopyServiceById",
			Router:           `/service/copy`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "GetServiceDefaultConfKay",
			Router:           `/service/defaultKey/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "GetServiceRelatedById",
			Router:           `/service/id/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "DeleteServiceAndAllRelatedByServiceId",
			Router:           `/service/id/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "GetAllServices",
			Router:           `/service/list/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "GetServiceByProjectForName",
			Router:           `/service/name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "GetServiceAllRelatedByPlatformId",
			Router:           `/service/platformId/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/service:ServiceController"] = append(beego.GlobalControllerRouter["controllers/service:ServiceController"],
		beego.ControllerComments{
			Method:           "AddServiceAllRelatedByService",
			Router:           `/service/related/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
