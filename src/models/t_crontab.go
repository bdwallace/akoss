package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Crontab struct {
	Id        int    `orm:"column(id);auto"`
	Name      string `orm:"column(name);unique"`
	Direction string `orm:"column(direction)"`
	Crontab   string `orm:"column(crontab)"`
	Auto      int    `orm:"column(auto)"`
	Api       string `orm:"column(api)"`
	Address   string `orm:"column(address)"`
	Script    string `orm:"column(script)"`
	Status    int16  `orm:"column(status)"`
	Result    string `orm:"column(result);type(text)"`

	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	Project *Project `orm:"rel(fk)"`
}

/*
	创建 project 外键约束
*/
func (c *SqlClass) CrontabCreateForeignKeyProject() {

	// 拼接两张表外键约束sql
	addForeignSqlStr := fmt.Sprintf("alter table %s ADD CONSTRAINT %s foreign key (%s) references %s(%s);", crontabTableName, crontabForProjectForeignKeyName, projectId, projectTableName, primaryKey)

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
func (t *Crontab) TableName() string {
	return "t_crontab"
}
*/

func init() {
	orm.RegisterModelWithPrefix("t_", new(Crontab))
}

// AddSystem insert a new System into database and returns
// last inserted Id on success.
func AddCrontab(m *Crontab) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetCrontabAll() (m []Crontab, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(crontabTableName).All(&m)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetCrontabByAuto(auto int) (m []Crontab, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(crontabTableName).Filter("auto", auto).All(&m)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetCrontabById(id int) (v *Crontab, err error) {
	o := orm.NewOrm()
	v = &Crontab{Id: id}

	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetCrontabByProjectlName(level int, name string) (m []Crontab, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(crontabTableName).Filter("level", level).Filter("name", name).All(&m)

	return
}

// UpdateSystem updates System by Id and returns error if
// the record to be updated doesn't exist
func UpCrontab(m *Crontab) (err error) {
	o := orm.NewOrm()
	v := Crontab{Id: m.Id}
	if err = o.Read(&v); err != nil {
		fmt.Println("error: UpCrontab Read", err)
		return
	}

	if _, err = o.Update(m); err != nil {
		fmt.Println("error: UpCrontab Update", err)
		return
	}
	return
}

// DeleteSystem deletes System by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCrontab(id int) error {
	o := orm.NewOrm()

	o.Begin()
	if _, err := o.QueryTable(crontabLogTableName).Filter("crontab_id", id).Delete(); err != nil {
		err = fmt.Errorf("error:  o.QueryTable(crontabLogTableName).Filter(\"crontab_id\",id).Delete() %s \n", err)
		fmt.Println(err)
		o.Rollback()
		return err
	}

	if _, err := o.Delete(&Crontab{Id: id}); err != nil {
		err = fmt.Errorf("error:   o.Delete(crontab) %s \n", err)
		fmt.Println(err)
		o.Rollback()
		return err
	}
	o.Commit()

	return nil
}
