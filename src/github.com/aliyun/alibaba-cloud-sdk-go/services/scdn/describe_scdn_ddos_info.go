package scdn

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

// DescribeScdnDdosInfo invokes the scdn.DescribeScdnDdosInfo API synchronously
func (client *Client) DescribeScdnDdosInfo(request *DescribeScdnDdosInfoRequest) (response *DescribeScdnDdosInfoResponse, err error) {
	response = CreateDescribeScdnDdosInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDdosInfoWithChan invokes the scdn.DescribeScdnDdosInfo API asynchronously
func (client *Client) DescribeScdnDdosInfoWithChan(request *DescribeScdnDdosInfoRequest) (<-chan *DescribeScdnDdosInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDdosInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDdosInfo(request)
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

// DescribeScdnDdosInfoWithCallback invokes the scdn.DescribeScdnDdosInfo API asynchronously
func (client *Client) DescribeScdnDdosInfoWithCallback(request *DescribeScdnDdosInfoRequest, callback func(response *DescribeScdnDdosInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDdosInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDdosInfo(request)
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

// DescribeScdnDdosInfoRequest is the request struct for api DescribeScdnDdosInfo
type DescribeScdnDdosInfoRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnDdosInfoResponse is the response struct for api DescribeScdnDdosInfo
type DescribeScdnDdosInfoResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	SecBandwidth     int    `json:"SecBandwidth" xml:"SecBandwidth"`
	ElasticBandwidth int    `json:"ElasticBandwidth" xml:"ElasticBandwidth"`
}

// CreateDescribeScdnDdosInfoRequest creates a request to invoke DescribeScdnDdosInfo API
func CreateDescribeScdnDdosInfoRequest() (request *DescribeScdnDdosInfoRequest) {
	request = &DescribeScdnDdosInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDdosInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeScdnDdosInfoResponse creates a response to parse from DescribeScdnDdosInfo response
func CreateDescribeScdnDdosInfoResponse() (response *DescribeScdnDdosInfoResponse) {
	response = &DescribeScdnDdosInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}