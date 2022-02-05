package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/backup:BackupController"] = append(beego.GlobalControllerRouter["controllers/backup:BackupController"],
        beego.ControllerComments{
            Method: "GetService",
            Router: `/backup/service`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/backup:BackupController"] = append(beego.GlobalControllerRouter["controllers/backup:BackupController"],
        beego.ControllerComments{
            Method: "BackupAkgoSql",
            Router: `/backup/sql`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
