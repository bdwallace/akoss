package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)



type HostConf struct {
	Id        		int    				`orm:"column(id);pk;auto"`
	ConfName 		string				`orm:"column(conf_name);size(300)"`
	Class			string				`orm:"column(class);size(100)"`
	LocalPath 		string 				`orm:"column(local_path);size(500)"`
	Deleted 		int					`orm:"column*(deleted);default(0)"`
	Project 		*Project			`orm:"rel(fk)"`
	Platform		*Platform			`orm:"rel(fk)"`
	Service			*Service			`orm:"rel(fk)"`
	Host			*Host				`orm:"rel(fk)"`
}





func init(){
	orm.RegisterModelWithPrefix("t_",new(HostConf))
}



//var NGINXLOCALPATH = ""

/*
	创建 hostConf 中 project_id
*/
func (c * SqlClass)HostConfCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",hostConfTableName, hostConfProjectForeignKeyName,projectId,projectTableName,primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil{
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}

/*
	创建 hostConf 中 platform_id
*/
func (c * SqlClass)HostConfCreateForeignKeyPlatform() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",hostConfTableName, hostConfPlatformForeignKeyName,platformId,platformTableName,primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil{
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}


/*
	创建 hostConf 中 service_id
*/
func (c * SqlClass)HostConfCreateForeignKeyService() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",hostConfTableName, hostConfServiceForeignKeyName, serviceId, serviceTableName, primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil{
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}


/*
	创建 hostConf 中 host_id
*/
func (c * SqlClass)HostConfCreateForeignKeyHost() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",hostConfTableName, hostConfHostForeignKeyName, hostId, hostTableName, primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
	if err != nil{
		if strings.Index(err.Error(), errDuplicate) != -1 {
			beego.Info("key is duplicate")
		} else {
			beego.Error(addForeignSqlStr, " : ", err.Error())
		}
	} else {
		beego.Info("key is create")
	}
}



/*
	通过 hostConf 查询关联表数据
*/
func GetHostConfAllRelated(hostConf *HostConf) (*HostConf, error){
	o := orm.NewOrm()

	if _, err := o.LoadRelated(hostConf,"Project"); err != nil{
		return nil, err
	}

	if _, err := o.LoadRelated(hostConf,"Platform"); err != nil{
		return nil, err
	}

	if _, err := o.LoadRelated(hostConf,"Service"); err != nil{
		return nil, err
	}

	if _, err := o.LoadRelated(hostConf,"Host"); err != nil{
		return nil, err
	}

	return hostConf, nil
}

func AddHostConf (hostConf *HostConf) (err error) {

	o := orm.NewOrm()

	_, err = o.Insert(hostConf)
	if err != nil{
		return
	}

	return
}

func SearchHostConfOfPlatformClassHost(platformId int, serviceClass string, hostId int) (hostConf []*HostConf, err error){

	o := orm.NewOrm()
	_, err = o.QueryTable(hostConfTableName).Filter("platform_id",platformId).Filter("class",serviceClass).Filter("host_id",hostId).Filter("deleted",0).RelatedSel().All(&hostConf)
	return
}

func SearchHostConfOfPlatformClass(platformId int, serviceClass string) (hostConf []*HostConf, err error){

	o := orm.NewOrm()
	_, err = o.QueryTable(hostConfTableName).Filter("platform_id",platformId).Filter("class",serviceClass).Filter("deleted",0).RelatedSel().All(&hostConf)
	return
}

func SearchHostConfOfPlatformHost(platformId int,  hostId int) (hostConf []*HostConf, err error){

	o := orm.NewOrm()
	_, err = o.QueryTable(hostConfTableName).Filter("platform_id",platformId).Filter("host_id",hostId).Filter("deleted",0).RelatedSel().All(&hostConf)
	return
}

func SearchHostConfOfPlatform(platformId int) (hostConf []*HostConf, err error){

	o := orm.NewOrm()
	_, err = o.QueryTable(hostConfTableName).Filter("platform_id",platformId).Filter("deleted",0).RelatedSel().All(&hostConf)
	return
}


func GetHostConfById(id int) (hostConf *HostConf, err error){

	o := orm.NewOrm()
	hostConf = &HostConf{}
	err = o.QueryTable(hostConfTableName).Filter("id",id).Filter("deleted",0).RelatedSel().One(hostConf)
	return
}


func GetHostConfByName(name string)(hostConf *HostConf, err error){
	o := orm.NewOrm()
	hostConf = &HostConf{}
	err = o.QueryTable(hostConfTableName).Filter("name",name).Filter("deleted",0).RelatedSel().One(hostConf)
	return
}

func GetHostConfByLocalPath(path string)(hostConf *HostConf, err error){
	o := orm.NewOrm()
	hostConf = &HostConf{}
	err = o.QueryTable(hostConfTableName).Filter("local_path",path).Filter("deleted",0).RelatedSel().One(hostConf)
	return
}

func DeleteHostConfByLocalPath(path string)(int,error){

	o := orm.NewOrm()
	hostConf := &HostConf{LocalPath:path}

	if err := o.QueryTable(hostConfTableName).Filter("local_path",path).Filter("deleted",0).RelatedSel().One(hostConf); err != nil{
		return -1, err
	}

	hostConf.Deleted = 1
	id, err := o.Update(hostConf)
	if err != nil {
		return -1, err
	}
	return int(id), nil
}
