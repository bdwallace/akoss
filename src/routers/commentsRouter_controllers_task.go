package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers/task:TaskController"] = append(beego.GlobalControllerRouter["controllers/task:TaskController"],
		beego.ControllerComments{
			Method:           "AddTask",
			Router:           `/task/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/task:TaskController"] = append(beego.GlobalControllerRouter["controllers/task:TaskController"],
		beego.ControllerComments{
			Method:           "GetTaskById",
			Router:           `/task/id/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/task:TaskController"] = append(beego.GlobalControllerRouter["controllers/task:TaskController"],
		beego.ControllerComments{
			Method:           "DeleteTask",
			Router:           `/task/id/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("id", param.IsRequired),
			),
			Filters: nil,
			Params:  nil})

	beego.GlobalControllerRouter["controllers/task:TaskController"] = append(beego.GlobalControllerRouter["controllers/task:TaskController"],
		beego.ControllerComments{
			Method:           "GetAllTask",
			Router:           `/task/list/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/task:TaskController"] = append(beego.GlobalControllerRouter["controllers/task:TaskController"],
		beego.ControllerComments{
			Method:           "GetTaskByServiceId",
			Router:           `/task/serviceId/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
