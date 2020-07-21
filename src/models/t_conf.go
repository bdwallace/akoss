package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


type Conf struct {

	Id        	int	                `orm:"column(id);pk;auto"`
	Name  		string            	`orm:"column(name);size(100)"`					      // 配置名称
	Value 		string     			`orm:"column(value);type(text)"`                      // 配置项， 以 k -> v 形式储存json

	CreatedAt              time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt              time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	Project  	*Project 			`orm:"rel(fk)"` // oneToMany relation  project 主表
	Services 	[]*Service			`orm:"reverse(many)"`

}


type ConfValue struct {
	Key 	string 		`json:"Key"`
	Value 	string		`json:"Value"`
}


func (c *Conf) TableUnique() [][]string {
	return [][]string{
		[]string{"Name", "Project"},
	}
}


/*
	创建 project 外键约束
*/
func (c * SqlClass)ConfCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",confTableName,confProjectForeignKeyName,projectId,projectTableName,primaryKey)

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


//////  projectConf  curd

/*
	初始化所有 project conf key
*/
func InitConfKey(projectId int)error{


	if err := InitAwsConfKey(projectId);err != nil{
		return err
	}

	if err := InitMysqlConfKey(projectId);err != nil{
		return err
	}

	if err := InitRedisConfKey(projectId);err != nil{
		return err
	}

	if err := InitKafkaConfKey(projectId);err != nil{
		return err
	}

	if err := InitESConfKey(projectId);err != nil{
		return err
	}

	if err := InitEmqttConfKey(projectId);err != nil{
		return err
	}

	if err := InitNacosConfKey(projectId);err != nil{
		return err
	}

	return nil
}

func init(){
	orm.RegisterModelWithPrefix("t_",new(Conf))
}

/*
	查询所有 project conf
*/
func GetAllConfs(projectId int)(c []*Conf, err error) {

	o := orm.NewOrm()

	if _, err := o.QueryTable(confTableName).Filter("project_id",projectId).RelatedSel().All(&c); err != nil {
		return nil, err
	}

	return
}


func SearchConfs(searchText string, projectId int)( c []*Conf, err error){

	cond := orm.NewCondition()
	// search := cond.And("project_id",projectId).And("name__icontains",searchText).Or("value__icontains",searchText).And("project_id",projectId)
	or := cond.Or("name__icontains",searchText).Or("value__icontains",searchText)
	search := cond.And("project_id", projectId).AndCond(or)

	o := orm.NewOrm()
	qs := o.QueryTable(confTableName)


	if _, err := qs.SetCond(search).RelatedSel().All(&c); err != nil{
		return nil, err
	}

	return
}


func GetConfById(id int)( *Conf,  error) {

	o := orm.NewOrm()
	p := &Conf{Id:id}
	if err := o.Read(p); err != nil {
		return nil,err
	}
	return p, nil
}



func GetConfByConfName(name string)( *Conf,  error) {

	o := orm.NewOrm()
	c := &Conf{Name:name}
	if err := o.Read(c,"name"); err != nil {
		return nil,err
	}
	return c, nil
}


func GetConfByConfNameAndProjectId(name string,projectId int)(c []*Conf, err error) {

	o := orm.NewOrm()

	if _, err = o.QueryTable(confTableName).Filter("name",name).Filter("project_id",projectId).RelatedSel().All(&c) ;err != nil{
		return c, err
	}
	return c, nil
}



func GetConfByProjectId(projectId int)(p []*Conf, err error) {

	o := orm.NewOrm()
	_ ,err = o.QueryTable(confTableName).Filter("project_id",projectId).RelatedSel().All(&p)
	if err != nil {
		return nil, err
	}

	return p,nil
}

func GetConfByServiceId(serviceId int)(p []*Conf, err error) {

	o := orm.NewOrm()
	_ ,err = o.QueryTable(confTableName).Filter("service_id", serviceId).All(p)
	if err != nil {
		return nil, err
	}

	return p,nil
}



/*
	增加project conf
*/
func AddConf(conf *Conf)(id int64, err error){

	o := orm.NewOrm()
	id, err = o.Insert(conf)
	if err != nil{
		return 0, err
	}

	return id, nil
}


/*
	修改 project conf
*/
func UpdateConf(conf *Conf)(num int64, err error){

o := orm.NewOrm()

	if num, err = o.Update(conf); err == nil {
		return num, err
	}
	return num, err
}


/*
	删除 project conf
*/
func DeleteConfById(id int)(bool, error){

	o := orm.NewOrm()
	v := Conf{Id: id}
	// ascertain id exists in the database
	if err := o.Read(&v); err != nil {
		return false, err
	}
	num, err := o.Delete(&Conf{Id: id})
	if err != nil {
		return false, err
	}

	fmt.Println("Number of records deleted in database:", num)

	return true, nil
}


//////////////
/*
	多对多操作
 */

/*
	通过 conf id 查询 conf 以及关联表数据
 */
func GetConfRelated(conf *Conf) (*Conf, error){

	o := orm.NewOrm()


	if _, err := o.LoadRelated(conf, "Services"); err != nil{
		return conf, err
	}

	return conf, nil
}

/*
	删除 conf 的 project 关联关系
*/
func DeleteConfRelatedServices(conf *Conf) (err error) {

	o := orm.NewOrm()
	// 获取 conf 关联的 projects， 直接全部清除
	m2mForProjcts := o.QueryM2M(conf, "Services")
	if _, err = m2mForProjcts.Clear(); err != nil{
		beego.Error("error: m2m.Clear(conf.Services): ", err)
		return
	}

	return
}



/*
	添加 conf 以及关联表 数据
 */
func UpdateConfRelatedServices(conf *Conf) (err error) {

	o := orm.NewOrm()

	// 校验 conf 数据 是否为 nil
	if len(conf.Services) != 0{
		// 建立 host 关系, 插入link_host表 host 数据
		m2m := o.QueryM2M(conf, "Services")
		_, err = m2m.Add(conf.Services)
		if err != nil{
			beego.Error("error: m2m.Add(conf.Services): ", err)
			return
		}
	}

	return
}
