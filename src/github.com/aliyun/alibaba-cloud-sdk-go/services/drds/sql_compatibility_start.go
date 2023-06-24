package drds

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

// SqlCompatibilityStart invokes the drds.SqlCompatibilityStart API synchronously
func (client *Client) SqlCompatibilityStart(request *SqlCompatibilityStartRequest) (response *SqlCompatibilityStartResponse, err error) {
	response = CreateSqlCompatibilityStartResponse()
	err = client.DoAction(request, response)
	return
}

// SqlCompatibilityStartWithChan invokes the drds.SqlCompatibilityStart API asynchronously
func (client *Client) SqlCompatibilityStartWithChan(request *SqlCompatibilityStartRequest) (<-chan *SqlCompatibilityStartResponse, <-chan error) {
	responseChan := make(chan *SqlCompatibilityStartResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SqlCompatibilityStart(request)
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

// SqlCompatibilityStartWithCallback invokes the drds.SqlCompatibilityStart API asynchronously
func (client *Client) SqlCompatibilityStartWithCallback(request *SqlCompatibilityStartRequest, callback func(response *SqlCompatibilityStartResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SqlCompatibilityStartResponse
		var err error
		defer close(result)
		response, err = client.SqlCompatibilityStart(request)
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

// SqlCompatibilityStartRequest is the request struct for api SqlCompatibilityStart
type SqlCompatibilityStartRequest struct {
	*requests.RpcRequest
	TargetVersion   string           `position:"Query" name:"TargetVersion"`
	DrdsInstanceId  string           `position:"Query" name:"DrdsInstanceId"`
	PerformanceTest requests.Boolean `position:"Query" name:"PerformanceTest"`
}

// SqlCompatibilityStartResponse is the response struct for api SqlCompatibilityStart
type SqlCompatibilityStartResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSqlCompatibilityStartRequest creates a request to invoke SqlCompatibilityStart API
func CreateSqlCompatibilityStartRequest() (request *SqlCompatibilityStartRequest) {
	request = &SqlCompatibilityStartRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "SqlCompatibilityStart", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSqlCompatibilityStartResponse creates a response to parse from SqlCompatibilityStart response
func CreateSqlCompatibilityStartResponse() (response *SqlCompatibilityStartResponse) {
	response = &SqlCompatibilityStartResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}