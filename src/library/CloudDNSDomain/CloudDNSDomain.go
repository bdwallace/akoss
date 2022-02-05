package CloudDNSDomain

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"library/aliDNS"
	"library/domainBackup"
	"models"
	"models/request"
)


type DomainAdapterDomain struct {
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

}


func DomainSearch(searchReq *request.DomainForSearchRequest) (resp interface{}, err error){

	if searchReq.Src == "" {
		return nil, fmt.Errorf("error:  DomainAdapterSearchOrBackUpLocal  ::  src or dest is empty")
	}

	if searchReq.Src == "ALi"{
		// domain search from ali domain
		resp := SearchAliDomain(searchReq)
		return resp, nil
	}

	if searchReq.Src == "TenCent" {
		// domain search from tencent domain

	}

	if searchReq.Src == "CloudFare"{
		// domain search from cloudfare domain

	}


	return
}



func DomainBackUpToLocal(backupReq *request.DomainForBackUpTolocalRequest) (err error) {

	if backupReq.Src == "" {
		return fmt.Errorf("error:  DomainAdapterSearchOrBackUpLocal  ::  src or dest is empty")
	}

	if backupReq.Src == "ALi"{
		// domain backup to local ali domain
		err = BackUpOrUpdateToLocalAliDomain(backupReq)
		if err != nil{
			return
		}
	}

	if backupReq.Src == "TenCent" {
		// domain backup to local tencent domain

	}

	if backupReq.Src == "CloudFare"{
		// domain backup to local cloudfare domain

	}

	return nil
}



func DomainImportToCloud(importReq *request.DomainForImportOtherCloudRequest) (resp interface{}, err error){

	if importReq.Dest == "ALi"{
		errMsg, err := ImportToAliDomain(importReq)
		if err != nil{
			return errMsg, err
		}
		return nil, nil
	}
	if importReq.Dest == "TenCent"{
		ImportToTenCentDomain()

	}
	if importReq.Dest == "CloudFare"{

	}

	return
}






func BackUpOrUpdateToLocalAliDomain (backupReq *request.DomainForBackUpTolocalRequest)(err error){

	client, err := alidns.NewClientWithAccessKey("cn-qingdao", backupReq.SrcKeyId, backupReq.SrcKeySecret)
	if err != nil {
		fmt.Print("error:  alidns.NewClientWithAccessKey is failed  ::  ", err)
		//c.SetJson(1, err, "获取域名失败")
		return
	}

	remoteDomains := make([]*models.AliDomain, 0)
	request := alidns.CreateDescribeDomainsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(100)
	remoteTotal := 100
	response := new(alidns.DescribeDomainsResponse)
	models.ConnDB()
	for i := 1; (100*i)-remoteTotal < 100; i++ {
		request.PageNumber = requests.NewInteger(i)
		response, err = client.DescribeDomains(request)
		if err != nil {
			fmt.Println("error:  DescribeDomains is failed  ::  ", err.Error())
			//c.SetJson(1, err, "获取域名失败")
			return
		}
		remoteTotal = int(response.TotalCount)

		backupDomian := &domainBackup.AnalysisImportDomain{
			Resp:      response,
			KeyId:     backupReq.SrcKeyId,
			KeySecret: backupReq.SrcKeySecret,
			Client:    client,
		}

		domains, err := domainBackup.AliCloudAnalysisResponse(backupDomian)
		if err != nil {
			fmt.Println("error:  domainBackupAll.AnalysisResponse  ::  ", err.Error())
			//c.SetJson(1, err, "获取域名失败")
			return err
		}
		remoteDomains = append(remoteDomains, domains...)
	}

	for _, updateDomain := range remoteDomains {
		if err = models.AddOrUpdateBackupDomain(updateDomain); err != nil {
			fmt.Println("error:  domainBackupAll.AddOrUpdateBackupDomain  ::  ", err.Error())
			//c.SetJson(1, err, "更新本地数据失败")
			return err
		}
	}

	if int(response.TotalCount) == len(remoteDomains) {

		// 备份，清空，
		tableName, err := models.CreateTableForBackupTheDomainWithAccKeyAndDate(backupReq.SrcKeyId)
		if err != nil {
			fmt.Println("error:  CreateTableForBackupTheDomainWithAccKeyAndDate  ::  ", err.Error())
			//c.SetJson(1, err, "本地创建备份表失败")
			return err
		}

		oldCount, err := models.CountTempDomainBackup(tableName)
		if err != nil {
			fmt.Printf("error:  domainBackupAll.CountTempDomainBackup is failed  ::  %s\n", err.Error())
			//c.SetJson(1, err, "域名未完全备份,请联系管理员")
			return err
		}
		if oldCount > 1 {
			if err = models.TruncateBackupTable(tableName); err != nil {
				fmt.Printf("error:  domainBackup.TruncateBackupTable is failed  tableName: %s ::  %s\n", tableName, err.Error())
			}
		}
		if err = models.InsertTempDomainBackupFormBackup(tableName); err != nil {

			fmt.Printf("error:  domainBackupAll.InsertTempDomainBackupFormBackup  table : %s ::  %s\n", tableName, err.Error())
			//c.SetJson(1, err, "域名备份失败")
			return err
		}
		newCount, err := models.CountTempDomainBackup(tableName)
		if err != nil {
			fmt.Printf("error:  domainBackupAll.CountTempDomainBackup is failed  ::  %s\n", err.Error())
			//c.SetJson(1, err, "域名未完全备份,请联系管理员")
			return err
		}
		countBackupTable, err := models.CountTempDomainBackup(models.AliDomainTableName)
		if err != nil {
			fmt.Printf("error:  domainBackupAll.t_domain_backup is failed  ::  %s\n", err.Error())
			//c.SetJson(1, err, "域名未完全备份,请联系管理员")
			return err
		}
		if newCount >= len(remoteDomains) && newCount == countBackupTable {
			// truncate
			if err = models.TruncateBackupTable(models.AliDomainTableName); err != nil {
				fmt.Printf("error:  domainBackupAll.TruncateBackupTable is failed  tableName: %s  ::  %s\n", models.AliDomainTableName, err.Error())
				//c.SetJson(1, err, "域名未完全备份,请联系管理员")
				return err
			}
		}

	} else {
		err = fmt.Errorf("error:   int(response.TotalCount) != len(resDomains) ")
		//c.SetJson(1, err, "域名备份数量不符,请联系管理员")
		return err
	}

	return nil

}

func BackUpOrUpdateToLocalTenCentDomain (){

}

func BackUpOrUpdateToLocalCloudFareDomain (){

}





func ImportToAliDomain (importReq *request.DomainForImportOtherCloudRequest)(errMsg string, err error){

	req := &domainBackup.ReqImportAliDomainRecord{
		KeyId:     importReq.DestKeyId,
		KeySecret: importReq.DestKeySecret,
		Domains:   importReq.AliDomain,
	}
	failedDomains, err := domainBackup.AliCloudImportDomain(req)
	if err != nil {
		if len(failedDomains) > 0 {
			fmt.Println("failed:  domainBackup.ImportDomain is failed  ::  ", failedDomains)
			errMsg = fmt.Sprintf("导入域名失败： %s\n", failedDomains)
			return
		}

		fmt.Println("error：  domainBackup.ImportDomain is failed  ::  ", err)
		return
	}

	return
}

func ImportToTenCentDomain (){

}

func ImportToCloudFareDomain (){

}


func SearchAliDomain (searchReq *request.DomainForSearchRequest)(*aliDNS.RespGetAliDNS){

	resp := new(aliDNS.RespGetAliDNS)
	var err error
	if searchReq.SearchDomain == "" {
		// 获取全部
		aliReq := new(aliDNS.ReqGetAliDNS)
		aliReq.KeyId = searchReq.SrcKeyId
		aliReq.KeySecret = searchReq.SrcKeySecret
		aliReq.SearchDomain = searchReq.SearchDomain
		aliReq.PageSize = searchReq.PageSize
		aliReq.PageNumber = searchReq.PageNum
		resp, err = aliDNS.GetAliDNS(aliReq)
		if err != nil {
			return nil
		}
	} else {
		// 搜索获取本地
		searchDomain, err := models.GetBackupByDomainName(searchReq.SearchDomain)
		if err != nil {
			// 域名不存在
			fmt.Printf("error:  GetBackupByDomainName is failed! domian:%s  ::  %s\n", searchReq.SearchDomain, err.Error())
			return nil
		}
		resp.Domains = append(resp.Domains, searchDomain)
		resp.Total = 1
		resp.PageNumber = 1
	}
	return resp

}

func SearchTenCentDomain (){

}

func SearchCloudFareDomain (){

}