package domainExImport

import (
	"controllers"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"library/CloudDNSDomain"
	"library/TenCentDNS"
	"models/request"

	"library/aliDNS"
	"models"
)

type DomainExImport struct {
	controllers.BaseController
}


// @Title search domain
// @Description search domain
// @Success 0 {id} int64
// @Failure 1 导出 domain 失败
// @Failure 2 User not found
// @router /searchAliDomain [get]
func (c *DomainExImport) SearchDomain() {


	src := c.GetString("src_class")
	keyId := c.GetString("src_key_id")
	keySecret := c.GetString("src_key_secret")
	domainName := c.GetString("domain_name")
	pageSize, err := c.GetInt("page_size", 20)
	if err != nil {
		fmt.Printf("error:  get page size is failed!  ::  %s\n", err.Error())
		c.SetJson(1, err, "获取 page size 失败")
		return
	}

	pageNumber, err := c.GetInt("page_number", 1)
	if err != nil {
		fmt.Printf("error:  get page number is failed!  ::  %s\n", err.Error())
		c.SetJson(1, err, "获取 page number 失败")
		return
	}
	if src == "" {
		fmt.Printf("error:  SearchDomain  ::  src class can't be empty")
		c.SetJson(1, err, "获取 page number 失败")
		return
	}
	searchReq := new(request.DomainForSearchRequest)
	searchReq.Src = src
	searchReq.SrcKeyId = keyId
	searchReq.SrcKeySecret = keySecret
	searchReq.SearchDomain = domainName
	searchReq.PageSize = pageSize
	searchReq.PageNum = pageNumber


	respSearch, err :=  CloudDNSDomain.DomainSearch(searchReq)
	if err != nil{
		fmt.Printf("error:  CloudDNSDomain.DomainSearch(searchReq) is failed!  ::  %s\n", err.Error())
		c.SetJson(1, err, "搜索 domain 失败")
		return
	}
	if resp, ok := respSearch.(*aliDNS.RespGetAliDNS); ok {
		resp.Class = "ALi"
		c.SetJson(0, map[string]interface{}{"item_total": resp.Total, "page_num": pageNumber, "domains": resp.Domains}, "")
		return
	}
	// if resp is tencent domain

	// if resp is cloudfare domain

}

// @Title 域名导入
// @Description alicloud domain import
// @Success 0 {id} int64
// @Failure 1 导入 domain 失败
// @Failure 2 User not found
// @router /domainImport [post]
func (c *DomainExImport) DomainImport() {

	src := c.GetString("src_class")
	dest := c.GetString("dest_class")
	destKeyId := c.GetString("dest_key_id")
	destKeySecret := c.GetString("dest_key_secret")
	if destKeyId == "" || destKeySecret == "" {
		c.SetJson(1, nil, "请输入目标 key id 和 key secret")
		return
	}
	importReq := new(request.DomainForImportOtherCloudRequest)
	importReq.Src = src
	importReq.Dest = dest
	importReq.DestKeyId = destKeyId
	importReq.DestKeySecret = destKeySecret


	if src == "ALi"	{
		importAliDomains := make([]*models.AliDomain, 0)
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &importAliDomains)
		if err != nil {
			fmt.Println("error：  DomainImport.json.Unmarshal is failed  ::  ", err)
			c.SetJson(1, err, "解析数据失败，请联系管理员")
			return
		}

		importReq.AliDomain = make([]*models.AliDomain,0)
		importReq.AliDomain = append(importReq.AliDomain,importAliDomains...)
	}
	if src == "TenCent" {

	}
	if src == "CloudFare" {

	}

	msg, err := CloudDNSDomain.DomainImportToCloud(importReq)
	if err != nil{
		if errMsg, ok := msg.(string); ok {
			c.SetJson(1, err, errMsg)
		} else {
			c.SetJson(1, err, "域名导入失败，请联系管理员")
		}
		return
	}

	c.SetJson(0, nil, "")
	return
}

// @Title 域名全量备份
// @Description alicloud domain backup all
// @Success 0 {id} int64
// @Failure 1 导出 domain 失败
// @Failure 2 User not found
// @router /domainBackupAll [get]
func (c *DomainExImport) DomainBackupAll() {

	src := c.GetString("src_class")
	keyId := c.GetString("src_key_id")
	keySecret := c.GetString("src_key_secret")

	backupReq := new(request.DomainForBackUpTolocalRequest)
	backupReq.Src = src
	backupReq.SrcKeyId = keyId
	backupReq.SrcKeySecret = keySecret

	err := CloudDNSDomain.DomainBackUpToLocal(backupReq)
	if err != nil{
		fmt.Println("Success：  成功备份全部域名")
		c.SetJson(1, err, "全量备份到本地失败，请联系管理员")
		return
	}else {
		fmt.Println("Success：  全量备份成功")
		c.SetJson(0, nil, "全量备份成功")
	}
	return
}

/*
// @Title 域名增量备份
// @Description alicloud domain backup incr
// @Success 0 {id} int64
// @Failure 1 导出 domain 失败
// @Failure 2 User not found
// @router /domainBackupIncr [post]
func (c *DomainExImport) DomainBackupIncr() {
	keyId := c.GetString("export_key_id")
	keySecret := c.GetString("export_key_secret")
	backupDomains := make([]*models.AliDomain, 0)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &backupDomains)
	if err != nil {
		fmt.Println("error：  DomainImport.json.Unmarshal is failed  ::  ", err)
		c.SetJson(1, err, "解析数据失败，请联系管理员")
		return
	}

	client, err := alidns.NewClientWithAccessKey("cn-qingdao", keyId, keySecret)
	if err != nil {
		fmt.Print("error:  alidns.NewClientWithAccessKey is failed  ::  ", err)
		err = fmt.Errorf("获取域名失败")
		return
	}

	request := alidns.CreateDescribeDomainsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(10)
	request.PageNumber = requests.NewInteger(1)
	for _, d := range backupDomains {
		if d.DomainName != "" {
			request.KeyWord = d.DomainName
			request.SearchMode = "LIKE"
		}
		response, err := client.DescribeDomains(request)
		if err != nil {
			fmt.Println("error:  DescribeDomains is failed  ::  ", err.Error())
			err = fmt.Errorf("获取域名失败")
			return
		}

		importDomian := &domainBackup.AnalysisImportDomain{
			Resp:      response,
			KeyId:     keyId,
			KeySecret: keySecret,
			Client:    client,
		}
		domains, err := domainBackup.AliCloudAnalysisResponse(importDomian)
		if err != nil {
			fmt.Println("error:  DomainBackupIncr.AnalysisResponse  ::  ", err.Error())
			err = fmt.Errorf("获取域名失败")
			return
		}
		for _, updateDomain := range domains {
			if err := models.AddOrUpdateBackupDomain(updateDomain); err != nil {
				fmt.Printf("error:  DomainBackupIncr.AddOrUpdateBackupDomain is failed! domain:%s  ::  %s\n", updateDomain.DomainName, err.Error())
				c.SetJson(1, err, "域名更新到本地失败,请联系管理员")
				return
			}
		}
	}

	resDomains := make([]string, 0)
	for _, resD := range backupDomains {
		resDomains = append(resDomains, resD.DomainName)
	}

	c.SetJson(0, resDomains, "")
	return

}
*/


// @Title TenCentCloudDNS Get Domain List
// @Description TenCentCloudDNS Get Domain List
// @Success 0 {id} int64
// @Failure 1 获取 TenCentCloudDNS Domain List 失败
// @Failure 2 User not found
// @router /tenCentDomain [get]
func (c * DomainExImport)TenCentCloutDomainList(){

	credential := common.NewCredential(
		"AKIDXv50g6GAoDMFsGPKzOXQ3QXOzKPEm46s",
		"Zty83kBNIxMVO1OWP8crsd8NNI0WqwmN",
	)
	TenCentDNS.TenCentDNSPodDescribeDomainList(credential)
	fmt.Println("end")
	c.SetJson(0,nil, "")
	return
}



// @Title TenCentCloudDNS Post Created Domain
// @Description TenCentCloudDNS Post Created Domain
// @Success 0 {id} int64
// @Failure 1 添加 TenCentCloudDNS Domain 失败
// @Failure 2 User not found
// @router /tenCentDomain [post]
func (c * DomainExImport)CreateTenCentCloutDomain(){

	credential := common.NewCredential(
		"AKIDXv50g6GAoDMFsGPKzOXQ3QXOzKPEm46s",
		"Zty83kBNIxMVO1OWP8crsd8NNI0WqwmN",
	)
	TenCentDNS.TenCentDNSPodCreateDomain(credential)
	fmt.Printf("end")
	c.SetJson(0,nil, "")
	return
}


// @Title TenCentCloudDNS Get Domain Details
// @Description TenCentCloudDNS Get Domain Details
// @Success 0 {id} int64
// @Failure 1 获取 TenCentCloudDNS Domain Details 失败
// @Failure 2 User not found
// @router /tenCentDomainDescribe [get]
func (c * DomainExImport)TenCentDescribeDomain(){

	credential := common.NewCredential(
		"AKIDXv50g6GAoDMFsGPKzOXQ3QXOzKPEm46s",
		"Zty83kBNIxMVO1OWP8crsd8NNI0WqwmN",
	)
/*	domain := c.GetString("domain")
	domainId, err := c.GetInt("domain_id")
	if err != nil{
		return
	}*/

	//TenCentDNS.DescribeDomain(credential, domain, domainId)
	TenCentDNS.DescribeDomain(credential, "fengniao668.com", 91281812)
	fmt.Printf("end")
	c.SetJson(0,nil, "")
	return
}






