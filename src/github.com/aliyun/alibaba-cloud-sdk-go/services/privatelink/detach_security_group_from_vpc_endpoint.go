package privatelink

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

// DetachSecurityGroupFromVpcEndpoint invokes the privatelink.DetachSecurityGroupFromVpcEndpoint API synchronously
func (client *Client) DetachSecurityGroupFromVpcEndpoint(request *DetachSecurityGroupFromVpcEndpointRequest) (response *DetachSecurityGroupFromVpcEndpointResponse, err error) {
	response = CreateDetachSecurityGroupFromVpcEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// DetachSecurityGroupFromVpcEndpointWithChan invokes the privatelink.DetachSecurityGroupFromVpcEndpoint API asynchronously
func (client *Client) DetachSecurityGroupFromVpcEndpointWithChan(request *DetachSecurityGroupFromVpcEndpointRequest) (<-chan *DetachSecurityGroupFromVpcEndpointResponse, <-chan error) {
	responseChan := make(chan *DetachSecurityGroupFromVpcEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachSecurityGroupFromVpcEndpoint(request)
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

// DetachSecurityGroupFromVpcEndpointWithCallback invokes the privatelink.DetachSecurityGroupFromVpcEndpoint API asynchronously
func (client *Client) DetachSecurityGroupFromVpcEndpointWithCallback(request *DetachSecurityGroupFromVpcEndpointRequest, callback func(response *DetachSecurityGroupFromVpcEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachSecurityGroupFromVpcEndpointResponse
		var err error
		defer close(result)
		response, err = client.DetachSecurityGroupFromVpcEndpoint(request)
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

// DetachSecurityGroupFromVpcEndpointRequest is the request struct for api DetachSecurityGroupFromVpcEndpoint
type DetachSecurityGroupFromVpcEndpointRequest struct {
	*requests.RpcRequest
	ClientToken     string           `position:"Query" name:"ClientToken"`
	EndpointId      string           `position:"Query" name:"EndpointId"`
	SecurityGroupId string           `position:"Query" name:"SecurityGroupId"`
	DryRun          requests.Boolean `position:"Query" name:"DryRun"`
}

// DetachSecurityGroupFromVpcEndpointResponse is the response struct for api DetachSecurityGroupFromVpcEndpoint
type DetachSecurityGroupFromVpcEndpointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachSecurityGroupFromVpcEndpointRequest creates a request to invoke DetachSecurityGroupFromVpcEndpoint API
func CreateDetachSecurityGroupFromVpcEndpointRequest() (request *DetachSecurityGroupFromVpcEndpointRequest) {
	request = &DetachSecurityGroupFromVpcEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "DetachSecurityGroupFromVpcEndpoint", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachSecurityGroupFromVpcEndpointResponse creates a response to parse from DetachSecurityGroupFromVpcEndpoint response
func CreateDetachSecurityGroupFromVpcEndpointResponse() (response *DetachSecurityGroupFromVpcEndpointResponse) {
	response = &DetachSecurityGroupFromVpcEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
