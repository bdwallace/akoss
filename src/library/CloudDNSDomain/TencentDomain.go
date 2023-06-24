package CloudDNSDomain

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tencentDNSpod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"models"
	"models/request"
	"strings"
)

func SearchTenCentDomain (searchReq *request.DomainForSearchRequest)(*RespCloudDomain, error){
	resp := new(RespCloudDomain)
	var err error
	searchDomain := make([]*models.TencentDomain,0)

	if searchReq.SearchDomain == "" {
		// 获取全部
		credential := common.NewCredential(
			searchReq.SrcKeyId,
			searchReq.SrcKeySecret,
		)
		if err = TenCentDNSPodDescribeDomainList(credential); err != nil {
			return nil, err
		}
		start := 0
		if searchReq.PageNum > 0 {
			start = (searchReq.PageNum - 1) * searchReq.PageSize
		}

		resp.TenCentDomains, err = models.GetAllTenCentDNSPodDomains(start, searchReq.PageSize - 1)
		if err != nil{
			return nil, err
		}
		resp.Total = len(resp.TenCentDomains)
		resp.PageNumber = searchReq.PageNum

	} else {
		// 搜索本地
		searchDomain, err = models.GetTenCentDNSPodDomainsByDomainName(searchReq.SearchDomain)
		if err != nil {
			// 域名不存在
			fmt.Printf("error:  GetTenCentDNSPodDomainsByDomainName is failed! domian:%s  ::  %s\n", searchReq.SearchDomain, err.Error())
			return nil, err
		}
		resp.TenCentDomains = append(resp.TenCentDomains, searchDomain...)
		resp.Total = len(searchDomain)
		resp.PageNumber = 1
	}

	return resp, nil
}

func BackUpOrUpdateToLocalTenCentDomain (backupReq *request.DomainForBackUpTolocalRequest)(err error){
	credential := new(common.Credential)
	credential.SecretId = backupReq.SrcKeyId
	credential.SecretKey = backupReq.SrcKeySecret
	return TenCentDNSPodDescribeDomainList(credential)
}

func ImportToTenCentDomain (importReq *request.DomainForImportOtherCloudRequest)(errMsg string, err error){

	req, err :=AnalysisRecordToTenCentDomainRecord(importReq)
	if err != nil {
		errMsg := fmt.Sprintf("域名解析或导入失败")
		return errMsg, err
	}
	errMsg, err = TenCentImportDomain(req)
	if err != nil || errMsg != ""{
		fmt.Printf("error:  ImportToTenCentDomain.TenCentImportDomain is failed  ::  %s\n",errMsg)
	}
	return errMsg,err
}

func TenCentDNSPodDescribeDomainList(credential *common.Credential)(err error){
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := tencentDNSpod.NewClient(credential, "", cpf)

	listRequest := tencentDNSpod.NewDescribeDomainListRequest()

	response, err := client.DescribeDomainList(listRequest)
	if err != nil {
		fmt.Printf("error:  client.DescribeDomainList is failed  ::  %s", err)
		return
	}

	if err = AnalysisTenCentDNSpodDomainToDatabase(response, credential); err != nil {
		return err
	}

	return nil
}

func TenCentDomainDescribeRecordList(credential *common.Credential, domain string)(domainDescribeJson string, err error){

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := tencentDNSpod.NewClient(credential, "", cpf)

	recordListRequest := tencentDNSpod.NewDescribeRecordListRequest()
	recordListRequest.Domain = common.StringPtr(domain)
	response, err := client.DescribeRecordList(recordListRequest)
	if err != nil {
		fmt.Printf("error:  client.DescribeRecordList is failed  ::  %s", err)
		return "", err
	}

	byteDomainRecord, err := json.Marshal(response.Response.RecordList)
	if err != nil{
		return "", err
	}
	domainDescribeJson = string(byteDomainRecord)

	return
}

func TenCentDNSPodCreateDomain(credential *common.Credential) {

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := tencentDNSpod.NewClient(credential, "", cpf)

	batchRequest := tencentDNSpod.NewCreateDomainBatchRequest()

	batchRequest.DomainList = common.StringPtrs([]string{"tcqp68.com" })

	response, err := client.CreateDomainBatch(batchRequest)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())

}

func AnalysisTenCentDNSpodDomainToDatabase(domain *tencentDNSpod.DescribeDomainListResponse, credential *common.Credential)(err error){

	for _, item := range domain.Response.DomainList{
		tencentDomain := new(models.TencentDomain)
		domainInfoByte, err := json.Marshal(item)
		if err != nil{
			fmt.Println("error:  AnalysisTenCentDNSpodDomainToDatabase is failed  ::  ",err)
			return err
		}
		//domainDesceribeJson, err := DescribeDomain(credential, *item.Name, int(*item.DomainId))
		desceribeRecordJson, err := TenCentDomainDescribeRecordList(credential, *item.Name)
		if err != nil{
			fmt.Println("error:  TenCentDNS.DescribeDomain is failed  ::  ",err)
			tencentDomain.DomainRecord = "NULL"
		}

		var strEffectiveDNS string
		for i, dns := range item.EffectiveDNS{
			if i % 2 == 1 {
				strEffectiveDNS += ", "
			}
			strEffectiveDNS += *dns
		}
		tencentDomain.DomainName = *item.Name
		tencentDomain.Dns = strEffectiveDNS
		tencentDomain.DomainSource = *item.Owner
		tencentDomain.DomainInfo = string(domainInfoByte)
		tencentDomain.DomainRecord = desceribeRecordJson

		err = models.AddOrUpdateTenCentDNSPodDomain(tencentDomain)
		if err != nil{
			return err
		}
	}
	return nil
}

func TenCentImportDomain(req *ReqImportTenCentDomainRecord)(errMsg string, err error){

	credential := new(common.Credential)
	credential.SecretId = req.KeyId
	credential.SecretKey = req.KeySecret
	domains := make([]string, 0)
	for _, item := range req.Domains {
		domains = append(domains, item.DomainName)
	}

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := tencentDNSpod.NewClient(credential, "", cpf)
	domainBatchRequest := tencentDNSpod.NewCreateDomainBatchRequest()
	domainBatchRequest.DomainList = common.StringPtrs(domains)

	_, err = client.CreateDomainBatch(domainBatchRequest)
	if err != nil{
		fmt.Println("error:  TenCentImportDomain.CreateDomainBatch is failed  ::  ",err)
		errMsg = "域名添加失败 "
		for _, item := range req.Domains{
			errMsg += item.DomainName
		}
		return errMsg, err
	}

	for _, domain := range req.Domains {
		for _, item := range domain.Record{
			err = AddTenCentRecord(credential, item)
			if err != nil{
				errInfo := err.Error()
				exists := strings.Index(errInfo,"记录已经存在")
				if exists > 0 {
					domainIndex := strings.Index(errMsg,domain.DomainName)
					if domainIndex > 0 {
						continue
					}
					errMsg += " " + domain.DomainName + " 记录已存在,"
				}
			}
		}
	}

	return errMsg, err
}

func AddTenCentRecord(credential *common.Credential,domain *DomainRecord)(err error){

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := tencentDNSpod.NewClient(credential, "", cpf)
	createRecordRequest := tencentDNSpod.NewCreateRecordRequest()
	createRecordRequest.Domain = common.StringPtr(domain.DomainName)
	createRecordRequest.RecordType = common.StringPtr(domain.RecordType)
	createRecordRequest.RecordLine = common.StringPtr(domain.RecordLine)
	createRecordRequest.Value = common.StringPtr(domain.RecordValue)

	_, err = client.CreateRecord(createRecordRequest)
	if err != nil{
		fmt.Printf("error：  DomainImport.AddTenCentRecord is failed, domain: %s  ::  %s\n", domain.DomainName, err.Error())
		return err
	}

	return
}

func AnalysisRecordToTenCentDomainRecord(importReq *request.DomainForImportOtherCloudRequest)(recordReq *ReqImportTenCentDomainRecord,err error){

	recordReq = new(ReqImportTenCentDomainRecord)
	recordReq.Domains = make([]*DomainInfo,0)

	recordReq.KeyId = importReq.DestKeyId
	recordReq.KeySecret = importReq.DestKeySecret

	if importReq.Src == "ALi"{
		for _, item := range importReq.AliDomain{
			aliRecord := make([]*models.DomainBackupRecord,0)
			tmp := new(DomainInfo)
			err = json.Unmarshal([]byte(item.DomainRecord), &aliRecord)
			if err != nil{
				fmt.Println("error:  AnalysisRecordToTenCentDomainRecord.Unmarshal is failed  ::  ",err)
				return nil, err
			}
			tmp.DomainName = item.DomainName
			tmp.Record = make([]*DomainRecord,0)
			for _, recordItem := range aliRecord {
				tmpRecord := new(DomainRecord)
				tmpRecord.DomainName = recordItem.DomainName
				tmpRecord.RecordType = recordItem.Type
				tmpRecord.RecordLine = recordItem.Line
				tmpRecord.RecordValue = recordItem.Value
				tmp.Record = append(tmp.Record, tmpRecord)
			}

			recordReq.Domains = append(recordReq.Domains, tmp)
		}
	}

	if importReq.Src == "TenCent"{
		for _, item := range importReq.TenCentDomain{
			tmp := new(DomainInfo)
			tenCentRecord := make([]*models.TenCentRecordListItem,0)
			err = json.Unmarshal([]byte(item.DomainRecord),&tenCentRecord)
			if err != nil{
				fmt.Println("error：  DomainImport.AddTenCentRecord unmarshal is failed  ::  ", err)
				err = fmt.Errorf("域名解析失败：%s，请联系管理员\n", item.DomainName)
				return nil, err
			}
			tmp.DomainName = item.DomainName
			tmp.Record = make([]*DomainRecord, 0)
			for _, recordItem := range tenCentRecord {
				tmpRecord := new(DomainRecord)
				tmpRecord.DomainName = item.DomainName
				tmpRecord.RecordType = *recordItem.Type
				tmpRecord.RecordLine = *recordItem.Line
				tmpRecord.RecordValue = *recordItem.Value
				tmp.Record = append(tmp.Record, tmpRecord)
			}
			recordReq.Domains = append(recordReq.Domains, tmp)
		}
	}


	if importReq.Src == "CloudFare" {

	}

	return
}