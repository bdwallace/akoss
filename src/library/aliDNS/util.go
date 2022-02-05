package aliDNS

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"library/domainBackup"
	"models"
	"sync"
)

type ReqGetAliDNS struct {
	PageSize     int
	PageNumber   int
	KeyId        string
	KeySecret    string
	SearchDomain string
}

type RespGetAliDNS struct {
	PageNumber 	int
	Total      	int
	Class 		string
	Domains  	[]*models.AliDomain
}

type ReqRemoteAliDomain struct {
	response *alidns.DescribeDomainsResponse
	client   *alidns.Client
	req      *ReqGetAliDNS
}

func GetAliDNS(req *ReqGetAliDNS) (resp *RespGetAliDNS, err error) {

	client, err := alidns.NewClientWithAccessKey("cn-qingdao", req.KeyId, req.KeySecret)
	if err != nil {
		fmt.Print("error:  alidns.NewClientWithAccessKey is failed  ::  ", err)
		err = fmt.Errorf("获取域名失败")
		return
	}

	request := alidns.CreateDescribeDomainsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(req.PageSize)
	request.PageNumber = requests.NewInteger(req.PageNumber)
	if req.SearchDomain != "" {
		request.KeyWord = req.SearchDomain
		request.SearchMode = "LIKE"
	}
	response, err := client.DescribeDomains(request)
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

	resp = new(RespGetAliDNS)
	resp.Total = int(response.TotalCount)
	resp.PageNumber = int(response.PageNumber)
	resp.Domains = res
	return
}

func GetRemoteAliDomains(remote *ReqRemoteAliDomain) (domains []*models.AliDomain, err error) {

	importDomian := &domainBackup.AnalysisImportDomain{
		Resp:      remote.response,
		KeyId:     remote.req.KeyId,
		KeySecret: remote.req.KeySecret,
		Client:    remote.client,
	}
	domains, err = domainBackup.AliCloudAnalysisResponse(importDomian)
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
