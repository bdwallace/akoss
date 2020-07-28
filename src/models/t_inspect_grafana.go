package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type InspectGrafana struct {
	Id             	int       				`orm:"column(id);auto"`
	Name      		string 					`orm:"column(name);size(100)"`
	Url 			string 					`orm:"column(url);size(800)"`
	Sessions 		string 					`orm:"column(sessions);size(800)"`
	WaitVisible		string					`orm:"column(wait_visible);size(200)"`
	Domain			string					`orm:"column(domain);size(200)"`

	CreatedAt              time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt              time.Time `orm:"column(updated_at);type(datetime);auto_now;"`
}

func (t *InspectGrafana) TableName() string {
	return inspectGrafanaTableName
}

func init(){
	orm.RegisterModel(new(InspectGrafana))
	// 	orm.RegisterModelWithPrefix("t_",new(InspectGrafana))
}



func AddInspectGrafana(t *InspectGrafana) (id int64, err error) {
	o := orm.NewOrm()
	if id, err = o.Insert(t); err == nil{
		return id,nil
	}

	return -1, err
}


func GetInspectGrafanaById(id int) (t *InspectGrafana, err error) {
	o := orm.NewOrm()
	t = &InspectGrafana{}

	err = o.QueryTable(inspectGrafanaTableName).Filter("id", id).RelatedSel().One(t)

	return
}

func GetAllInspectGrafana()(t []*InspectGrafana,err error){

	o := orm.NewOrm()
	if _, err = o.QueryTable(inspectGrafanaTableName).All(&t);err != nil{
		return nil,err
	}
	return t,nil

}


func UpdateInspectGrafanaById(t *InspectGrafana) (id int64 ,err error) {
	o := orm.NewOrm()

	if id, err = o.Update(t); err == nil {
		return -1,err
	}
	return id, nil
}


func DeleteInspectGrafanaById(id int) (num int64, err error) {

	o := orm.NewOrm()
	if num, err = o.Delete(&InspectGrafana{Id: id}); err == nil {
		return num, nil
	}
	return -1, err
}

