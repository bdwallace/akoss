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

func BlackListAndDenyUserAgentCP(service *models.Service, blackList bool, denyUserAgent bool) (err error) {

	for _, host := range service.Hosts {

		baseCom := components.BaseComponents{}
		baseCom.SetProject(service.Project)
		baseCom.SetService(service)
		baseCom.SetHost(service.Hosts)
		baseDocker := new(components.BaseDocker)
		baseDocker.SetBaseComponents(baseCom)

		allCmds := make([]string, 0)
		//docker cp 命令
		cc := fmt.Sprintf("/usr/bin/env docker -H tcp://%s:%s ", host.UseIp, "2375")

		if blackList {
			blackListDir := fmt.Sprintf("black_list/%s/%s", service.Project.Alias, service.Name)
			err = common.Mkdir(blackListDir)
			if err != nil {
				fmt.Println("error: MkdirH5ClackList common.Mkdir ", err)
				return
			}
			srcFilePath := fmt.Sprintf("%s/deny.conf", blackListDir)

			// echo > deny.conf
			echoCmd := fmt.Sprintf(`echo '%s' > %s`, service.BlackList, srcFilePath)
			s, cmdErr := gopubssh.CommandLocal(echoCmd, components.SSHTIMEOUT)
			if cmdErr != nil {
				if cmdErr.Error() == "cmd time out" {
					s.Result = fmt.Sprintf("%s: %s  cmd time out", echoCmd, host.UseIp)
					beego.Error(s.Result)
					return cmdErr
				}
				return cmdErr
			}

			allCmds = append(allCmds, fmt.Sprintf("%s cp %s %s:/etc/nginx/conf.d/deny.conf", cc, srcFilePath, baseDocker.BaseComponents.Docker.Name))
		}

		if denyUserAgent {
			denyUserAgent := fmt.Sprintf("deny_user_agent/%s/%s", service.Project.Alias, service.Name)
			err = common.Mkdir(denyUserAgent)
			if err != nil {
				fmt.Println("error: MkdirH5ClackList common.Mkdir ", err)
				return
			}
			srcDenyUserAgentFilePath := fmt.Sprintf("%s/user_agent_deny.conf", denyUserAgent)
			// echo > deny.conf
			echoCmd := fmt.Sprintf(`echo '%s' > %s`, service.DenyUserAgent, srcDenyUserAgentFilePath)
			s, cmdErr := gopubssh.CommandLocal(echoCmd, components.SSHTIMEOUT)
			if cmdErr != nil {
				if cmdErr.Error() == "cmd time out" {
					s.Result = fmt.Sprintf("%s: %s  cmd time out", echoCmd, host.UseIp)
					beego.Error(s.Result)
					return cmdErr
				}
				return cmdErr
			}

			allCmds = append(allCmds, fmt.Sprintf("%s cp %s %s:/etc/nginx/user_agent_deny.conf", cc, srcDenyUserAgentFilePath, baseDocker.BaseComponents.Docker.Name))
		}

		allCmds = append(allCmds, fmt.Sprintf("sleep 1s; %s exec -i %s nginx -s reload", cc, baseDocker.BaseComponents.Docker.Name))
		cmd := strings.Join(allCmds, " ; ")
		s, cmdErr := gopubssh.CommandLocal(cmd, components.SSHTIMEOUT)
		if cmdErr != nil {
			if cmdErr.Error() == "cmd time out" {
				s.Result = fmt.Sprintf("%s: %s  cmd time out", cmd, host.UseIp)
				beego.Error(s.Result)
			}
			return cmdErr
		}
	}
	return
}
