package domainExImport

import (
	"controllers"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"library/domainBackup"
	"models"
	"sync"
)

type DomainExImport struct {
	controllers.BaseController
}


// @Title 域名导出
// @Description alicloud domain export
// @Success 0 {id} int64
// @Failure 1 导出 domain 失败
// @Failure 2 User not found
// @router /domainExport [get]
func (c *DomainExImport)DomainExport() {

	keyId := c.GetString("export_key_id")
	keySecret := c.GetString("export_key_secret")
	pageSize, err := c.GetInt("page_size",20)
	start := 0

	if err != nil{
		fmt.Printf("error:  get page size is failed!  ::  %s\n",err.Error())
		c.SetJson(1,err,"获取 page size 失败")
		return
	}

	pageNumber, err := c.GetInt("page_number",1)
	if err != nil{
		fmt.Printf("error:  get page number is failed!  ::  %s\n",err.Error())
		c.SetJson(1,err,"获取 page number 失败")
		return
	}

	client, err := alidns.NewClientWithAccessKey("cn-qingdao", keyId, keySecret)
	if err != nil{
		fmt.Print("error:  alidns.NewClientWithAccessKey is failed  ::  ",err)
		c.SetJson(1,err,"获取域名失败")
		return
	}

	request := alidns.CreateDescribeDomainsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(pageSize)
	request.PageNumber = requests.NewInteger(pageNumber)
	response, err := client.DescribeDomains(request)
	if err != nil {
		fmt.Println("error:  DescribeDomains is failed  ::  ", err.Error())
		c.SetJson(1, err, "获取域名失败")
		return
	}

	importDomian := &domainBackup.AnalysisImportDomain{
		Resp: response,
		KeyId: keyId,
		KeySecret: keySecret,
		Client: client,
	}
	domains, err := domainBackup.AnalysisResponse(importDomian)
	if err != nil{
		fmt.Println("error:  domainBackup.AnalysisResponse  ::  ", err.Error())
		c.SetJson(1, err, "获取域名失败")
		return
	}

	var wgM sync.WaitGroup
	wgM.Add(len(domains))
	for i, _ := range domains {
		go func(index int) {
			if err := models.AddOrUpdateBackupDomain(domains[index]); err != nil{
				fmt.Printf("error:  AddOrUpdateBackupDomain is failed! domain:%s  ::  %s\n",domains[index].DomainName, err.Error())
			}
			wgM.Done()
		}(i)
	}
	wgM.Wait()

	if pageNumber > 0 {
		start = (pageNumber - 1) * pageSize
	}
	localBackupDomains, err := models.GetBackupDomainByPageBreak(start, pageSize - 1)
	if err != nil{
		fmt.Println("error:  domainBackup.GetBackupDomainByPageBreak  ::  ", err.Error())
		c.SetJson(1, err, "获取本地备份域名失败")
		return
	}

	total := response.TotalCount
	pageNumber = int(response.PageNumber)

	c.SetJson(0,map[string]interface{}{"item_total": total,"page_num": pageNumber, "domains": localBackupDomains},"")
	return
}



// @Title 域名导入
// @Description alicloud domain import
// @Success 0 {id} int64
// @Failure 1 导入 domain 失败
// @Failure 2 User not found
// @router /domainImport [post]
func (c *DomainExImport)DomainImport() {

	keyId := c.GetString("import_key_id")
	keySecret := c.GetString("import_key_secret")
	importDomains := make([]*models.DomainBackup,0)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &importDomains)
	if err != nil{
		fmt.Println("error：  DomainImport.json.Unmarshal is failed  ::  ",err)
		c.SetJson(1, err, "解析数据失败，请联系管理员")
		return
	}
	req := &domainBackup.ReqImportDomainRecord{
		KeyId: keyId,
		KeySecret: keySecret,
		Domains: importDomains,
	}
	failedDomains, err := domainBackup.ImportDomain(req)
	if err != nil{
		fmt.Println("error：  domainBackup.ImportDomain is failed  ::  ",err)
		c.SetJson(1, err, err.Error())
		return
	}
	if len(failedDomains) > 0 {
		fmt.Println("failed:  domainBackup.ImportDomain is failed  ::  ",failedDomains)
		resp := fmt.Sprintf("域名已存在： %s\n",failedDomains)
		c.SetJson(1, nil, resp)
		return
	}

	c.SetJson(0,nil,"")
	return
}