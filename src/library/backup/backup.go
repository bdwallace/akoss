package backup

import (
	"fmt"
	"io/ioutil"
	"library/common"
	"library/components"
	"strings"

	// "library/myoss"
	"library/aws"
	"models"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


type Mysql struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Table    string
}


func BackupSql() (backPath string, err error) {


	var mysql Mysql
	mysql.Host = beego.AppConfig.String("mysqlhost")
	mysql.Port = beego.AppConfig.String("mysqlport")
	mysql.User = beego.AppConfig.String("mysqluser")
	mysql.Password = beego.AppConfig.String("mysqlpass")
	mysql.Database = beego.AppConfig.String("mysqldb")

	mydate := time.Now().Format("20060102")
	mytime := time.Now().Format("20060102-150405")

	hostName, _ := os.Hostname()
	backPath = fmt.Sprintf("./upload/backup/%s/gopubsql/%s/gopubsql.%s.sql", hostName, mydate, mytime)
	backDir, _ := path.Split(backPath)
	err = common.Mkdir(backDir)
	if err != nil {
		fmt.Println("error: BackupSql 新目录备份目录失败  ", err)
		return
	}
	err = mysql.Backup(backPath)
	if err != nil {
		beego.Error(fmt.Sprintf("mysqldump fail: %s", err))
		return
	}

	// remote := fmt.Sprintf("ops-backup/%s/gopubsql/%s/gopubsql.%s.sql", hostName, mydate, mytime)
	// err = myoss.Upload(remote, backPath)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// 亚马逊云S3备份
	remote := fmt.Sprintf("%s/gopubsql/%s/gopubsql.%s.sql", hostName, mydate, mytime)
	var myAws aws.BaseAws
	myAws.Alias = beego.AppConfig.String("backupAlias")
	myAws.Region = beego.AppConfig.String("backupAwsRegion")
	myAws.InitSession()

	err = myAws.UploadS3(backPath, remote)
	if err != nil {
		beego.Error(fmt.Sprintf("oss updata fail: %s", err))
	}

	err = os.Remove(backPath)
	if err != nil {
		beego.Error(fmt.Sprintf("rm backpath fail: %s", err))
	}

	return
}


func (c *Mysql) Backup(sqlPath string) (err error) {
	var cmd *exec.Cmd
	var cmd_str string

	if c.Table == "" {
		cmd_str = "mysqldump" + " --opt" + " -h" + c.Host + " -P" + c.Port + " -u" + c.User + " -p" + c.Password + " " + c.Database + " >" + sqlPath
	} else {
		cmd_str = "mysqldump" + " --opt" + " -h" + c.Host + " -P" + c.Port + " -u" + c.User + " -p" + c.Password + " " + c.Database + " " + c.Table + " >" + sqlPath
	}
	cmd = exec.Command("/bin/sh", "-c", cmd_str)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("error: Backup  初始化命令管道失败  ", err)
		return err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("error: Backup 执行命令失败  ", err)
		return err
	}

	_, err = ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("error: Backup 获取命令结果失败  ",err)
		return err
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("error: Backup 等待任务执行完  ",err)
		return err
	}

	return
}




func BackupService(serviceClass string) (data []map[string]interface{}, backPath string, err error) {
	data,err = GetService(serviceClass)
	if err != nil{
		return nil,"",err
	}

	f := excelize.NewFile()

	err = f.SetColWidth("Sheet1", "A", "A", 30)
	err = f.SetColWidth("Sheet1", "B", "B", 30)
	err = f.SetColWidth("Sheet1", "C", "B", 30)
	err = f.SetColWidth("Sheet1", "D", "D", 30)
	err = f.SetColWidth("Sheet1", "E", "E", 80)
	err = f.SetColWidth("Sheet1", "F", "F", 80)
	err = f.SetColWidth("Sheet1", "G", "G", 80)
	values := map[string]string{"A1": "项目名", "B1": "第1套tag", "C1": "第2套tag", "D1": "第三套tag", "E1": "第一套详细状态", "F1": "第二套详细状态", "G1": "第三套详细状态"}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}

	cell := ""
	for index, row := range data {
		x := index + 2
		cell, _ = excelize.CoordinatesToCellName(1, x)
		f.SetCellValue("Sheet1", cell, row["name"])

		cell, _ = excelize.CoordinatesToCellName(2, x)
		f.SetCellValue("Sheet1", cell, row["PsImage1"])

		cell, _ = excelize.CoordinatesToCellName(3, x)
		f.SetCellValue("Sheet1", cell, row["PsImage2"])

		cell, _ = excelize.CoordinatesToCellName(4, x)
		f.SetCellValue("Sheet1", cell, row["PsImage3"])


		style, _ := f.NewStyle(`{"alignment":{"wrap_text":true}}`)

		cell, _ = excelize.CoordinatesToCellName(5, x)
		f.SetCellValue("Sheet1", cell, row["ServiceStatus1"])
		f.SetCellStyle("Sheet1", cell, cell, style)

		cell, _ = excelize.CoordinatesToCellName(6, x)
		f.SetCellValue("Sheet1", cell, row["ServiceStatus2"])
		f.SetCellStyle("Sheet1", cell, cell, style)

		cell, _ = excelize.CoordinatesToCellName(7, x)
		f.SetCellValue("Sheet1", cell, row["ServiceStatus3"])
		f.SetCellStyle("Sheet1", cell, cell, style)


	}

	mydate := time.Now().Format("20060102")
	mytime := time.Now().Format("20060102-150405")

	hostName, _ := os.Hostname()
	backPath = fmt.Sprintf("./upload/backup/%s/service/%s/%s/service.%s.xlsx", hostName, serviceClass, mydate, mytime)
	backDir, _ := path.Split(backPath)
	err = common.Mkdir(backDir)
	if err != nil {
		fmt.Println("error: BackupService 新目录备份目录失败  ", err)
		return
	}
	err = f.SaveAs(backPath)
	if err != nil {
		beego.Error(fmt.Sprintf("save %s fail: %s", backPath, err))
	}

	// 阿里云oss备份
	// remote := fmt.Sprintf("ops-backup/%s/service/%s/%s/service.%s.xlsx", hostName, serviceClass, mydate, mytime)
	// err = myoss.Upload(remote, backPath)
	// if err != nil {
	// 	beego.Error(fmt.Sprintf("oss updata fail: %s", err))
	// }

	// 亚马逊云S3备份
	remote := fmt.Sprintf("%s/tag/%s/%s/service.%s.xlsx", hostName, serviceClass, mydate, mytime)
	var myAws aws.BaseAws
	myAws.Alias = beego.AppConfig.String("backupAlias")
	myAws.Region = beego.AppConfig.String("backupAwsRegion")
	myAws.InitSession()

	err = myAws.UploadS3(backPath, remote)
	if err != nil {
		beego.Error(fmt.Sprintf("s3 updata fail: %s", err))
	}

	err = os.Remove(backPath)
	if err != nil {
		beego.Error(fmt.Sprintf("rm backpath fail: %s", err))
	}

	return
}

type ServiceRelatedProject struct{
	ServiceName string
	ServiceTag string
	ServiceRelatedProjectName string
}



func GetServiceTag(serviceClass string) (serviceObj [][]ServiceRelatedProject,  err error) {

	serviceNameList := new(orm.ParamsList)
	if serviceClass	== "java" {
		serviceNameList, err = models.GetServiceClassJava()
	}else {
		serviceNameList, err = models.GetServiceClassOther()
	}
	if err != nil{
		return nil,err
	}

	wildcardServiceName := make(map[string]int, 0)
	for _, serviceName := range *serviceNameList {
		strServiceName := serviceName.(string)
		if strings.Index(strServiceName,"_") > 0{
			serviceName = strings.Replace(strServiceName,"_","-",-1)
		}
		wildcardServiceName[serviceName.(string)] += 1
	}

	for serviceName, num := range wildcardServiceName {

		services := make([]*models.Service, 0)

		services, _ = models.GetServiceByServiceName(serviceName)
		oldServiceName := ""
		if num == 2 && strings.Index(serviceName,"-") > 0{
			oldServiceName = strings.Replace(serviceName,"-","_",-1)
		}
		services2, _ := models.GetServiceByServiceName(oldServiceName)
		services = append(services, services2...)

		sameNameOfServices := make([]ServiceRelatedProject,0)

		for _, s := range services {
			serviceRelatedProjectOjb := ServiceRelatedProject {
				ServiceName:s.Name,
				ServiceTag:s.Tag,
				ServiceRelatedProjectName:s.Project.Name,
			}

			if serviceName == "" {
				fmt.Println("error:  serviceName == nil")
			}

			sameNameOfServices = append(sameNameOfServices,serviceRelatedProjectOjb)
		}
		serviceObj = append(serviceObj, sameNameOfServices)
	}

	return
}



//
//func GetServiceTag(serviceClass string) (serviceObj [][]ServiceRelatedProject,  err error) {
//
//	serviceNameList := new(orm.ParamsList)
//	if serviceClass	== "java" {
//		serviceNameList, err = models.GetServiceClassJava()
//	}else {
//		serviceNameList, err = models.GetServiceClassOther()
//	}
//	if err != nil{
//		return nil,err
//	}
//
//	// 获取所有service,  根据 serviceName 获取service
//	//serviceObj = make(map[string][]ServiceRelatedProject)
//	//serviceObj = make(map[string][]ServiceRelatedProject)
//
//	wildcardServiceName := make(map[string]int, 0)
//	for _, serviceName := range *serviceNameList {
//		strServiceName := serviceName.(string)
//		serviceName = strings.Replace(strServiceName,"_","-",-1)
//		wildcardServiceName[serviceName.(string)] += 1
//	}
//
//	//for _, serviceName := range *serviceNameList	{
//	for serviceName, num := range wildcardServiceName {
//
//		services := make([]*models.Service, 0)
//
//	/*
//		if value, ok := serviceName.(string); ok {
//			services, _ = models.GetServiceByServiceName(value)
//		}
//	*/
//		if num == 2 {
//			services, _ = models.GetServiceByServiceName(serviceName)
//			oldServiceName := strings.Replace(serviceName,"-","_",-1)
//			services2, _ := models.GetServiceByServiceName(oldServiceName)
//			services = append(services, services2...)
//		}
//
//		sameNameOfServices := make([]ServiceRelatedProject,0)
//		serviceNum := len(services)
//		for i, s := range services {
//
//			serviceRelatedProjectOjb := ServiceRelatedProject {
//				ServiceName:s.Name,
//				ServiceTag:s.Tag,
//				ServiceRelatedProjectName:s.Project.Name,
//			}
//
//			// 第一套补位
//			// 只有一个service,并且为第二套, 补位第一套
//			if serviceNum == 1 && i == 0 && services[0].Project.Name == "第二套" {
//
//				tempServiceRelatedProjectOjb := ServiceRelatedProject {
//					ServiceName: s.Name,
//					ServiceTag: "---",
//					ServiceRelatedProjectName: "第二套",
//				}
//
//				// 补位第一套
//				sameNameOfServices = append(sameNameOfServices,tempServiceRelatedProjectOjb)
//			}
//
///*
//			// 只有一个service,并且为第三套, 补位第一套
//			if serviceNum == 1 && i == 0 && s.Project.Name == "第三套" {
//
//			}
//			// 有两个service, 为第二套,第三套, 补位第一套
//			if serviceNum == 2 && i == 0 && services[0].Project.Name == "第二套"{
//
//				if  serviceNum == 2 && i == 1 && services[1].Project.Name == "第三套"{
//
//				}
//			}
//
//			// 第二套补位
//			// 有一个service,为第一套,无需补位
//			// 有一个service,为第三套, 补位第一套,第二套,
//			// 有两个service,为第一套,第三套, 补位第二套
//*/
//			sameNameOfServices = append(sameNameOfServices,serviceRelatedProjectOjb)
//
//		}
//		serviceObj = append(serviceObj, sameNameOfServices)
//	}
//
//	return
//}


func GetService(serviceClass string) (data []map[string]interface{},err error) {

	serviceNameList := new(orm.ParamsList)
	if serviceClass	== "java" {
		serviceNameList, err = models.GetServiceClassJava()
	}else {
		serviceNameList, err = models.GetServiceClassOther()
	}
	if err != nil{
		return nil,err
	}

	for _, serviceName := range *serviceNameList {
		services := make([]*models.Service,0)
		if value, ok := serviceName.(string); ok {
			services, _ = models.GetServiceByServiceName(value)
		}

		var dataIndex map[string]interface{}
		dataIndex = make(map[string]interface{})
		dataIndex["name"] = serviceName
		dataIndex["PsImage1"] = ""
		dataIndex["ServiceStatus1"] = ""
		dataIndex["PsImage2"] = ""
		dataIndex["ServiceStatus2"] = ""
		dataIndex["PsImage3"] = ""
		dataIndex["ServiceStatus3"] = ""

		for _, service := range services {
			service, err := models.GetServiceAllRelated(service)
			if err != nil{
				return nil,err
			}

			if len(service.Hosts) == 0 {
				continue
			}

			s := components.BaseComponents{}
			s.SetProject(service.Project)
			s.SetService(service)
			s.SetHost(service.Hosts)

			d := components.BaseDocker{}
			d.SetBaseComponents(s)

			res, _ := d.DockerPs("")

			PsImage := fmt.Sprintf("PsImage%d", service.Project.Id)
			ServiceStatus := fmt.Sprintf("ServiceStatus%d", service.Project.Id)
			var PsImageList []string
			var ServiceStatusList []string
			for _, re := range res {
				if !common.InList(strings.TrimSpace(re["ps_image"]), PsImageList) {
					PsImageList = append(PsImageList, strings.TrimSpace(re["ps_image"]))
				}
				ServiceStatusList = append(ServiceStatusList, fmt.Sprintf("%20s:%30s:%20s:%20s", re["ip"], strings.TrimSpace(re["ps_image"]), re["ps_status"], re["ps_created_at"]))
			}
			dataIndex[PsImage] = strings.Join(PsImageList, ",")
			dataIndex[ServiceStatus] = strings.Join(ServiceStatusList, "\n")
		}
		data = append(data, dataIndex)
	}

	return
}

