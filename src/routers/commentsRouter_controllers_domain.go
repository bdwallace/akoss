package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "AddDomain",
			Router:           `/domain`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "GetDomainByProjectForClass",
			Router:           `/domain/class/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "DomainHealthCheck",
			Router:           `/domain/health/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "DeleteDomainById",
			Router:           `/domain/id/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "GetDomainById",
			Router:           `/domain/id/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "GetDomainList",
			Router:           `/domain/list/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "GetAllMonitor",
			Router:           `/domain/monitor/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "GetDomainByPlatformAndClass",
			Router:           `/domain/platformAndClass`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "GetDomainByProjectsId",
			Router:           `/domain/projectId/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/domain:DomainController"] = append(beego.GlobalControllerRouter["controllers/domain:DomainController"],
		beego.ControllerComments{
			Method:           "GetDomainByServiceId",
			Router:           `/domain/serviceId/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
