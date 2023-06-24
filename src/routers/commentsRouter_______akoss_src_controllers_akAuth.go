package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/akAuth:AkAuthController"] = append(beego.GlobalControllerRouter["controllers/akAuth:AkAuthController"],
        beego.ControllerComments{
            Method: "AkAuth",
            Router: `/akAuth`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
