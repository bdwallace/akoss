package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Tencentdnspod struct {
	Id 				 int			`orm:"column(id);pk;auto"`
	DomainId         int			`orm:"column(domain_id);size(100)"`
	Name             string  		`orm:"column(name);size(100)"`
	Status           string  		`orm:"column(status);size(100)"`
	TTL              int     		`orm:"column(ttl);size(100)"`
	CNAMESpeedup     string  		`orm:"column(cname_speedup);size(100)"`
	DNSStatus        string			`orm:"column(dns_status);size(100)"`
	Grade            string			`orm:"column(grade);size(100)"`
	GroupId          int			`orm:"column(group_id);size(100)"`
	SearchEnginePush string			`orm:"column(search_engine_push);size(100)"`
	Remark           string			`orm:"column(remark);size(100)"`
	Punycode         string			`orm:"column(punycode);size(100)"`
	EffectiveDNS     string			`orm:"column(effective_dns);size(100)"`
	GradeLevel       int			`orm:"column(grade_level);size(100)"`
	GradeTitle       string			`orm:"column(grade_title);size(100)"`
	IsVip            string			`orm:"column(is_vip);size(100)"`
	VipStartAt       string			`orm:"column(vip_start_at);size(100)"`
	VipEndAt         string			`orm:"column(vip_end_at);size(100)"`
	VipAutoRenew     string			`orm:"column(vip_auto_renew);size(100)"`
	RecordCount      int			`orm:"column(record_count);size(100)"`
	CreatedOn        string			`orm:"column(created_on);size(100)"`
	UpdatedOn        string			`orm:"column(updated_on);size(100)"`
	Owner            string			`orm:"column(owner);size(100)"`
	DomainDescribe	 string 		`orm:"column(domain_describe);type(text)"`
}

type DomainDescribe struct {
	DomainId     *uint64   `json:"DomainId,omitempty" name:"DomainId"`
	Status       *string   `json:"Status,omitempty" name:"Status"`
	Grade        *string   `json:"Grade,omitempty" name:"Grade"`
	GroupId      *uint64   `json:"GroupId,omitempty" name:"GroupId"`
	IsMark       *string   `json:"IsMark,omitempty" name:"IsMark"`
	TTL          *uint64   `json:"TTL,omitempty" name:"TTL"`
	CnameSpeedup *string   `json:"CnameSpeedup,omitempty" name:"CnameSpeedup"`
	Remark       *string   `json:"Remark,omitempty" name:"Remark"`
	Punycode     *string   `json:"Punycode,omitempty" name:"Punycode"`
	DnsStatus    *string   `json:"DnsStatus,omitempty" name:"DnsStatus"`
	DnspodNsList []*string `json:"DnspodNsList,omitempty" name:"DnspodNsList"`
	Domain       *string   `json:"Domain,omitempty" name:"Domain"`
	GradeLevel   *uint64   `json:"GradeLevel,omitempty" name:"GradeLevel"`
	UserId       *uint64   `json:"UserId,omitempty" name:"UserId"`
	IsVip        *string   `json:"IsVip,omitempty" name:"IsVip"`
	Owner        *string   `json:"Owner,omitempty" name:"Owner"`
	GradeTitle   *string   `json:"GradeTitle,omitempty" name:"GradeTitle"`
	CreatedOn    *string   `json:"CreatedOn,omitempty" name:"CreatedOn"`
	UpdatedOn    *string   `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	Uin          *string   `json:"Uin,omitempty" name:"Uin"`
	ActualNsList []*string `json:"ActualNsList,omitempty" name:"ActualNsList"`
	RecordCount  *uint64   `json:"RecordCount,omitempty" name:"RecordCount"`
	OwnerNick    *string   `json:"OwnerNick,omitempty" name:"OwnerNick"`
}
func init() {
	orm.RegisterModelWithPrefix("t_", new(Tencentdnspod))
}

func AddOrUpdateTenCentDNSPod(tenCentDomain *Tencentdnspod) (err error) {

	o := orm.NewOrm()
	r := new(Tencentdnspod)
	err = o.QueryTable(TenCentDNSPodTableName).Filter("name", tenCentDomain.Name).One(r)

	if r.Id > 0 {
		// update
		tenCentDomain.Id = r.Id
		if _, err = o.Update(tenCentDomain); err != nil {
			return
		}
	} else if err.Error() == "<QuerySeter> no row found" {
		// add
		_, err = o.Insert(tenCentDomain)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("error:  Add or update TenCentDNSPod is failed!  domian: %s\n", tenCentDomain.Name)
		return
	}

	return nil
}


func GetAllTenCentDNSPodDomains() ([]*Tencentdnspod, error) {
	o := orm.NewOrm()
	TenCentDNSPods := make([]*Tencentdnspod,0)
	_, err := o.QueryTable(TenCentDNSPodTableName).OrderBy("id").All(&TenCentDNSPods )
	if err != nil {
		return nil, err
	}

	return TenCentDNSPods , nil
}

func GetTenCentDNSPodDomainsByDomainName(domainName string) ([]*Tencentdnspod, error) {
	o := orm.NewOrm()

	r := make([]*Tencentdnspod,0)
	if _, err := o.QueryTable(TenCentDNSPodTableName).Filter("name__contains", domainName).All(&r); err != nil {
		return nil, err
	}

	return r, nil
}