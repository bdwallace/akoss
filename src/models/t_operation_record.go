package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OperationRecord struct {
	Id        	int    		`orm:"column(id);auto"`
	Status    	int16  		`orm:"column(status);null"`
	Action    	int16  		`orm:"column(action);null"`
	Class    	string  	`orm:"column(class);null"`			// build start stop dockerPs
	Command   	string 		`orm:"column(command);type(text);null"`
	Duration  	int    		`orm:"column(duration);null"`
	Memo      	string 		`orm:"column(memo);type(text);null"`
	Count 		int 		`orm:"column(Count);null"`

	CreatedAt              time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt              time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	User         	*User				`orm:"rel(fk)"`
	Service			*Service			`orm:"rel(fk)"`
	Host			*Host				`orm:"rel(fk)"`
}



func init() {
	orm.RegisterModelWithPrefix("t_",new(OperationRecord))
}


/*
	添加 user id 外键约束
*/
func (c *SqlClass)OperationRecordCreateForeignKeyUser() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",operationRecordTableName,operationRecordForUserForeignKeyName,userId,userTableName,primaryKey)

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
	添加 service id 外键约束
*/
func (c *SqlClass)OperationRecordCreateForeignKeyService() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", operationRecordTableName, operationRecordForServiceForeignKeyName, serviceId, serviceTableName, primaryKey)

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




func AddOperationRecord(r *OperationRecord) (id int64, err error) {
	o := orm.NewOrm()
	if id, err = o.Insert(r); err == nil{
		return id,nil
	}

	return -1, err
}


func GetAllOperationRecord()(r []*OperationRecord,err error){

	o := orm.NewOrm()
	if _, err = o.QueryTable(operationRecordTableName).All(&r);err != nil{
		return nil,err
	}
	return r,nil

}

func GetOperationRecordById(id int) (r *OperationRecord, err error) {
	o := orm.NewOrm()
	r = &OperationRecord{}

	err = o.QueryTable(operationRecordTableName).Filter("id", id).RelatedSel().One(r)

	return
}


func GetOperationRecordByServiceId(serviceId int) (r []*OperationRecord, err error) {

	o := orm.NewOrm()
	_ ,err = o.QueryTable(operationRecordTableName).Filter("service_id",serviceId).RelatedSel().All(&r)
	if err != nil {
		return nil, err
	}
	return r, nil

}

func GetOperationRecordByUserId(userId int) (r []*OperationRecord, err error) {

	o := orm.NewOrm()
	_ ,err = o.QueryTable(operationRecordTableName).Filter("user_id",userId).RelatedSel().All(&r)
	if err != nil {
		return nil, err
	}
	return r,nil

}



func UpdateOperationRecordById(r *OperationRecord) (id int64 ,err error) {
	o := orm.NewOrm()

	//fmt.Printf("update task id: %d  task cmd: %s",m.Id,m.Cmd)
	if id, err = o.Update(r); err == nil {
		return -1,err
	}
	return id, nil
}


func DeleteOperationRecord(id int) (num int64, err error) {

	o := orm.NewOrm()
	if num, err = o.Delete(&OperationRecord{Id: id}); err == nil {
		return num, nil
	}
	return -1, err
}


func GetOperationRecordByServiceIdClass(serviceId int, class string) (r []*OperationRecord, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(operationRecordTableName).Filter("service_id",serviceId).Filter("class", class).OrderBy("-id").Limit(1).All(&r)
	if err != nil {
		return r, err
	}
	return r, nil
}


// DeleteRecord deletes Record by Id and returns error if
// the record to be deleted doesn't exist
func GetOperationRecordCount(serviceId int, class string) (list *orm.ParamsList, err error) {
	list = new(orm.ParamsList)
	o := orm.NewOrm()
	_, err = o.QueryTable(operationRecordTableName).Filter("service_id", serviceId).Filter("class", class).GroupBy("count").OrderBy("-count").Limit(20).ValuesFlat(list, "count")

	return
}



// DeleteRecord deletes Record by Id and returns error if
// the record to be deleted doesn't exist
func GetOperationRecordListCount(serviceId int, class string, count int) (list []*OperationRecord, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(operationRecordTableName).Filter("service_id", serviceId).Filter("class", class).Filter("count", count).All(&list)

	return
}


