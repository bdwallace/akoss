package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Link struct {
	Id        int    `orm:"column(id);auto"`
	Name      string `orm:"column(name);"`
	Direction string `orm:"column(direction)"`
	Link      string `orm:"column(link);"`
	Class     string `orm:"column(class)"`

	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add;"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);auto_now;"`

	// Project 	*Project	`orm:"rel(fk)"`
	// Platform	*Platform	`orm:"rel(fk)"`
	Projects []*Project `orm:"rel(m2m)"`
	// Platforms	[]*Platform	`orm:"rel(m2m)"`
}

func (t *Link) TableName() string {
	// return "t_link"
	return linkTableName
}

func init() {
	orm.RegisterModel(new(Link))
	// orm.RegisterModelWithPrefix("t_",new(Link))
}

// AddSystem insert a new System into database and returns
// last inserted Id on success.
func AddLink(o orm.Ormer, s *Link) (id int64, err error) {
	if o == nil {
		o = orm.NewOrm()
	}

	id, err = o.Insert(s)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetLinkAll() (m []*Link, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(linkTableName).All(&m)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetLinkById(id int) (s *Link, err error) {
	o := orm.NewOrm()
	s = &Link{}
	err = o.QueryTable(linkTableName).Filter("id", id).One(s)

	return
}

// GetSystemById retrieves System by Id. Returns error if
// Id doesn't exist
func GetLinkByProjectlName(level int, name string) (m []Link, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(linkTableName).Filter("level", level).Filter("name", name).All(&m)

	return
}

// UpdateSystem updates System by Id and returns error if
// the record to be updated doesn't exist
func UpdateLink(o orm.Ormer, s *Link) (err error) {
	if o == nil {
		o = orm.NewOrm()
	}

	_, err = o.Update(s)

	return
}

// DeleteSystem deletes System by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLink(o orm.Ormer, s *Link) (err error) {
	if o == nil {
		o = orm.NewOrm()
	}

	_, err = o.Delete(s)

	return
}

//////////////
/*
	多对多操作
*/

//通过 link 查询 projects 关联表数据
func GetLinkRelatedProjects(link *Link) (err error) {

	o := orm.NewOrm()

	if _, err = o.LoadRelated(link, "Projects"); err != nil {
		return
	}

	return
}

/*
	删除 link 的 project 关联关系
*/
func DeleteLinkRelatedProjects(o orm.Ormer, link *Link) (err error) {

	if o == nil {
		o = orm.NewOrm()
	}
	// 获取 link 关联的 projects， 直接全部清除
	m2mForProjcts := o.QueryM2M(link, "Projects")
	if _, err = m2mForProjcts.Clear(); err != nil {
		return
	}

	return
}

/*
	添加 link 以及关联表 数据
*/
func UpdateLinkRelatedProjects(o orm.Ormer, link *Link) (err error) {

	if o == nil {
		o = orm.NewOrm()
	}

	// 校验 conf 数据 是否为 nil
	if len(link.Projects) != 0 {
		// 建立 host 关系, 插入link_host表 host 数据
		m2m := o.QueryM2M(link, "Projects")
		_, err = m2m.Add(link.Projects)
		if err != nil {
			fmt.Println("error: m2m.Add(link.Projects): ", err)
			return
		}
	}

	return
}
