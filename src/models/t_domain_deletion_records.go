package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"library/common"
	"strings"
	"time"
)

type DeleteDomainReco struct {
	Id        		int    				`orm:"column(id);pk;auto"`
	Domain     		string 				`orm:"column(domain);size(100);unique"`
	Class     		string 				`orm:"column(class);size(100)"`
	Comment 		string 				`orm:"column(comment);type(text)"`
	UserName		string 				`orm:"column(user_name);size(200)"`

	CreatedAt 		time.Time 			`orm:"column(created_at);type(datetime);auto_now_add;"`

	Project 		*Project			`orm:"rel(fk)"`
}

func init(){
	orm.RegisterModelWithPrefix("t_",new(DeleteDomainReco))
}


/*
	创建 project 外键约束
*/

func (c * SqlClass)DeleteDomainRecordsCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",domainDeleteRecordsTableName,domainDeleteRecordsFroProjectForeignKeyName,projectId,projectTableName,primaryKey)

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


func AddDomainRecord(d *DeleteDomainReco)(id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(d)
	if err != nil{
		return
	}
	return
}

func GetAllDomainRecords(projectId int)(res []*DeleteDomainReco, err error) {

	if _, err = o.QueryTable(domainDeleteRecordsTableName).Filter("project_id",projectId).RelatedSel().All(&res);err != nil{
		return
	}

	return
}

func SearchDomainRecords(projectId int, start, length int, class, searchText string)(total int64, res []*DeleteDomainReco, err error){

	base := fmt.Sprintf("SELECT `t_delete_domain_reco`.`domain`, `t_delete_domain_reco`.`id`, `t_delete_domain_reco`.`user_name`, `t_delete_domain_reco`.`class`, `t_delete_domain_reco`.`created_at` , `t_delete_domain_reco`.`comment` FROM `%s` `t_delete_domain_reco`", domainDeleteRecordsTableName)

	where := "WHERE"
	if projectId != 0 {
		where = fmt.Sprintf("%s `t_delete_domain_reco`.`project_id` = %d", where, projectId)
	}

	if class != "NOT_FILTER" {
		where = fmt.Sprintf("%s AND `t_delete_domain_reco`.`class` in (%s)", where, common.AddStringListSingle(class))
	}
	if searchText != "" {
		where = fmt.Sprintf("%s AND (`t_delete_domain_reco`.`domain` LIKE '%%%s%%' OR `t_delete_domain_reco`.`user_name` LIKE '%%%s%%' OR `t_delete_domain_reco`.`comment` LIKE '%%%s%%')", where, searchText, searchText, searchText)
	}

	sql := fmt.Sprintf("%s %s", base, where)

	o := orm.NewOrm()

	var count []orm.Params
	o.Raw("SELECT count(id) FROM (" + sql + ") as temp").Values(&count)
	if len(count) > 0 {
		total = int64(common.GetInt(count[0]["count(id)"]))
	}

	_, err = o.Raw(sql + " order by id desc "+ " LIMIT ?,? ", start, length).QueryRows(&res)

	return
}


func GetDomainRecoByProjectForClass(projectId int) (list *orm.ParamsList, err error) {
	list = new(orm.ParamsList)
	o := orm.NewOrm()
	_ ,err = o.QueryTable(domainDeleteRecordsTableName).Filter("project_id",projectId).GroupBy("class").OrderBy("class").ValuesFlat(list, "class")
	return
}
