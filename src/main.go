package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"library/aws"
	"library/crontab"
	"models"
	"os"
	"os/signal"
	_ "routers"
	"syscall"
)

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
		if v == "-docker" {
			beego.BConfig.RunMode = "docker"
			models.Syncdb()
		}
	}
}



func init() {
	//初始化数据库
	initArgs()
	//beego.Info("开始启动")
	//连接MySQL
	dbUser := beego.AppConfig.String("mysqluser")
	dbPass := beego.AppConfig.String("mysqlpass")
	dbHost := beego.AppConfig.String("mysqlhost")
	dbPort := beego.AppConfig.String("mysqlport")
	dbName := beego.AppConfig.String("mysqldb")
	if beego.BConfig.RunMode == "docker" {
		if os.Getenv("MYSQL_USER") != "" {
			dbUser = os.Getenv("MYSQL_USER")
		}
		if os.Getenv("MYSQL_PASS") != "" {
			dbPass = os.Getenv("MYSQL_PASS")
		}
		if os.Getenv("MYSQL_HOST") != "" {
			dbHost = os.Getenv("MYSQL_HOST")
		}
		if os.Getenv("MYSQL_PORT") != "" {
			dbPort = os.Getenv("MYSQL_PORT")
		}
		if os.Getenv("MYSQL_DB") != "" {
			dbName = os.Getenv("MYSQL_DB")
		}
		if os.Getenv("JenkinsUserName") != "" {
			beego.AppConfig.Set("JenkinsUserName", os.Getenv("JenkinsUserName"))
		}
		if os.Getenv("JenkinsPwd") != "" {
			beego.AppConfig.Set("JenkinsPwd", os.Getenv("JenkinsPwd"))
		}
	}

	maxIdleConn, _ := beego.AppConfig.Int("mysql_max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("mysql_max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FShanghai"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)
/*
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
*/
	//设置日志
	fn := "logs/run.log"
	if _, err := os.Stat(fn); err != nil {
		if os.IsNotExist(err) {
			os.Create(fn)
		}
	}
	beego.SetLogger("file", `{"filename":"`+fn+`"}`)
	if beego.BConfig.RunMode == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}


	orm.RunSyncdb("default",false,true)


	// 创建project platform service等表关联关系
	models.CreateAkTables()
}

func handleSignals(c chan os.Signal) {
	switch <-c {
	case syscall.SIGINT, syscall.SIGTERM:

		beego.Info("Shutdown quickly, bye...")
	case syscall.SIGQUIT:
		beego.Info("Shutdown gracefully, bye...")
		// do graceful shutdown
	}
	os.Exit(0)
}

func main() {
	//获取全局panic
	defer func() {
		if err := recover(); err != nil {
			beego.Error("Panic error:", err)
		}
	}()

	//热启动
	graceful, _ := beego.AppConfig.Bool("Graceful")
	if graceful {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go handleSignals(sigs)
	}
	//API自动化文档
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	}

	//beego.Info(beego.BConfig.RunMode)

	// 初始化 ~/.aws/credentials 文件
	aws.InitCredentials()

	// 初始化 crontab
	err := crontab.Crontab()
	if err != nil {
		beego.Error("Panic error:", err)
	}

	beego.SetStaticPath("/upload", "upload")


	/*
		if beego.BConfig.RunMode != "docker" {
			init_sever.Start()
		}
	*/
	beego.Run()
}
