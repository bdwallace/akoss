package CloudDNSDomain

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"models"
	"models/request"
	"strings"
	"sync"
)


func SearchAliDomain (searchReq *request.DomainForSearchRequest)(*RespCloudDomain,error){

	resp := new(RespCloudDomain)
	var err error

	if searchReq.SearchDomain == "" {
		// 获取全部
		aliReq := new(ReqGetCloudDomain)
		aliReq.KeyId = searchReq.SrcKeyId
		aliReq.KeySecret = searchReq.SrcKeySecret
		aliReq.SearchDomain = searchReq.SearchDomain
		aliReq.PageSize = searchReq.PageSize
		aliReq.PageNumber = searchReq.PageNum
		resp, err = GetAliDNS(aliReq)
		if err != nil {
			return nil, err
		}
	} else {
		// 搜索获取本地
		searchDomain, err := models.GetBackupByDomainName(searchReq.SearchDomain)
		if err != nil {
			// 域名不存在
			fmt.Printf("error:  GetBackupByDomainName is failed! domian:%s  ::  %s\n", searchReq.SearchDomain, err.Error())
			return nil, err
		}
		resp.AliDomains = append(resp.AliDomains, searchDomain)
		resp.Total = 1
		resp.PageNumber = 1
	}
	return resp, nil

}

func BackUpOrUpdateToLocalAliDomain (backupReq *request.DomainForBackUpTolocalRequest)(err error){

	client, err := alidns.NewClientWithAccessKey("cn-qingdao", backupReq.SrcKeyId, backupReq.SrcKeySecret)
	if err != nil {
		fmt.Print("error:  alidns.NewClientWithAccessKey is failed  ::  ", err)
		//c.SetJson(1, err, "获取域名失败")
		return
	}

	remoteDomains := make([]*models.AliDomain, 0)
	describeDomainsRequest := alidns.CreateDescribeDomainsRequest()
	describeDomainsRequest.Scheme = "https"
	describeDomainsRequest.PageSize = requests.NewInteger(100)
	remoteTotal := 100
	response := new(alidns.DescribeDomainsResponse)
	models.ConnDB()
	for i := 1; (100*i)-remoteTotal < 100; i++ {
		describeDomainsRequest.PageNumber = requests.NewInteger(i)
		response, err = client.DescribeDomains(describeDomainsRequest)
		if err != nil {
			fmt.Println("error:  DescribeDomains is failed  ::  ", err.Error())
			return
		}
		remoteTotal = int(response.TotalCount)

		backupDomian := &AnalysisImportDomain{
			Resp:      response,
			KeyId:     backupReq.SrcKeyId,
			KeySecret: backupReq.SrcKeySecret,
			Client:    client,
		}

		domains, err := AliCloudAnalysisResponse(backupDomian)
		if err != nil {
			fmt.Println("error:  domainBackupAll.AnalysisResponse  ::  ", err.Error())
			return err
		}
		remoteDomains = append(remoteDomains, domains...)
	}

	for _, updateDomain := range remoteDomains {
		if err = models.AddOrUpdateBackupDomain(updateDomain); err != nil {
			fmt.Println("error:  domainBackupAll.AddOrUpdateBackupDomain  ::  ", err.Error())
			return err
		}
	}

	if int(response.TotalCount) == len(remoteDomains) {

		// 备份，清空，
		tableName, err := models.CreateTableForBackupTheDomainWithAccKeyAndDate(backupReq.SrcKeyId)
		if err != nil {
			fmt.Println("error:  CreateTableForBackupTheDomainWithAccKeyAndDate  ::  ", err.Error())
			return err
		}

		oldCount, err := models.CountTempDomainBackup(tableName)
		if err != nil {
			fmt.Printf("error:  domainBackupAll.CountTempDomainBackup is failed  ::  %s\n", err.Error())
			return err
		}
		if oldCount > 1 {
			if err = models.TruncateBackupTable(tableName); err != nil {
				fmt.Printf("error:  domainBackup.TruncateBackupTable is failed  tableName: %s ::  %s\n", tableName, err.Error())
			}
		}
		if err = models.InsertTempDomainBackupFormBackup(tableName); err != nil {

			fmt.Printf("error:  domainBackupAll.InsertTempDomainBackupFormBackup  table : %s ::  %s\n", tableName, err.Error())
			return err
		}
		newCount, err := models.CountTempDomainBackup(tableName)
		if err != nil {
			fmt.Printf("error:  domainBackupAll.CountTempDomainBackup is failed  ::  %s\n", err.Error())
			return err
		}
		countBackupTable, err := models.CountTempDomainBackup(models.AliDomainTableName)
		if err != nil {
			fmt.Printf("error:  domainBackupAll.t_domain_backup is failed  ::  %s\n", err.Error())
			return err
		}
		if newCount >= len(remoteDomains) && newCount == countBackupTable {
			// truncate
			if err = models.TruncateBackupTable(models.AliDomainTableName); err != nil {
				fmt.Printf("error:  domainBackupAll.TruncateBackupTable is failed  tableName: %s  ::  %s\n", models.AliDomainTableName, err.Error())
				return err
			}
		}

	} else {
		err = fmt.Errorf("error:   int(response.TotalCount) != len(resDomains) ")
		return err
	}

	return nil

}

func BackupAliDomainToLocal(domains []*models.AliDomain) {
	var wgM sync.WaitGroup
	wgM.Add(len(domains))
	for i, _ := range domains {
		go func(index int) {
			if err := models.AddOrUpdateBackupDomain(domains[index]); err != nil {
				fmt.Printf("error:  AddOrUpdateBackupDomain is failed! domain:%s  ::  %s\n", domains[index].DomainName, err.Error())
				return
			}
			wgM.Done()
		}(i)
	}
	wgM.Wait()
}

func ImportToAliDomain (importReq *request.DomainForImportOtherCloudRequest)(errMsg string, err error){
	if importReq.Src == "TenCent"{
		errMsg = "导入Ali域名失败，请联系管理员"
		err = fmt.Errorf("导入Ali域名失败，请联系管理员")
		return
	}
	if importReq.Src == "ALi"{
		req := &ReqImportAliDomainRecord{
			KeyId:     importReq.DestKeyId,
			KeySecret: importReq.DestKeySecret,
			Domains:   importReq.AliDomain,
		}
		failedDomains, err := AliCloudImportDomain(req)
		if err != nil {
			if len(failedDomains) > 0 {
				fmt.Println("failed:  domainBackup.ImportDomain is failed  ::  ", failedDomains)
				errMsg = fmt.Sprintf("导入域名失败： %s\n", failedDomains)
				return errMsg, err
			}
			fmt.Println("error：  domainBackup.ImportDomain is failed  ::  ", err)
			return errMsg, err
		}
	}
	return
}

func GetAliDNS(req *ReqGetCloudDomain) (resp *RespCloudDomain, err error) {

	client, err := alidns.NewClientWithAccessKey("cn-qingdao", req.KeyId, req.KeySecret)
	if err != nil {
		fmt.Print("error:  alidns.NewClientWithAccessKey is failed  ::  ", err)
		err = fmt.Errorf("获取域名失败")
		return
	}

	createDescribeDomainsRequest := alidns.CreateDescribeDomainsRequest()
	createDescribeDomainsRequest.Scheme = "https"
	createDescribeDomainsRequest.PageSize = requests.NewInteger(req.PageSize)
	createDescribeDomainsRequest.PageNumber = requests.NewInteger(req.PageNumber)
	if req.SearchDomain != "" {
		createDescribeDomainsRequest.KeyWord = req.SearchDomain
		createDescribeDomainsRequest.SearchMode = "LIKE"
	}
	response, err := client.DescribeDomains(createDescribeDomainsRequest)
	if err != nil {
		fmt.Println("error:  DescribeDomains is failed  ::  ", err.Error())
		err = fmt.Errorf("获取域名失败")
		return
	}

	// get remote ali domain
	remote := &ReqRemoteAliDomain{
		response: response,
		client:   client,
		req:      req,
	}
	domains, err := GetRemoteAliDomains(remote)
	if err != nil {
		return
	}
	BackupAliDomainToLocal(domains)

	start := 0
	if req.PageNumber > 0 {
		start = (req.PageNumber - 1) * req.PageSize
	}

	res, err := models.GetBackupDomainByPageBreak(start, req.PageSize-1)
	if err != nil {
		fmt.Println("error:  domainBackup.GetBackupDomainByPageBreak  ::  ", err.Error())
		err = fmt.Errorf("获取本地备份域名失败")
		return nil, err
	}

	resp = new(RespCloudDomain)
	resp.Total = int(response.TotalCount)
	resp.PageNumber = int(response.PageNumber)
	resp.AliDomains = res
	return
}

func GetRemoteAliDomains(remote *ReqRemoteAliDomain) (domains []*models.AliDomain, err error) {

	importDomian := &AnalysisImportDomain{
		Resp:      remote.response,
		KeyId:     remote.req.KeyId,
		KeySecret: remote.req.KeySecret,
		Client:    remote.client,
	}
	domains, err = AliCloudAnalysisResponse(importDomian)
	if err != nil {
		fmt.Println("error:  domainBackup.AnalysisResponse  ::  ", err.Error())
		err = fmt.Errorf("获取域名失败")
		return
	}

	/*
		if remote.req.SearchDomain != ""{
			domain, err := models.GetBackupByDomainName(remote.req.SearchDomain)
			if err != nil{
				fmt.Println("error:  domainBackup.GetBackupDomainByPageBreak  ::  ", err.Error())
				err = fmt.Errorf("获取本地备份域名失败")
				return nil, err
			}
			resDomains = append(resDomains,domain)
		} else {
			start := 0
			if remote.req.PageNumber > 0 {
				start = (remote.req.PageNumber - 1) * remote.req.PageSize
			}
			resDomains, err = models.GetBackupDomainByPageBreak(start, remote.req.PageSize - 1)
			if err != nil{
				fmt.Println("error:  domainBackup.GetBackupDomainByPageBreak  ::  ", err.Error())
				err = fmt.Errorf("获取本地备份域名失败")
				return nil, err
			}
		}
	*/
	return
}

func AliCloudAnalysisResponse(importDomain *AnalysisImportDomain) (backupDomains []*models.AliDomain, err error) {

	var wgA sync.WaitGroup
	wgA.Add(len(importDomain.Resp.Domains.Domain))
	for i, _ := range importDomain.Resp.Domains.Domain {
		tmp := new(models.AliDomain)
		info, err := json.Marshal(importDomain.Resp.Domains.Domain[i])
		if err != nil {
			fmt.Printf("error:  AnalysisResponse.json.Marshal(Domian) is failed domainName: %s  ::  %s\n", importDomain.Resp.Domains.Domain[i].DomainName, err.Error())
			return nil, err
		}

		b, err := json.Marshal(importDomain.Resp.Domains.Domain[i].DnsServers.DnsServer)
		if err != nil {
			fmt.Printf("error:  AnalysisResponse.json.Marshal(DnsServer) is failed domainName: %s  ::  %s\n", importDomain.Resp.Domains.Domain[i].DomainName, err.Error())
			return nil, err
		}

		importDomainRecords := &ImportAliDomainRecords{
			keyId:      importDomain.KeyId,
			keySecret:  importDomain.KeySecret,
			domainName: importDomain.Resp.Domains.Domain[i].DomainName,
			client:     importDomain.Client,
		}
		record, err := AliCloudDomainRecords(importDomainRecords)
		if err != nil {
			fmt.Printf("error:  AnalysisResponse.DomainRecords is failed domainName: %s  ::  %s\n", importDomain.Resp.Domains.Domain[i].DomainName, err.Error())
			record = ""
		}

		tmp.Dns = string(b)
		tmp.DomainName = importDomain.Resp.Domains.Domain[i].DomainName
		tmp.InstanceEndTime = importDomain.Resp.Domains.Domain[i].InstanceEndTime
		tmp.DomainSource = importDomain.KeyId
		tmp.DomainInfo = string(info)
		tmp.DomainRecord = record
		backupDomains = append(backupDomains, tmp)
		wgA.Done()
	}
	wgA.Wait()

	return
}

func AliCloudDomainRecords(importDomainRecords *ImportAliDomainRecords) (domainRecords string, err error) {
	//client, err := alidns.NewClientWithAccessKey("cn-hangzhou",importDomainRecords.keyId, importDomainRecords.keySecret)

	recordsRequest := alidns.CreateDescribeDomainRecordsRequest()
	recordsRequest.Scheme = "https"
	recordsRequest.DomainName = importDomainRecords.domainName
	response, err := importDomainRecords.client.DescribeDomainRecords(recordsRequest)
	if err != nil {
		fmt.Println("error:  DescribeDomainRecords is failed  ::  ", err.Error())
		return "", err
	}
	domainRecordsJson, err := json.Marshal(response.DomainRecords.Record)
	if err != nil {
		fmt.Println("error:  json.Marshal(response.DomainRecords.Record) is failed  ::  ", err.Error())
		return "", err
	}

	return string(domainRecordsJson), err
}

func AliCloudImportDomain(req *ReqImportAliDomainRecord) (existsDomainLlist []string, err error) {

	client, err := alidns.NewClientWithAccessKey("cn-qingdao", req.KeyId, req.KeySecret)
	if err != nil {
		fmt.Println("error：  DomainImport.NewClientWithAccessKey is failed  ::  ", err)
		err = fmt.Errorf("导入域名失败，请联系管理员")
		return
	}
	for _, domain := range req.Domains {
		addDomainRequest := alidns.CreateAddDomainRequest()
		addDomainRequest.Scheme = "https"
		addDomainRequest.DomainName = domain.DomainName
		response := new(alidns.AddDomainResponse)

		response, err = client.AddDomain(addDomainRequest)
		if err != nil {
			if strings.Index(err.Error(), "domain name already exists") > 0 {
				fmt.Println("Failed: error：  DomainImport.AddDomain is failed  ::  ", addDomainRequest.DomainName)
				err = fmt.Errorf("域名已存在：%s \n", addDomainRequest.DomainName)
				existsDomainLlist = append(existsDomainLlist, addDomainRequest.DomainName)
			} else {
				fmt.Println("error：  DomainImport.AddDomain is failed  ::  ", err)
				err = fmt.Errorf("导入域名失败：%s，请联系管理员\n", addDomainRequest.DomainName)
				existsDomainLlist = append(existsDomainLlist, addDomainRequest.DomainName)
			}

			continue
		}
		if response.DomainName != "" {

			reqStrSlice := make([]*models.DomainBackupRecord, 0)
			if err := json.Unmarshal([]byte(domain.DomainRecord), &reqStrSlice); err != nil {
				fmt.Println("error：  DomainImport.RecordRequest unmarshal is failed  ::  ", err)
				err = fmt.Errorf("域名解析失败：%s，请联系管理员\n", addDomainRequest.DomainName)
				return existsDomainLlist, err
			}

			for _, item := range reqStrSlice {
				domainRecordReq := alidns.CreateAddDomainRecordRequest()
				domainRecordReq.Scheme = "https"
				domainRecordReq.RR = item.RR
				domainRecordReq.Line = item.Line
				domainRecordReq.Type = item.Type
				domainRecordReq.Value = item.Value
				domainRecordReq.DomainName = item.DomainName
				_, err := client.AddDomainRecord(domainRecordReq)
				if err != nil {
					fmt.Printf("error：  DomainImport.AddDomainRecord is failed, domain: %s  ::  %s\n", domainRecordReq.DomainName, err.Error())
					err = fmt.Errorf("导入域名解析失败：%s，请联系管理员\n", addDomainRequest.DomainName)
					return existsDomainLlist, err
				}
			}
		}
	}

	return
}
