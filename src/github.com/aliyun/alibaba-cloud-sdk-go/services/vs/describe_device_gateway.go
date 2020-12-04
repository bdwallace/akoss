package vs

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

// DescribeDeviceGateway invokes the vs.DescribeDeviceGateway API synchronously
func (client *Client) DescribeDeviceGateway(request *DescribeDeviceGatewayRequest) (response *DescribeDeviceGatewayResponse, err error) {
	response = CreateDescribeDeviceGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDeviceGatewayWithChan invokes the vs.DescribeDeviceGateway API asynchronously
func (client *Client) DescribeDeviceGatewayWithChan(request *DescribeDeviceGatewayRequest) (<-chan *DescribeDeviceGatewayResponse, <-chan error) {
	responseChan := make(chan *DescribeDeviceGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDeviceGateway(request)
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

// DescribeDeviceGatewayWithCallback invokes the vs.DescribeDeviceGateway API asynchronously
func (client *Client) DescribeDeviceGatewayWithCallback(request *DescribeDeviceGatewayRequest, callback func(response *DescribeDeviceGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeviceGatewayResponse
		var err error
		defer close(result)
		response, err = client.DescribeDeviceGateway(request)
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

// DescribeDeviceGatewayRequest is the request struct for api DescribeDeviceGateway
type DescribeDeviceGatewayRequest struct {
	*requests.RpcRequest
	ClientIp string           `position:"Query" name:"ClientIp"`
	Id       string           `position:"Query" name:"Id"`
	ShowLog  string           `position:"Query" name:"ShowLog"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	Expire   requests.Integer `position:"Query" name:"Expire"`
}

// DescribeDeviceGatewayResponse is the response struct for api DescribeDeviceGateway
type DescribeDeviceGatewayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Host      string `json:"Host" xml:"Host"`
	Port      int64  `json:"Port" xml:"Port"`
	Protocol  string `json:"Protocol" xml:"Protocol"`
	Token     string `json:"Token" xml:"Token"`
}

// CreateDescribeDeviceGatewayRequest creates a request to invoke DescribeDeviceGateway API
func CreateDescribeDeviceGatewayRequest() (request *DescribeDeviceGatewayRequest) {
	request = &DescribeDeviceGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeDeviceGateway", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDeviceGatewayResponse creates a response to parse from DescribeDeviceGateway response
func CreateDescribeDeviceGatewayResponse() (response *DescribeDeviceGatewayResponse) {
	response = &DescribeDeviceGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
