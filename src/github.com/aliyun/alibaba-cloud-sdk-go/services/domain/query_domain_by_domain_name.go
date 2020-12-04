package domain

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryDomainByDomainName invokes the domain.QueryDomainByDomainName API synchronously
// api document: https://help.aliyun.com/api/domain/querydomainbydomainname.html
func (client *Client) QueryDomainByDomainName(request *QueryDomainByDomainNameRequest) (response *QueryDomainByDomainNameResponse, err error) {
	response = CreateQueryDomainByDomainNameResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainByDomainNameWithChan invokes the domain.QueryDomainByDomainName API asynchronously
// api document: https://help.aliyun.com/api/domain/querydomainbydomainname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainByDomainNameWithChan(request *QueryDomainByDomainNameRequest) (<-chan *QueryDomainByDomainNameResponse, <-chan error) {
	responseChan := make(chan *QueryDomainByDomainNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainByDomainName(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryDomainByDomainNameWithCallback invokes the domain.QueryDomainByDomainName API asynchronously
// api document: https://help.aliyun.com/api/domain/querydomainbydomainname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainByDomainNameWithCallback(request *QueryDomainByDomainNameRequest, callback func(response *QueryDomainByDomainNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainByDomainNameResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainByDomainName(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QueryDomainByDomainNameRequest is the request struct for api QueryDomainByDomainName
type QueryDomainByDomainNameRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// QueryDomainByDomainNameResponse is the response struct for api QueryDomainByDomainName
type QueryDomainByDomainNameResponse struct {
	*responses.BaseResponse
	UserId                       string                           `json:"UserId" xml:"UserId"`
	DomainName                   string                           `json:"DomainName" xml:"DomainName"`
	InstanceId                   string                           `json:"InstanceId" xml:"InstanceId"`
	RegistrationDate             string                           `json:"RegistrationDate" xml:"RegistrationDate"`
	ExpirationDate               string                           `json:"ExpirationDate" xml:"ExpirationDate"`
	RegistrantOrganization       string                           `json:"RegistrantOrganization" xml:"RegistrantOrganization"`
	RegistrantName               string                           `json:"RegistrantName" xml:"RegistrantName"`
	Email                        string                           `json:"Email" xml:"Email"`
	UpdateProhibitionLock        string                           `json:"UpdateProhibitionLock" xml:"UpdateProhibitionLock"`
	TransferProhibitionLock      string                           `json:"TransferProhibitionLock" xml:"TransferProhibitionLock"`
	DomainNameProxyService       bool                             `json:"DomainNameProxyService" xml:"DomainNameProxyService"`
	Premium                      bool                             `json:"Premium" xml:"Premium"`
	EmailVerificationStatus      int                              `json:"EmailVerificationStatus" xml:"EmailVerificationStatus"`
	EmailVerificationClientHold  bool                             `json:"EmailVerificationClientHold" xml:"EmailVerificationClientHold"`
	RealNameStatus               string                           `json:"RealNameStatus" xml:"RealNameStatus"`
	RegistrantUpdatingStatus     string                           `json:"RegistrantUpdatingStatus" xml:"RegistrantUpdatingStatus"`
	TransferOutStatus            string                           `json:"TransferOutStatus" xml:"TransferOutStatus"`
	RegistrantType               string                           `json:"RegistrantType" xml:"RegistrantType"`
	DomainNameVerificationStatus string                           `json:"DomainNameVerificationStatus" xml:"DomainNameVerificationStatus"`
	RequestId                    string                           `json:"RequestId" xml:"RequestId"`
	ZhRegistrantOrganization     string                           `json:"ZhRegistrantOrganization" xml:"ZhRegistrantOrganization"`
	ZhRegistrantName             string                           `json:"ZhRegistrantName" xml:"ZhRegistrantName"`
	RegistrationDateLong         int64                            `json:"RegistrationDateLong" xml:"RegistrationDateLong"`
	ExpirationDateLong           int64                            `json:"ExpirationDateLong" xml:"ExpirationDateLong"`
	DnsList                      DnsListInQueryDomainByDomainName `json:"DnsList" xml:"DnsList"`
}

// CreateQueryDomainByDomainNameRequest creates a request to invoke QueryDomainByDomainName API
func CreateQueryDomainByDomainNameRequest() (request *QueryDomainByDomainNameRequest) {
	request = &QueryDomainByDomainNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainByDomainName", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDomainByDomainNameResponse creates a response to parse from QueryDomainByDomainName response
func CreateQueryDomainByDomainNameResponse() (response *QueryDomainByDomainNameResponse) {
	response = &QueryDomainByDomainNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
