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

// QueryRegistrantProfiles invokes the domain.QueryRegistrantProfiles API synchronously
// api document: https://help.aliyun.com/api/domain/queryregistrantprofiles.html
func (client *Client) QueryRegistrantProfiles(request *QueryRegistrantProfilesRequest) (response *QueryRegistrantProfilesResponse, err error) {
	response = CreateQueryRegistrantProfilesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRegistrantProfilesWithChan invokes the domain.QueryRegistrantProfiles API asynchronously
// api document: https://help.aliyun.com/api/domain/queryregistrantprofiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRegistrantProfilesWithChan(request *QueryRegistrantProfilesRequest) (<-chan *QueryRegistrantProfilesResponse, <-chan error) {
	responseChan := make(chan *QueryRegistrantProfilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRegistrantProfiles(request)
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

// QueryRegistrantProfilesWithCallback invokes the domain.QueryRegistrantProfiles API asynchronously
// api document: https://help.aliyun.com/api/domain/queryregistrantprofiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRegistrantProfilesWithCallback(request *QueryRegistrantProfilesRequest, callback func(response *QueryRegistrantProfilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRegistrantProfilesResponse
		var err error
		defer close(result)
		response, err = client.QueryRegistrantProfiles(request)
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

// QueryRegistrantProfilesRequest is the request struct for api QueryRegistrantProfiles
type QueryRegistrantProfilesRequest struct {
	*requests.RpcRequest
	RegistrantProfileId      requests.Integer `position:"Query" name:"RegistrantProfileId"`
	PageNum                  requests.Integer `position:"Query" name:"PageNum"`
	PageSize                 requests.Integer `position:"Query" name:"PageSize"`
	RealNameStatus           string           `position:"Query" name:"RealNameStatus"`
	Lang                     string           `position:"Query" name:"Lang"`
	Email                    string           `position:"Query" name:"Email"`
	ZhRegistrantOrganization string           `position:"Query" name:"ZhRegistrantOrganization"`
	RegistrantType           string           `position:"Query" name:"RegistrantType"`
	RegistrantProfileType    string           `position:"Query" name:"RegistrantProfileType"`
	DefaultRegistrantProfile requests.Boolean `position:"Query" name:"DefaultRegistrantProfile"`
	RegistrantOrganization   string           `position:"Query" name:"RegistrantOrganization"`
	UserClientIp             string           `position:"Query" name:"UserClientIp"`
}

// QueryRegistrantProfilesResponse is the response struct for api QueryRegistrantProfiles
type QueryRegistrantProfilesResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	TotalItemNum       int                `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum     int                `json:"CurrentPageNum" xml:"CurrentPageNum"`
	TotalPageNum       int                `json:"TotalPageNum" xml:"TotalPageNum"`
	PageSize           int                `json:"PageSize" xml:"PageSize"`
	PrePage            bool               `json:"PrePage" xml:"PrePage"`
	NextPage           bool               `json:"NextPage" xml:"NextPage"`
	RegistrantProfiles RegistrantProfiles `json:"RegistrantProfiles" xml:"RegistrantProfiles"`
}

// CreateQueryRegistrantProfilesRequest creates a request to invoke QueryRegistrantProfiles API
func CreateQueryRegistrantProfilesRequest() (request *QueryRegistrantProfilesRequest) {
	request = &QueryRegistrantProfilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryRegistrantProfiles", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryRegistrantProfilesResponse creates a response to parse from QueryRegistrantProfiles response
func CreateQueryRegistrantProfilesResponse() (response *QueryRegistrantProfilesResponse) {
	response = &QueryRegistrantProfilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
