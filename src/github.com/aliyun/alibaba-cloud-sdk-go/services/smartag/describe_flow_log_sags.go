package smartag

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

// DescribeFlowLogSags invokes the smartag.DescribeFlowLogSags API synchronously
func (client *Client) DescribeFlowLogSags(request *DescribeFlowLogSagsRequest) (response *DescribeFlowLogSagsResponse, err error) {
	response = CreateDescribeFlowLogSagsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowLogSagsWithChan invokes the smartag.DescribeFlowLogSags API asynchronously
func (client *Client) DescribeFlowLogSagsWithChan(request *DescribeFlowLogSagsRequest) (<-chan *DescribeFlowLogSagsResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowLogSagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlowLogSags(request)
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

// DescribeFlowLogSagsWithCallback invokes the smartag.DescribeFlowLogSags API asynchronously
func (client *Client) DescribeFlowLogSagsWithCallback(request *DescribeFlowLogSagsRequest, callback func(response *DescribeFlowLogSagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowLogSagsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlowLogSags(request)
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

// DescribeFlowLogSagsRequest is the request struct for api DescribeFlowLogSags
type DescribeFlowLogSagsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FlowLogId            string           `position:"Query" name:"FlowLogId"`
}

// DescribeFlowLogSagsResponse is the response struct for api DescribeFlowLogSags
type DescribeFlowLogSagsResponse struct {
	*responses.BaseResponse
	RequestId  string                    `json:"RequestId" xml:"RequestId"`
	TotalCount int                       `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                       `json:"PageSize" xml:"PageSize"`
	Sags       SagsInDescribeFlowLogSags `json:"Sags" xml:"Sags"`
}

// CreateDescribeFlowLogSagsRequest creates a request to invoke DescribeFlowLogSags API
func CreateDescribeFlowLogSagsRequest() (request *DescribeFlowLogSagsRequest) {
	request = &DescribeFlowLogSagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeFlowLogSags", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeFlowLogSagsResponse creates a response to parse from DescribeFlowLogSags response
func CreateDescribeFlowLogSagsResponse() (response *DescribeFlowLogSagsResponse) {
	response = &DescribeFlowLogSagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
