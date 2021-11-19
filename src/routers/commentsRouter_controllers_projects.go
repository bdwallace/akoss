package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers/projects:ProjectController"] = append(beego.GlobalControllerRouter["controllers/projects:ProjectController"],
		beego.ControllerComments{
			Method:           "AddProject",
			Router:           `/project/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/projects:ProjectController"] = append(beego.GlobalControllerRouter["controllers/projects:ProjectController"],
		beego.ControllerComments{
			Method:           "GetProjectByAlias",
			Router:           `/project/alias/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/projects:ProjectController"] = append(beego.GlobalControllerRouter["controllers/projects:ProjectController"],
		beego.ControllerComments{
			Method:           "GetProjectById",
			Router:           `/project/id/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/projects:ProjectController"] = append(beego.GlobalControllerRouter["controllers/projects:ProjectController"],
		beego.ControllerComments{
			Method:           "DeleteProjectById",
			Router:           `/project/id/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/projects:ProjectController"] = append(beego.GlobalControllerRouter["controllers/projects:ProjectController"],
		beego.ControllerComments{
			Method:           "GetAllProject",
			Router:           `/project/list/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
