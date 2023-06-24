package retailcloud

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

// ListAppCmsGroups invokes the retailcloud.ListAppCmsGroups API synchronously
// api document: https://help.aliyun.com/api/retailcloud/listappcmsgroups.html
func (client *Client) ListAppCmsGroups(request *ListAppCmsGroupsRequest) (response *ListAppCmsGroupsResponse, err error) {
	response = CreateListAppCmsGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppCmsGroupsWithChan invokes the retailcloud.ListAppCmsGroups API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listappcmsgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAppCmsGroupsWithChan(request *ListAppCmsGroupsRequest) (<-chan *ListAppCmsGroupsResponse, <-chan error) {
	responseChan := make(chan *ListAppCmsGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppCmsGroups(request)
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

// ListAppCmsGroupsWithCallback invokes the retailcloud.ListAppCmsGroups API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listappcmsgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAppCmsGroupsWithCallback(request *ListAppCmsGroupsRequest, callback func(response *ListAppCmsGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppCmsGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListAppCmsGroups(request)
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

// ListAppCmsGroupsRequest is the request struct for api ListAppCmsGroups
type ListAppCmsGroupsRequest struct {
	*requests.RpcRequest
	AppId      requests.Integer `position:"Query" name:"AppId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	EnvId      requests.Integer `position:"Query" name:"EnvId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListAppCmsGroupsResponse is the response struct for api ListAppCmsGroups
type ListAppCmsGroupsResponse struct {
	*responses.BaseResponse
	Code       int      `json:"Code" xml:"Code"`
	ErrorMsg   string   `json:"ErrorMsg" xml:"ErrorMsg"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount int64    `json:"TotalCount" xml:"TotalCount"`
	Data       []string `json:"Data" xml:"Data"`
}

// CreateListAppCmsGroupsRequest creates a request to invoke ListAppCmsGroups API
func CreateListAppCmsGroupsRequest() (request *ListAppCmsGroupsRequest) {
	request = &ListAppCmsGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListAppCmsGroups", "", "")
	request.Method = requests.GET
	return
}

// CreateListAppCmsGroupsResponse creates a response to parse from ListAppCmsGroups response
func CreateListAppCmsGroupsResponse() (response *ListAppCmsGroupsResponse) {
	response = &ListAppCmsGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
