package confManagement

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"library/common"
	"library/components"
	gopubssh "library/ssh"
	"models"
	"models/response"
	"os"
	"strconv"
	"strings"
)

var CONFDIR = "/etc/nginx"
var NGDIR = "/nginx"
var NGCONF = "/nginx.conf"
var NGCONF_D = "/conf.d/"

var USRLOCAL = "/usr/local"
var ORNGDIR = "/openresty/nginx"
var OPNGCONF = "/conf/nginx.conf"
var WAFCONF = "/conf/waf/conf.lua"
var ACCRDSLUA = "/lua/access_by_redis.lua"

var LOCALBASECONFDIR = "/tmp/akoss_conf_management"

func RegisterBaseComm(project *models.Project, service *models.Service)*components.BaseDocker{

	baseCom := components.BaseComponents{}
	baseCom.SetProject(project)
	baseCom.SetService(service)
	baseDocker := new(components.BaseDocker)
	baseDocker.SetBaseComponents(baseCom)
	return  baseDocker
}


func FileExist(path string)bool{
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func GetConfAndCpToLocal(confRelation * response.ConfRelation, docker *components.BaseDocker)(err error){

	cpDesDir := fmt.Sprintf("%s/%s",LOCALBASECONFDIR,confRelation.ServiceName)
	if err = MakeDirOfLocal(cpDesDir); err != nil{
		return err
	}

	cpSrcDir := fmt.Sprintf("%s:%s",
		docker.BaseComponents.Docker.Name,
		CONFDIR,
	)
	cpOpenrestyConfDir := fmt.Sprintf("%s:%s",
		docker.BaseComponents.Docker.Name,
		"/usr/local/openresty/nginx",
	)

	cpOpenrestyLuaDir := fmt.Sprintf("%s:%s",
		docker.BaseComponents.Docker.Name,
		"/usr/local/openresty/nginx",
	)
	openrestyLocalConfDir := cpDesDir + ORNGDIR + "/conf"
	if err = MakeDirOfLocal(openrestyLocalConfDir); err != nil{
		return
	}
	openrestyLocalWafDir := openrestyLocalConfDir + "/waf"
	if err = MakeDirOfLocal(openrestyLocalWafDir); err != nil{
		return
	}

	openrestyLocalLuaDir := cpDesDir + ORNGDIR + "/lua"
	if err = MakeDirOfLocal(openrestyLocalLuaDir); err != nil{
		return
	}


	openresyCmd := make([]string,0)
	NGCPCmd := fmt.Sprintf("/usr/bin/env docker -H tcp://%s:2375 cp %s %s",confRelation.HostUseIp,cpSrcDir,cpDesDir)
	OpenrestyNGConfCPCmd := fmt.Sprintf("/usr/bin/env docker -H tcp://%s:2375 cp %s%s %s",confRelation.HostUseIp,cpOpenrestyConfDir,OPNGCONF,openrestyLocalConfDir)
	OpenrestyWafConfCPCmd := fmt.Sprintf("/usr/bin/env docker -H tcp://%s:2375 cp %s%s %s",confRelation.HostUseIp,cpOpenrestyConfDir,WAFCONF,openrestyLocalWafDir)
	OpenrestyLuaCPCmd := fmt.Sprintf("/usr/bin/env docker -H tcp://%s:2375 cp %s%s %s",confRelation.HostUseIp,cpOpenrestyLuaDir,ACCRDSLUA,openrestyLocalLuaDir)
	openresyCmd = append(openresyCmd, OpenrestyNGConfCPCmd, OpenrestyWafConfCPCmd, OpenrestyLuaCPCmd)

	s, cmdErr := gopubssh.CommandLocal(NGCPCmd,components.SSHTIMEOUT)
	if cmdErr != nil {
		if cmdErr.Error() == "cmd time out" {
			s.Result = fmt.Sprintf("cmd time out  ::  %s\n", NGCPCmd)
			beego.Error(s.Result)
		}
	}
	if FileExist(cpDesDir + "/nginx"){
		fmt.Println("SUCCESS:  ",NGCPCmd)
	}else {
		fmt.Println("FAILED:  ",NGCPCmd)
	}


	for _, cmd := range openresyCmd {
		// openresty conf  and  lua
		gopubssh.CommandLocal(cmd, components.SSHTIMEOUT)
		if cmdErr != nil {
			if cmdErr.Error() == "cmd time out" {
				s.Result = fmt.Sprintf("cmd time out  ::  %s\n", cmd)
				beego.Error(s.Result)
			}
		}
	}



	if err = ReadConfFile(confRelation,cpDesDir); err != nil{
		return
	}
	return

}

func MakeDirOfLocal(cpDesDir string)(err error){

	if FileExist(cpDesDir){
		// remove
		if err = os.RemoveAll(cpDesDir); err != nil{
			fmt.Printf("error:  MakeDirOfLocal -> RemoveAll is failed  %s  ::  %s\n",cpDesDir,err)
			return err
		}
	}

	// mkdir
	if err = common.Mkdir(cpDesDir); err != nil{
		fmt.Println("error:  MakeDirOfLocal -> Mkdir is failed  ::  ",err)
		return
	}
	return nil
}

func ReadConfFile(confRelation *response.ConfRelation, cpDesDir string)(err error) {

	// read nginx.conf
	if FileExist(cpDesDir + NGDIR + NGCONF) {
		if err = ReadNginxConfInfo(cpDesDir + NGDIR + NGCONF,confRelation);err != nil {
			return
		}
	}

	// read nginx.d/*
	if FileExist(cpDesDir + NGDIR + NGCONF_D) {
		if err = ReadConfD(cpDesDir + NGDIR + NGCONF_D, confRelation); err != nil{
			return
		}
	}

	// read openresty *.conf

	if FileExist(cpDesDir + ORNGDIR) {
		if err = ReadOpenrestyConfs(cpDesDir + ORNGDIR, confRelation); err != nil{
			return
		}
	}



	return

}


func ReadOpenrestyConfs(openrestyNGDir string, confRelation *response.ConfRelation)(err error){

	// nginx.conf
	rdOpenNGPath := fmt.Sprintf("%s%s",openrestyNGDir,OPNGCONF)

	// waf/conf.lua
	rdOpenWafConfPath := fmt.Sprintf("%s%s",openrestyNGDir,WAFCONF)

	// access_by_redis.lua
	rdOpenAccRdsPath := fmt.Sprintf("%s%s",openrestyNGDir,ACCRDSLUA)

	filePaths := make([]string,0)
	filePaths = append(filePaths, rdOpenNGPath, rdOpenWafConfPath, rdOpenAccRdsPath)

	ReadOpenrestyFile(filePaths,confRelation)


	return
}




func ReadOpenrestyFile(filePath []string, confRelation *response.ConfRelation){

	for _, path := range filePath {

		if FileExist(path) {
			conf := new(response.ConfFile)
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return
			}
			conf.ConfInfo = string(buf)
			index := strings.Index(path, ORNGDIR)

			conf.ConfFileName = USRLOCAL + path[index:]
			conf.ConfFullPath = path
			conf.ConfServiceId = confRelation.ServiceId
			conf.HostUseIp = confRelation.HostUseIp
			confRelation.Confs = append(confRelation.Confs, conf)
		}
	}


}



func ReadNginxConfInfo(filePath string, confRelation *response.ConfRelation)(err error){
	conf := new(response.ConfFile)
	buf, err := ioutil.ReadFile(filePath)
	if err != nil{
		return
	}
	conf.ConfInfo = string(buf)
	conf.ConfFileName = CONFDIR + NGCONF
	conf.ConfFullPath = filePath
	conf.ConfServiceId = confRelation.ServiceId
	conf.HostUseIp = confRelation.HostUseIp

	confRelation.Confs = append(confRelation.Confs,conf)
	return
}


func ReadConfD(dir string, confRelation *response.ConfRelation) (err error){


	fileInfoList, _ := ioutil.ReadDir(dir)
	for _, file :=  range fileInfoList {
		conf := new(response.ConfFile)
		name := file.Name()
		buf, err := ioutil.ReadFile(dir + name)
		if err != nil{
			return err
		}
		conf.ConfFileName = CONFDIR + NGCONF_D + name
		conf.ConfFullPath = dir + name
		conf.ConfInfo = string(buf)
		conf.ConfServiceId = confRelation.ServiceId
		conf.HostUseIp = confRelation.HostUseIp

		confRelation.Confs = append(confRelation.Confs,conf)
	}

	return
}



//////////////////////////////

func WriteConfToLocal(confEdit * models.ConfEdit)(err error){

	if FileExist(confEdit.Path){
		// remove
		if err = os.Remove(confEdit.Path); err != nil{
			fmt.Printf("error:  WriteConfToLocal -> Remove file is failed  %s  ::  %s\n", confEdit.Path,err)
			return err
		}
	}

	f, err := os.OpenFile(confEdit.Path, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 436)
	if err != nil{
		return
	}
	defer f.Close()
	n := 0
	buf := []byte(confEdit.Info)
	n, err = f.Write(buf)
	if err != nil {
		if n < len(buf){
			err = io.ErrShortWrite
		}
		return err
	}
	return
}


func DockerCPConfFileToContainer(confEdit *models.ConfEdit) (err error){

	// get service
	id, err := strconv.Atoi(confEdit.ServiceId)
	if err != nil{
		return
	}
	service, err := models.GetServiceById(id)
	if err != nil{
		return
	}
	p, err := models.GetProjectById(service.Project.Id)
	if err != nil{
		return
	}

	docker := RegisterBaseComm(p,service)

	cpDesPath := ""
	if 0 < strings.Index(confEdit.Name,"openresty"){
		cpDesPath = fmt.Sprintf("%s:%s",
			docker.BaseComponents.Docker.Name,
			confEdit.Name,
		)
	} else {
		cpDesPath = fmt.Sprintf("%s:%s",
			docker.BaseComponents.Docker.Name,
			confEdit.Name,
		)
	}




	dockerCPCmd := fmt.Sprintf("/usr/bin/env docker -H tcp://%s:2375 cp %s %s",confEdit.IP,confEdit.Path,cpDesPath)

	s, cmdErr := gopubssh.CommandLocal(dockerCPCmd,components.SSHTIMEOUT)
	if cmdErr != nil {
		if cmdErr.Error() == "cmd time out" {
			s.Result = fmt.Sprintf("cmd time out  ::  %s\n", dockerCPCmd)
			beego.Error(s.Result)
		}
		return cmdErr
	}

	ReloadNginx(confEdit, docker)

	return
}


func ReloadNginx(confEdit *models.ConfEdit, docker *components.BaseDocker)(err error){

	reloadCmd := fmt.Sprintf("/usr/bin/env docker -H tcp://%s exec -i %s nginx -t && /usr/bin/env docker -H tcp://%s exec -i %s nginx -s reload",
		confEdit.IP,
		docker.BaseComponents.Docker.Name,
		confEdit.IP,
		docker.BaseComponents.Docker.Name,
	)


	s, cmdErr := gopubssh.CommandLocal(reloadCmd, components.SSHTIMEOUT)
	if cmdErr != nil {
		if cmdErr.Error() == "cmd time out" {
			s.Result = fmt.Sprintf("cmd time out  ::  %s\n", reloadCmd)
			beego.Error(s.Result)
		}
		return cmdErr
	}else {
		fmt.Println("SUCCESS:  reload nginx is success ")
	}

	return
}
