package deploycontrollers

import (
	"fmt"
	"github.com/gaoyue1989/sshexec"
	"library/components"
	"models"
	"strings"
	"time"
)


/*
	以平台的形式创建发布单 (发布前端服务 h5 download ...)
	以服务的形式创建发布单 (发布后端服务 java)
 */


/*
	初始化发布对象
	project service host task user docker等
*/
func (c *DeployController)initComponents(task *models.Task) *components.BaseDocker {

	s := components.BaseComponents{}
	s.SetTask(task)
	s.SetProject(task.Project)
/*
	if len(task.Service.Platforms) != 0 && len(task.Service.Platforms) > 1{
		s.SetPlatform(task.Service.Platforms[0])
	}
*/
	s.SetService(task.Service)
	s.SetHost(task.Hosts)
	s.SetUser(task.User)

	docker := new(components.BaseDocker)
	docker.SetBaseComponents(s)

	return docker
}


/*
	发布前校验发布单状态
*/
func (c *DeployController) DeployCheck(deploy *models.Deploy) {

	if deploy == nil || deploy.Id == 0 {
		c.SetJson(1, nil, "Parameter error")
		return
	}
	if deploy.IsRun == 1 {
		c.SetJson(1, nil, "此多项目任务正在上线")
		return
	}

	deploy.IsRun = 1
	//deploy.Count = deploy.Count + 1
}

/*
	创建 前段服务发布命令
 */

/*
func (c *DeployController)createFrontEndDockerCmd(deploy *models.Deploy,s models.Service,i int)*models.Deploy{

	task := new(models.Task)
	task.Service = &s
	task.Project = s.Project
	task.Hosts = s.Hosts
	task.User = deploy.User
	deploy.Tasks = append(deploy.Tasks,task)

	docker := c.initComponents(deploy.Tasks[i])
	fmt.Println("task id: ",deploy.Tasks[i].Id)
	fmt.Printf("docker: %+v \n",docker)
	err := docker.CreateDockerCmd(deploy.Tasks[i],deploy.Count)
	if err != nil{
		c.SetJson(1, nil, "创建 docker cmd 失败")
		return nil
	}

	fmt.Println("===========")
	fmt.Printf("c.cmd: \n%s\n",docker.Cmds)
	fmt.Println("===========")
	t, err := models.GetTaskByDeployIdAndServiceId(int64(deploy.Id),int64(s.Id))
	if err != nil{
		c.SetJson(1, nil, "创建 docker cmd 获取 deploy 相关 tasks 失败")
		return nil
	}

	t.Cmd = docker.Cmds

	id, err := models.UpdateTaskById(t)
	if err != nil{
		c.SetJson(1, nil, "创建 docker cmd 插入 deploy 相关 task cmd 失败")
		return nil
	}
	deploy.Tasks[i].Id = int(id)
	return deploy
}

*/



func (c *DeployController) getCheckCmd(cmd string) (checkCmd string){

	begin := strings.Index(cmd,"check")
	end := strings.Index(cmd,"#")
	checkCmd = cmd[begin+6:end]

	return
}


/*
	解析 每一条发布命令
*/
func (c *DeployController) getEveryoneCmd(cmd string) (retCmds []string){

	begin := strings.Index(cmd,"[")
	end := strings.Index(cmd,"]")
	cmd1 := cmd[begin:end]
	cmds := strings.Split(cmd1,"##")
	cmds2 := cmds[:len(cmds) - 1]

	for _, value := range cmds2{
		v := strings.Trim(value,"[")
		v = strings.Trim(v,"\n")
		v = strings.Trim(v," ")
		retCmds = append(retCmds, v)
	}

	return
}

/*
	解析发布命令
*/
func (c *DeployController) AnalyzeDockerCmd(allCmd string,class string) (checkResCmd string, pullResCmds []string,runResCmds []string,domainResCmds []string){

	pullIndex := strings.Index(allCmd,"pull")
	runIndex := strings.Index(allCmd,"run")
	domainIndex := strings.Index(allCmd,"domain")

	if class == "java"{
		checkIndex := strings.Index(allCmd,"check")
		if checkIndex >= 0 {
			checkAllCmd := allCmd[checkIndex:pullIndex]
			checkResCmd = c.getCheckCmd(checkAllCmd)
		}
	}

	pullAllCmd := allCmd[pullIndex:runIndex]
	pullResCmds = c.getEveryoneCmd(pullAllCmd)


	if class != "java"{
		if domainIndex == -1{
			runAllCmd := allCmd[runIndex:]
			runResCmds = c.getEveryoneCmd(runAllCmd)

			return "",pullResCmds, runResCmds,nil
		}else {
			runAllCmd := allCmd[runIndex:domainIndex]
			runResCmds = c.getEveryoneCmd(runAllCmd)
			domainAllCmd := allCmd[domainIndex:]
			domainResCmds = c.getEveryoneCmd(domainAllCmd)

			return "",pullResCmds, runResCmds,domainResCmds
		}
	}else {
		runAllCmd := allCmd[runIndex:]
		runResCmds = c.getEveryoneCmd(runAllCmd)

		return checkResCmd,pullResCmds, runResCmds,nil
	}
}


type deployCmd struct {
	Host			string
	CheckCmd 		string
	PullCmds 		string
	RunCmds 		string
	DomainCmds		string
}




//普通上线任务
/*
	服务发布
*/
func (c *DeployController)releaseHandling(ch chan int,task *models.Task,deploy *models.Deploy,hosts []*models.Host) {

	checkCmd, pullCmds, runCmds, domainCmds := c.AnalyzeDockerCmd(task.Cmd,deploy.Class)

	var err error

	// docker pull
	for _, cmd := range pullCmds{

		var pullResId int64
		_, isDelpoy := c.CheckIsDeploy(cmd,hosts)
		if !isDelpoy{
			continue
		}

		if pullResId, err = c.RunDockerCmd(task, "pull",cmd); err != nil && pullResId > 0 {
			if err := UpdateRecordActionAndCount(pullResId,120,deploy.Count); err != nil {
				fmt.Println("error:  UpdateRecordCount: ",err)
				return
			}

			c.SetJson(1, nil, "运行 docker pull 失败")
			return
		}
		if pullResId > 0 {
			if err := UpdateRecordActionAndCount(pullResId,30,deploy.Count); err != nil{
				fmt.Println("error:  UpdateRecordCount: ",err)
				return
			}
		}
	}

	var runResId int64
	// docker run
	for i, cmd := range runCmds{

		_, isDelpoy := c.CheckIsDeploy(cmd,hosts)
		if !isDelpoy{
			continue
		}

		if runResId, err = c.RunDockerCmd(task, "run",cmd); err != nil && runResId > 0{
			if err := UpdateRecordActionAndCount(runResId,120,deploy.Count); err != nil{
				fmt.Println("error:  UpdateRecordCount: ",err)
				return
			}

			c.SetJson(1, nil, "运行 docker run 失败")
			return
		}

		if runResId > 0 {
			if err := UpdateRecordActionAndCount(runResId,60,deploy.Count); err != nil{
				fmt.Println("error:  UpdateRecordCount: ",err)
				return
			}
		}


		if i == 0 && deploy.Class == "java" && checkCmd != ""{
			if runResId, err = c.RunDockerCmd(task, "check",checkCmd); err != nil && runResId > 0 {
				if err := UpdateRecordActionAndCount(runResId,120,deploy.Count); err != nil{
					fmt.Println("error:  UpdateRecordCount: ",err)
					return
				}
				c.SetJson(1, err, "主机发布失败: ")
				return
			}
			if runResId > 0 {
				if err := UpdateRecordActionAndCount(runResId,60,deploy.Count); err != nil {
					fmt.Println("error:  UpdateRecordCount: ",err)
					return
				}
			}
		}

		// 更新服务最新tag
		// 因为发布次数会有不确定性，所以更新service的tag和last_tag不能在发布的时候更新
		// 只能在确定task的tag和生成命令时更新service的tag和last_tag
		// if err := c.UpdateServiceTagAndLastTag(task);err != nil{
		// 	c.SetJson(1, nil, "发布 service 更新 tag 失败")
		// 	return
		// }

	}

	// domain
	if deploy.Class != "java" {

		var domainResId int64
		for _, cmd := range domainCmds{
			_,isDelpoy := c.CheckIsDeploy(cmd,hosts)
			if !isDelpoy{
				continue
			}
			if domainResId, err = c.RunDockerCmd(task,"domain",cmd); err != nil && domainResId > 0{
				if err := UpdateRecordActionAndCount(domainResId,120,deploy.Count); err != nil{
					fmt.Println("error:  UpdateRecordCount: ",err)
					return
				}
				c.SetJson(1, nil, "运行 docker domain 失败")
				return
			}
		}
		if domainResId > 0{
			if err := UpdateRecordActionAndCount(domainResId,100,deploy.Count); err != nil{
				fmt.Println("error:  UpdateRecordCount: ",err)
				return
			}
		}
		if domainResId == 0{
			if err := UpdateRecordActionAndCount(runResId,100,deploy.Count); err != nil{
				fmt.Println("error:  UpdateRecordCount: ",err)
				return
			}
		}
	}else {
		if runResId > 0{
			if err := UpdateRecordActionAndCount(runResId,100,deploy.Count); err != nil{
				fmt.Println("error:  UpdateRecordCount: ",err)
				return
			}
		}
	}

	ch <- 1
}




func (c *DeployController)UpdateServiceTagAndLastTag(task *models.Task)(err error){

	// 更新服务最新tag
	task.Service.LastTag = task.Service.Tag
	task.Service.Tag = task.Tag
	//time.Now().Format("2006-01-02 15:04:05")
	task.Service.LastTagUpdateAt = time.Now()
	if _, err = models.UpdateServiceById(task.Service);err != nil{
		return err
	}
	return
}



func (c *DeployController)CheckIsDeploy(cmd string,hosts []*models.Host)(useIp string,isDeploy bool){

	//start := strings.Index(cmd,":")
	//tempCmd:= cmd[start+1:]
	//end := strings.Index(tempCmd,":")
	isDeploy = false
	hostIp := ""
	//start := strings.Index(cmd,"\n")
	end := strings.Index(cmd,":")



	if end == -1 {
		return "",isDeploy
	}else {
		hostIp = cmd[:end]
	}

	if hostIp == ""{
		return "",isDeploy
	}

	for _, h := range hosts {
		if h.UseIp == hostIp{
			isDeploy = true
			return hostIp,isDeploy
		}
	}
	return hostIp,isDeploy
}


/*
	执行 docker 命令
*/
func (c *DeployController)RunDockerCmd(task * models.Task ,cmdType string, dockerCmd string)(resId int64,err error){

	hostIndex := strings.Index(dockerCmd,":")
	host := dockerCmd[:hostIndex]
	cmd := dockerCmd[hostIndex + 1:]
	//fmt.Println("============")
	//fmt.Println("run cmd:  ",cmd)
	//fmt.Println("============")
	docker := new(components.BaseDocker)
	docker.BaseComponents.User = c.TaskUser
	docker.BaseComponents.Task = c.Task


	var sshExec sshexec.ExecResult

	// check
	if cmdType == "check" {
		sshExec, resId, err = docker.BaseComponents.RunLocalCommandHostLog(task, cmd, 120, host,"add")
		if err != nil{
			fmt.Println("error docker check ",err)
			fmt.Println("sshEcec result : ",sshExec.Result)
			return resId, err
		}
	}


	//docker pull
	if cmdType == "pull"{
		sshExec, resId, err = docker.BaseComponents.RunLocal(task, cmd, 0, host,"add")
		if err != nil{
			fmt.Println("error docker pull ",err)
			fmt.Println("sshEcec result : ",sshExec.Result)
			return resId, err
		}
	}

	// docker run
	if cmdType == "run"	{
		sshExec, resId, err = docker.BaseComponents.RunLocal(task, cmd, 60, host,"add")
		if err != nil{
			fmt.Println("error docker run ",err)
			fmt.Println("sshEcec result : ",sshExec.Result)
			return resId, err
		}

	}

	// domain docker run
	if cmdType == "domain"	{
		sshExec, resId, err = docker.BaseComponents.RunLocal(task, cmd, 60, host,"add")
		if err != nil{
			fmt.Println("error domain docker cmd ",err)
			fmt.Println("sshEcec result : ",sshExec.Result)
			return resId, err
		}
	}

	return resId, err
}


func UpdateRecordAction(resId int64, action int16)error{
	recordId := int(resId)

	re, err := models.GetRecordById(recordId)
	if err != nil{
		return err
	}
	re.Action = action

	if err := models.UpdateRecordById(re);err != nil{
		return err
	}
	return nil
}


func UpdateRecordCount(resId int64, count int)error{
	recordId := int(resId)
	reCount := int(count)
	re, err := models.GetRecordById(recordId)
	if err != nil{
		return err
	}
	re.Count = reCount

	if err := models.UpdateRecordById(re);err != nil{
		return err
	}
	return nil
}


func UpdateRecordActionAndCount(resId int64, action int16, count int) (err error){
	if err = UpdateRecordAction(resId,action); err != nil{
		err = fmt.Errorf("error: UpdateRecordAction %d  %s\n",action,err)
		return err
	}

	if err = UpdateRecordCount(resId,count);err != nil{
		err = fmt.Errorf("error: UpdateRecordCount %d  %s\n",count,err)
		return err
	}
	return
}
