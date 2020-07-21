package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["controllers/host:HostCheckController"] = append(beego.GlobalControllerRouter["controllers/host:HostCheckController"],
        beego.ControllerComments{
            Method: "AwsCheckStop",
            Router: `/awsHost/awsCheckStop/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostCheckController"] = append(beego.GlobalControllerRouter["controllers/host:HostCheckController"],
        beego.ControllerComments{
            Method: "AwsRsyncUpdate",
            Router: `/awsHost/awsRsyncUpdate/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostCheckController"] = append(beego.GlobalControllerRouter["controllers/host:HostCheckController"],
        beego.ControllerComments{
            Method: "AwsSetStopTime",
            Router: `/awsHost/awsSetStopTime/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostCheckController"] = append(beego.GlobalControllerRouter["controllers/host:HostCheckController"],
        beego.ControllerComments{
            Method: "AwsStart",
            Router: `/awsHost/awsStart/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostCheckController"] = append(beego.GlobalControllerRouter["controllers/host:HostCheckController"],
        beego.ControllerComments{
            Method: "AwsStatus",
            Router: `/awsHost/awsStatus/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostCheckController"] = append(beego.GlobalControllerRouter["controllers/host:HostCheckController"],
        beego.ControllerComments{
            Method: "AwsStop",
            Router: `/awsHost/awsStop/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostCheckController"] = append(beego.GlobalControllerRouter["controllers/host:HostCheckController"],
        beego.ControllerComments{
            Method: "AwsStopOnlineHost",
            Router: `/awsHost/awsStopOnlineHost/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostCheckController"] = append(beego.GlobalControllerRouter["controllers/host:HostCheckController"],
        beego.ControllerComments{
            Method: "AwsRsync",
            Router: `/awsHost/awsrsync/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostController"] = append(beego.GlobalControllerRouter["controllers/host:HostController"],
        beego.ControllerComments{
            Method: "AddHost",
            Router: `/host`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostController"] = append(beego.GlobalControllerRouter["controllers/host:HostController"],
        beego.ControllerComments{
            Method: "CopyHostById",
            Router: `/host/copy`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostController"] = append(beego.GlobalControllerRouter["controllers/host:HostController"],
        beego.ControllerComments{
            Method: "GetHostDockerPs",
            Router: `/host/dockerPs/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostController"] = append(beego.GlobalControllerRouter["controllers/host:HostController"],
        beego.ControllerComments{
            Method: "GetHostById",
            Router: `/host/id/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostController"] = append(beego.GlobalControllerRouter["controllers/host:HostController"],
        beego.ControllerComments{
            Method: "DeleteHostById",
            Router: `/host/id/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostController"] = append(beego.GlobalControllerRouter["controllers/host:HostController"],
        beego.ControllerComments{
            Method: "GetAllHost",
            Router: `/host/list/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["controllers/host:HostController"] = append(beego.GlobalControllerRouter["controllers/host:HostController"],
        beego.ControllerComments{
            Method: "GetHostByProjectsId",
            Router: `/host/projectId/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
