package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers/operationRecord:OperationRecordController"] = append(beego.GlobalControllerRouter["controllers/operationRecord:OperationRecordController"],
		beego.ControllerComments{
			Method:           "RecordCount",
			Router:           `/recordoperation/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["controllers/operationRecord:OperationRecordController"] = append(beego.GlobalControllerRouter["controllers/operationRecord:OperationRecordController"],
		beego.ControllerComments{
			Method:           "RecordListCount",
			Router:           `/recordoperation/listcount`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
