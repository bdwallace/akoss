package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "AddDeploy",
            Router: `/deploy`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "AddDeployFromBuild",
            Router: `/deploy/build`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "DeleteDeploy",
            Router: `/deploy/id/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("id", param.IsRequired),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "GetDeployById",
            Router: `/deploy/id/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "GetAllDeploy",
            Router: `/deploy/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "GetDeployByProjectId",
            Router: `/deploy/projectId/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "GetDeployByProjectIdUserId",
            Router: `/deploy/projectId/userId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "DeployService",
            Router: `/deploy/service`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "GetDeployByServiceId",
            Router: `/deploy/serviceId/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/deploy:DeployController"] = append(beego.GlobalControllerRouter["controllers/deploy:DeployController"],
        beego.ControllerComments{
            Method: "AddDeployTagCmd",
            Router: `/deploy/tagcmd`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
