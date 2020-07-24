package components

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"library/common"
	gopubssh "library/ssh"
	"models"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/gaoyue1989/sshexec"
)

const SSHTIMEOUT = 3600
const SSHWorker = 10

const SSHREMOTETIMEOUT = 600


type BaseComponents struct {
	Project  	*models.Project
	Platform 	*models.Platform
	Service 	*models.Service
	Hosts 		[]*models.Host
	Task    	*models.Task
	User    	*models.User
	Docker  	Docker
	Repo		Repo
}

type Docker struct {
	Name	string
	Port 	string
	Base    string
}

type Repo struct {
	User					string
	Pass					string
	AddressImage			string
	AddressApi 				string
}

func (c *BaseComponents) SetProject(project *models.Project) {
	c.Project = project
}

func (c *BaseComponents) SetService(service *models.Service) {
	c.Service = service
	c.Docker = Docker{}

	if c.Service.DockerName == "" {
		if  c.Project.Alias != ""{
			c.Docker.Name = c.Service.Name + "_" + c.Project.Alias
		} else {
			c.Docker.Name = c.Service.Name + "_" + service.Project.Alias
		}

	} else {
		c.Docker.Name = service.DockerName
	}

	c.Docker.Port = beego.AppConfig.String("dockerport")
	c.Docker.Base = beego.AppConfig.String("dockerbase")

	c.Repo.User = beego.AppConfig.String("repouser")
	c.Repo.Pass = beego.AppConfig.String("repopass")

	c.Repo.AddressImage = beego.AppConfig.String("repoimage")
	c.Repo.AddressApi = beego.AppConfig.String("repoapi")
}

func (c *BaseComponents) SetHost(hosts []*models.Host) {
	c.Hosts = append(c.Hosts, hosts...)
}

func (c *BaseComponents) SetTask(task *models.Task) {
	c.Task = task
}

func (c *BaseComponents) SetPlatform(platform *models.Platform) {
	c.Platform = platform
}


func (c *BaseComponents) SetUser(user *models.User) {
	c.User = user
}


func (c *BaseComponents) GetProject()(project *models.Project) {
	return c.Project
}

func (c *BaseComponents) GetService()(service *models.Service) {
	return c.Service
}

func (c *BaseComponents) GetHost()(hosts []*models.Host) {
	return c.Hosts
}

func (c *BaseComponents) GetTask()(task *models.Task) {
	return c.Task
}

func (c *BaseComponents) GetUser()(user *models.User) {
	return c.User
}



/**
 * 万能的本地命令
 *
 */
/*func (c *BaseComponents) Local(command string, timeout int, sql string, host string) (sshexec.ExecResult, error) {
	if timeout == 0 {
		timeout = SSHTIMEOUT
	}
	s, err := gopubssh.CommandLocal(command, timeout)
	if err != nil && err.Error() == "cmd time out" {
		s.Result = fmt.Sprintf("%s: cmd time out", host)
		beego.Error(s.Result)
	}

	if sql != "" {
		duration := common.GetInt(s.EndTime.Sub(s.StartTime).Seconds())
		//createdAt := time.Now()
		status := 1
		if s.Error != nil {
			status = 0
		}
		c.SaveSql(command, duration, status, s, host, sql)
	}

	return s, err
}
*/

/**
 * 万能的本地命令,非发布命令，在build,download,start,restart,stop时使用
 *
 */
func (c *BaseComponents) LocalOtherNull(command string, action int, re *models.OperationRecord) (err error) {
	//createdAt := time.Now()
	// err = c.SaveSqlToOperationRecord(command, 0, status, action, "", operationRecord.Host, operationRecord.Class)
	// if err != nil{
	// 	return
	// }
	re.Id = 0
	re.Status = 1
	re.Action = int16(action)
	re.Command = command
	re.Duration = 0
	re.User = c.User
	_, err = models.AddOperationRecord(re)
	if err != nil {
		beego.Error(err)
		return err
	}
	return
}


/**
 * 万能的本地命令,非发布命令，在build,download,start,restart,stop时使用
 *
 */
func (c *BaseComponents) LocalOther(command string, timeout int, action int, re *models.OperationRecord) (sshexec.ExecResult, error) {
	if timeout == 0 {
		timeout = SSHTIMEOUT
	}
	s, err := gopubssh.CommandLocal(command, timeout)
	if err != nil && err.Error() == "cmd time out" {
		s.Result = fmt.Sprintf("%s: cmd time out", re.Host.UseIp)
		beego.Error(s.Result)
	}
	duration := common.GetInt(s.EndTime.Sub(s.StartTime).Seconds())
	//createdAt := time.Now()
	status := 1
	if s.Error != nil {
		status = 0
	}
	// err = c.SaveSqlToOperationRecord(command, duration, status, action, s, operationRecord.Host, operationRecord.Class)
	// if err != nil{
	// 	return s, err
	// }
	re.Id = 0
	re.Status = 1
	re.Action = int16(action)
	re.Command = command
	sResult, _ := json.Marshal(s)
	re.Memo = string(sResult)
	re.Duration = duration
	re.Status = int16(status)
	re.User = c.User
	_, err = models.AddOperationRecord(re)
	if err != nil {
		beego.Error(err)
		return s, err
	}

	return s, err
}


/**
 * 万能的本地命令
 *
 */
func (c *BaseComponents) RunLocal(task *models.Task,command string, timeout int, host string, sql string) (s sshexec.ExecResult, resId int64,err error) {
	if timeout == 0 {
		timeout = SSHTIMEOUT
	}
	s, cmdErr := gopubssh.CommandLocal(command, timeout)
	if cmdErr != nil && cmdErr.Error() == "cmd time out" {
		s.Result = fmt.Sprintf("%s: %s  cmd time out",command, host)
		beego.Error(s.Result)
	}


	if sql != "" {
		duration := common.GetInt(s.EndTime.Sub(s.StartTime).Seconds())
		status := 1
		if s.Error != nil {
			status = 0
		}

		resId,err = c.SaveSql(task, command, duration, status, s, host, sql)
		if err != nil{
			return s, resId,err
		}
	}

	if cmdErr != nil{
		return s,resId, cmdErr
	}
	return s, resId, err
}


func (c *BaseComponents) RunDockerPs(command string, timeout int, host string)(s sshexec.ExecResult, err error){

	if timeout == 0 {
		timeout = SSHTIMEOUT
	}
	s, err = gopubssh.CommandLocal(command, timeout)
	if err != nil && err.Error() == "cmd time out" {
		s.Result = fmt.Sprintf("%s: cmd time out", host)
		beego.Error(s.Result)
		return
	}

	return
}




func (c *BaseComponents) SaveSql(task *models.Task, command string, duration int, status int, value interface{}, host string, sql string) (int64,error){
	beego.Info(value)
	if duration < 0 {
		duration = 0
	}

	re := models.Record{}
	re.Status = int16(status)
	//re.Action = c.Task.Action
	re.Command = command
	re.Duration = duration
	sResult, _ := json.Marshal(value)
	re.Memo = string(sResult)
	re.CreatedAt = time.Now()
	re.Host = host
	//re.Count = c.Task.Deploy.Count
	re.Deploy = task.Deploy
	re.User = task.User
	re.Task = task

	if sql == "up" {
		err := models.UpdateRecordByTaskHost(&re)
		if err != nil {
			beego.Error(err)
			return -1, err
		}
	}

	id, err := models.AddRecord(&re)
	if err != nil {
		beego.Error(err)
		return -1, err
	}
	return id,err
}






func (c *BaseComponents) SaveSqlToOperationRecord(command string, duration int, status int, action int, value interface{}, host *models.Host,class string) (err error) {
	beego.Info(value)
	if duration < 0 {
		duration = 0
	}
	re := models.OperationRecord{}
	re.Status = int16(status)
	re.Action = int16(action)
	re.Class = class
	re.Command = command
	re.Duration = duration
	sResult, _ := json.Marshal(value)
	re.Memo = string(sResult)
	re.Host = host
	re.User = c.User
	re.Service = c.Service

	_, err = models.AddOperationRecord(&re)
	if err != nil {
		beego.Error(err)
		return err
	}

	return nil

}


func LineChangeHost(url string, host *models.Host, line string) (err error) {

	hosts,err := models.GetHostRelated(host)
	if err != nil {
		return
	}

	for _, service := range hosts.Services{
		port := service.Port
		hostPort := fmt.Sprintf("%s:%s", host.UseIp, port)
		err = LineChange(url, hostPort, line)
		if err != nil {
			return
		}
	}

	return

}


// ????
func LineChange(url string, hostPort string, line string) (err error) {

	lineData, err := Line(url)
	if err != nil {
		return
	}

	lineList := strings.Split(lineData, ",")
	if line == "on" {
		for common.InListIndex(hostPort, lineList) != -1 {
			i := common.InListIndex(hostPort, lineList)
			lineList = append(lineList[:i], lineList[i+1:]...)
		}
	}

	if line == "off" {
		if common.InListIndex(hostPort, lineList) == -1 {
			lineList = append(lineList, hostPort)
		}
	}

	reqPost := httplib.Post(url)
	reqPost.Param("dataId", "grayReleaseConfig.properties")
	reqPost.Param("group", "DEFAULT_GROUP")
	reqPost.Param("type", "properties")
	content1 := "grayRelease.enable=true\r\n"
	content2 := "grayRelease.grayLogEnabled=true\r\n"
	content3 := "grayRelease.grayClientIpList=\r\n"
	content4 := "grayRelease.grayServerList="
	content := content1 + content2 + content3 + content4 + strings.Trim(strings.Join(lineList, ","), ",")
	reqPost.Param("content", content)
	resultPost, err := reqPost.String()
	if err != nil {
		return
	}
	beego.Info(fmt.Sprintf("%s %s %s", hostPort, line, resultPost))

	return

}


func Line(url string) (lineStr string, err error) {
	req := httplib.Get(url)
	req.Param("dataId", "grayReleaseConfig.properties")
	req.Param("group", "DEFAULT_GROUP")
	result, err := req.String()
	beego.Info(fmt.Sprintf("curl: %s", url))

	resultList := strings.Split(result, "=")
	lineStr = resultList[len(resultList)-1]
	beego.Info(fmt.Sprintf("lineStr: %s", lineStr))

	return
}




/**
* 执行本地宿主机命令,并在record表记录执行的IP
 */
func (c *BaseComponents) RunLocalCommandHostLog(task *models.Task, command string, timeout int, host string, sql string) (sshexec.ExecResult,int64, error) {
	if timeout == 0 {
		timeout = SSHTIMEOUT
	}
	s, err := gopubssh.CommandLocal(command, timeout)
	if s.Result == "" {
		s.Result = "not key in logs!"
		s.Error = fmt.Errorf("not key in logs!")
		err = fmt.Errorf("not key in logs!")
	}

	//获取执行时间
	duration := common.GetInt(s.EndTime.Sub(s.StartTime).Seconds())
	status := 1
	if s.Error != nil {
		status = 0
	}

	resId, err := c.SaveSql(task, command,duration,status,s,host,sql)
	if err != nil{
		return s,resId, err
	}

	return s, resId, err
}




func GetMqttHealth(domain *models.Domain) (health string, err error){

	confs, err := models.GetConfByConfNameAndProjectId("emqtt",domain.Project.Id)
	if err != nil{
		//fmt.Println("error: ",err)
		return health, err
	}

	confValues := make([]*models.ConfValue,0)
	if err := json.Unmarshal([]byte(confs[0].Value),&confValues); err != nil{
		//fmt.Println("error: ",err)
		return health, err
	}

	broker := "tcp://" + domain.Domain + ":" + domain.Port
	mqttUserName := ""
	mqttPwd := ""

	for _, v := range confValues {
		if v.Key == "MQTT_USER" {
			mqttUserName = v.Value
		}

		if v.Key == "MQTT_PWD"{
			mqttPwd = v.Value
		}
	}

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts = opts.SetClientID("mqtt-client")
	opts = opts.SetUsername(mqttUserName)
	opts = opts.SetPassword(mqttPwd)
	opts = opts.SetCleanSession(true)
	opts = opts.SetAutoReconnect(true)

	mqttClient := mqtt.NewClient(opts)
	token := mqttClient.Connect()
	if token.Wait() && token.Error() != nil{
		err = token.Error()
		return "", err
	}

	if resIsConn := mqttClient.IsConnected(); resIsConn{
		health = "200"
	}

	return health, err
}
