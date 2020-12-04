package sgw

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

// DescribeGatewayFileShares invokes the sgw.DescribeGatewayFileShares API synchronously
func (client *Client) DescribeGatewayFileShares(request *DescribeGatewayFileSharesRequest) (response *DescribeGatewayFileSharesResponse, err error) {
	response = CreateDescribeGatewayFileSharesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGatewayFileSharesWithChan invokes the sgw.DescribeGatewayFileShares API asynchronously
func (client *Client) DescribeGatewayFileSharesWithChan(request *DescribeGatewayFileSharesRequest) (<-chan *DescribeGatewayFileSharesResponse, <-chan error) {
	responseChan := make(chan *DescribeGatewayFileSharesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGatewayFileShares(request)
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

// DescribeGatewayFileSharesWithCallback invokes the sgw.DescribeGatewayFileShares API asynchronously
func (client *Client) DescribeGatewayFileSharesWithCallback(request *DescribeGatewayFileSharesRequest, callback func(response *DescribeGatewayFileSharesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGatewayFileSharesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGatewayFileShares(request)
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

// DescribeGatewayFileSharesRequest is the request struct for api DescribeGatewayFileShares
type DescribeGatewayFileSharesRequest struct {
	*requests.RpcRequest
	Refresh       requests.Boolean `position:"Query" name:"Refresh"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	IndexId       string           `position:"Query" name:"IndexId"`
	GatewayId     string           `position:"Query" name:"GatewayId"`
}

// DescribeGatewayFileSharesResponse is the response struct for api DescribeGatewayFileShares
type DescribeGatewayFileSharesResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Success    bool       `json:"Success" xml:"Success"`
	Code       string     `json:"Code" xml:"Code"`
	Message    string     `json:"Message" xml:"Message"`
	FileShares FileShares `json:"FileShares" xml:"FileShares"`
}

// CreateDescribeGatewayFileSharesRequest creates a request to invoke DescribeGatewayFileShares API
func CreateDescribeGatewayFileSharesRequest() (request *DescribeGatewayFileSharesRequest) {
	request = &DescribeGatewayFileSharesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeGatewayFileShares", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGatewayFileSharesResponse creates a response to parse from DescribeGatewayFileShares response
func CreateDescribeGatewayFileSharesResponse() (response *DescribeGatewayFileSharesResponse) {
	response = &DescribeGatewayFileSharesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
