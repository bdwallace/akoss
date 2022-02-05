package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/cloud:CloudController"] = append(beego.GlobalControllerRouter["controllers/cloud:CloudController"],
        beego.ControllerComments{
            Method: "AddCloud",
            Router: `/cloud/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/cloud:CloudController"] = append(beego.GlobalControllerRouter["controllers/cloud:CloudController"],
        beego.ControllerComments{
            Method: "CopyCloud",
            Router: `/cloud/copy/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/cloud:CloudController"] = append(beego.GlobalControllerRouter["controllers/cloud:CloudController"],
        beego.ControllerComments{
            Method: "GetCloudById",
            Router: `/cloud/id/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/cloud:CloudController"] = append(beego.GlobalControllerRouter["controllers/cloud:CloudController"],
        beego.ControllerComments{
            Method: "DeleteCloud",
            Router: `/cloud/id/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/cloud:CloudController"] = append(beego.GlobalControllerRouter["controllers/cloud:CloudController"],
        beego.ControllerComments{
            Method: "GetAllCloud",
            Router: `/cloud/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/cloud:CloudController"] = append(beego.GlobalControllerRouter["controllers/cloud:CloudController"],
        beego.ControllerComments{
            Method: "SyncCloud",
            Router: `/cloud/sync/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
