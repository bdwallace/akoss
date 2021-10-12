package confManagement

import (
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"library/common"
	"library/components"
	gopubssh "library/ssh"
	"models"
	"models/request"
	"models/response"
	"os"
	"sync"
)


func RegisterBaseComm(project *models.Project, platform *models.Platform, service *models.Service)*components.BaseDocker{

	baseCom := components.BaseComponents{}
	baseCom.SetProject(project)
	baseCom.SetPlatform(platform)
	baseCom.SetService(service)
	baseDocker := new(components.BaseDocker)
	baseDocker.SetBaseComponents(baseCom)
	return  baseDocker
}



//////////////////////////////////////////////////////////////////////////////////////////////

var CHCP = make(chan int)
var CHRD = make(chan int)

type Producer struct {
	Project			*request.ProjectConfNginx
	Platform 		*request.PlatformConfNginx
	Service 		*request.ServiceConfNginx
	Host			*request.HostConfNginx
	Conf 			*request.ConfFile
	Wg 				sync.WaitGroup

}

type Consumer struct {
	//RespConfOfProject 			*response.ProjectConfNginx
}



var CONFDIR = "/etc/nginx"
var NGDIR = "/nginx"
var NGCONF = "/nginx.conf"
var NGCONF_D = "/conf.d/"
var LOCALBASECONFDIR = "/tmp/akoss_conf_management"

func (pr *Producer) GetAllFrontServicesNginxConfToLocal()(err error){
	if err = pr.GetAllPlatforms(); err != nil{
		return
	}

	return
}


func (pr *Producer) GetAllPlatforms() (err error){

	pr.Project.Platforms = make([]*request.PlatformConfNginx,0,0)

	// get platforms
	srcPlatforms, err := models.GetPlatformByProjectId(pr.Project.ProjectId)
	if err != nil{
		fmt.Println("error:  GetPlatforms -> GetPlatformByProjectId is failed  ::  ",err)
		return err
	}

	for _, srcP := range srcPlatforms{

		tmpReqPlatform := new(request.PlatformConfNginx)
		tmpReqPlatform.PlatformId = srcP.Id
		tmpReqPlatform.PlatformName = srcP.Name
		pr.Platform = tmpReqPlatform
		if err = pr.GetServices(tmpReqPlatform, srcP); err != nil{
			return err
		}
		pr.Project.Platforms = append(pr.Project.Platforms,tmpReqPlatform)
	}


	return

}


func (pr *Producer) GetServices(desP * request.PlatformConfNginx, srcP *models.Platform) (err error){

	desP.Services = make([]*request.ServiceConfNginx,0, 0)

	// get services of platforms
	srcP, err =  models.GetPlatformAndServiceRelated(srcP)
	if err != nil{
		fmt.Println("error:  GetServices -> GetPlatformAndServiceRelated")
		return err
	}
	for _, s := range srcP.Services {
		tmpReqService := new(request.ServiceConfNginx)
		tmpReqService.ServiceId = s.Id
		tmpReqService.ServiceName = s.Name
		tmpReqService.ServiceClass = s.Class
		tmpReqService.BaseDocker = RegisterBaseComm(srcP.Project,srcP,s)
		pr.Service = tmpReqService
		if err = pr.GetHosts(tmpReqService, s); err != nil{
			return err
		}
		desP.Services = append(desP.Services, tmpReqService)
	}

	return
}

func (pr *Producer) GetHosts(desS *request.ServiceConfNginx, srcS *models.Service)(err error){

	desS.Hosts = make([]*request.HostConfNginx, 0, 0)

	srcS, err = models.GetServiceRelatedForHosts(srcS)
	if err != nil{
		fmt.Println("error:  GetHosts -> GetServiceRelatedFforHosts is failed  ::  ",err)
		return err
	}
	// get hosts
	for _, h := range srcS.Hosts {
		tmpRepHost := new(request.HostConfNginx)
		tmpRepHost.HostId = h.Id
		tmpRepHost.UseIp = h.UseIp
		tmpRepHost.HostName = h.Name
		// cp hosts confs
		pr.Host = tmpRepHost
		//pr.GetConf(tmpRepHost)
		if err = pr.CPHostNginxConfDir(); err != nil{
			return err
		}

		desS.Hosts = append(desS.Hosts, tmpRepHost)
	}
	pr.Wg.Wait()

	return
}

// docker cp xxxxx:/etc/nginx /tmp/akoss_conf_management/....
func (pr *Producer) CPHostNginxConfDir ()(err error){

	cpDesDir := fmt.Sprintf("%s/%s/%s/%s/%s",
		LOCALBASECONFDIR,
		pr.Project.ProjectAlias,
		pr.Platform.PlatformName,
		pr.Service.ServiceName,
		pr.Host.UseIp,
	)


	cpSrcDir := fmt.Sprintf("%s:%s",
		pr.Service.BaseDocker.BaseComponents.Docker.Name,
		CONFDIR,
	)

	// exist
	if FileExist(cpDesDir){
		// remove
		if err = os.RemoveAll(cpDesDir); err != nil{
			fmt.Printf("error:  CPHostNginxConfdir -> RemoveAll is failed  %s  ::  %s\n",cpDesDir,err)
			return err
		}
		id, err := models.DeleteHostConfByLocalPath(cpDesDir)
		if err != nil || id <= 0 {
			fmt.Printf("error:  CPHostNginxConfdir -> DeleteHostConfByLocalPath is failed  id: %d   %s  ::  %s\n",id ,cpDesDir, err)
			return err
		}
	}

	// mkdir
	if err = common.Mkdir(cpDesDir); err != nil{
		fmt.Println("error:  CPHostNginxConfdir -> Mkdir is failed  ::  ",err)
		return
	}

	dockerCPCmd := fmt.Sprintf("/usr/bin/env docker -H tcp://%s:2375 cp %s %s",pr.Host.UseIp,cpSrcDir,cpDesDir)

	pr.Wg.Add(1)
	go func() {
		defer pr.Wg.Done()
		s, cmdErr := gopubssh.CommandLocal(dockerCPCmd,components.SSHTIMEOUT)
		if cmdErr != nil {
			if cmdErr.Error() == "cmd time out" {
				s.Result = fmt.Sprintf("cmd time out  ::  %s\n", dockerCPCmd)
				beego.Error(s.Result)
			}
		}
		if FileExist(cpDesDir + "/nginx"){
			fmt.Println("SUCCESS:  ",dockerCPCmd)
			if err = pr.SaveMetaData(cpDesDir); err != nil{
				fmt.Printf("error:  save metadata of nginx conf is failed  %s  %s  ::  %s \n",pr.Service.ServiceName,pr.Host.UseIp,err)
			}
		}else {
			fmt.Println("FAILED:  ",dockerCPCmd)
		}
	}()

	return
}




func (pr *Producer) GetNGConfMetaData(desH * request.HostConfNginx){
	cf := new(request.ConfFile)
	cf.FilePath = fmt.Sprintf("%s%s%s",LOCALBASECONFDIR,CONFDIR,NGCONF)
	cf.FileName = "nginx.conf"
	cf.FileDir = fmt.Sprintf("%s%s",LOCALBASECONFDIR,CONFDIR)

	desH.NginxConfs = append(desH.NginxConfs, cf)
}


func (pr *Producer) SaveMetaData (cpDesDir string)(err error){

	hostConf := new(models.HostConf)
	hostConf.Project = &models.Project{Id:pr.Project.ProjectId}
	hostConf.Platform = &models.Platform{Id:pr.Platform.PlatformId}
	hostConf.Service = &models.Service{Id:pr.Service.ServiceId}
	hostConf.Host = &models.Host{Id:pr.Host.HostId}
	hostConf.ConfName = pr.Service.ServiceName
	hostConf.Class = pr.Service.ServiceClass
	hostConf.LocalPath = cpDesDir

	if err = models.AddHostConf(hostConf); err != nil{
		fmt.Printf("error:  insert host conf to database is failed  %s  ::  %s\n",hostConf.ConfName, err.Error())
		return
	}

	return
}


func FileExist(path string)bool{
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}


//////////////////////


func (c *Consumer) SearchNGConf(platformId int, serviceClass string, hostId int)(hostConf []*models.HostConf, err error){


	// three search options
	if platformId > 0 && serviceClass != "" && hostId > 0{
		return models.SearchHostConfOfPlatformClassHost(platformId,serviceClass,hostId)
	}

	// two search options
	if platformId > 0 && serviceClass != ""{
		return models.SearchHostConfOfPlatformClass(platformId,serviceClass)
	}

	if platformId > 0 && hostId > 0{
		return models.SearchHostConfOfPlatformHost(platformId,hostId)
	}

	if platformId > 0{
		return models.SearchHostConfOfPlatform(platformId)
	}


	return nil, fmt.Errorf("no search options")
}




func (pr *Producer) GetNGConfDMetaData(desH * request.HostConfNginx){
	dirConfD := fmt.Sprintf("%s%s%s",LOCALBASECONFDIR,CONFDIR,NGCONF_D)
	fileInfoList, _ := ioutil.ReadDir(dirConfD)

	for _, file := range fileInfoList{
		fmt.Println("filename:  ", file.Name())
	}
}


func (c *Consumer)GetNginxConf (hostConf *models.HostConf) ([]*response.ConfFile, error) {

	res, err := c.ReadConfFile(hostConf)
	if err != nil{
		return nil, err
	}
	return res, nil

}


func (c *Consumer) ReadConfFile(hostConf *models.HostConf)(res []*response.ConfFile, err error) {

	// read nginx.conf
	if FileExist(hostConf.LocalPath + NGDIR + NGCONF){
		confFileNG, err := c.ReadNginxConfInfo(hostConf.LocalPath +  NGDIR + NGCONF, hostConf)
		if err != nil{
			return nil, err
		}
		res = append(res, confFileNG)

	}

	// read nginx.d/*
	if FileExist(hostConf.LocalPath + NGDIR + NGCONF_D) {
		resConfDFile, err := ReadConfD(hostConf.LocalPath + NGDIR + NGCONF_D, hostConf)
		if err != nil{
			return nil, err
		}
		res = append(res, resConfDFile...)
	}

	return res, nil
}



func ReadConfD(dir string, hostConf * models.HostConf) ([]*response.ConfFile, error){

	resConfFiles := make([]*response.ConfFile,0,0)

	fileInfoList, _ := ioutil.ReadDir(dir)
	for _, file :=  range fileInfoList {
		name := file.Name()
		tmpConf := new(response.ConfFile)
		tmpConf.ConfId = hostConf.Id
		tmpConf.ConfFileName += "/nginx/conf.d/" + name
		buf, err := ioutil.ReadFile(dir + name)
		if err != nil{
			return nil, err
		}
		tmpConf.ConfInfo = string(buf)

		tmpConf.ConfRelation.ProjectId = hostConf.Project.Id
		tmpConf.ConfRelation.ProjectName = hostConf.Project.Name
		tmpConf.ConfRelation.PlatformId = hostConf.Platform.Id
		tmpConf.ConfRelation.PlatformName = hostConf.Platform.Name
		tmpConf.ConfRelation.ServiceId = hostConf.Service.Id
		tmpConf.ConfRelation.ServiceName = hostConf.Service.Name
		tmpConf.ConfRelation.ServiceClass = hostConf.Class
		tmpConf.ConfRelation.HostId = hostConf.Host.Id
		tmpConf.ConfRelation.HostIp = hostConf.Host.UseIp


		resConfFiles = append(resConfFiles, tmpConf)
	}

	return resConfFiles, nil

}

func (c *Consumer)ReadNginxConfInfo(filePath string,hostConf *models.HostConf)(*response.ConfFile,error){
	confFile := new(response.ConfFile)
	buf, err := ioutil.ReadFile(filePath)
	if err != nil{
		return nil,err
	}

	confFile.ConfId = hostConf.Id
	confFile.ConfFileName = NGDIR + NGCONF
	confFile.ConfInfo = string(buf)

	confFile.ConfRelation.ProjectId = hostConf.Project.Id
	confFile.ConfRelation.ProjectName = hostConf.Project.Name
	confFile.ConfRelation.PlatformId = hostConf.Platform.Id
	confFile.ConfRelation.PlatformName = hostConf.Platform.Name
	confFile.ConfRelation.ServiceId = hostConf.Service.Id
	confFile.ConfRelation.ServiceName = hostConf.Service.Name
	confFile.ConfRelation.ServiceClass = hostConf.Class
	confFile.ConfRelation.HostId = hostConf.Host.Id
	confFile.ConfRelation.HostIp = hostConf.Host.UseIp


	return confFile,nil
}


