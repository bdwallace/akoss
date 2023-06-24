package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/link:LinkController"] = append(beego.GlobalControllerRouter["controllers/link:LinkController"],
        beego.ControllerComments{
            Method: "AddLink",
            Router: `/link/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/link:LinkController"] = append(beego.GlobalControllerRouter["controllers/link:LinkController"],
        beego.ControllerComments{
            Method: "CopyLink",
            Router: `/link/copy/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/link:LinkController"] = append(beego.GlobalControllerRouter["controllers/link:LinkController"],
        beego.ControllerComments{
            Method: "GetLinkById",
            Router: `/link/id/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/link:LinkController"] = append(beego.GlobalControllerRouter["controllers/link:LinkController"],
        beego.ControllerComments{
            Method: "DeleteLink",
            Router: `/link/id/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/link:LinkController"] = append(beego.GlobalControllerRouter["controllers/link:LinkController"],
        beego.ControllerComments{
            Method: "GetAllLink",
            Router: `/link/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
