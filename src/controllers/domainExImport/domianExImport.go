package domainExImport

import (
	"controllers"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"library/aliDNS"
	"library/domainBackup"
	"models"
)

type DomainExImport struct {
	controllers.BaseController
}

// @Title 域名导出
// @Description alicloud domain export
// @Success 0 {id} int64
// @Failure 1 导出 domain 失败
// @Failure 2 User not found
// @router /searchAliDomain [get]
func (c *DomainExImport) DomainExport() {

	keyId := c.GetString("export_key_id")
	keySecret := c.GetString("export_key_secret")
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

	resp := new(aliDNS.RespGetAliDNS)

	if domainName == "" {
		// 获取全部
		req := new(aliDNS.ReqGetAliDNS)
		req.KeyId = keyId
		req.KeySecret = keySecret
		req.SearchDomain = domainName
		req.PageSize = pageSize
		req.PageNumber = pageNumber
		resp, err = aliDNS.GetAliDNS(req)
		if err != nil {
			c.SetJson(1, err, err.Error())
			return
		}
	} else {
		// 搜索获取本地
		searchDomain, err := models.GetBackupByDomainName(domainName)
		if err != nil {
			fmt.Printf("error:  GetBackupByDomainName is failed! domian:%s  ::  %s\n", domainName, err.Error())
			c.SetJson(1, err, "该域名不存在")
			return
		}
		resp.Domains = append(resp.Domains, searchDomain)
		resp.Total = 1
		pageNumber = 1
	}

	c.SetJson(0, map[string]interface{}{"item_total": resp.Total, "page_num": pageNumber, "domains": resp.Domains}, "")
	return
}

// @Title 域名导入
// @Description alicloud domain import
// @Success 0 {id} int64
// @Failure 1 导入 domain 失败
// @Failure 2 User not found
// @router /domainImport [post]
func (c *DomainExImport) DomainImport() {

	keyId := c.GetString("dest_key_id")
	keySecret := c.GetString("dest_key_secret")
	if keyId == "" || keySecret == "" {
		c.SetJson(1, nil, "请输入目标 key id 和 key secret")
		return
	}

	importDomains := make([]*models.DomainBackup, 0)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &importDomains)
	if err != nil {
		fmt.Println("error：  DomainImport.json.Unmarshal is failed  ::  ", err)
		c.SetJson(1, err, "解析数据失败，请联系管理员")
		return
	}
	req := &domainBackup.ReqImportDomainRecord{
		KeyId:     keyId,
		KeySecret: keySecret,
		Domains:   importDomains,
	}
	failedDomains, err := domainBackup.ImportDomain(req)
	if err != nil {
		if len(failedDomains) > 0 {
			fmt.Println("failed:  domainBackup.ImportDomain is failed  ::  ", failedDomains)
			resp := fmt.Sprintf("导入域名失败： %s\n", failedDomains)
			c.SetJson(1, nil, resp)
			return
		}

		fmt.Println("error：  domainBackup.ImportDomain is failed  ::  ", err)
		c.SetJson(1, err, err.Error())
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

	keyId := c.GetString("export_key_id")
	keySecret := c.GetString("export_key_secret")

	client, err := alidns.NewClientWithAccessKey("cn-qingdao", keyId, keySecret)
	if err != nil {
		fmt.Print("error:  alidns.NewClientWithAccessKey is failed  ::  ", err)
		c.SetJson(1, err, "获取域名失败")
		return
	}

	remoteDomains := make([]*models.DomainBackup, 0)
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
			c.SetJson(1, err, "获取域名失败")
			return
		}
		remoteTotal = int(response.TotalCount)

		backupDomian := &domainBackup.AnalysisImportDomain{
			Resp:      response,
			KeyId:     keyId,
			KeySecret: keySecret,
			Client:    client,
		}

		domains, err := domainBackup.AnalysisResponse(backupDomian)
		if err != nil {
			fmt.Println("error:  domainBackupAll.AnalysisResponse  ::  ", err.Error())
			c.SetJson(1, err, "获取域名失败")
			return
		}
		remoteDomains = append(remoteDomains, domains...)
	}

	for _, updateDomain := range remoteDomains {
		if err = models.AddOrUpdateBackupDomain(updateDomain); err != nil {
			fmt.Println("error:  domainBackupAll.AddOrUpdateBackupDomain  ::  ", err.Error())
			c.SetJson(1, err, "更新本地数据失败")
			return
		}
	}

	if int(response.TotalCount) == len(remoteDomains) {

		// 备份，清空，
		tableName, err := models.CreateTableForBackupTheDomainWithAccKeyAndDate(keyId)
		if err != nil {
			fmt.Println("error:  CreateTableForBackupTheDomainWithAccKeyAndDate  ::  ", err.Error())
			c.SetJson(1, err, "本地创建备份表失败")
			return
		}

		oldCount, err := models.CountTempDomainBackup(tableName)
		if err != nil {
			fmt.Printf("error:  domainBackupAll.CountTempDomainBackup is failed  ::  %s\n", err.Error())
			c.SetJson(1, err, "域名未完全备份,请联系管理员")
			return
		}
		if oldCount > 1 {
			if err = models.TruncateBackupTable(tableName); err != nil {
				fmt.Printf("error:  domainBackup.TruncateBackupTable is failed  tableName: %s ::  %s\n", tableName, err.Error())
			}
		}
		if err = models.InsertTempDomainBackupFormBackup(tableName); err != nil {

			fmt.Printf("error:  domainBackupAll.InsertTempDomainBackupFormBackup  table : %s ::  %s\n", tableName, err.Error())
			c.SetJson(1, err, "域名备份失败")
			return
		}
		newCount, err := models.CountTempDomainBackup(tableName)
		if err != nil {
			fmt.Printf("error:  domainBackupAll.CountTempDomainBackup is failed  ::  %s\n", err.Error())
			c.SetJson(1, err, "域名未完全备份,请联系管理员")
			return
		}
		countBackupTable, err := models.CountTempDomainBackup(models.DomainBackupTableName)
		if err != nil {
			fmt.Printf("error:  domainBackupAll.t_domain_backup is failed  ::  %s\n", err.Error())
			c.SetJson(1, err, "域名未完全备份,请联系管理员")
			return
		}
		if newCount >= len(remoteDomains) && newCount == countBackupTable {
			// truncate
			if err = models.TruncateBackupTable(models.DomainBackupTableName); err != nil {
				fmt.Printf("error:  domainBackupAll.TruncateBackupTable is failed  tableName: %s  ::  %s\n", models.DomainBackupTableName, err.Error())
				c.SetJson(1, err, "域名未完全备份,请联系管理员")
				return
			}
		}

	} else {
		fmt.Printf("error:   int(response.TotalCount) != len(resDomains) ")
		c.SetJson(1, err, "域名备份数量不符,请联系管理员")
		return
	}

	fmt.Println("Success：  成功备份全部域名")
	c.SetJson(0, nil, "备份成功")
	return
}

// @Title 域名增量备份
// @Description alicloud domain backup incr
// @Success 0 {id} int64
// @Failure 1 导出 domain 失败
// @Failure 2 User not found
// @router /domainBackupIncr [post]
func (c *DomainExImport) DomainBackupIncr() {
	keyId := c.GetString("export_key_id")
	keySecret := c.GetString("export_key_secret")
	backupDomains := make([]*models.DomainBackup, 0)
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
		domains, err := domainBackup.AnalysisResponse(importDomian)
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
