package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/deleteDomainReco:DeleteDomainRecoController"] = append(beego.GlobalControllerRouter["controllers/deleteDomainReco:DeleteDomainRecoController"],
        beego.ControllerComments{
            Method: "GetDomainReco",
            Router: `/domainReco`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deleteDomainReco:DeleteDomainRecoController"] = append(beego.GlobalControllerRouter["controllers/deleteDomainReco:DeleteDomainRecoController"],
        beego.ControllerComments{
            Method: "GetDomainRecoByProjectForClass",
            Router: `/domainReco/class`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
