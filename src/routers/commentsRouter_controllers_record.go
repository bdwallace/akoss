package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/record:RecordController"] = append(beego.GlobalControllerRouter["controllers/record:RecordController"],
        beego.ControllerComments{
            Method: "RecordCount",
            Router: `/record/count`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/record:RecordController"] = append(beego.GlobalControllerRouter["controllers/record:RecordController"],
        beego.ControllerComments{
            Method: "RecordList",
            Router: `/record/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/record:RecordController"] = append(beego.GlobalControllerRouter["controllers/record:RecordController"],
        beego.ControllerComments{
            Method: "RecordListCount",
            Router: `/record/listcount`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
