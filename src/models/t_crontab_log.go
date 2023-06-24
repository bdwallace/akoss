package models

import (
	// "errors"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CrontabLog struct {
	Id          int    `orm:"column(id);auto"`
	CrontabName string `orm:"column(crontab_name)"`
	Auto        int    `orm:"column(auto)"`
	Status      int16  `orm:"column(status)"`
	Result      string `orm:"column(result);type(text)"`

	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	Crontab *Crontab `orm:"rel(fk)"`
}

/*
	创建 crontab 外键约束
*/
func (c *SqlClass) CrontabLogCreateForeignKeyCrontab() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", crontabLogTableName, crontabLogForCrontabForeignKeyName, crontabId, crontabTableName, primaryKey)

	err := c.CreateForeignKey(addForeignSqlStr)
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

// Auto 1：为自动，0：手动

/*
func (t *CrontabLog) TableName() string {
	return "t_crontab_log"
}
*/
func init() {
	orm.RegisterModelWithPrefix("t_", new(CrontabLog))
}

// AddSystem insert a new System into database and returns
// last inserted Id on success.
func AddCrontabLog(m *CrontabLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetCrontabLogAll() (m []CrontabLog, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(crontabLogTableName).All(&m)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetCrontabLogAllPage(start int, length int) (count int64, m []*CrontabLog) {
	o := orm.NewOrm()
	count, _ = o.QueryTable(crontabLogTableName).Count()
	o.QueryTable(crontabLogTableName).Limit(length, start).OrderBy("-id").All(&m)

	return
}

func GetCrontabLogByCrontabId(crontabId int) (c []CrontabLog, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(crontabLogTableName).Filter("crontab_id", crontabId).RelatedSel().All(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetCrontabLogById(id int) (v *CrontabLog, err error) {
	o := orm.NewOrm()
	v = &CrontabLog{Id: id}

	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateSystem updates System by Id and returns error if
// the record to be updated doesn't exist
func UpCrontabLog(m *CrontabLog) (err error) {
	o := orm.NewOrm()
	v := CrontabLog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err != nil {
		fmt.Println("error: UpCrontabLog Read", err)
		return
	}

	if _, err = o.Update(m); err != nil {
		fmt.Println("error: UpCrontabLog Update", err)
		return
	}
	return
}

// DeleteSystem deletes System by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCrontabLog(id int) (err error) {
	o := orm.NewOrm()
	v := CrontabLog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CrontabLog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
