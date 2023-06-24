package cloudcallcenter

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

// ListVnNavigationScripts invokes the cloudcallcenter.ListVnNavigationScripts API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listvnnavigationscripts.html
func (client *Client) ListVnNavigationScripts(request *ListVnNavigationScriptsRequest) (response *ListVnNavigationScriptsResponse, err error) {
	response = CreateListVnNavigationScriptsResponse()
	err = client.DoAction(request, response)
	return
}

// ListVnNavigationScriptsWithChan invokes the cloudcallcenter.ListVnNavigationScripts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listvnnavigationscripts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVnNavigationScriptsWithChan(request *ListVnNavigationScriptsRequest) (<-chan *ListVnNavigationScriptsResponse, <-chan error) {
	responseChan := make(chan *ListVnNavigationScriptsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVnNavigationScripts(request)
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

// ListVnNavigationScriptsWithCallback invokes the cloudcallcenter.ListVnNavigationScripts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listvnnavigationscripts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVnNavigationScriptsWithCallback(request *ListVnNavigationScriptsRequest, callback func(response *ListVnNavigationScriptsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVnNavigationScriptsResponse
		var err error
		defer close(result)
		response, err = client.ListVnNavigationScripts(request)
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

// ListVnNavigationScriptsRequest is the request struct for api ListVnNavigationScripts
type ListVnNavigationScriptsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	CategoryId string           `position:"Query" name:"CategoryId"`
}

// ListVnNavigationScriptsResponse is the response struct for api ListVnNavigationScripts
type ListVnNavigationScriptsResponse struct {
	*responses.BaseResponse
	RequestId         string             `json:"RequestId" xml:"RequestId"`
	TotalCount        int64              `json:"TotalCount" xml:"TotalCount"`
	PageNumber        int                `json:"PageNumber" xml:"PageNumber"`
	PageSize          int                `json:"PageSize" xml:"PageSize"`
	NavigationScripts []NavigationScript `json:"NavigationScripts" xml:"NavigationScripts"`
}

// CreateListVnNavigationScriptsRequest creates a request to invoke ListVnNavigationScripts API
func CreateListVnNavigationScriptsRequest() (request *ListVnNavigationScriptsRequest) {
	request = &ListVnNavigationScriptsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListVnNavigationScripts", "", "")
	request.Method = requests.GET
	return
}

// CreateListVnNavigationScriptsResponse creates a response to parse from ListVnNavigationScripts response
func CreateListVnNavigationScriptsResponse() (response *ListVnNavigationScriptsResponse) {
	response = &ListVnNavigationScriptsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
