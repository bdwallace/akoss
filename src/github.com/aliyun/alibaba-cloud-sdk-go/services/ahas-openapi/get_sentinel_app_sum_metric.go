package ahas_openapi

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

// GetSentinelAppSumMetric invokes the ahas_openapi.GetSentinelAppSumMetric API synchronously
// api document: https://help.aliyun.com/api/ahas-openapi/getsentinelappsummetric.html
func (client *Client) GetSentinelAppSumMetric(request *GetSentinelAppSumMetricRequest) (response *GetSentinelAppSumMetricResponse, err error) {
	response = CreateGetSentinelAppSumMetricResponse()
	err = client.DoAction(request, response)
	return
}

// GetSentinelAppSumMetricWithChan invokes the ahas_openapi.GetSentinelAppSumMetric API asynchronously
// api document: https://help.aliyun.com/api/ahas-openapi/getsentinelappsummetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSentinelAppSumMetricWithChan(request *GetSentinelAppSumMetricRequest) (<-chan *GetSentinelAppSumMetricResponse, <-chan error) {
	responseChan := make(chan *GetSentinelAppSumMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSentinelAppSumMetric(request)
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

// GetSentinelAppSumMetricWithCallback invokes the ahas_openapi.GetSentinelAppSumMetric API asynchronously
// api document: https://help.aliyun.com/api/ahas-openapi/getsentinelappsummetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSentinelAppSumMetricWithCallback(request *GetSentinelAppSumMetricRequest, callback func(response *GetSentinelAppSumMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSentinelAppSumMetricResponse
		var err error
		defer close(result)
		response, err = client.GetSentinelAppSumMetric(request)
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

// GetSentinelAppSumMetricRequest is the request struct for api GetSentinelAppSumMetric
type GetSentinelAppSumMetricRequest struct {
	*requests.RpcRequest
	EndTime        string `position:"Query" name:"EndTime"`
	StartTime      string `position:"Query" name:"StartTime"`
	AppName        string `position:"Query" name:"AppName"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
	Namespace      string `position:"Query" name:"Namespace"`
}

// GetSentinelAppSumMetricResponse is the response struct for api GetSentinelAppSumMetric
type GetSentinelAppSumMetricResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Code       string     `json:"Code" xml:"Code"`
	Success    bool       `json:"Success" xml:"Success"`
	Message    string     `json:"Message" xml:"Message"`
	MetricData MetricData `json:"MetricData" xml:"MetricData"`
}

// CreateGetSentinelAppSumMetricRequest creates a request to invoke GetSentinelAppSumMetric API
func CreateGetSentinelAppSumMetricRequest() (request *GetSentinelAppSumMetricRequest) {
	request = &GetSentinelAppSumMetricRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "GetSentinelAppSumMetric", "ahas", "openAPI")
	return
}

// CreateGetSentinelAppSumMetricResponse creates a response to parse from GetSentinelAppSumMetric response
func CreateGetSentinelAppSumMetricResponse() (response *GetSentinelAppSumMetricResponse) {
	response = &GetSentinelAppSumMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
