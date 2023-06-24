package CloudDNSDomain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"models"
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

type ReqImportTenCentDomainRecord struct {
	KeyId     		string
	KeySecret 		string
	Domains 		[]*DomainInfo
}
type DomainInfo struct {
	DomainName		string
	Record 			[]*DomainRecord
}
type DomainRecord struct {
	DomainName		string
	RecordType 		string
	RecordLine		string
	RecordValue		string
}

/*
	createRecordRequest.Domain = common.StringPtr(domain.DomainName)
	createRecordRequest.RecordType = common.StringPtr(*domainRecord.Type)
	createRecordRequest.RecordLine = common.StringPtr(*domainRecord.Line)
	createRecordRequest.Value = common.StringPtr(*domainRecord.Value)
 */


type ReqGetCloudDomain struct {
	PageSize     int
	PageNumber   int
	KeyId        string
	KeySecret    string
	SearchDomain string
}

type RespCloudDomain struct {
	PageNumber 			int
	Total      			int
	AliDomains 			[]*models.AliDomain
	TenCentDomains		[]*models.TencentDomain
}

type ReqRemoteAliDomain struct {
	response *alidns.DescribeDomainsResponse
	client   *alidns.Client
	req      *ReqGetCloudDomain
}


