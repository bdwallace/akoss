package vod

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

// DescribeVodDomainDetail invokes the vod.DescribeVodDomainDetail API synchronously
func (client *Client) DescribeVodDomainDetail(request *DescribeVodDomainDetailRequest) (response *DescribeVodDomainDetailResponse, err error) {
	response = CreateDescribeVodDomainDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodDomainDetailWithChan invokes the vod.DescribeVodDomainDetail API asynchronously
func (client *Client) DescribeVodDomainDetailWithChan(request *DescribeVodDomainDetailRequest) (<-chan *DescribeVodDomainDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeVodDomainDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodDomainDetail(request)
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

// DescribeVodDomainDetailWithCallback invokes the vod.DescribeVodDomainDetail API asynchronously
func (client *Client) DescribeVodDomainDetailWithCallback(request *DescribeVodDomainDetailRequest, callback func(response *DescribeVodDomainDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodDomainDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodDomainDetail(request)
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

// DescribeVodDomainDetailRequest is the request struct for api DescribeVodDomainDetail
type DescribeVodDomainDetailRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeVodDomainDetailResponse is the response struct for api DescribeVodDomainDetail
type DescribeVodDomainDetailResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	DomainDetail DomainDetail `json:"DomainDetail" xml:"DomainDetail"`
}

// CreateDescribeVodDomainDetailRequest creates a request to invoke DescribeVodDomainDetail API
func CreateDescribeVodDomainDetailRequest() (request *DescribeVodDomainDetailRequest) {
	request = &DescribeVodDomainDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodDomainDetail", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodDomainDetailResponse creates a response to parse from DescribeVodDomainDetail response
func CreateDescribeVodDomainDetailResponse() (response *DescribeVodDomainDetailResponse) {
	response = &DescribeVodDomainDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
