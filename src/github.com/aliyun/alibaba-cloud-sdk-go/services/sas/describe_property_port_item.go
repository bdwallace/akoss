package sas

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

// DescribePropertyPortItem invokes the sas.DescribePropertyPortItem API synchronously
func (client *Client) DescribePropertyPortItem(request *DescribePropertyPortItemRequest) (response *DescribePropertyPortItemResponse, err error) {
	response = CreateDescribePropertyPortItemResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePropertyPortItemWithChan invokes the sas.DescribePropertyPortItem API asynchronously
func (client *Client) DescribePropertyPortItemWithChan(request *DescribePropertyPortItemRequest) (<-chan *DescribePropertyPortItemResponse, <-chan error) {
	responseChan := make(chan *DescribePropertyPortItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePropertyPortItem(request)
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

// DescribePropertyPortItemWithCallback invokes the sas.DescribePropertyPortItem API asynchronously
func (client *Client) DescribePropertyPortItemWithCallback(request *DescribePropertyPortItemRequest, callback func(response *DescribePropertyPortItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePropertyPortItemResponse
		var err error
		defer close(result)
		response, err = client.DescribePropertyPortItem(request)
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

// DescribePropertyPortItemRequest is the request struct for api DescribePropertyPortItem
type DescribePropertyPortItemRequest struct {
	*requests.RpcRequest
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Port        string           `position:"Query" name:"Port"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ForceFlush  requests.Boolean `position:"Query" name:"ForceFlush"`
}

// DescribePropertyPortItemResponse is the response struct for api DescribePropertyPortItem
type DescribePropertyPortItemResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	PageInfo      PageInfo       `json:"PageInfo" xml:"PageInfo"`
	PropertyItems []PropertyItem `json:"PropertyItems" xml:"PropertyItems"`
}

// CreateDescribePropertyPortItemRequest creates a request to invoke DescribePropertyPortItem API
func CreateDescribePropertyPortItemRequest() (request *DescribePropertyPortItemRequest) {
	request = &DescribePropertyPortItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribePropertyPortItem", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePropertyPortItemResponse creates a response to parse from DescribePropertyPortItem response
func CreateDescribePropertyPortItemResponse() (response *DescribePropertyPortItemResponse) {
	response = &DescribePropertyPortItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}