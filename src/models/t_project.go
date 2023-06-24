package models

import "C"
import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id           	int    `orm:"column(id);pk;auto"`
	Name         	string `orm:"column(name);size(100);unique"`  // 项目名称，创建后不可修改
	Alias        	string `orm:"column(alias);size(100);unique"` // 项目别名
	Nacos1       	string `orm:"column(nacos_1);size(100);null"`
	Nacos1UserName 	string `orm:"column(user_name_nacos_1);size(100);null"`
	Nacos1Pwd	 	string `orm:"column(password_nacos_1);size(100);null"`

	Nacos2       	string `orm:"column(nacos_2);size(100);null"`
	Nacos2UserName 	string `orm:"column(user_name_nacos_2);size(100);null"`
	Nacos2Pwd	 	string `orm:"column(password_nacos_2);size(100);null"`

	Nacos3       	string `orm:"column(nacos_3);size(100);null"`
	Nacos3UserName 	string `orm:"column(user_name_nacos_3);size(100);null"`
	Nacos3Pwd	 	string `orm:"column(password_nacos_3);size(100);null"`

	AwsKeyId     	string `orm:"column(aws_key_id);size(100);null"`
	AwsKeySecret 	string `orm:"column(aws_key_secret);size(100);null"`
	AwsRegion    	string `orm:"column(aws_region);size(100);null"`

	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);auto_now;"`
}

type Nacos struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func init() {
	orm.RegisterModelWithPrefix("t_", new(Project))

	//orm.RegisterModel(new(Project),new(Conf))
}

// // 创建 conf project外键约束
// func (c * SqlClass)ProjectCreateForeignKey(){

// 	// 拼接两张表外键约束sql
// 	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",confTableName,confProjectForeignKeyName,projectId,projectTableName,primaryKey)

// 	// 创建 project project_conf 外键约束
// 	if err := c.CreateForeignKey(addForeignSqlStr); err != nil{
// 		errStr := DefaultFailOnError("project foreign key sql exec",err)
// 		fmt.Println(errStr)
// 		return
// 	}else {
// 		fmt.Println("project tables foreign key relationship created successfully")
// 	}
// }

//////  init project xxx conf key
/*

	以下 InitXXXConfKey 类型方法,
	在project创建时,关联project_id并初始化key
	将固定模板key插入mysql,
	对应value为"",
	设置默认ConfName字段,

*/

/*
	Aws conf
*/
func InitAwsConfKey(projectId int) error {

	// awsConf := make(map[string]string,0)
	// awsConf["AwsKeyId"] = ""
	// awsConf["AwsKeySecret"] = ""
	// awsConf["AwsRegion"] = ""
	var confData []ConfValue
	confData = append(confData, ConfValue{Key: "AwsKeyId", Value: ""})
	confData = append(confData, ConfValue{Key: "AwsKeySecret", Value: ""})
	confData = append(confData, ConfValue{Key: "AwsRegion", Value: ""})

	b, err := json.Marshal(confData)
	if err != nil {
		fmt.Println("error: json.Marshal(awsConf) ", err)
		return err
	}
	// 转 string
	bStr := string(b)

	// init project_conf obj
	proConf := &Conf{
		//  需要修改 ConfName 数据库字段 null
		Name:    "aws",
		Value:   bStr,
		Project: &Project{Id: projectId},
	}

	res, err := AddConf(proConf)
	if err != nil {
		fmt.Println("error: InitAwsConfKey  ", err)
		return err
	}

	fmt.Println("InitAwsConfKey success res: ", res)

	return nil
}

/*
	mysql conf
*/
func InitMysqlConfKey(projectId int) error {

	// mysqlConf := make(map[string]string,0)
	// mysqlConf["spring.datasource.url"] = ""
	// mysqlConf["spring.datasource.username"] = ""
	// mysqlConf["spring.datasource.password"] = ""
	var confData []ConfValue
	confData = append(confData, ConfValue{Key: "spring.datasource.url", Value: ""})
	confData = append(confData, ConfValue{Key: "spring.datasource.username", Value: ""})
	confData = append(confData, ConfValue{Key: "spring.datasource.password", Value: ""})

	b, err := json.Marshal(confData)
	if err != nil {
		fmt.Println("error: json.Marshal(mysqlConf) ", err)
		return err
	}
	// 转 string
	bStr := string(b)

	// init project_conf obj
	proConf := &Conf{
		//  需要修改 ConfName 数据库字段 null
		Name:    "mysql",
		Value:   bStr,
		Project: &Project{Id: projectId},
	}

	res, err := AddConf(proConf)
	if err != nil {
		fmt.Println("error: InitMysqlConfKey  ", err)
		return err
	}

	fmt.Println("InitMysqlConfKey success res: ", res)

	return nil
}

/*
	redis conf
*/
func InitRedisConfKey(projectId int) error {

	// redisConf := make(map[string]string,0)
	// redisConf["ak.wildcard.redis.host"] = ""
	// redisConf["ak.wildcard.redis.port"] = ""
	var confData []ConfValue
	confData = append(confData, ConfValue{Key: "ak.wildcard.redis.host", Value: ""})
	confData = append(confData, ConfValue{Key: "ak.wildcard.redis.port", Value: ""})

	b, err := json.Marshal(confData)
	if err != nil {
		fmt.Println("error: json.Marshal(redisConf) ", err)
		return err
	}

	// 转 string
	bStr := string(b)

	// init project_conf obj
	proConf := &Conf{
		Name:    "redis",
		Value:   bStr,
		Project: &Project{Id: projectId},
	}

	res, err := AddConf(proConf)
	if err != nil {
		fmt.Println("error: InitRedisConfKey  ", err)
		return err
	}

	fmt.Println("InitRedisConfKey success res: ", res)

	return nil

}

/*
	kafka conf
*/
func InitKafkaConfKey(projectId int) error {

	// kafkaConf := make(map[string]string,0)
	// kafkaConf["spring.kafka.bootstrap-servers"] = ""
	var confData []ConfValue
	confData = append(confData, ConfValue{Key: "spring.kafka.bootstrap-servers", Value: ""})

	b, err := json.Marshal(confData)
	if err != nil {
		fmt.Println("error: json.Marshal(kafkaConf) ", err)
		return err
	}
	// 转 string
	bStr := string(b)

	// init project_conf obj
	proConf := &Conf{
		//  需要修改 ConfName 数据库字段 null
		Name:    "kafka",
		Value:   bStr,
		Project: &Project{Id: projectId},
	}

	res, err := AddConf(proConf)
	if err != nil {
		fmt.Println("error: InitKafkaConfKey  ", err)
		return err
	}

	fmt.Println("InitKafkaConfKey success res: ", res)

	return nil
}

/*
	ES conf
*/
func InitESConfKey(projectId int) error {

	// ESConf := make(map[string]string,0)
	// ESConf["elasticsearch.http.host"] = ""
	var confData []ConfValue
	confData = append(confData, ConfValue{Key: "elasticsearch.http.host", Value: ""})

	b, err := json.Marshal(confData)
	if err != nil {
		fmt.Println("error: json.Marshal(ESConf) ", err)
		return err
	}
	// 转 string
	bStr := string(b)

	// init project_conf obj
	proConf := &Conf{
		//  需要修改 ConfName 数据库字段 null
		Name:    "es",
		Value:   bStr,
		Project: &Project{Id: projectId},
	}

	res, err := AddConf(proConf)
	if err != nil {
		fmt.Println("error: InitESConfKey  ", err)
		return err
	}

	fmt.Println("InitESConfKey success res: ", res)

	return nil
}

/*
	Emqtt conf
*/
func InitEmqttConfKey(projectId int) error {

	// emqttConf := make(map[string]string,0)
	// emqttConf["emqtt.broker"] = ""
	// emqttConf["emqtt.username"] = ""
	// emqttConf["emqtt.password"] = ""
	var confData []ConfValue
	confData = append(confData, ConfValue{Key: "emqtt.broker", Value: ""})
	confData = append(confData, ConfValue{Key: "emqtt.username", Value: ""})
	confData = append(confData, ConfValue{Key: "emqtt.password", Value: ""})

	b, err := json.Marshal(confData)
	if err != nil {
		fmt.Println("error: json.Marshal(emqttConf) ", err)
		return err
	}
	// 转 string
	bStr := string(b)

	// init project_conf obj
	proConf := &Conf{
		//  需要修改 ConfName 数据库字段 null
		Name:    "emqtt",
		Value:   bStr,
		Project: &Project{Id: projectId},
	}

	res, err := AddConf(proConf)
	if err != nil {
		fmt.Println("error: InitEmqttConfKey  ", err)
		return err
	}

	fmt.Println("InitEmqttConfKey success res: ", res)

	return nil
}

/*
	Nacos conf
*/
func InitNacosConfKey(projectId int) error {

	// nacosConf := make(map[string]string,0)
	// nacosConf["nacos.config.server-addr"] = ""
	var confData []ConfValue
	confData = append(confData, ConfValue{Key: "nacos.config.server-addr", Value: ""})

	b, err := json.Marshal(confData)
	if err != nil {
		fmt.Println("error: json.Marshal(nacosConf) ", err)
		return err
	}
	// 转 string
	bStr := string(b)

	// init project_conf obj
	proConf := &Conf{
		//  需要修改 ConfName 数据库字段 null
		Name:    "nacos",
		Value:   bStr,
		Project: &Project{Id: projectId},
	}

	res, err := AddConf(proConf)
	if err != nil {
		fmt.Println("error: InitNacosConfKey  ", err)
		return err
	}

	fmt.Println("InitNacosConfKey success res: ", res)

	return nil
}

//////  project
/*
	查询所有 project
*/
func GetAllProject() (p []Project, err error) {

	o := orm.NewOrm()
	if _, err = o.QueryTable(projectTableName).All(&p); err != nil {
		return nil, err
	}

	return p, err
}

/*
	通过 id  查询 project
*/
func GetProjectById(id int) (p *Project, err error) {

	o := orm.NewOrm()
	p = &Project{Id: id}
	if err = o.Read(p); err == nil {
		return p, nil
	}
	return nil, err
}

/*
	通过 id  查询 project
*/
func GetProjectByName(name string) (p *Project, err error) {

	o := orm.NewOrm()
	p = &Project{Name: name}
	if err = o.Read(p,"name"); err == nil {
		return p, nil
	}
	return nil, err
}

/*
	通过 alias 查询 project
*/
func GetProjectByAlias(alias string) (p *Project, err error) {

	o := orm.NewOrm()
	p = &Project{Alias: alias}
	if err = o.Read(p, "Alias"); err == nil {
		return p, nil
	}
	return nil, err
}

func GetProjectAll() (m []*Project, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(projectTableName).All(&m)
	return
}

/*
	增加project
*/
func AddProject(p *Project) (id int64, err error) {

	o := orm.NewOrm()

	id, err = o.Insert(p)
	if err != nil {
		return
	}

	return
}

/*
	修改 project alias
*/
func UpdateProject(p *Project) (num int64, err error) {

	o := orm.NewOrm()

	num, err = o.Update(p)
	return
}

func DeleteProjectById(id int) (bool, error) {

	o := orm.NewOrm()
	v := Project{Id: id}
	// ascertain id exists in the database
	if err := o.Read(&v); err != nil {
		//err = fmt.Errorf("error: project_id is not exists: %s",err)
		return false, err
	}

	if _, err := o.Delete(&Project{Id: id}); err != nil {

		err = fmt.Errorf("error: delete for project_id  %s", err)
		return false, err
	}

	fmt.Println("Number of records deleted in database:")
	return true, nil
}

func GetNacosByProjectId(id int) (nacos []*Nacos, err error) {

	nacos = make([]*Nacos, 0)

	o := orm.NewOrm()
	p := &Project{Id: id}
	if err = o.Read(p); err != nil {
		return nil, err
	}

	n1 := new(Nacos)
	n1.Name = "nacos-1"
	n1.Value = p.Nacos1
	nacos = append(nacos, n1)
	if len(p.Nacos2) > 6 {
		n2 := new(Nacos)
		n2.Name = "nacos-2"
		n2.Value = p.Nacos2
		nacos = append(nacos, n2)
	}
	if len(p.Nacos3) > 6 {
		n3 := new(Nacos)
		n3.Name = "nacos-3"
		n3.Value = p.Nacos3
		nacos = append(nacos, n3)
	}


	return
}
