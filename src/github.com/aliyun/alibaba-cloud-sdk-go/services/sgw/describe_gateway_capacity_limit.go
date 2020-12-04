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

// DescribeGatewayCapacityLimit invokes the sgw.DescribeGatewayCapacityLimit API synchronously
func (client *Client) DescribeGatewayCapacityLimit(request *DescribeGatewayCapacityLimitRequest) (response *DescribeGatewayCapacityLimitResponse, err error) {
	response = CreateDescribeGatewayCapacityLimitResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGatewayCapacityLimitWithChan invokes the sgw.DescribeGatewayCapacityLimit API asynchronously
func (client *Client) DescribeGatewayCapacityLimitWithChan(request *DescribeGatewayCapacityLimitRequest) (<-chan *DescribeGatewayCapacityLimitResponse, <-chan error) {
	responseChan := make(chan *DescribeGatewayCapacityLimitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGatewayCapacityLimit(request)
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

// DescribeGatewayCapacityLimitWithCallback invokes the sgw.DescribeGatewayCapacityLimit API asynchronously
func (client *Client) DescribeGatewayCapacityLimitWithCallback(request *DescribeGatewayCapacityLimitRequest, callback func(response *DescribeGatewayCapacityLimitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGatewayCapacityLimitResponse
		var err error
		defer close(result)
		response, err = client.DescribeGatewayCapacityLimit(request)
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

// DescribeGatewayCapacityLimitRequest is the request struct for api DescribeGatewayCapacityLimit
type DescribeGatewayCapacityLimitRequest struct {
	*requests.RpcRequest
	SizeInGB      requests.Integer `position:"Query" name:"SizeInGB"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	GatewayId     string           `position:"Query" name:"GatewayId"`
}

// DescribeGatewayCapacityLimitResponse is the response struct for api DescribeGatewayCapacityLimit
type DescribeGatewayCapacityLimitResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	Success            bool   `json:"Success" xml:"Success"`
	Code               string `json:"Code" xml:"Code"`
	Message            string `json:"Message" xml:"Message"`
	FileNumber         int64  `json:"FileNumber" xml:"FileNumber"`
	FileSystemSizeInTB int64  `json:"FileSystemSizeInTB" xml:"FileSystemSizeInTB"`
	IsMetadataSeparate bool   `json:"IsMetadataSeparate" xml:"IsMetadataSeparate"`
}

// CreateDescribeGatewayCapacityLimitRequest creates a request to invoke DescribeGatewayCapacityLimit API
func CreateDescribeGatewayCapacityLimitRequest() (request *DescribeGatewayCapacityLimitRequest) {
	request = &DescribeGatewayCapacityLimitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeGatewayCapacityLimit", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGatewayCapacityLimitResponse creates a response to parse from DescribeGatewayCapacityLimit response
func CreateDescribeGatewayCapacityLimitResponse() (response *DescribeGatewayCapacityLimitResponse) {
	response = &DescribeGatewayCapacityLimitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
