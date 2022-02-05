package domainBackup

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	tencentDNSpod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"models"
	"strings"
	"sync"
)

type AnalysisImportDomain struct {
	Resp      *alidns.DescribeDomainsResponse
	KeyId     string
	KeySecret string
	Client    *alidns.Client
}

type ImportAliDomainRecords struct {
	keyId      string
	keySecret  string
	domainName string
	client     *alidns.Client
}

type ReqImportAliDomainRecord struct {
	KeyId     string
	KeySecret string
	Domains   []*models.AliDomain
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

		tmp.DnsServer = string(b)
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

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"
	request.DomainName = importDomainRecords.domainName
	response, err := importDomainRecords.client.DescribeDomainRecords(request)
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
		request := alidns.CreateAddDomainRequest()
		request.Scheme = "https"
		request.DomainName = domain.DomainName
		response := new(alidns.AddDomainResponse)

		response, err = client.AddDomain(request)
		if err != nil {
			if strings.Index(err.Error(), "domain name already exists") > 0 {
				fmt.Println("Failed: error：  DomainImport.AddDomain is failed  ::  ", request.DomainName)
				err = fmt.Errorf("域名已存在：%s \n", request.DomainName)
				existsDomainLlist = append(existsDomainLlist, request.DomainName)
			} else {
				fmt.Println("error：  DomainImport.AddDomain is failed  ::  ", err)
				err = fmt.Errorf("导入域名失败：%s，请联系管理员\n", request.DomainName)
				existsDomainLlist = append(existsDomainLlist, request.DomainName)
			}

			continue
		}
		if response.DomainName != "" {

			reqStrSlice := make([]*models.DomainBackupRecord, 0)
			if err := json.Unmarshal([]byte(domain.DomainRecord), &reqStrSlice); err != nil {
				fmt.Println("error：  DomainImport.RecordRequest unmarshal is failed  ::  ", err)
				err = fmt.Errorf("域名解析失败：%s，请联系管理员\n", request.DomainName)
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
					err = fmt.Errorf("导入域名解析失败：%s，请联系管理员\n", request.DomainName)
					return existsDomainLlist, err
				}
			}
		}
	}

	return
}



func AnalysisTenCentDNSpodDomainToDatabase(domain *tencentDNSpod.DescribeDomainListResponse)(err error){

	for _, item := range domain.Response.DomainList{
		tmp := new(models.Tencentdnspod)
		var strEffectiveDNS string
		for i, effe := range item.EffectiveDNS{
			if i > 0 {
				strEffectiveDNS += ", "
			}
			strEffectiveDNS += *effe
		}
		tmp.DomainId = int(*item.DomainId)
		tmp.Name = *item.Name
		tmp.Status = *item.Status
		tmp.TTL = int(*item.TTL)
		tmp.CNAMESpeedup = *item.CNAMESpeedup
		tmp.DNSStatus = *item.DNSStatus
		tmp.Grade = *item.Grade
		tmp.GroupId = int(*item.GroupId)
		tmp.SearchEnginePush = *item.SearchEnginePush
		tmp.Remark = *item.Remark
		tmp.Punycode = *item.Punycode
		tmp.EffectiveDNS = strEffectiveDNS
		tmp.GradeLevel = int(*item.GradeLevel)
		tmp.GradeTitle = *item.GradeTitle
		tmp.IsVip = *item.IsVip
		tmp.VipStartAt = *item.VipStartAt
		tmp.VipEndAt = *item.VipEndAt
		tmp.VipAutoRenew = *item.VipAutoRenew
		tmp.RecordCount = int(*item.RecordCount)
		tmp.CreatedOn = *item.CreatedOn
		tmp.UpdatedOn = *item.UpdatedOn
		tmp.Owner = *item.Owner

		err = models.AddOrUpdateTenCentDNSPod(tmp)
		if err != nil{
			return
		}
	}
	return
}