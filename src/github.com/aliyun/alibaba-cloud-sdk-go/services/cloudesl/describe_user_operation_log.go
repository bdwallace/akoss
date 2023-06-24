package cloudesl

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

// DescribeUserOperationLog invokes the cloudesl.DescribeUserOperationLog API synchronously
func (client *Client) DescribeUserOperationLog(request *DescribeUserOperationLogRequest) (response *DescribeUserOperationLogResponse, err error) {
	response = CreateDescribeUserOperationLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserOperationLogWithChan invokes the cloudesl.DescribeUserOperationLog API asynchronously
func (client *Client) DescribeUserOperationLogWithChan(request *DescribeUserOperationLogRequest) (<-chan *DescribeUserOperationLogResponse, <-chan error) {
	responseChan := make(chan *DescribeUserOperationLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserOperationLog(request)
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

// DescribeUserOperationLogWithCallback invokes the cloudesl.DescribeUserOperationLog API asynchronously
func (client *Client) DescribeUserOperationLogWithCallback(request *DescribeUserOperationLogRequest, callback func(response *DescribeUserOperationLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserOperationLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserOperationLog(request)
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

// DescribeUserOperationLogRequest is the request struct for api DescribeUserOperationLog
type DescribeUserOperationLogRequest struct {
	*requests.RpcRequest
	OperateUserId requests.Integer `position:"Query" name:"OperateUserId"`
	StoreId       string           `position:"Query" name:"StoreId"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	FromDate      string           `position:"Query" name:"FromDate"`
	ItemId        requests.Integer `position:"Query" name:"ItemId"`
	ToDate        string           `position:"Query" name:"ToDate"`
	EslBarCode    string           `position:"Query" name:"EslBarCode"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	ItemBarCode   string           `position:"Query" name:"ItemBarCode"`
	ItemTitle     string           `position:"Query" name:"ItemTitle"`
	OperateStatus string           `position:"Query" name:"OperateStatus"`
	Reverse       requests.Boolean `position:"Query" name:"Reverse"`
	OperateType   string           `position:"Query" name:"OperateType"`
}

// DescribeUserOperationLogResponse is the response struct for api DescribeUserOperationLog
type DescribeUserOperationLogResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	Success       bool          `json:"Success" xml:"Success"`
	Message       string        `json:"Message" xml:"Message"`
	ErrorCode     string        `json:"ErrorCode" xml:"ErrorCode"`
	TotalCount    int           `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int           `json:"PageNumber" xml:"PageNumber"`
	PageSize      int           `json:"PageSize" xml:"PageSize"`
	OperationLogs OperationLogs `json:"OperationLogs" xml:"OperationLogs"`
}

// CreateDescribeUserOperationLogRequest creates a request to invoke DescribeUserOperationLog API
func CreateDescribeUserOperationLogRequest() (request *DescribeUserOperationLogRequest) {
	request = &DescribeUserOperationLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2018-08-01", "DescribeUserOperationLog", "cloudesl", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUserOperationLogResponse creates a response to parse from DescribeUserOperationLog response
func CreateDescribeUserOperationLogResponse() (response *DescribeUserOperationLogResponse) {
	response = &DescribeUserOperationLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
