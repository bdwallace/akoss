package models

import (
	"fmt"
	"library/common"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Host struct {
	Id        		int    				`orm:"column(id);pk;auto"`
	Name      		string 				`orm:"column(name);size(100);unique"`
	PrivateIp 		string 				`orm:"column(private_ip);size(100);unique"`
	PublicIp  		string 				`orm:"column(public_ip);size(100)"`
	UseIp			string				`orm:"column(use_ip)"`
	InstanceId 		string 				`orm:"column(instance_id);size(200)"`	  //  ++
	Region     		string 				`orm:"column(region);size(200)"`			// ++
	StartTime  		string 				`orm:"column(start_time);size(100)"`		// ++
	EndTime    		string 				`orm:"column(end_time);size(100)"`		// ++

	CreatedAt              time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt              time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	Project 		*Project			`orm:"rel(fk)"`
	Services 		[]*Service			`orm:"reverse(many)"`

}
// UsePublic 0:表示不使用公网IP,1:表示使用公网IP
/*const(
	// 设置表信息
	hostMainTableName = "t_host"
	hostMainField = "id"
	//hostSubField = "service_id"
	//hostForeignKeyName = "service_id"
	//
	hostProjectForeignKeyName = "host_project_id"
	hostProjectfield = "project_id"
)
*/
// hostMainTableName,serviceProjectForeignKeyName,hostProjectfield,projectMainTableName,projectMainField)

func init(){
	orm.RegisterModelWithPrefix("t_",new(Host))
}


/*
	创建 project 外键约束
*/

func (c * SqlClass)HostCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",hostTableName,hostProjectForeignKeyName,projectId,projectTableName,primaryKey)

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



//////  host curd

func AddHost(h *Host) (id int64, err error) {

	o := orm.NewOrm()

	id, err = o.Insert(h)
	if err != nil{
		return
	}

	return
}



/*
	查询所有 host
*/
func GetAllHost(projectId int) (h []*Host, err error) {

	o := orm.NewOrm()

	if _, err = o.QueryTable(hostTableName).Filter("project_id",projectId).RelatedSel().All(&h);err != nil{
		return nil, err
	}

	return
}



func SearchHosts(projectId int, awsRegion string, searchText string)(h []*Host, err error){
	base := fmt.Sprintf("SELECT `T0`.`id`, `T0`.`name`, `T0`.`private_ip`, `T0`.`public_ip`, `T0`.`use_ip`, `T0`.`region`, `T0`.`instance_id`, `T0`.`start_time`, `T0`.`end_time` FROM `%s` `T0`", hostTableName)

	where := "WHERE"
	if projectId != 0 {
		where = fmt.Sprintf("%s `T0`.`project_id` = %d", where, projectId)
	}
	if awsRegion != "NOT_FILTER" {
		where = fmt.Sprintf("%s AND `T0`.`region` in (%s)", where, common.AddStringListSingle(awsRegion))
	}

	if searchText != "" {
		// where = fmt.Sprintf("%s AND (`T0`.`name` LIKE '%%%s%%' OR `T0`.`domain` LIKE '%%%s%%' OR `t_service`.`name` LIKE '%%%s%%' OR `t_platform`.`name` LIKE '%%%s%%')", where, searchText, searchText, searchText, searchText)
		where = fmt.Sprintf("%s AND (`T0`.`name` LIKE '%%%s%%' OR `T0`.`private_ip` LIKE '%%%s%%' OR `T0`.`public_ip` LIKE '%%%s%%' OR `T0`.`region` LIKE '%%%s%%')", where, searchText, searchText, searchText, searchText)
	}

	sql := fmt.Sprintf("%s %s", base, where)

	o := orm.NewOrm()

	// var count []orm.Params
	// o.Raw("SELECT count(id) FROM (" + sql + ") as temp").Values(&count)
	// if len(count) > 0 {
	// 	total = common.GetInt(count[0]["count(id)"])
	// }

	// _, err = o.Raw(sql + " LIMIT ?,?", start, length).QueryRows(&d)
	_, err = o.Raw(sql).QueryRows(&h)

	return
}


/*
	通过 id  查询 host
*/
func GetHostById(id int) (h *Host, err error) {

	o := orm.NewOrm()
	h = &Host{}
	err = o.QueryTable(hostTableName).Filter("id", id).RelatedSel().One(h)
	return

}

/*
	通过 project_id  查询 host
*/
func GetHostByProjectId(project_id int) (h []*Host, err error) {

	o := orm.NewOrm()
	_ ,err = o.QueryTable(hostTableName).Filter("project_id",project_id).RelatedSel().All(&h)
	if err != nil {
		fmt.Println("error: o.QueryTable(hostMainTableName) by project_id  ",err)
		return nil, err
	}

	return h,err

}

/*
	通过 use_ip 查询 host
*/
func GetHostByUseIp(useIp string) (h *Host, err error) {

	o := orm.NewOrm()
	err = o.QueryTable(hostTableName).Filter("use_ip",useIp).RelatedSel().One(&h)
	if err != nil {
		fmt.Println("error: o.QueryTable(hostMainTableName) by project_id  ",err)
		return nil, err
	}

	return h,err
}





func UpdateHostById(h *Host)(id int64, err error){

	o := orm.NewOrm()

	if id, err = o.Update(h); err != nil{
		return
	}
	return
}


func UpdateHostByHost(m *Host) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func DeleteHostById(id int)(bool, error){

	o := orm.NewOrm()
	v := Host{Id: id}
	//ascertain id exists in the database
	if err := o.Read(&v); err != nil {
		return false,err
	}

	if _, err := o.Delete(&Host{Id: id}); err != nil {
		return false, err
	}

	return true, nil
}


func GetHostRelated(host *Host) (*Host, error) {

	o := orm.NewOrm()

	if _, err := o.LoadRelated(host, "Services"); err != nil {
		return nil, err
	}

	return host, nil
}


func GetHostRelatedProject(host *Host) (err error) {

	o := orm.NewOrm()

	if _, err = o.LoadRelated(host, "Project"); err != nil {
		return
	}

	return
}


func GetHostStopCheck(projectId int) (m []*Host, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(hostTableName).Filter("project_id", projectId).Exclude("instance_id__exact", "").Exclude("start_time__exact", "").Exclude("end_time__exact", "").All(&m)
	return
}




func SetHostStopTime(m []*Host) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	for _, host := range m {
		_, err = o.Update(host, "start_time", "end_time")
		if err != nil {
			err = o.Rollback()
			return
		}
	}

	err = o.Commit()
	return
}
