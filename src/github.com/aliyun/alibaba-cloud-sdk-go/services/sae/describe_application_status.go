package sae

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

// DescribeApplicationStatus invokes the sae.DescribeApplicationStatus API synchronously
func (client *Client) DescribeApplicationStatus(request *DescribeApplicationStatusRequest) (response *DescribeApplicationStatusResponse, err error) {
	response = CreateDescribeApplicationStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApplicationStatusWithChan invokes the sae.DescribeApplicationStatus API asynchronously
func (client *Client) DescribeApplicationStatusWithChan(request *DescribeApplicationStatusRequest) (<-chan *DescribeApplicationStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeApplicationStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApplicationStatus(request)
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

// DescribeApplicationStatusWithCallback invokes the sae.DescribeApplicationStatus API asynchronously
func (client *Client) DescribeApplicationStatusWithCallback(request *DescribeApplicationStatusRequest, callback func(response *DescribeApplicationStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApplicationStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeApplicationStatus(request)
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

// DescribeApplicationStatusRequest is the request struct for api DescribeApplicationStatus
type DescribeApplicationStatusRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// DescribeApplicationStatusResponse is the response struct for api DescribeApplicationStatus
type DescribeApplicationStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeApplicationStatusRequest creates a request to invoke DescribeApplicationStatus API
func CreateDescribeApplicationStatusRequest() (request *DescribeApplicationStatusRequest) {
	request = &DescribeApplicationStatusRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeApplicationStatus", "/pop/v1/sam/app/describeApplicationStatus", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeApplicationStatusResponse creates a response to parse from DescribeApplicationStatus response
func CreateDescribeApplicationStatusResponse() (response *DescribeApplicationStatusResponse) {
	response = &DescribeApplicationStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
