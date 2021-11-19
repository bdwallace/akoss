package models

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SqlClass struct {
	DbUser      string
	DbPass      string
	DbHost      string
	DbPort      string
	DbName      string
	MaxIdleConn int
	MaxOpenConn int
	DbLink      string
	DB          *sql.DB
}

func DefaultFailOnError(errLocation string, err error) string {
	errInfo := fmt.Sprintf("error:  %s: %s", errLocation, err)
	return errInfo
}

const (
	errDuplicate = "Error 1022: Can't write; duplicate key in table"

	primaryKey = "id"

	// project
	projectTableName = "t_project"
	projectId        = "project_id"

	// conf
	confTableName             = "t_conf"
	confProjectForeignKeyName = "conf_project_id"

	// service
	serviceTableName              = "t_service"
	serviceProjectForeignKeyName  = "service_project_id"
	servicePlatformForeignKeyName = "service_platform_id"
	serviceHostForeignKeyName     = "service_host_id"
	serviceId                     = "service_id"

	// host
	hostTableName             = "t_host"
	hostId                    = "host_id"
	hostProjectForeignKeyName = "host_project_id"

	// t_service_t_conf
	serviceConfTableName                = "t_service_t_confs"
	serviceConfForConfForeignKeyName    = "service_conf_for_conf_id"
	serviceConfForConfId                = "t_conf_id"
	serviceConfForServiceForeignKeyName = "service_conf_for_service_id"
	serviceConfForServiceId             = "t_service_id"

	// t_service_t_host
	serviceHostTableName                = "t_service_t_hosts"
	serviceHostForHostForeignKeyName    = "service_host_for_host_id"
	serviceHostForHostId                = "t_host_id"
	serviceHostForServiceForeignKeyName = "service_host_for_service_id"
	serviceHostForServiceId             = "t_service_id"

	// t_deploy
	deployTableName              = "t_deploy"
	deployId                     = "deploy_id"
	deployUserForeignKeyName     = "deploy_user_id"
	deployProjectForeignKeyName  = "deploy_project_id"
	deployPlatformForeignKeyName = "deploy_platform_id"

	// t_task
	taskTableName             = "t_task"
	taskId                    = "task_id"
	taskProjectForeignKeyName = "task_project_id"
	taskServiceForeignKeyName = "task_service_id"
	taskDeployForeignKeyName  = "task_deploy_id"
	taskUserForeignKeyName    = "task_user_id"

	// t_task_t_host
	taskHostTableName             = "t_task_t_hosts"
	taskHostForHostForeignKeyName = "task_host_for_host_id"
	taskHostForHostId             = "t_host_id"
	taskHostForTaskForeignKeyName = "task_host_for_task_id"
	taskHostForTaskId             = "t_task_id"

	// t_record
	recordTableName            = "t_record"
	recordTaskForeignKeyName   = "record_task_id"
	recordUserForeignKeyName   = "record_user_id"
	recordDeployForeignKeyName = "record_deploy_id"

	// t_platform
	platformTableName             = "t_platform"
	platformId                    = "platform_id"
	platformProjectForeignKeyName = "platform_project_id"

	// t_platform_t_services
	platformServicesTableName                 = "t_service_t_platforms"
	platformServicesForPlatformForeignKeyName = "service_platform_for_platform_id"
	platformServicesForPlatformId             = "t_platform_id"
	platformServicesForServiceForeignKeyName  = "service_platform_for_Service_id"
	platformServicesForServiceId              = "t_service_id"

	// t_domain
	domainTableName             = "t_domain"
	domainId                    = "domain_id"
	domainProjectForeignKeyName = "domain_project_id"

	// t_service_t_domain
	serviceDomainTableName                = "t_service_t_domains"
	serviceDomainForDomainForeignKeyName  = "service_domain_for_domain_id"
	serviceDomainForDomainId              = "t_domain_id"
	serviceDomainForServiceForeignKeyName = "service_domain_for_service_id"
	serviceDomainForServiceId             = "t_service_id"

	// t_domain_t_platform
	domainPlatformTableName                 = "t_domain_t_platforms"
	domainPlatformForDomainForeignKeyName   = "domain_platform_for_domain_id"
	domainPlatformForDomainId               = "t_domain_id"
	domainPlatformForPlatformForeignKeyName = "domain_platform_for_platform_id"
	domainPlatformForPlatformId             = "t_platform_id"

	// t_operation_record
	operationRecordTableName                = "t_operation_record"
	operationRecordForUserForeignKeyName    = "operation_record_for_user_id"
	operationRecordForServiceForeignKeyName = "operation_record_for_service_id"

	// t_crontab
	crontabTableName                = "t_crontab"
	crontabId                       = "crontab_id"
	crontabForProjectForeignKeyName = "crontab_for_project_id"

	// t_crontab_log
	crontabLogTableName                = "t_crontab_log"
	crontabLogId                       = "crontab_log_id"
	crontabLogForCrontabForeignKeyName = "crontab_log_for_crontab_id"

	// t_domain_delete_records
	domainDeleteRecordsTableName                = "t_delete_domain_reco"
	domainDeleteRecordsFroProjectForeignKeyName = "delete_domian_reco_for_project_id"

	// t_domain_backup
	DomainBackupTableName = "t_domain_backup"

	//confTableName,confProjectForeignKeyName,projectId,projectTableName,primaryKey

	// t_link t_link和t_project使用新方法自动添加外键
	linkTableName = "t_link"
	// linkId = "link_id"

	// t_link_t_project
	// linkProjectTableName = "t_link_t_project"
	// linkProjectForLinkForeignKeyName = "link_project_for_link_id"
	// linkProjectForLinkId = "t_link_id"
	// linkProjectForProjectForeignKeyName = "link_project_for_project_id"
	// linkProjectForProjectId = "t_project_idw

	// t_cloud 各种云帐号信息
	cloudTableName = "t_cloud"

	// user
	userTableName = "user"
	userId        = "user_id"

	//inspect
	inspectGrafanaTableName = "t_inspect_grafana"
)

var DomainBackupTableNameWithAccessKeyAndDate string

func GetDBConnInfo() *SqlClass {

	//连接MySQL
	sqlObj := new(SqlClass)

	sqlObj.DbUser = beego.AppConfig.String("mysqluser")
	sqlObj.DbPass = beego.AppConfig.String("mysqlpass")
	sqlObj.DbHost = beego.AppConfig.String("mysqlhost")
	sqlObj.DbPort = beego.AppConfig.String("mysqlport")
	sqlObj.DbName = beego.AppConfig.String("mysqldb")
	if beego.BConfig.RunMode == "docker" {
		if os.Getenv("MYSQL_USER") != "" {
			sqlObj.DbUser = os.Getenv("MYSQL_USER")
		}
		if os.Getenv("MYSQL_PASS") != "" {
			sqlObj.DbPass = os.Getenv("MYSQL_PASS")
		}
		if os.Getenv("MYSQL_HOST") != "" {
			sqlObj.DbHost = os.Getenv("MYSQL_HOST")
		}
		if os.Getenv("MYSQL_PORT") != "" {
			sqlObj.DbPort = os.Getenv("MYSQL_PORT")
		}
		if os.Getenv("MYSQL_DB") != "" {
			sqlObj.DbName = os.Getenv("MYSQL_DB")
		}
		if os.Getenv("JenkinsUserName") != "" {
			beego.AppConfig.Set("JenkinsUserName", os.Getenv("JenkinsUserName"))
		}
		if os.Getenv("JenkinsPwd") != "" {
			beego.AppConfig.Set("JenkinsPwd", os.Getenv("JenkinsPwd"))
		}
	}

	sqlObj.MaxIdleConn, _ = beego.AppConfig.Int("mysql_max_idle_conn")
	sqlObj.MaxOpenConn, _ = beego.AppConfig.Int("mysql_max_open_conn")
	sqlObj.DbLink = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", sqlObj.DbUser, sqlObj.DbPass, sqlObj.DbHost, sqlObj.DbPort, sqlObj.DbName) + "&loc=Asia%2FShanghai"
	//utils.Display("dbLink", dbLink)

	fmt.Println(sqlObj.DbLink)
	return sqlObj
}

func (c *SqlClass) ConncetDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", c.DbLink)
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*
	检测 表中 指定外键字段是否存在
*/
/*
func (c * SqlClass)CheckForeignKeyExists(foreignKeyName string)bool{

	checkForeignKeyExists := fmt.Sprintf("SELECT * FROM `information_schema`.`KEY_COLUMN_USAGE` where constraint_name='%s'",foreignKeyName)
	if _, err := c.DB.Exec(checkForeignKeyExists); err == nil{
		return true
	}

	return false
}
*/

/*
	创建外键约束
*/
func (c *SqlClass) CreateForeignKey(addForeignSqlStr string) error {
	// 创建表 外键关联
	if _, err := c.DB.Exec(addForeignSqlStr); err != nil {
		DefaultFailOnError(" foreign key sql exec", err)
		return err
	}

	return nil
}

/*
	创建外键约束,新形式，无需在models里逐个定义增加外键的方法
*/
func (c *SqlClass) CreateForeignKeyNew(foreignType string, tbMain string, tbForeign string) error {
	// 创建表 外键关联
	if foreignType == "m2m" {
		//定义多对多关联表，如： `t_link`和`t_project`,多对多表则为`t_link_t_projects`。
		tableName := fmt.Sprintf("%s_%ss", tbMain, tbForeign)
		//定义主键和副键的外键名
		keyNameForMain := fmt.Sprintf("%s_%s_for_%s", tbMain, tbForeign, tbMain)
		keyNameForForeign := fmt.Sprintf("%s_%s_for_%s", tbMain, tbForeign, tbForeign)
		//定义主键和副键的字段名，统一为`${tablename}_id`
		keyIdForMain := fmt.Sprintf("%s_id", tbMain)
		keyIdForForeign := fmt.Sprintf("%s_id", tbForeign)

		sqlMain := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(id);", tableName, keyNameForMain, keyIdForMain, tbMain)
		c.ExecForeignSql(sqlMain)

		sqlForeign := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(id);", tableName, keyNameForForeign, keyIdForForeign, tbForeign)
		c.ExecForeignSql(sqlForeign)
	} else {
		keyName := fmt.Sprintf("%s_for_%s", tbMain, tbForeign)
		keyId := fmt.Sprintf("%s_id", tbForeign)
		sql := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(id);", tbMain, keyName, keyId, tbForeign)
		c.ExecForeignSql(sql)
	}

	return nil
}

/*
	创建外键约束的执行SQL
*/
func (c *SqlClass) ExecForeignSql(addForeignSqlStr string) {
	// 创建表 外键关联
	_, err := c.DB.Exec(addForeignSqlStr)
	if err != nil {
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}

func CreateAkTables() {

	// 链接 mysql
	var err error
	sqlObj := GetDBConnInfo()
	if sqlObj.DB, err = sqlObj.ConncetDB(); err != nil {
		errStr := DefaultFailOnError("err: connectDB ", err)
		fmt.Println(errStr)
		return
	}
	defer sqlObj.DB.Close()

	// t_service 表
	sqlObj.ServiceCreateForeignKeyProject()
	//sqlObj.ServiceCreateForeignKeyPlantform()

	// t_conf 表
	sqlObj.ConfCreateForeignKeyProject()

	// t_host 表
	sqlObj.HostCreateForeignKeyProject()

	// t_service_t_host
	sqlObj.ServiceHostCreateForeignKeyForHost()
	sqlObj.ServiceHostCreateForeignKeyForService()

	// t_service_t_conf
	sqlObj.ServiceConfCreateForeignKeyForConf()
	sqlObj.ServiceConfCreateForeignKeyForService()

	// t_service_t_domain
	sqlObj.ServiceDomainCreateForeignKeyForDomain()
	sqlObj.ServiceDomainCreateForeignKeyForService()

	// t_deploy
	sqlObj.DeployCreateForeignKeyProject()
	sqlObj.DeployCreateForeignKeyUser()
	// sqlObj.DeployCreateForeignKeyPlatform()

	// t_task
	sqlObj.TaskCreateForeignKeyProject()
	sqlObj.TaskCreateForeignKeyUser()
	sqlObj.TaskCreateForeignKeyDeploy()
	sqlObj.TaskCreateForeignKeyService()
	// sqlObj.TaskCreateForeignKeyPlatform()

	// t_task_t_host
	sqlObj.TaskHostCreateForeignKeyForHost()
	sqlObj.TaskHostCreateForeignKeyForTask()

	// t_record
	sqlObj.RecordCreateForeignKeyUser()
	sqlObj.RecordCreateForeignKeyTask()
	sqlObj.RecordCreateForeignKeyDeploy()

	// t_domain
	sqlObj.DomainCreateForeignKeyProject()
	// sqlObj.DomainCreateForeignKeyPlatform()

	// t_domain_platform
	sqlObj.DomainPlatformCreateForeignKeyForDomain()
	sqlObj.DomainPlatformCreateForeignKeyForPlatform()

	// t_platform
	sqlObj.PlatformCreateForeignKeyProject()

	// t_platform_service
	sqlObj.PlatformServiceCreateForeignKeyForPlatform()
	sqlObj.PlatformServiceCreateForeignKeyForService()

	// t_operation_record
	sqlObj.OperationRecordCreateForeignKeyService()
	sqlObj.OperationRecordCreateForeignKeyUser()

	// t_crontab
	sqlObj.CrontabCreateForeignKeyProject()

	// t_link_t_project
	// sqlObj.LinkProjectCreateForeignKeyForLink()
	// sqlObj.LinkProjectCreateForeignKeyForProject()
	// 新t_link 使用新的方式添加外键
	sqlObj.CreateForeignKeyNew("m2m", "t_link", "t_project")

	// t_crontab_log
	sqlObj.CrontabLogCreateForeignKeyCrontab()

	// t_deleteDomain_reco
	sqlObj.DeleteDomainRecordsCreateForeignKeyProject()
	/*
		// t_host_conf
		sqlObj.HostConfCreateForeignKeyProject()
		sqlObj.HostConfCreateForeignKeyPlatform()
		sqlObj.HostConfCreateForeignKeyService()
		sqlObj.HostConfCreateForeignKeyHost()
	*/

}

// 解析conf key -> value
func AnalyzeConf(confValue string) map[string]string {
	/*
		m := make(map[string]string)
		var key []string
		if -1 != strings.Index(confValue,","){
			key = strings.Split(confValue,",")
		}
		for _,tempV := range key {
			v1 := strings.TrimLeft(tempV,"{")
			v := strings.TrimRight(v1,"}")

			index := strings.Index(v,":")

			key := v[:index]
			value := v[index+1:]

			m[key] = value
		}
	*/
	v1 := strings.TrimLeft(confValue, "{")
	v2 := strings.TrimRight(v1, "}")

	fmt.Println("v2 : ", v2)
	m := make(map[string]string)
	var key []string
	if -1 != strings.Index(confValue, ",") {
		key = strings.Split(confValue, ",")
	} else {
		key = append(key, v2)
	}

	for _, v := range key {

		index := strings.Index(v, ":")

		key := v[:index]
		value := v[index+1:]

		m[key] = value
	}

	return m
}

func OrmBegin() (o orm.Ormer, err error) {
	o = orm.NewOrm()
	err = o.Begin()
	return
}

func OrmCommit(o orm.Ormer) (err error) {
	err = o.Commit()
	return
}

func OrmRollback(o orm.Ormer) (err error) {
	err = o.Rollback()
	return
}
