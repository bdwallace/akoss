package monitorcontrollers

import (
	"controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gaoyue1989/sshexec"
	gopubssh "library/ssh"
	"models"
	"strings"
	"sync"
)

type MonitorController struct {
	controllers.BaseController
}



// @Title monitor
// @Description 检查 monitor  nod_exporter 容器，若没有则启动该容器。
// @Param   projectId     	query     int 		true       "project id"
// @Failure 1 获取 AwsStatus by instanceId and region 失败
// @Failure 2 User not found
// @router /monitor/ [get]
func (c *MonitorController)Monitor(){

	projects, err := models.GetAllProject()
	if err != nil{
		c.SetJson(1, nil, "error: get all project id")
		return
	}

	for _, project := range projects {

		hosts, err := models.GetAllHost(project.Id)
		if err != nil{
			c.SetJson(1, nil, "error: get all host by project id")
			return
		}

		waitgroup := new(sync.WaitGroup)
		for _, host := range hosts{
			waitgroup.Add(2)

			go DeployCadvisor(host.UseIp, waitgroup)
			go DeployNodeExporter(host.UseIp, waitgroup)
		}

		waitgroup.Wait()
		c.SetJson(0, nil , "")
		return
	}

}

func monitorContainerUpdate(host, dockerName string) {

	dockerPort := beego.AppConfig.String("dockerport")
	repoUser := beego.AppConfig.String("repouser")
	repoPasswd := beego.AppConfig.String("repopass")
	repoImageAddress :=  beego.AppConfig.String("repoimage")

	deployCmd := fmt.Sprintf("docker -H tcp://%s:%s login -u %s -p %s %s ; docker -H tcp://%s:%s container update --name %s --restart=always -d ", host, dockerPort, repoUser, repoPasswd, repoImageAddress, host, dockerPort,dockerName)

	if err := runDeployMonitor(host, "docker container node_exporter", deployCmd); err != nil{
		return
	}
}


// 检查容器是否已启动
func monitorCheckRes(host string, dockerName string) (sshexec.ExecResult, error) {
	checkResultCmd := fmt.Sprintf( "docker -H tcp://%s ps | grep %s",host,dockerName)

	s, cmdErr := gopubssh.CommandLocal(checkResultCmd, 60)
	if cmdErr != nil && cmdErr.Error() == "cmd time out" {
		s.Result = fmt.Sprintf("%s :  ps | grep cadvisor cmd time out: %s ",host, cmdErr)
		fmt.Println(s.Result)
		return s, cmdErr
	}

	return s, nil
}

// docker run 部署容器
func runDeployMonitor(host, dockerName, deployCmd string) error{

	s, cmdErr := gopubssh.CommandLocal(deployCmd,60)
	if cmdErr != nil && cmdErr.Error() == "cmd time out" {
		s.Result = fmt.Sprintf("%s : deploy %s cmd time out: %s ",host, dockerName, cmdErr)
		fmt.Println(s.Result)
		return cmdErr
	}

	return nil
}

// 部署  cadvisor
func DeployCadvisor(host string, waitgroup * sync.WaitGroup) {

	defer waitgroup.Done()

	monitorContainerUpdate(host,"cadvisor")

	// 检查是否启动 cadvisor 容器
	s, err := monitorCheckRes(host,"cadvisor")
	if err != nil{
		fmt.Printf("error: %s cadvisor monitor check res %s\n",host,err.Error())
	}


	if strings.Index(s.Result, "cadvisor") > 0 {
		fmt.Printf("%s:: exec check cadvisor is exist\n",host)
		return
	} else {

		// 运行启动 cadvisor 命令
		//fmt.Println("exec cadvisor is not exist")
		fmt.Printf("%s:: exec deploy cadvisor ... \n",host)

		dockerPort := beego.AppConfig.String("dockerport")

		dockerRun := fmt.Sprintf( "docker -H tcp://%s:%s run -itd",host, dockerPort)
		deployCmd := dockerRun +
			" --name cadvisor" +
			" --restart=always" +
			" --log-driver json-file" +
			" --log-opt max-size=2g" +
			" -p 7878:8080" +
			" -v /dev/disk/:/dev/disk:ro" +
			" -v /var/lib/docker/:/var/lib/docker:ro" +
			" -v /sys:/sys:ro" +
			" -v /var/run:/var/run:rw" +
			" -v /:/rootfs:ro" +
			" google/cadvisor:v0.33.0"

		if err := runDeployMonitor(host, "cadvisor", deployCmd); err != nil {
			return
		}

		s2, err := monitorCheckRes(host,"cadvisor")
		if err != nil{
			fmt.Printf("error: %s cadvisor monitor check again  res %s\n",host,err.Error())
		}

		// 再次检查是否启动 cadvisor 容器
		if strings.Index(s2.Result, "cadvisor") > 0 {
			//fmt.Printf("%s:: exec check cadvisor is exist\n",host)
			return
		} else {
			//fmt.Println("exec cadvisor is not exist")
			fmt.Printf("%s:: exec deploy cadvisor failed!!!\n",host)
			return
		}
	}
}


// 部署 node exporter
func DeployNodeExporter(host string, waitgroup * sync.WaitGroup) {

	defer waitgroup.Done()

	monitorContainerUpdate(host,"node_exporter")

	s, err := monitorCheckRes(host, "node_exporter")
	if err != nil{
		fmt.Printf("error: %s node_exporter monitor check res %s\n",host,err)
	}

	if strings.Index(s.Result, "node_exporter") > 0 {
		fmt.Printf("%s:: exec check node_exporter is exist\n",host)
		return
	} else {

		// 运行启动 cadvisor 命令
		//fmt.Println("exec node_exporter is not exist")
		fmt.Printf("%s:: exec deploy node_exporter ... \n",host)
		//docker -H tcp://${remote} login -u docker -p r0JT67Ud6Ybt @REG_ADDRESS@ ; docker -H tcp://${remote} run --name node-exporter -d \
		//dockerRun := fmt.Sprintf( "docker -H tcp://%s run -itd",host)
		dockerPort := beego.AppConfig.String("dockerport")
		repoUser := beego.AppConfig.String("repouser")
		repoPasswd := beego.AppConfig.String("repopass")
		repoImageAddress :=  beego.AppConfig.String("repoimage")

		// user   pass  addressimage

		dockerRun := fmt.Sprintf("docker -H tcp://%s:%s login -u %s -p %s %s ; docker -H tcp://%s:%s run --name node-exporter --restart=always -d ", host, dockerPort, repoUser, repoPasswd, repoImageAddress, host, dockerPort)
		deployCmd := dockerRun +
			" -v \"/proc:/host/proc\"" +
			" -v \"/sys:/host/sys\"" +
			" -v \"/:/rootfs\"" +
			" -v /var/run/dbus/system_bus_socket:/var/run/dbus/system_bus_socket" +
			" --net=host" +
			" --pid=\"host\"" +
			" harbor.one2.cc/devops/node-exporter:v2.0" +
			" --path.procfs /host/proc" +
			" --path.sysfs /host/sys" +
			" --collector.systemd" +
			" --collector.tcpstat" +
			" --collector.filesystem.ignored-mount-points" +
			" \"^/(sys|proc|dev|host|etc)($|/)\""

		if err := runDeployMonitor(host, "node_exporter", deployCmd); err != nil{
			return
		}
		s2, err := monitorCheckRes(host, "node_exporter")
		if err != nil{
			fmt.Printf("error: %s:: node_exporter monitor check again res %s\n",host,err)
		}

		// 再次检查是否启动 cadvisor 容器
		if strings.Index(s2.Result, "node_exporter") > 0 {
			//fmt.Printf("%s:: exec check node_exporter is exist\n",host)
			return
		} else {
			fmt.Printf("%s:: exec deploy node_exporter failed!!!\n",host)
			return
		}
	}
}

