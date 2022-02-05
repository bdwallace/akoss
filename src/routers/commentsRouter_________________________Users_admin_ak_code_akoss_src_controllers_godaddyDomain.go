package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/godaddyDomain:GodaddyControllers"] = append(beego.GlobalControllerRouter["controllers/godaddyDomain:GodaddyControllers"],
        beego.ControllerComments{
            Method: "GetAllDomain",
            Router: `/godaddy`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/godaddyDomain:GodaddyControllers"] = append(beego.GlobalControllerRouter["controllers/godaddyDomain:GodaddyControllers"],
        beego.ControllerComments{
            Method: "ModifyGodaddyDomainDNS",
            Router: `/godaddy`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
