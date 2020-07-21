package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "AddPlatform",
            Router: `/platform`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "GetPlatformById",
            Router: `/platform/id/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "DeletePlatformId",
            Router: `/platform/id/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "GetAllPlatform",
            Router: `/platform/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "GetPlatformByProjectForName",
            Router: `/platform/name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "GetPlatformByProjectsId",
            Router: `/platform/projectId/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "DeletePlatformCheck",
            Router: `/platform/relatedCheck/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "DeletePlatformAndServiceRelatedByPlatformId",
            Router: `/platformAndService/id/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "GetPlatformAndServiceRelatedByPlatformId",
            Router: `/platformAndService/platformId/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/platform:PlatformController"] = append(beego.GlobalControllerRouter["controllers/platform:PlatformController"],
        beego.ControllerComments{
            Method: "AddPlatformAndServiceRelated",
            Router: `/platformAndService/related/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
