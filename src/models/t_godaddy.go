package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)



type Godaddy struct {
	Id				                int			    `orm:"column(id);pk;auto"`
	Domain							string			`orm:"column(domain);size(100)"`
	DomainId						string			`orm:"column(domain_id);size(100)"`
	Expires							string			`orm:"column(expires);size(100)"`
	Status							string			`orm:"column(status);size(100)"`
}



func init() {
	orm.RegisterModelWithPrefix("t_", new(Godaddy))
}


func AddOrUpdateGodaddy(godaddy *Godaddy) (err error) {

	o := orm.NewOrm()
	r := new(Godaddy)
	err = o.QueryTable(GoDaddyTableName).Filter("domain", godaddy.Domain).One(r)

	if r.Id > 0 {
		// update
		godaddy.Id = r.Id
		if _, err = o.Update(godaddy); err != nil {
			return
		}
	} else if err.Error() == "<QuerySeter> no row found" {
		// add
		_, err = o.Insert(godaddy)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("error:  Add or update godaddy is failed!  domian: %s\n", godaddy.Domain)
		return
	}

	return nil
}


func GetAllGodaddyDomains() ([]*Godaddy, error) {

	o := orm.NewOrm()
	resGodaddy := make([]*Godaddy,0)
	_, err := o.QueryTable(GoDaddyTableName).OrderBy("id").All(&resGodaddy)
	if err != nil {
		return nil, err
	}

	return resGodaddy, nil
}


func GetGodaddyDomainByDomainName(domainName string) ([]*Godaddy, error) {
	o := orm.NewOrm()

	r := make([]*Godaddy,0)
	if _, err := o.QueryTable(GoDaddyTableName).Filter("domain__contains", domainName).All(&r); err != nil {
		return nil, err
	}

	return r, nil
}