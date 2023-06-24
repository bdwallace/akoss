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

// DescribeRiskCheckSummary invokes the sas.DescribeRiskCheckSummary API synchronously
func (client *Client) DescribeRiskCheckSummary(request *DescribeRiskCheckSummaryRequest) (response *DescribeRiskCheckSummaryResponse, err error) {
	response = CreateDescribeRiskCheckSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRiskCheckSummaryWithChan invokes the sas.DescribeRiskCheckSummary API asynchronously
func (client *Client) DescribeRiskCheckSummaryWithChan(request *DescribeRiskCheckSummaryRequest) (<-chan *DescribeRiskCheckSummaryResponse, <-chan error) {
	responseChan := make(chan *DescribeRiskCheckSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRiskCheckSummary(request)
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

// DescribeRiskCheckSummaryWithCallback invokes the sas.DescribeRiskCheckSummary API asynchronously
func (client *Client) DescribeRiskCheckSummaryWithCallback(request *DescribeRiskCheckSummaryRequest, callback func(response *DescribeRiskCheckSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRiskCheckSummaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeRiskCheckSummary(request)
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

// DescribeRiskCheckSummaryRequest is the request struct for api DescribeRiskCheckSummary
type DescribeRiskCheckSummaryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp                   string           `position:"Query" name:"SourceIp"`
	Lang                       string           `position:"Query" name:"Lang"`
	ResourceDirectoryAccountId string           `position:"Query" name:"ResourceDirectoryAccountId"`
}

// DescribeRiskCheckSummaryResponse is the response struct for api DescribeRiskCheckSummary
type DescribeRiskCheckSummaryResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	RiskCheckSummary RiskCheckSummary `json:"RiskCheckSummary" xml:"RiskCheckSummary"`
}

// CreateDescribeRiskCheckSummaryRequest creates a request to invoke DescribeRiskCheckSummary API
func CreateDescribeRiskCheckSummaryRequest() (request *DescribeRiskCheckSummaryRequest) {
	request = &DescribeRiskCheckSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeRiskCheckSummary", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRiskCheckSummaryResponse creates a response to parse from DescribeRiskCheckSummary response
func CreateDescribeRiskCheckSummaryResponse() (response *DescribeRiskCheckSummaryResponse) {
	response = &DescribeRiskCheckSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}