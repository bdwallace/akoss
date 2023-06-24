package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)


type TencentDomain struct {
	Id              	int    		`orm:"column(id);pk;auto"`
	DomainName      	string 		`orm:"column(domain_name);size(100)"`
	Dns			       	string 		`orm:"column(dns);size(100)"`
	DomainSource    	string 		`orm:"column(domain_source);size(100)"`
	DomainInfo      	string 		`orm:"column(domain_info);type(text)"`			// Tencentdnspod
	DomainRecord	  	string 		`orm:"column(domain_record);type(text)"`		// DomainDescribe
}

type Tencentdnspod struct {
	DomainId         *uint64   `json:"DomainId,omitempty" name:"DomainId"`
	Name             *string   `json:"Name,omitempty" name:"Name"`
	Status           *string   `json:"Status,omitempty" name:"Status"`
	TTL              *uint64   `json:"TTL,omitempty" name:"TTL"`
	CNAMESpeedup     *string   `json:"CNAMESpeedup,omitempty" name:"CNAMESpeedup"`
	DNSStatus        *string   `json:"DNSStatus,omitempty" name:"DNSStatus"`
	Grade            *string   `json:"Grade,omitempty" name:"Grade"`
	GroupId          *uint64   `json:"GroupId,omitempty" name:"GroupId"`
	SearchEnginePush *string   `json:"SearchEnginePush,omitempty" name:"SearchEnginePush"`
	Remark           *string   `json:"Remark,omitempty" name:"Remark"`
	Punycode         *string   `json:"Punycode,omitempty" name:"Punycode"`
	EffectiveDNS     []*string `json:"EffectiveDNS,omitempty" name:"EffectiveDNS"`
	GradeLevel       *uint64   `json:"GradeLevel,omitempty" name:"GradeLevel"`
	GradeTitle       *string   `json:"GradeTitle,omitempty" name:"GradeTitle"`
	IsVip            *string   `json:"IsVip,omitempty" name:"IsVip"`
	VipStartAt       *string   `json:"VipStartAt,omitempty" name:"VipStartAt"`
	VipEndAt         *string   `json:"VipEndAt,omitempty" name:"VipEndAt"`
	VipAutoRenew     *string   `json:"VipAutoRenew,omitempty" name:"VipAutoRenew"`
	RecordCount      *uint64   `json:"RecordCount,omitempty" name:"RecordCount"`
	CreatedOn        *string   `json:"CreatedOn,omitempty" name:"CreatedOn"`
	UpdatedOn        *string   `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	Owner            *string   `json:"Owner,omitempty" name:"Owner"`
}


type TenCentRecordListItem struct {
	RecordId      *uint64 `json:"RecordId,omitempty" name:"RecordId"`
	Value         *string `json:"Value,omitempty" name:"Value"`
	Status        *string `json:"Status,omitempty" name:"Status"`
	UpdatedOn     *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	Name          *string `json:"Name,omitempty" name:"Name"`
	Line          *string `json:"Line,omitempty" name:"Line"`
	LineId        *string `json:"LineId,omitempty" name:"LineId"`
	Type          *string `json:"Type,omitempty" name:"Type"`
	Weight        *uint64 `json:"Weight,omitempty" name:"Weight"`
	MonitorStatus *string `json:"MonitorStatus,omitempty" name:"MonitorStatus"`
	Remark        *string `json:"Remark,omitempty" name:"Remark"`
	TTL           *uint64 `json:"TTL,omitempty" name:"TTL"`
	MX            *uint64 `json:"MX,omitempty" name:"MX"`
}
/*
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
}*/
func init() {
	orm.RegisterModelWithPrefix("t_", new(TencentDomain))
}

func AddOrUpdateTenCentDNSPodDomain(tenCentDomain *TencentDomain) (err error) {

	o := orm.NewOrm()
	r := new(TencentDomain)
	err = o.QueryTable(TenCentDomainTableName).Filter("domain_name", tenCentDomain.DomainName).One(r)

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
		err = fmt.Errorf("error:  Add or update TenCentDNSPod is failed!  domian: %s\n", tenCentDomain.DomainName)
		return
	}

	return nil
}


/*func GetAllTenCentDNSPodDomains() ([]*Tencentdnspod, error) {
	o := orm.NewOrm()
	TenCentDNSPods := make([]*Tencentdnspod,0)
	_, err := o.QueryTable(TenCentDNSPodTableName).OrderBy("id").All(&TenCentDNSPods )
	if err != nil {
		return nil, err
	}

	return TenCentDNSPods , nil
}*/

func GetTenCentDNSPodDomainsByDomainName(domainName string) ([]*TencentDomain, error) {
	o := orm.NewOrm()

	r := make([]*TencentDomain,0)
	_, err := o.QueryTable(TenCentDomainTableName).Filter("domain_name__contains", domainName).All(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetAllTenCentDNSPodDomains(start, length int)([]*TencentDomain, error){
	o := orm.NewOrm()
	res := make([]*TencentDomain, 0, 60)

	if _, err := o.QueryTable(TenCentDomainTableName).Limit(length, start).OrderBy("id").All(&res); err != nil {
		return nil, err
	}
	return res, nil
}