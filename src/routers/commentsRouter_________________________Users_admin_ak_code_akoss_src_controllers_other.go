package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/other:OtherController"] = append(beego.GlobalControllerRouter["controllers/other:OtherController"],
        beego.ControllerComments{
            Method: "AppConfig",
            Router: `/other/appconfig`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
