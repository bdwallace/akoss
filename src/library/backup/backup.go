package backup

import (
	"fmt"
	"io/ioutil"
	"library/common"
	"library/components"

	// "library/myoss"
	"library/aws"
	"models"
	"os"
	"os/exec"
	"path"
	"strings"
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
		fmt.Println("--新目录备份目录失败---:", err)
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
	fmt.Println("---cmd---:", cmd_str)
	cmd = exec.Command("/bin/sh", "-c", cmd_str)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("--1.初始化命令管道失败--")
		return err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("--2.执行命令失败--")
		return err
	}

	_, err = ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("--3.获取命令结果失败--")
		return err
	}

	if err := cmd.Wait(); err != nil {
 		fmt.Println("--4.等待任务执行完--")
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
		fmt.Println("--新目录备份目录失败---:", err)
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

func GetService(serviceClass string) (data []map[string]interface{},err error) {

	serviceNameList := new(orm.ParamsList)
	if serviceClass	== "java" {
		serviceNameList, err = models.GetServiceClassJava()
	}else {
		serviceNameList, err = models.GetServiceClassOther()
	}
	//
	fmt.Println("==============")
	fmt.Println("serviceNameList: ",serviceNameList)


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
			fmt.Printf("service: %+v\n",service)
			for _, h := range service.Hosts{
				fmt.Printf("host: %+v\n",h)
			}

			if len(service.Hosts) == 0 {
				continue
			}

			fmt.Printf("service.project: %+v\n",service.Project)
			s := components.BaseComponents{}
			s.SetProject(service.Project)
			s.SetService(service)
			s.SetHost(service.Hosts)

			d := components.BaseDocker{}
			d.SetBaseComponents(s)

			fmt.Println("s.project: ",s.Project)
			fmt.Println("d.project: ",d.BaseComponents.Project.Alias)

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

