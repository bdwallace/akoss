package TenCentDNS

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tencentDNSpod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"library/domainBackup"
)


func TenCentDNSPodDescribeDomainList(credential *common.Credential)(err error){
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := tencentDNSpod.NewClient(credential, "", cpf)

	request := tencentDNSpod.NewDescribeDomainListRequest()

	response, err := client.DescribeDomainList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}

	if err = domainBackup.AnalysisTenCentDNSpodDomainToDatabase(response); err != nil {
		return err
	}

	return nil
}

func DescribeDomain (credential *common.Credential, domain string, domainId int){

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := tencentDNSpod.NewClient(credential, "", cpf)

	request := tencentDNSpod.NewDescribeDomainRequest()

	request.Domain = common.StringPtr(domain)
	request.DomainId = common.Uint64Ptr(uint64(domainId))

	response, err := client.DescribeDomain(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())

}



func TenCentDNSPodCreateDomain(credential *common.Credential) {

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, _ := tencentDNSpod.NewClient(credential, "", cpf)

	request := tencentDNSpod.NewCreateDomainBatchRequest()

	request.DomainList = common.StringPtrs([]string{ "tcqp68.com" })

	response, err := client.CreateDomainBatch(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())

}
