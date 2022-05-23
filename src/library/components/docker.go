package components

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library/common"
	"models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type BaseConf struct {
	DockerUser string
	DockerPass string
	DockerPort string
	DockerRepo string
	RepoTag    string
	DockerTag  string
	ImagePath  string
	DockerName string
}

type CMD struct {
	BaseParam    *BaseConf
	PublicParam  []string
	DefaultParam string
	Port         []string
}

type BaseDocker struct {
	BaseComponents BaseComponents
	*CMD
	serviceAllCmd string
	Cmds          string
	DockerCmd     CmdObj
}

type CmdObj struct {
	PullCmd  string `json:"pull"`
	RunCmd   string `json:"run"`
	CheckCmd string `json:"check"`
}

func (c *BaseDocker) SetBaseComponents(b BaseComponents) {
	c.BaseComponents = b
}

///////////
/*
docker run
*/

/*
初始化 基础参数
app.conf  docker 等配置
*/
func (c *BaseDocker) InitBaseParam() error {

	if c.BaseComponents.Service.DockerName == "" {
		c.BaseParam.DockerName = c.BaseComponents.Service.Name + "_" + c.BaseComponents.Project.Alias
	} else {
		c.BaseParam.DockerName = c.BaseComponents.Service.DockerName
	}

	return nil
}

/*
初始化 app.conf 默认参数
*/
func (c *BaseDocker) InitDefaultParam() error {

	c.DefaultParam = beego.AppConfig.String("dockerRunBase")

	if c.DefaultParam == "" {

		return fmt.Errorf("error:  serviceId: %d InitDefaultParam nil value field", c.BaseComponents.Service.Id)
	}
	return nil
}

/*
初始化 端口格式
*/
func (c *BaseDocker) InitPort() error {

	service, err := models.GetServiceById(c.BaseComponents.Service.Id)
	if err != nil {
		return err
	}

	slicePort := strings.Split(service.Port, ",")
	c.Port = append(c.Port, slicePort...)

	return nil
}

func (c *BaseDocker) GetPortForDockerServiceHealth() (port string, have443 int) {
	if c.BaseComponents.Service.Port == "" {
		return "", -1
	}

	ports := strings.Split(c.BaseComponents.Service.Port, ",")
	have443 = strings.Index(c.BaseComponents.Service.Port, "443")

	if have443 > -1 {
		for _, p := range ports {
			if strings.Index(p, "443") > -1 {
				port = strings.Split(p, ":")[0]
				break
			}
		}
	} else {
		index := strings.Index(ports[0], ":")
		port = ports[0][:index]
	}

	return

}

func (c *BaseDocker) DockerServiceHealth(host *models.Host, port string, have443 int) (health string, url string) {

	urlUse := ""
	//  前端服务
	if c.BaseComponents.Service.Health == "" {
		//  前端 https
		if have443 > 0 {
			url = fmt.Sprintf("https://%s:%s/", host.PublicIp, port)
			urlUse = fmt.Sprintf("https://%s:%s/", host.UseIp, port)
		} else {
			//  前端 http
			url = fmt.Sprintf("http://%s:%s/", host.PublicIp, port)
			urlUse = fmt.Sprintf("http://%s:%s/", host.UseIp, port)
		}
	} else {
		//  后端服务
		url = fmt.Sprintf("http://%s:%s%s", host.PublicIp, port, c.BaseComponents.Service.Health)
		urlUse = fmt.Sprintf("http://%s:%s%s", host.UseIp, port, c.BaseComponents.Service.Health)
	}

	response, err := common.HttpGet(urlUse, nil)
	if err != nil {
		health = err.Error()
		return
	}

	if strings.Index(url, "/actuator/health") != -1 {
		// 后端服务
		body, _ := ioutil.ReadAll(response.Body)
		if index := strings.Index(string(body), "UP"); index > 0 {
			health = "200"
		} else {
			health = "-1"
		}

	} else {
		// 前端服务
		health = common.IntToStr(response.StatusCode)
	}
	return
}

/*
执行 docker ps 查看容器和上下线状态
*/
func (c *BaseDocker) DockerPs(lineData string) (res []map[string]string, err error) {

	var lineList []string
	if c.BaseComponents.Service.UseNacos == "" && c.BaseComponents.Project.Nacos1 != ""{
		c.BaseComponents.Service.UseNacos = c.BaseComponents.Project.Nacos1
		if err := models.UpdateServiceAndRelated(c.BaseComponents.Service); err != nil {
			return nil, err
		}
	}

	if lineData == "" && c.BaseComponents.Service.UseNacos != ""{
		url := "http://" + c.BaseComponents.Service.UseNacos + "/nacos/v1/cs/configs"
		lineData, err = LineGet(url)
		if err != nil {
	  		return
	  	}
	}

	if lineData != "----" {
		lineList = strings.Split(lineData, ",")
	}


	for _, host := range c.BaseComponents.Hosts {

		PsStatus := "------"
		PsCreatedAt := "------"
		PsImage := "------"
		Line := "------"

		if c.BaseComponents.Service.Class == "java" {
			hostPort := host.PrivateIp + ":" + strings.Split(strings.Split(c.BaseComponents.Service.Port, ",")[0], ":")[0]
			if common.InList(hostPort, lineList) {
				Line = "off"
			} else {
				Line = "on"
			}
		}
		cmdTls := ""
		dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
		if dockerTlsPath != ""{
			cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
		}


		cmd := fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s ps -a -f name='^%s$' --format {{.Status}}#{{.CreatedAt}}#{{.Image}}",cmdTls, host.UseIp, c.BaseComponents.Service.DockerPort, c.BaseComponents.Docker.Name)
		//beego.Info(cmd)

		s, err := c.BaseComponents.RunDockerPs(cmd, 3, host.UseIp)
		if err != nil {
			stErr := fmt.Sprintf("%s", err)
			if stErr == "cmd time out" {
				PsStatus = "超时"
			}
		} else {
			if s.Result != "" {
				arr := strings.Split(s.Result, "#")
				PsStatus = arr[0]
				PsCreatedAt = arr[1][:19]
				if len(strings.Split(arr[2], ":")) == 2 {
					PsImage = strings.Split(arr[2], ":")[1]
				} else if len(strings.Split(arr[2], ":")) == 3 {
					PsImage = strings.Split(arr[2], ":")[2]
				} else {
					PsImage = arr[2]
				}
			}
		}

		p, have443 := c.GetPortForDockerServiceHealth()
		health, url := c.DockerServiceHealth(host, p, have443)

		res = append(res, map[string]string{"project_name:":c.BaseComponents.Project.Name, "host_id": common.IntToStr(host.Id), "ip": host.PrivateIp, "ip_pub": host.PublicIp, "ip_show": host.UseIp, "ps_status": PsStatus, "ps_created_at": PsCreatedAt, "ps_image": PsImage, "line": Line, "health": health, "url": url, "service_id": strconv.Itoa(c.BaseComponents.Service.Id), "service_name": c.BaseComponents.Service.Name})

	}
	return res, err
}




func (c *BaseDocker) createDockerPull(dockerTag string, host *models.Host) (cmdPull string, err error) {

	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}

	cc := fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s ",cmdTls, host.UseIp, c.BaseComponents.Service.DockerPort)
	cmds := []string{}
	cmds = append(cmds, "date +%Y%m%d-%H%M%S-%s")
	cmds = append(cmds, fmt.Sprintf("%s login -u %s -p %s %s", cc, c.BaseComponents.Repo.User, c.BaseComponents.Repo.Pass, c.BaseComponents.Repo.AddressImage))
	cmds = append(cmds, fmt.Sprintf("%s pull %s ", cc, dockerTag))
	cmd := strings.Join(cmds, " && ")

	cmdPull = fmt.Sprintf("%s##\n", cmd)
	return
}

/*
执行 docker run 启动镜像
记录日志 等操作
*/
func (c *BaseDocker) createDockerRun(dockerTag string, host *models.Host, ports string, public string, platformParam string) (cmdRun string, err error) {

	//private, err := c.AnalyzePrivateParam(host.Id)
	private, err := c.AnalyzePrivateParam(host.Id)
	if err != nil {
		return "", err
	}
	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}

	var cmds []string
	cc := fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s ",cmdTls, host.UseIp, c.BaseComponents.Service.DockerPort)
	cmds = append(cmds, "date +%Y%m%d-%H%M%S-%s")

	cmds = append(cmds, fmt.Sprintf("%s stop %s ", cc, c.BaseComponents.Docker.Name))
	cmds = append(cmds, fmt.Sprintf("%s rm -v %s ", cc, c.BaseComponents.Docker.Name))

	// if private == ""{
	// 	cmds = append(cmds, fmt.Sprintf("%s run -itd --name %s %s %s %s %s %s", cc, c.BaseComponents.Docker.Name, ports, c.BaseComponents.Docker.Base, public, platformParam, dockerTag))
	// }else {
	// 判断如果是java项目则默认使用--net-host模式,并默认添加eureka注册信息
	if c.BaseComponents.Service.Class == "java" {
		private = fmt.Sprintf("%s -e EUREKA_INSTANCE_IP-ADDRESS=%s %s", c.BaseComponents.Service.DockerNetwork, host.PrivateIp, private)
	}
	cmds = append(cmds, fmt.Sprintf("%s run -itd --name %s %s %s %s %s %s %s", cc, c.BaseComponents.Docker.Name, ports, c.BaseComponents.Docker.Base, private, public, platformParam, dockerTag))
	// }

	cmd := strings.Join(cmds, "; ")

	cmdRun = fmt.Sprintf("%s##\n", cmd)
	return
}

func (c *BaseDocker) AnalyzePortsAndFindHttps(ports string) (usePort int) {

	index := strings.Index(ports, "443")
	if index == -1 {
		usePort = 80
	} else {
		usePort = 443
	}
	return
}

func (c *BaseDocker) GetDomainByPlatformAndClass(platformId int, class string) (resDomains []*models.Domain, err error) {

	platform, err := models.GetPlatformById(platformId)
	if err != nil {
		err = fmt.Errorf("error: GetPlatformById %s\n", err)
		return nil, err
	}

	resPlatform, err := models.GetPlatformAndDomainRelated(platform)
	if err != nil {
		err = fmt.Errorf("error: GetPlatformAndDomainRelated %s\n", err)
		return nil, err
	}

	resDomains = make([]*models.Domain, 0)
	for _, domain := range resPlatform.Domains {
		if domain.Class == class {
			resDomains = append(resDomains, domain)
		}
	}

	return resDomains, err
}

func (c *BaseDocker) CreateDomainDockerCmd(host *models.Host) (cmd string, err error) {

	resDomains, err := c.GetDomainByPlatformAndClass(c.BaseComponents.Platform.Id, c.BaseComponents.Service.Class)
	if err != nil {
		return "", err
	}
	if len(resDomains) == 0 {
		err = fmt.Errorf("There is no domain name associated with this platform\n")
		return "", nil
	}
	// c.BaseComponents.Service.Domains = append(c.BaseComponents.Service.Domains,resDomains...)

	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}

	cc := fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s ",cmdTls, host.UseIp, c.BaseComponents.Service.DockerPort)

	cpCmds := make([]string, 0)
	// 拼接 命令
	cpCmds = append(cpCmds, "date +%Y%m%d-%H%M%S-%s")
	// 域名配置
	domainConf := ""
	//  nginx 路径
	nginxPath := ""
	// nginx cmd
	cmdsNginx := make([]string, 0)
	cmdsNginx = append(cmdsNginx, "date +%Y%m%d-%H%M%S-%s")

	domainDir := fmt.Sprintf("upload/domain/%s/%s", c.BaseComponents.Project.Alias, c.BaseComponents.Service.Name)
	//if _, err = os.Stat(domainDir); os.IsNotExist(err) {
	err = common.Mkdir(domainDir)
	if err != nil {
		fmt.Println("error: CreateDomainDockerCmd common.Mkdir ", err)
		return "", err
	}

	ports := c.BaseComponents.Task.Service.Port

	port := c.AnalyzePortsAndFindHttps(ports)

	// https
	if port == 443 {
		// 拼接 nginx path
		nginxPath = fmt.Sprintf("%s/%s.vhost.443.conf", domainDir, c.BaseComponents.Docker.Name)
		// 拼接命令	  nginxPath  name
		cpCmds = append(cpCmds, fmt.Sprintf("%s cp %s %s:/etc/nginx/conf.d/vhost.443.conf", cc, nginxPath, c.BaseComponents.Docker.Name))

		//  查询 domaniAndPlantform  多个域名,

		// 获取 domain id
		for _, domain := range resDomains {

			//  拼接 domain.Conf
			domainConf = domainConf + domain.Conf

			// 创建 cmdCrtKey  证书
			cmdsCrtKey := []string{}
			cmdsCrtKey = append(cmdsCrtKey, "date +%Y%m%d-%H%M%S-%s")

			// 拼接 crtPath
			crtPath := fmt.Sprintf("%s/%s.crt", domainDir, domain.Domain)

			// 拼接 KeyPath
			keyPath := fmt.Sprintf("%s/%s.key", domainDir, domain.Domain)

			// echo ...
			cmdsCrtKey = append(cmdsCrtKey, fmt.Sprintf(`echo '%s' > %s`, domain.Crt, crtPath))
			cmdsCrtKey = append(cmdsCrtKey, fmt.Sprintf(`echo '%s' > %s`, domain.Key, keyPath))

			// 拼接 cp 命令
			cpCmds = append(cpCmds, fmt.Sprintf("%s cp %s/%s.crt %s:/etc/pki/tls/%s.crt", cc, domainDir, domain.Domain, c.BaseComponents.Docker.Name, domain.Domain))
			cpCmds = append(cpCmds, fmt.Sprintf("%s cp %s/%s.key %s:/etc/pki/tls/%s.key", cc, domainDir, domain.Domain, c.BaseComponents.Docker.Name, domain.Domain))

			cmdsCrtKey = append(cmdsCrtKey, fmt.Sprintf(`echo '%s' > %s`, domainConf, nginxPath))
			cmdCrtKey := strings.Join(cmdsCrtKey, " ; ")

			// 运行命令 下载 domain 证书
			_, _, err = c.BaseComponents.RunLocal(c.BaseComponents.Task, cmdCrtKey, 0, host.UseIp, "")
			if err != nil {
				return "", err
			}
		}
	}

	// http
	if port == 80 {
		nginxPath = fmt.Sprintf("%s/%s.vhost.80.conf", domainDir, c.BaseComponents.Service.Name)
		cpCmds = append(cpCmds, fmt.Sprintf("%s cp %s/%s.vhost.80.conf %s:/etc/nginx/conf.d/vhost.80.conf", cc, domainDir, c.BaseComponents.Docker.Name, c.BaseComponents.Docker.Name))
		for _, domain := range c.BaseComponents.Service.Domains {
			domainConf = domainConf + domain.Conf80
		}
	}

	// 必须使用echo '$s',用单引号将nginx内容括起来，使用""",shell会将配置内容"$host$request_uri"内容转义。
	cmdsNginx = append(cmdsNginx, fmt.Sprintf(`echo '%s' > %s`, domainConf, nginxPath))

	cmdNginx := strings.Join(cmdsNginx, " ; ")
	// _, err = c.baseComponents.Local(cmdNginx, 0, "add", ip)
	_, _, err = c.BaseComponents.RunLocal(c.BaseComponents.Task, cmdNginx, 0, host.UseIp, "")
	if err != nil {
		return "", err
	}

	cpCmds = append(cpCmds, fmt.Sprintf("sleep 1s; %s restart %s ", cc, c.BaseComponents.Docker.Name))
	cmd = strings.Join(cpCmds, " ; ")

	cmd = fmt.Sprintf("%s:%s##\n", host.UseIp, cmd)

	return cmd, nil
}

func (c *BaseDocker) CreateH5BlackListCmd(host *models.Host) (cmd string, err error) {

	blackListDir := fmt.Sprintf("black_list/%s/%s", c.BaseComponents.Project.Alias, c.BaseComponents.Service.Name)

	err = common.Mkdir(blackListDir)
	if err != nil {
		fmt.Println("error: MkdirH5ClackList common.Mkdir ", err)
		return
	}

	srcFilePath := fmt.Sprintf("%s/deny.conf", blackListDir)

	// echo > deny.conf
	echoCmd := fmt.Sprintf(`echo '%s' > %s`, c.BaseComponents.Service.BlackList, srcFilePath)
	_, _, err = c.BaseComponents.RunLocal(c.BaseComponents.Task, echoCmd, 0, host.UseIp, "")
	if err != nil {
		return
	}
	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}


	//docker cp 命令
	blackListCmds := make([]string, 0)
	cc := fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s ",cmdTls, host.UseIp, c.BaseComponents.Service.DockerPort)
	blackListCmds = append(blackListCmds, fmt.Sprintf("%s cp %s %s:/etc/nginx/conf.d/deny.conf", cc, srcFilePath, c.BaseComponents.Docker.Name))
	blackListCmds = append(blackListCmds, fmt.Sprintf("sleep 1s; %s restart %s ", cc, c.BaseComponents.Docker.Name))
	//blackListCmds = append(blackListCmds, fmt.Sprintf("sleep 1s;  %s exec -i %s nginx -s reload", cc, c.BaseComponents.Docker.Name))
	cmd = strings.Join(blackListCmds, " ; ")
	cmd = fmt.Sprintf("%s:%s##\n", host.UseIp, cmd)

	return
}

func (c *BaseDocker) checkCreateCmdParam(task *models.Task) (err error) {

	if task == nil {
		err = fmt.Errorf("error:  create docker cmd task is nil")
		return
	}
	if c == nil {
		err = fmt.Errorf("error:  create docker cmd BaseDocker is nil")
		return
	}
	if c.BaseComponents.Task.Tag == "" {
		err = fmt.Errorf("error:  create docker cmd task tag is nil")
		return
	}
	if c.BaseComponents.Service.ImagePath == "" {
		err = fmt.Errorf("error:  create docker cmd service.ImagePath is nil")
		return
	}

	if c.BaseComponents.Repo.AddressImage == "" {
		err = fmt.Errorf("error:  create docker cmd repo.AddressImage is nil")
		return
	}
	return
}

func (c *BaseDocker) CreateDockerCmd(task *models.Task, count int, serviceClass string) (err error) {

	dockerTag := fmt.Sprintf("%s/%s:%s", c.BaseComponents.Repo.AddressImage, c.BaseComponents.Service.ImagePath, c.BaseComponents.Task.Tag)
	serviceCmdPullAll := make([]string, 0)
	serviceCmdRunAll := make([]string, 0)
	domainCmdAll := make([]string, 0)
	blackListCmd := ""
	checkDeployHostsResultCmd := ""
	//isCheck := false

	// analyze public param
	public, err := c.AnalyzePublicParam()
	if err != nil {
		beego.Error("AnalyzePublicParam error!")
		return
	}

	ports := ""
	for _, v := range strings.Split(c.BaseComponents.Service.Port, ",") {
		ports = ports + " -p " + v
	}

	platformParam := ""
	var encryptionDomain string
	if serviceClass != "java" && len(task.Service.Platforms) != 0 {

		/*
			// 添加 playfrom param
			if c.BaseComponents.Platform != nil{

				// 指定加速域名
				if len(task.Service.Domains) != 0 {
					encryptionDomain = task.Service.Domains[0].Domain
				}else {
					//  查询 平台所关联的域名,  再筛选 域名类型,  判断域名列表是否有加速域名,如果没有 制定第0个域名
					encryptionDomain, err = c.GetEncryptionDomain(serviceClass)
					if err != nil{
						err = fmt.Errorf("error: GetEncryptionDomain  %s\n ",err)
						return err
					}
				}
				if encryptionDomain != ""{
					encryptionDomain = "wss://" + encryptionDomain
				}
				platformParam, err = c.AnalyzePlatformParam(encryptionDomain)
				if err != nil {
					beego.Error("AnalyzePlatformParam error!")
					return
				}
			}
		*/
		// 添加 playfrom param
		if c.BaseComponents.Platform != nil {

			// 指定加速域名
			if len(task.Service.Domains) != 0 {
				encryptionDomain = "wss://" + task.Service.Domains[0].Domain
			} else {
				encryptionDomain = ""
			}

			platformParam, err = c.AnalyzePlatformParam(encryptionDomain)
			if err != nil {
				beego.Error("AnalyzePlatformParam error!")
				return
			}
		}

	}

	//   add other service class
	m := make(map[string]bool, 0)
	m["h5"] = true
	m["h5-proxy"] = true
	m["h5-site"] = true

	//var domainCmd string
	// 遍历hosts  创建所有主机的 pull run 命令
	for i, host := range c.BaseComponents.Hosts {

		// create docker pull
		cmdPull, err := c.createDockerPull(dockerTag, host)
		if err != nil {
			beego.Error("error: create docker pull cmd: ", err)
			return err
		}
		serviceCmdPull := fmt.Sprintf("%s:%s", host.UseIp, cmdPull)
		serviceCmdPullAll = append(serviceCmdPullAll, serviceCmdPull)

		// create docker run cmd
		cmdRun, err := c.createDockerRun(dockerTag, host, ports, public, platformParam)
		if err != nil {
			beego.Error("error: create docker run cmd: ", err)
			return err
		}
		serviceCmdRun := fmt.Sprintf("%s:%s", host.UseIp, cmdRun)
		serviceCmdRunAll = append(serviceCmdRunAll, serviceCmdRun)

		// create check docker run log cmd
		if i == 0 && c.BaseComponents.Service.LogKeyword != "" {
			checkDeployHostsResultCmd = c.CreateCheckDeployHostsResultCmd(host)
			checkDeployHostsResultCmd = fmt.Sprintf("%s##", checkDeployHostsResultCmd)
			checkDeployHostsResultCmd = fmt.Sprintf("%s:%s", host.UseIp, checkDeployHostsResultCmd)
			//isCheck = true
		}

		// 判断服务有绑定域名
		if c.BaseComponents.Service.Class != "java" && len(c.BaseComponents.Service.Platforms) != 0 {
			domainCmd, err := c.CreateDomainDockerCmd(host)
			if err != nil {
				beego.Error("error: create domain docker cmd: ", err)
				return err
			}
			if domainCmd != "" {
				domainCmdAll = append(domainCmdAll, domainCmd)
			}
		}

		if m[c.BaseComponents.Service.Class] && c.BaseComponents.Service.BlackList != "" {
			blackListCmd, err = c.CreateH5BlackListCmd(host)
		}
	}

	if len(domainCmdAll) < 1 {
/*
		if isCheck {
			c.Cmds = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", "docker-check", checkDeployHostsResultCmd, "docker-pull", serviceCmdPullAll, "docker-run", serviceCmdRunAll)
		} else {
			c.Cmds = fmt.Sprintf("%s\n%s\n%s\n%s\n", "docker-pull", serviceCmdPullAll, "docker-run", serviceCmdRunAll)
		}
*/
		c.Cmds = fmt.Sprintf("%s\n%s\n%s\n%s\n", "docker-pull", serviceCmdPullAll, "docker-run", serviceCmdRunAll)

	} else {
/*
		if isCheck {
			c.Cmds = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n", "docker-check", checkDeployHostsResultCmd, "docker-pull", serviceCmdPullAll, "docker-run", serviceCmdRunAll, "docker-domain", domainCmdAll)
		} else {
			c.Cmds = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", "docker-pull", serviceCmdPullAll, "docker-run", serviceCmdRunAll, "docker-domain", domainCmdAll)
		}
*/
		c.Cmds = fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", "docker-pull", serviceCmdPullAll, "docker-run", serviceCmdRunAll, "docker-domain", domainCmdAll)
	}

	// add other service class

	if m[c.BaseComponents.Service.Class] && blackListCmd != "" {
		c.Cmds += fmt.Sprintf("%s\n[%s]", "h5-black-list", blackListCmd)
	}

	task.Count = count
	task.Cmd = c.Cmds

	return
}

type DomainNode struct {
	Domain    string
	IsQuicken int
}

//  查询 平台所关联的域名,  再筛选 域名类型,  判断域名列表是否有加速域名,如果没有 制定第0个域名
func (c *BaseDocker) GetEncryptionDomain(serviceClass string) (encryptionDomain string, err error) {

	Platform, err := models.GetPlatformAndDomainRelated(c.BaseComponents.Platform)
	if err != nil {
		err := fmt.Errorf("error: GetPlatformAndDomainRelated err   %s", err)
		return "", err
	}

	quickenDomainList := make([]*DomainNode, 0)
	normalDomainList := make([]*DomainNode, 0)

	for _, domain := range Platform.Domains {

		if domain.Class != serviceClass {
			continue
		}

		domainNode := new(DomainNode)
		domainNode.Domain = domain.Domain
		domainNode.IsQuicken = domain.Quicken

		if domainNode.IsQuicken == 1 {
			quickenDomainList = append(quickenDomainList, domainNode)
		} else {
			normalDomainList = append(normalDomainList, domainNode)
		}
	}

	if len(quickenDomainList) != 0 {
		encryptionDomain = quickenDomainList[0].Domain
	} else {
		if len(normalDomainList) != 0 {
			encryptionDomain = normalDomainList[0].Domain
		}
	}

	return encryptionDomain, nil
}

func (c *BaseDocker) CreateCheckDeployHostsResultCmd(host *models.Host) (checkDeployHostsResultCmd string) {

	Keyword := c.BaseComponents.Service.LogKeyword

	if Keyword != "" {
		cmdTls := ""
		dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
		if dockerTlsPath != ""{
			cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
		}

		dockerManager := fmt.Sprintf("tcp://%s:%s", host.UseIp, c.BaseComponents.Service.DockerPort)
		cc := fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s ",cmdTls, host.UseIp, c.BaseComponents.Service.DockerPort)
		date := "date +%Y%m%d-%H%M%S-%s"
		cmdAwk := fmt.Sprintf("awk '/%s/{print $0;system(\"ps x |grep \\\"docker %s -H %s logs -f %s\\\"|grep -Ev \\\"awk|grep\\\" |cut -b -6|xargs kill\")}'", Keyword,cmdTls, dockerManager, c.BaseComponents.Docker.Name)

		checkDeployHostsResultCmd = fmt.Sprintf("%s ; %s logs -f %s|%s", date, cc, c.BaseComponents.Docker.Name, cmdAwk)

	}

	return
}

/*
解析 platform value 配置
*/
func (c *BaseDocker) AnalyzePlatformParam(encryptionDomain string) (resPlatformValue string, err error) {

	platformValues := make([]*models.PlatformValue, 0)
	fmt.Printf("platformValue: %s\n",c.BaseComponents.Platform.Value)
	if err := json.Unmarshal([]byte(c.BaseComponents.Platform.Value), &platformValues); err != nil {
		fmt.Printf("error:  json.Unmarshal  ::  %s\n",c.BaseComponents.Platform.Value)
		return resPlatformValue, err
	}

	var EncryptKey string
	var EncryptValue string
	var VUE_APP_NAME string        ///  平台参数
	var VUE_APP_PLATFORM_ID string ///  平台参数中
	var keyStr string

	for _, v := range platformValues {

		if "CHANNEL_NAME" == v.Key {
			VUE_APP_NAME = v.Value
		}
		if "CHANNEL_ID" == v.Key {
			VUE_APP_PLATFORM_ID = v.Value
		}
		if "ENCRYPTED_KEY" == v.Key {
			keyStr = v.Value
		}

		if v.Key == "CHANNEL_NAME" || v.Key == "KEYWORDS" || v.Key == "META_DESC" {
			v.Value = "'" + v.Value + "'"
		}
		subPara := fmt.Sprintf("-e %s=%s ", v.Key, v.Value)
		resPlatformValue = resPlatformValue + subPara
	}

	// 平台参数加密
	if VUE_APP_NAME != "" && VUE_APP_PLATFORM_ID != "" && keyStr != "" {
		vue := make(map[string]string, 0)

		emqttConf, err := models.GetConfByConfName("emqtt")
		if err != nil {
			err = fmt.Errorf("error:  GetConfByConfName err ", err)
			return "", err
		}

		emqttUser, emqttPwd := c.GetPublicEncryptParam(emqttConf.Value)

		vue["VUE_APP_MQTT_USERNAME"] = emqttUser
		vue["VUE_APP_MQTT_PASSWORD"] = emqttPwd
		vue["VUE_APP_NAME"] = VUE_APP_NAME                 ///  平台参数中
		vue["VUE_APP_PLATFORM_ID"] = VUE_APP_PLATFORM_ID   ///  平台参数中
		vue["VUE_APP_MQTTT_BROKER_URL"] = encryptionDomain //  取 h5服务关联的第一个域名 做mqtt的 broker url

		vueJson, _ := json.Marshal(vue) ///  转 json
		vueStr := string(vueJson)

		vueEncryptValue, err := common.DesECBEncrypt(vueStr, keyStr)
		if err != nil {
			return "", err
		}

		EncryptKey = "VUE_APP_CONFIG_ENCRYPTED_STR"
		EncryptValue = vueEncryptValue
	}

	// 追加 平台加密参数
	if EncryptKey != "" && EncryptValue != "" {
		subPara := fmt.Sprintf("-e %s=%s ", EncryptKey, EncryptValue)

		resPlatformValue = fmt.Sprintf("%s %s ", resPlatformValue, subPara)
	}

	return
}

/*
解析 public conf 配置
*/
func (c *BaseDocker) GetPublicEncryptParam(emqttPublicParam string) (emqttUser string, emqttPwd string) {

	confValue := emqttPublicParam
	values := strings.Split(confValue, "},{")
	var k string
	var v string

	for _, value := range values {
		value1 := strings.Trim(value, "[")
		value1 = strings.Trim(value1, "]")
		value1 = strings.Trim(value1, "{")
		value1 = strings.Trim(value1, "}")
		value2s := strings.Split(value1, ",")

		tempMap := make(map[string]string, 0)
		for _, pair := range value2s {

			index := strings.Index(pair, ":")
			tempK := pair[:index]
			tempV := pair[index+1:]
			tempK = strings.Trim(tempK, "\"")
			tempV = strings.Trim(tempV, "\"")
			tempK = strings.TrimSpace(tempK)
			tempV = strings.TrimSpace(tempV)

			tempMap[tempK] = tempV
			if _, ok := tempMap["Key"]; ok {
				k = tempMap["Key"]
			}

			if _, ok := tempMap["Value"]; ok {
				v = tempMap["Value"]
			}
		}

		if v, ok := tempMap["Key"]; ok {
			kk := v
			k = strings.Trim(kk, "\"")
		}
		if v, ok := tempMap["Value"]; ok {
			vv := v
			v = strings.Trim(vv, "\"")
		}

		if k == "emqtt.username" {
			emqttUser = v
		}

		if k == "emqtt.password" {
			emqttPwd = v
		}
	}
	return
}

/*
解析 public conf 配置
*/

func (c *BaseDocker) AnalyzePublicParam() (public string, err error) {

	confObj := make([]*models.ConfValue, 0)
	for _, conf := range c.BaseComponents.Service.Confs {
		if err := json.Unmarshal([]byte(conf.Value), &confObj); err != nil {
			return "", err
		}
		for _, v := range confObj {
			subPara := fmt.Sprintf("-e %s=%s ", v.Key, v.Value)
			public = public + subPara
		}
	}

	if c.BaseComponents.Service.Class == "java" {
		defConf, err := models.GetConfByConfNameAndProjectId("def_other", c.BaseComponents.Service.Project.Id)
		if err != nil {
			err = fmt.Errorf("error:  GetConfByConfName err %s\n", err)
		}

		if len(defConf) == 1 {
			var defConfValue []*models.ConfValue
			err = json.Unmarshal([]byte(defConf[0].Value), &defConfValue)
			for _, v := range defConfValue {
				subPara := fmt.Sprintf("-e %s=%s ", v.Key, v.Value)
				public = public + subPara
			}
		}
	}

	return
}

func (c *BaseDocker) AnalyzeServiceValueHostId(serviceValue string) string {

	index := strings.LastIndex(serviceValue, ":")
	hostId := serviceValue[index+1:]
	hostId = strings.Trim(hostId, "}]")
	hostId = strings.Trim(hostId, "[{")
	hostId = strings.Trim(hostId, "\"")

	return hostId
}

/*
解析 service_Value 公共 配置
*/

func (c *BaseDocker) AnalyzePrivateParam(hostId int) (private string, err error) {

	serviceValue := make([]*models.ServiceValue, 0)
	if c.BaseComponents.Service.Value == "" {
		return "", nil
	}
	value := c.BaseComponents.Service.Value
	if err := json.Unmarshal([]byte(value), &serviceValue); err != nil {
		return "", err
	}
	for _, v := range serviceValue {
		if v.HostId != 0 && v.HostId != hostId {
			continue
		}
		private = fmt.Sprintf("%s %s", private, v.Value)
	}

	return
}

/**
* build镜像
 */
func (c *BaseDocker) Build(operationRecord *models.OperationRecord) (repoTag string, err error) {
	name := c.BaseComponents.Service.Name

	dir := fmt.Sprintf("upload/task/%s/%s", name, name)
	zip := fmt.Sprintf("upload/task/%s/%s.zip", name, name)

	dockerManager := fmt.Sprintf("tcp://%s:%s", operationRecord.Host.UseIp, c.BaseComponents.Service.DockerPort)
	repoTag = fmt.Sprintf("%s_%s", name, time.Now().Format("20060102_150405"))

	dockerFile := "FROM harbor.one2.cc/devops/nginx_akgo:1.11.13\nADD html /data/www/html"

	err = common.RmMkdir(dir)
	if err != nil {
		fmt.Printf("error: Build common.RmMkdir \"%s\"  %s\n ", dir, err)
		return
	}
	err = c.BaseComponents.LocalOtherNull(fmt.Sprintf("新建初始化目录：%s", dir), 10, operationRecord)
	if err != nil {
		return
	}

	dockerFilePath := fmt.Sprintf("%s/Dockerfile", dir)
	err = common.FileWriter(dockerFilePath, dockerFile)
	if err != nil {
		fmt.Println("error: Build common.FileWriter   ", err)
		return
	}
	err = c.BaseComponents.LocalOtherNull(fmt.Sprintf("新建Dockerfile：%s", dockerFilePath), 20, operationRecord)
	if err != nil {
		return
	}

	cmd_unzip := fmt.Sprintf("/usr/bin/env unzip -o %s -d %s", zip, dir)
	_, err = c.BaseComponents.LocalOther(cmd_unzip, 0, 40, operationRecord)
	if err != nil {
		return
	}
	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}


	cmd_login := fmt.Sprintf("/usr/bin/env docker %s -H %s login -u %s -p %s %s",cmdTls, dockerManager, c.BaseComponents.Repo.User, c.BaseComponents.Repo.Pass, c.BaseComponents.Repo.AddressImage)
	_, err = c.BaseComponents.LocalOther(cmd_login, 0, 60, operationRecord)
	if err != nil {
		return
	}

	cmd_build := fmt.Sprintf("/usr/bin/env docker %s -H %s build --pull %s -t  %s/%s:%s",cmdTls, dockerManager, dir, c.BaseComponents.Repo.AddressImage, c.BaseComponents.Service.ImagePath, repoTag)
	_, err = c.BaseComponents.LocalOther(cmd_build, 0, 80, operationRecord)
	if err != nil {
		return
	}

	cmd_date := "date +%Y%m%d-%H%M%S-%s"
	cmd_push := fmt.Sprintf("%s;/usr/bin/env docker %s -H %s push %s/%s:%s", cmdTls, cmd_date, dockerManager, c.BaseComponents.Repo.AddressImage, c.BaseComponents.Service.ImagePath, repoTag)
	_, err = c.BaseComponents.LocalOther(cmd_push, 0, 100, operationRecord)
	if err != nil {
		return
	}

	return
}

/**
* 控制docker重启
 */
func (c *BaseDocker) Restart(operationRecord *models.OperationRecord) error {
	//dockermanager := beego.AppConfig.String("dockerManager")
	//c.BaseComponents.Service.DockerPort

	name := ""
	if c.BaseComponents.Service.DockerName == "" {
		name = c.BaseComponents.Service.Name + "_" + c.BaseComponents.Project.Alias
	} else {
		name = c.BaseComponents.Service.DockerName
	}
	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}

	var OutTime = 60
	cmds := []string{}
	cmds = append(cmds, "date +%Y%m%d-%H%M%S-%s")
	if c.BaseComponents.Service.KillTime == "" {
		cmds = append(cmds, fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s restart %s",cmdTls, operationRecord.Host.UseIp, c.BaseComponents.Service.DockerPort, name))
	} else {
		cmds = append(cmds, fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s restart -t=%s %s",cmdTls, operationRecord.Host.UseIp, c.BaseComponents.Service.DockerPort, c.BaseComponents.Service.KillTime, name))
		OutTime = common.StrToInt(c.BaseComponents.Service.KillTime) + 60
	}

	cmd := strings.Join(cmds, " ; ")
	_, err := c.BaseComponents.LocalOther(cmd, OutTime, 100, operationRecord)

	if err != nil {
		beego.Error(fmt.Sprintf("%s restart %s: %s", operationRecord.Host.UseIp, name, err))
	}

	return err
}

/**
* 控制docker关闭
 */
func (c *BaseDocker) Stop(operationRecord *models.OperationRecord) error {
	//dockermanager := beego.AppConfig.String("dockerManager")
	//c.BaseComponents.Service.DockerPort

	name := ""
	if c.BaseComponents.Service.DockerName == "" {
		name = c.BaseComponents.Service.Name + "_" + c.BaseComponents.Project.Alias
	} else {
		name = c.BaseComponents.Service.DockerName
	}
	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}

	var OutTime = 60
	cmds := []string{}
	cmds = append(cmds, "date +%Y%m%d-%H%M%S-%s")
	if c.BaseComponents.Service.KillTime == "" {
		cmds = append(cmds, fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s stop %s",cmdTls, operationRecord.Host.UseIp, c.BaseComponents.Service.DockerPort, name))
	} else {
		cmds = append(cmds, fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s stop -t=%s %s", cmdTls, operationRecord.Host.UseIp, c.BaseComponents.Service.DockerPort, c.BaseComponents.Service.KillTime, name))
		OutTime = common.StrToInt(c.BaseComponents.Service.KillTime) + 60
	}

	cmd := strings.Join(cmds, " ; ")
	_, err := c.BaseComponents.LocalOther(cmd, OutTime, 100, operationRecord)

	if err != nil {
		beego.Error(fmt.Sprintf("%s stop %s: %s", operationRecord.Host.UseIp, name, err))
	}

	return err
}

/**
* 控制docker nginx reload 重载
 */
func (c *BaseDocker) Reload(operationRecord *models.OperationRecord) error {
	//dockermanager := beego.AppConfig.String("dockerManager")

	name := ""
	if c.BaseComponents.Service.DockerName == "" {
		name = c.BaseComponents.Service.Name + "_" + c.BaseComponents.Project.Alias
	} else {
		name = c.BaseComponents.Service.DockerName
	}
	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}


	var OutTime = 60
	cmds := []string{}
	cmds = append(cmds, "date +%Y%m%d-%H%M%S-%s")
	cmds = append(cmds, fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s exec -i %s nginx -t", cmdTls, operationRecord.Host.UseIp, c.BaseComponents.Service.DockerPort, name))
	cmds = append(cmds, fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s exec -i %s nginx -s reload", cmdTls, operationRecord.Host.UseIp, c.BaseComponents.Service.DockerPort, name))

	cmd := strings.Join(cmds, " && ")
	_, err := c.BaseComponents.LocalOther(cmd, OutTime, 100, operationRecord)

	if err != nil {
		beego.Error(fmt.Sprintf("%s reload %s: %s", operationRecord.Host.UseIp, name, err))
	}

	return err
}

/**
* 下载站之类的应用下载ZIP包
 */
func (c *BaseDocker) DownloadZip(operationRecord *models.OperationRecord) (string, error) {
	dockermanager := beego.AppConfig.String("dockerManager")

	name := ""
	if c.BaseComponents.Docker.Name == "" {
		name = c.BaseComponents.Service.Name + "_" + c.BaseComponents.Project.Alias
	} else {
		name = c.BaseComponents.Docker.Name
	}

	dir := fmt.Sprintf("upload/download/%d/%s", c.BaseComponents.Service.Project.Id, c.BaseComponents.Service.Name)
	err := common.RmMkdir(dir)
	if err != nil {
		fmt.Println("error: DownloadZip common.RmMkdir  ", dir, "  ", err)
		return "", err
	}
	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}


	cmds := []string{}
	cmds = append(cmds, "date +%Y%m%d-%H%M%S-%s")

	dirCp := fmt.Sprintf("%s/html", dir)
	cmds = append(cmds, fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s cp %s:%s %s",cmdTls, operationRecord.Host.UseIp, dockermanager, name, c.BaseComponents.Service.ReleaseTo, dirCp)) ///  添加ReleaseTo

	cmd := strings.Join(cmds, " ; ")
	_, err = c.BaseComponents.LocalOther(cmd, 0, 100, operationRecord)
	if err != nil {
		beego.Error(fmt.Sprintf("%s cp %s: %s", operationRecord.Host.UseIp, name, err))
	}

	pathZip := fmt.Sprintf("%s/%s.zip", dir, c.BaseComponents.Service.Name)
	err = common.Zip(dirCp, pathZip)
	if err != nil {
		beego.Error(fmt.Sprintf("zip is fail!"))
	}

	return pathZip, err
}

/**
* 检查docker的项目某hots的ps数
 */
func (c *BaseDocker) GetStatusHost(host string) (res []map[string]string, err error) {

	dockerPort := beego.AppConfig.String("dockerport")

	PsName := "------"
	PsStatus := "------"
	PsCreatedAt := "------"
	PsImage := "------\n"
	cmdTls := ""
	dockerTlsPath := c.BaseComponents.Service.DockerTlsPath
	if dockerTlsPath != ""{
		cmdTls = fmt.Sprintf( " --tls --tlscacert %s/cacert.pem --tlscert %s/certs/dk.crt  --tlskey %s/private/dk.key ",dockerTlsPath,dockerTlsPath,dockerTlsPath)
	}


	cmd := fmt.Sprintf("/usr/bin/env docker %s -H tcp://%s:%s ps -a --format {{.Names}}#{{.Status}}#{{.CreatedAt}}#{{.Image}}",cmdTls, host, dockerPort)

	s, _, err := c.BaseComponents.RunLocal(c.BaseComponents.Task, cmd, 3, host, "")
	if err != nil {
		stErr := fmt.Sprintf("%s", err)
		if stErr == "cmd time out" {
			PsStatus = "超时"
		}
		res = append(res, map[string]string{"ps_name": PsName, "ps_status": PsStatus, "ps_created_at": PsCreatedAt, "ps_image": PsImage})

	} else {
		if s.Result != "" {
			for _, result := range strings.Split(s.Result, "\n") {
				fmt.Println("__", result)
				if result == "" {
					continue
				}
				arr := strings.Split(result, "#")
				PsName = arr[0]
				PsStatus = arr[1]
				PsCreatedAt = arr[2][:19]
				if len(strings.Split(arr[3], ":")) == 2 {
					PsImage = strings.Split(arr[3], ":")[1]
				} else if len(strings.Split(arr[3], ":")) == 3 {
					PsImage = strings.Split(arr[3], ":")[2]
				} else {
					PsImage = arr[3]
				}

				res = append(res, map[string]string{"ps_name": PsName, "ps_status": PsStatus, "ps_created_at": PsCreatedAt, "ps_image": PsImage})
			}
		} else {
			PsStatus = "无容器"
			res = append(res, map[string]string{"ps_name": PsName, "ps_status": PsStatus, "ps_created_at": PsCreatedAt, "ps_image": PsImage})
		}

	}

	return res, err
}
