package openanalytics_open

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

// OpenDataLakeAnalyticsService invokes the openanalytics_open.OpenDataLakeAnalyticsService API synchronously
func (client *Client) OpenDataLakeAnalyticsService(request *OpenDataLakeAnalyticsServiceRequest) (response *OpenDataLakeAnalyticsServiceResponse, err error) {
	response = CreateOpenDataLakeAnalyticsServiceResponse()
	err = client.DoAction(request, response)
	return
}

// OpenDataLakeAnalyticsServiceWithChan invokes the openanalytics_open.OpenDataLakeAnalyticsService API asynchronously
func (client *Client) OpenDataLakeAnalyticsServiceWithChan(request *OpenDataLakeAnalyticsServiceRequest) (<-chan *OpenDataLakeAnalyticsServiceResponse, <-chan error) {
	responseChan := make(chan *OpenDataLakeAnalyticsServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenDataLakeAnalyticsService(request)
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

// OpenDataLakeAnalyticsServiceWithCallback invokes the openanalytics_open.OpenDataLakeAnalyticsService API asynchronously
func (client *Client) OpenDataLakeAnalyticsServiceWithCallback(request *OpenDataLakeAnalyticsServiceRequest, callback func(response *OpenDataLakeAnalyticsServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenDataLakeAnalyticsServiceResponse
		var err error
		defer close(result)
		response, err = client.OpenDataLakeAnalyticsService(request)
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

// OpenDataLakeAnalyticsServiceRequest is the request struct for api OpenDataLakeAnalyticsService
type OpenDataLakeAnalyticsServiceRequest struct {
	*requests.RpcRequest
	InternetChargeType string `position:"Query" name:"InternetChargeType"`
}

// OpenDataLakeAnalyticsServiceResponse is the response struct for api OpenDataLakeAnalyticsService
type OpenDataLakeAnalyticsServiceResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorInfo string `json:"ErrorInfo" xml:"ErrorInfo"`
	Result    string `json:"Result" xml:"Result"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOpenDataLakeAnalyticsServiceRequest creates a request to invoke OpenDataLakeAnalyticsService API
func CreateOpenDataLakeAnalyticsServiceRequest() (request *OpenDataLakeAnalyticsServiceRequest) {
	request = &OpenDataLakeAnalyticsServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "OpenDataLakeAnalyticsService", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOpenDataLakeAnalyticsServiceResponse creates a response to parse from OpenDataLakeAnalyticsService response
func CreateOpenDataLakeAnalyticsServiceResponse() (response *OpenDataLakeAnalyticsServiceResponse) {
	response = &OpenDataLakeAnalyticsServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
