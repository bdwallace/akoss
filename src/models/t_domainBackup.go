package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// table
type DomainBackup struct {
	Id				        	int	                `orm:"column(id);pk;auto"`
	DomainName					string				`orm:"column(domain_name);size(100)"`
	InstanceEndTime				string				`orm:"column(instance_end_time);size(100)"`
	DnsServer 					string				`orm:"column(dns_server);size(100)"`
	DomainInfo 					string				`orm:"column(domain_info);type(text)"`
	DomainRecord				string				`orm:"column(domain_record);type(text)"`
}

type DomainBackupRecord struct {
	Value      string			 `json:"Value" xml:"Value"`
	TTL        int64 			 `json:"TTL" xml:"TTL"`
	Remark     string			 `json:"Remark" xml:"Remark"`
	DomainName string			 `json:"DomainName" xml:"DomainName"`
	RR         string			 `json:"RR" xml:"RR"`
	Priority   int64 			 `json:"Priority" xml:"Priority"`
	RecordId   string			 `json:"RecordId" xml:"RecordId"`
	Status     string			 `json:"Status" xml:"Status"`
	Locked     bool  			 `json:"Locked" xml:"Locked"`
	Weight     int   			 `json:"Weight" xml:"Weight"`
	Line       string			 `json:"Line" xml:"Line"`
	Type       string			 `json:"Type" xml:"Type"`
}


func init(){
	orm.RegisterModelWithPrefix("t_",new(DomainBackup))
}


func GetBackupById(id int)( *DomainBackup,  error) {

	o := orm.NewOrm()
	domainBackup := &DomainBackup{Id:id}
	if err := o.Read(domainBackup); err != nil {
		return nil,err
	}
	return domainBackup, nil
}

func GetBackupByDomainName(domainName string)( *DomainBackup,  error) {
	o := orm.NewOrm()

	r := new(DomainBackup)
	if err := o.QueryTable(DomainBackupTableName).Filter("domain_name", domainName).One(r); err != nil{

		return nil,err
	}

	return r, nil
}

func GetBackupDomainByPageBreak(start, length int)([]*DomainBackup, error) {

	o := orm.NewOrm()
	resDomainBackup := make([]*DomainBackup,0,60)
	_, err := o.QueryTable(DomainBackupTableName).Limit(length,start).OrderBy("id").All(&resDomainBackup)
	if err != nil{
		return nil, err
	}

	return resDomainBackup, nil
}

func AddOrUpdateBackupDomain(domainBackup *DomainBackup)(err error){

	o := orm.NewOrm()
	r := new(DomainBackup)
	err = o.QueryTable(DomainBackupTableName).Filter("domain_name", domainBackup.DomainName).One(r)

	if r.Id > 0 {
		// update
		domainBackup.Id = r.Id
		if _, err = o.Update(domainBackup); err == nil {
			return
		}
	} else if err.Error() == "<QuerySeter> no row found" {
		// add
		_, err = o.Insert(domainBackup)
		if err != nil{
			return
		}
	} else {
		err = fmt.Errorf("error:  Add or update domain backup is failed!  domian: %s\n",domainBackup.DomainName)
		return
	}

	return nil
}
