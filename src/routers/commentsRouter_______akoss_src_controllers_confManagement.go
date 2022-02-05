package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/confManagement:HostConfNGController"] = append(beego.GlobalControllerRouter["controllers/confManagement:HostConfNGController"],
        beego.ControllerComments{
            Method: "ModifyNginxConf",
            Router: `/confManage/nginx`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/confManagement:HostConfNGController"] = append(beego.GlobalControllerRouter["controllers/confManagement:HostConfNGController"],
        beego.ControllerComments{
            Method: "SearchSericeHost",
            Router: `/confManage/searchServiceHost`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
