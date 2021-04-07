package blacklist

import (
	"fmt"
	"github.com/astaxie/beego"
	"library/common"
	"library/components"
	gopubssh "library/ssh"
	"models"
	"strings"
)

func H5BlackListCP(service *models.Service)(err error) {

	for _, host := range service.Hosts{

		blackListDir := fmt.Sprintf("black_list/%s/%s",service.Project.Alias, service.Name)

		err = common.Mkdir(blackListDir)
		if err != nil {
			fmt.Println("error: MkdirH5ClackList common.Mkdir ",err)
			return
		}

		srcFilePath := fmt.Sprintf("%s/deny.conf", blackListDir)

		// echo > deny.conf
		echoCmd :=fmt.Sprintf(`echo '%s' > %s`, service.BlackList,srcFilePath)
		s, cmdErr := gopubssh.CommandLocal(echoCmd, components.SSHTIMEOUT)
		if cmdErr != nil {
			if cmdErr.Error() == "cmd time out" {
				s.Result = fmt.Sprintf("%s: %s  cmd time out",echoCmd, host.UseIp)
				beego.Error(s.Result)
				return cmdErr
			}
			return cmdErr
		}


		baseCom := components.BaseComponents{}
		baseCom.SetProject(service.Project)
		baseCom.SetService(service)
		baseCom.SetHost(service.Hosts)
		baseDocker := new(components.BaseDocker)
		baseDocker.SetBaseComponents(baseCom)


		//docker cp 命令
		blackListCmds := make([]string,0)
		blackListCmds = append(blackListCmds, "date +%Y%m%d-%H%M%S-%s")
		cc := fmt.Sprintf("/usr/bin/env docker -H tcp://%s:%s ", host.UseIp,"2375")
		blackListCmds = append(blackListCmds, fmt.Sprintf("%s cp %s %s:/etc/nginx/conf.d/deny.conf", cc, srcFilePath, baseDocker.BaseComponents.Docker.Name))

		blackListCmds = append(blackListCmds, fmt.Sprintf("sleep 1s; %s restart %s ", cc, baseDocker.BaseComponents.Docker.Name))
		cmd := strings.Join(blackListCmds, " ; ")

		s, cmdErr = gopubssh.CommandLocal(cmd, components.SSHTIMEOUT)
		if cmdErr != nil{
			if cmdErr.Error() == "cmd time out" {
				s.Result = fmt.Sprintf("%s: %s  cmd time out",cmd, host.UseIp)
				beego.Error(s.Result)
			}
			return cmdErr
		}
	}
	return
}
