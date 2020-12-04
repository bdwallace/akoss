package waf_openapi

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

// DescribeCertMatchStatus invokes the waf_openapi.DescribeCertMatchStatus API synchronously
func (client *Client) DescribeCertMatchStatus(request *DescribeCertMatchStatusRequest) (response *DescribeCertMatchStatusResponse, err error) {
	response = CreateDescribeCertMatchStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCertMatchStatusWithChan invokes the waf_openapi.DescribeCertMatchStatus API asynchronously
func (client *Client) DescribeCertMatchStatusWithChan(request *DescribeCertMatchStatusRequest) (<-chan *DescribeCertMatchStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeCertMatchStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCertMatchStatus(request)
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

// DescribeCertMatchStatusWithCallback invokes the waf_openapi.DescribeCertMatchStatus API asynchronously
func (client *Client) DescribeCertMatchStatusWithCallback(request *DescribeCertMatchStatusRequest, callback func(response *DescribeCertMatchStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCertMatchStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeCertMatchStatus(request)
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

// DescribeCertMatchStatusRequest is the request struct for api DescribeCertMatchStatus
type DescribeCertMatchStatusRequest struct {
	*requests.RpcRequest
	Certificate string `position:"Query" name:"Certificate"`
	PrivateKey  string `position:"Query" name:"PrivateKey"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	Domain      string `position:"Query" name:"Domain"`
	Lang        string `position:"Query" name:"Lang"`
}

// DescribeCertMatchStatusResponse is the response struct for api DescribeCertMatchStatus
type DescribeCertMatchStatusResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	MatchStatus bool   `json:"MatchStatus" xml:"MatchStatus"`
}

// CreateDescribeCertMatchStatusRequest creates a request to invoke DescribeCertMatchStatus API
func CreateDescribeCertMatchStatusRequest() (request *DescribeCertMatchStatusRequest) {
	request = &DescribeCertMatchStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DescribeCertMatchStatus", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCertMatchStatusResponse creates a response to parse from DescribeCertMatchStatus response
func CreateDescribeCertMatchStatusResponse() (response *DescribeCertMatchStatusResponse) {
	response = &DescribeCertMatchStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
