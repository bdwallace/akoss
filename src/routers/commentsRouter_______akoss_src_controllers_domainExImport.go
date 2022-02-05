package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"] = append(beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"],
        beego.ControllerComments{
            Method: "DomainBackupAll",
            Router: `/domainBackupAll`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"] = append(beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"],
        beego.ControllerComments{
            Method: "DomainImport",
            Router: `/domainImport`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"] = append(beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"],
        beego.ControllerComments{
            Method: "SearchDomain",
            Router: `/searchAliDomain`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"] = append(beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"],
        beego.ControllerComments{
            Method: "TenCentCloutDomainList",
            Router: `/tenCentDomain`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"] = append(beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"],
        beego.ControllerComments{
            Method: "CreateTenCentCloutDomain",
            Router: `/tenCentDomain`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"] = append(beego.GlobalControllerRouter["controllers/domainExImport:DomainExImport"],
        beego.ControllerComments{
            Method: "TenCentDescribeDomain",
            Router: `/tenCentDomainDescribe`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
