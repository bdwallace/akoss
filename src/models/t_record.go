package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Record struct {
	Id        	int    		`orm:"column(id);auto"`
	Status    	int16  		`orm:"column(status)"`
	Action    	int16  		`orm:"column(action);null"`
	Command   	string 		`orm:"column(command);type(text);null"`
	Duration  	int    		`orm:"column(duration);null"`
	Memo      	string 		`orm:"column(memo);type(text);null"`
	Host        string    	`orm:"column(host);null"`
	Count 		int 		`orm:"column(Count);null"`

	CreatedAt              time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt              time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	User         	*User				`orm:"rel(fk)"`
	Task         	*Task				`orm:"rel(fk)"`
	Deploy			*Deploy				`orm:"rel(fk)"`
}

// func (t *Record) TableName() string {
// 	return "record"
// }

func init() {
	orm.RegisterModelWithPrefix("t_",new(Record))
}


/*
	添加 user id 外键约束
*/
func (c *SqlClass)RecordCreateForeignKeyUser() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",recordTableName,recordUserForeignKeyName,userId,userTableName,primaryKey)

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
	添加 task id 外键约束
*/
func (c *SqlClass)RecordCreateForeignKeyTask() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",recordTableName,recordTaskForeignKeyName,taskId,taskTableName,primaryKey)

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
	创建 project 外键约束
*/
func (c * SqlClass)RecordCreateForeignKeyDeploy() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf( "alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);",recordTableName,recordDeployForeignKeyName,deployId,deployTableName,primaryKey)

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



// AddRecord insert a new Record into database and returns
// last inserted Id on success.
func AddRecord(m *Record) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRecordById retrieves Record by Id. Returns error if
// Id doesn't exist
func GetRecordById(id int) (v *Record, err error) {
	o := orm.NewOrm()
	v = &Record{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRecord retrieves all Record matches certain condition. Returns empty list if
// no records exist
func GetAllRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Record))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Record
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRecord updates Record by Id and returns error if
// the record to be updated doesn't exist
func UpdateRecordById(m *Record) (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(m); err != nil {
		return
	}

	return
}

// UpdateRecord updates Record by Id and returns error if
// the record to be updated doesn't exist
func UpdateRecordByTaskHost(m *Record) (err error) {
	o := orm.NewOrm()
	v := Record{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(v); err != nil {
		fmt.Println("error: UpdateRecordByTaskHost Read", err)
		return
	}

	if _, err = o.Update(m); err != nil {
		fmt.Println("error: UpdateRecordByTaskHost Update", err)
		return
	}
	return
}

// DeleteRecord deletes Record by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRecord(id int) (err error) {
	o := orm.NewOrm()
	v := Record{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Record{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// DeleteRecord deletes Record by Id and returns error if
// the record to be deleted doesn't exist
func GetRecordCount(deployId int) (list *orm.ParamsList, err error) {
	list = new(orm.ParamsList)
	o := orm.NewOrm()
	_, err = o.QueryTable(recordTableName).Filter("deploy_id", deployId).GroupBy("count").OrderBy("-count").ValuesFlat(list, "count")

	return
}



// DeleteRecord deletes Record by Id and returns error if
// the record to be deleted doesn't exist
func GetRecordListCount(deployId int, count int) (list []*Record, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(recordTableName).Filter("deploy_id", deployId).Filter("count", count).All(&list)

	return
}

