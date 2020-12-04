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

// DescribeRiskCheckResult invokes the sas.DescribeRiskCheckResult API synchronously
func (client *Client) DescribeRiskCheckResult(request *DescribeRiskCheckResultRequest) (response *DescribeRiskCheckResultResponse, err error) {
	response = CreateDescribeRiskCheckResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRiskCheckResultWithChan invokes the sas.DescribeRiskCheckResult API asynchronously
func (client *Client) DescribeRiskCheckResultWithChan(request *DescribeRiskCheckResultRequest) (<-chan *DescribeRiskCheckResultResponse, <-chan error) {
	responseChan := make(chan *DescribeRiskCheckResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRiskCheckResult(request)
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

// DescribeRiskCheckResultWithCallback invokes the sas.DescribeRiskCheckResult API asynchronously
func (client *Client) DescribeRiskCheckResultWithCallback(request *DescribeRiskCheckResultRequest, callback func(response *DescribeRiskCheckResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRiskCheckResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeRiskCheckResult(request)
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

// DescribeRiskCheckResultRequest is the request struct for api DescribeRiskCheckResult
type DescribeRiskCheckResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Lang            string           `position:"Query" name:"Lang"`
	AssetType       string           `position:"Query" name:"AssetType"`
	QueryFlag       string           `position:"Query" name:"QueryFlag"`
	GroupId         requests.Integer `position:"Query" name:"GroupId"`
	ItemIds         *[]string        `position:"Query" name:"ItemIds"  type:"Repeated"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	RiskLevel       string           `position:"Query" name:"RiskLevel"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	Name            string           `position:"Query" name:"Name"`
	Status          string           `position:"Query" name:"Status"`
}

// DescribeRiskCheckResultResponse is the response struct for api DescribeRiskCheckResult
type DescribeRiskCheckResultResponse struct {
	*responses.BaseResponse
	RequestId   string                      `json:"RequestId" xml:"RequestId"`
	PageCount   int                         `json:"PageCount" xml:"PageCount"`
	Count       int                         `json:"Count" xml:"Count"`
	PageSize    int                         `json:"PageSize" xml:"PageSize"`
	TotalCount  int                         `json:"TotalCount" xml:"TotalCount"`
	CurrentPage int                         `json:"CurrentPage" xml:"CurrentPage"`
	List        []RiskCheckResultForDisplay `json:"List" xml:"List"`
}

// CreateDescribeRiskCheckResultRequest creates a request to invoke DescribeRiskCheckResult API
func CreateDescribeRiskCheckResultRequest() (request *DescribeRiskCheckResultRequest) {
	request = &DescribeRiskCheckResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeRiskCheckResult", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRiskCheckResultResponse creates a response to parse from DescribeRiskCheckResult response
func CreateDescribeRiskCheckResultResponse() (response *DescribeRiskCheckResultResponse) {
	response = &DescribeRiskCheckResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
