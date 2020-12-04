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

// DescribeGatewayFileStatus invokes the sgw.DescribeGatewayFileStatus API synchronously
func (client *Client) DescribeGatewayFileStatus(request *DescribeGatewayFileStatusRequest) (response *DescribeGatewayFileStatusResponse, err error) {
	response = CreateDescribeGatewayFileStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGatewayFileStatusWithChan invokes the sgw.DescribeGatewayFileStatus API asynchronously
func (client *Client) DescribeGatewayFileStatusWithChan(request *DescribeGatewayFileStatusRequest) (<-chan *DescribeGatewayFileStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeGatewayFileStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGatewayFileStatus(request)
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

// DescribeGatewayFileStatusWithCallback invokes the sgw.DescribeGatewayFileStatus API asynchronously
func (client *Client) DescribeGatewayFileStatusWithCallback(request *DescribeGatewayFileStatusRequest, callback func(response *DescribeGatewayFileStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGatewayFileStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeGatewayFileStatus(request)
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

// DescribeGatewayFileStatusRequest is the request struct for api DescribeGatewayFileStatus
type DescribeGatewayFileStatusRequest struct {
	*requests.RpcRequest
	FilePath      string `position:"Query" name:"FilePath"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	IndexId       string `position:"Query" name:"IndexId"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// DescribeGatewayFileStatusResponse is the response struct for api DescribeGatewayFileStatus
type DescribeGatewayFileStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateDescribeGatewayFileStatusRequest creates a request to invoke DescribeGatewayFileStatus API
func CreateDescribeGatewayFileStatusRequest() (request *DescribeGatewayFileStatusRequest) {
	request = &DescribeGatewayFileStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeGatewayFileStatus", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGatewayFileStatusResponse creates a response to parse from DescribeGatewayFileStatus response
func CreateDescribeGatewayFileStatusResponse() (response *DescribeGatewayFileStatusResponse) {
	response = &DescribeGatewayFileStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
