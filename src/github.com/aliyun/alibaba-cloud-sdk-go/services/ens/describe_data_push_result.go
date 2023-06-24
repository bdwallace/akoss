package ens

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

// DescribeDataPushResult invokes the ens.DescribeDataPushResult API synchronously
func (client *Client) DescribeDataPushResult(request *DescribeDataPushResultRequest) (response *DescribeDataPushResultResponse, err error) {
	response = CreateDescribeDataPushResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataPushResultWithChan invokes the ens.DescribeDataPushResult API asynchronously
func (client *Client) DescribeDataPushResultWithChan(request *DescribeDataPushResultRequest) (<-chan *DescribeDataPushResultResponse, <-chan error) {
	responseChan := make(chan *DescribeDataPushResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataPushResult(request)
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

// DescribeDataPushResultWithCallback invokes the ens.DescribeDataPushResult API asynchronously
func (client *Client) DescribeDataPushResultWithCallback(request *DescribeDataPushResultRequest, callback func(response *DescribeDataPushResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataPushResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataPushResult(request)
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

// DescribeDataPushResultRequest is the request struct for api DescribeDataPushResult
type DescribeDataPushResultRequest struct {
	*requests.RpcRequest
	MaxDate      string           `position:"Query" name:"MaxDate"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	RegionIds    string           `position:"Query" name:"RegionIds"`
	MinDate      string           `position:"Query" name:"MinDate"`
	DataVersions string           `position:"Query" name:"DataVersions"`
	AppId        string           `position:"Query" name:"AppId"`
	DataNames    string           `position:"Query" name:"DataNames"`
}

// DescribeDataPushResultResponse is the response struct for api DescribeDataPushResult
type DescribeDataPushResultResponse struct {
	*responses.BaseResponse
	RequestId   string                              `json:"RequestId" xml:"RequestId"`
	TotalCount  int                                 `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int                                 `json:"PageNumber" xml:"PageNumber"`
	PageSize    int                                 `json:"PageSize" xml:"PageSize"`
	PushResults PushResultsInDescribeDataPushResult `json:"PushResults" xml:"PushResults"`
}

// CreateDescribeDataPushResultRequest creates a request to invoke DescribeDataPushResult API
func CreateDescribeDataPushResultRequest() (request *DescribeDataPushResultRequest) {
	request = &DescribeDataPushResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeDataPushResult", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataPushResultResponse creates a response to parse from DescribeDataPushResult response
func CreateDescribeDataPushResultResponse() (response *DescribeDataPushResultResponse) {
	response = &DescribeDataPushResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}