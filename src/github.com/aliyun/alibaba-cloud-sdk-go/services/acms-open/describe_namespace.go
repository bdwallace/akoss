package acms_open

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

// DescribeNamespace invokes the acms_open.DescribeNamespace API synchronously
// api document: https://help.aliyun.com/api/acms-open/describenamespace.html
func (client *Client) DescribeNamespace(request *DescribeNamespaceRequest) (response *DescribeNamespaceResponse, err error) {
	response = CreateDescribeNamespaceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNamespaceWithChan invokes the acms_open.DescribeNamespace API asynchronously
// api document: https://help.aliyun.com/api/acms-open/describenamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNamespaceWithChan(request *DescribeNamespaceRequest) (<-chan *DescribeNamespaceResponse, <-chan error) {
	responseChan := make(chan *DescribeNamespaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNamespace(request)
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

// DescribeNamespaceWithCallback invokes the acms_open.DescribeNamespace API asynchronously
// api document: https://help.aliyun.com/api/acms-open/describenamespace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNamespaceWithCallback(request *DescribeNamespaceRequest, callback func(response *DescribeNamespaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNamespaceResponse
		var err error
		defer close(result)
		response, err = client.DescribeNamespace(request)
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

// DescribeNamespaceRequest is the request struct for api DescribeNamespace
type DescribeNamespaceRequest struct {
	*requests.RoaRequest
	NamespaceId string `position:"Query" name:"NamespaceId"`
}

// DescribeNamespaceResponse is the response struct for api DescribeNamespace
type DescribeNamespaceResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Code      string    `json:"Code" xml:"Code"`
	Message   string    `json:"Message" xml:"Message"`
	Namespace Namespace `json:"Namespace" xml:"Namespace"`
}

// CreateDescribeNamespaceRequest creates a request to invoke DescribeNamespace API
func CreateDescribeNamespaceRequest() (request *DescribeNamespaceRequest) {
	request = &DescribeNamespaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("acms-open", "2020-02-06", "DescribeNamespace", "/diamond-ops/pop/namespace", "acms", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeNamespaceResponse creates a response to parse from DescribeNamespace response
func CreateDescribeNamespaceResponse() (response *DescribeNamespaceResponse) {
	response = &DescribeNamespaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
