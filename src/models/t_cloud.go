package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Cloud struct {
	Id           int     `orm:"column(id);auto"`
	Name         string  `orm:"column(name);"`
	Class        string  `orm:"column(class);"`
	Account      string  `orm:"column(account);"`
	Direction    string  `orm:"column(direction)"`
	Link         string  `orm:"column(link);"`
	Balance      float64 `orm:"column(balance)"`
	AccessKeyId  string  `orm:"column(access_key_id)"`
	AccessSecret string  `orm:"column(access_secret);type(text)'"`

	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);auto_now;"`
}

func (t *Cloud) TableName() string {
	// return "t_Cloud"
	return cloudTableName
}

func init() {
	orm.RegisterModel(new(Cloud))
	// orm.RegisterModelWithPrefix("t_",new(Cloud))
}

// AddSystem insert a new System into database and returns
// last inserted Id on success.
func AddCloud(o orm.Ormer, s *Cloud) (id int64, err error) {
	if o == nil {
		o = orm.NewOrm()
	}

	id, err = o.Insert(s)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetCloudAll() (m []*Cloud, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(cloudTableName).All(&m)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetCloudById(id int) (s *Cloud, err error) {
	o := orm.NewOrm()
	s = &Cloud{}
	err = o.QueryTable(cloudTableName).Filter("id", id).One(s)

	return
}

// UpdateSystem updates System by Id and returns error if
// the record to be updated doesn't exist
func UpdateCloud(o orm.Ormer, s *Cloud) (err error) {
	if o == nil {
		o = orm.NewOrm()
	}

	_, err = o.Update(s)

	return
}

// DeleteSystem deletes System by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCloud(o orm.Ormer, s *Cloud) (err error) {
	if o == nil {
		o = orm.NewOrm()
	}

	_, err = o.Delete(s)

	return
}
