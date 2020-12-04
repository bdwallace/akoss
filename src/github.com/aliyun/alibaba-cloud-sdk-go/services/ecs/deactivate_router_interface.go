package ecs

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

// DeactivateRouterInterface invokes the ecs.DeactivateRouterInterface API synchronously
func (client *Client) DeactivateRouterInterface(request *DeactivateRouterInterfaceRequest) (response *DeactivateRouterInterfaceResponse, err error) {
	response = CreateDeactivateRouterInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// DeactivateRouterInterfaceWithChan invokes the ecs.DeactivateRouterInterface API asynchronously
func (client *Client) DeactivateRouterInterfaceWithChan(request *DeactivateRouterInterfaceRequest) (<-chan *DeactivateRouterInterfaceResponse, <-chan error) {
	responseChan := make(chan *DeactivateRouterInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeactivateRouterInterface(request)
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

// DeactivateRouterInterfaceWithCallback invokes the ecs.DeactivateRouterInterface API asynchronously
func (client *Client) DeactivateRouterInterfaceWithCallback(request *DeactivateRouterInterfaceRequest, callback func(response *DeactivateRouterInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeactivateRouterInterfaceResponse
		var err error
		defer close(result)
		response, err = client.DeactivateRouterInterface(request)
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

// DeactivateRouterInterfaceRequest is the request struct for api DeactivateRouterInterface
type DeactivateRouterInterfaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string           `position:"Query" name:"RouterInterfaceId"`
}

// DeactivateRouterInterfaceResponse is the response struct for api DeactivateRouterInterface
type DeactivateRouterInterfaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeactivateRouterInterfaceRequest creates a request to invoke DeactivateRouterInterface API
func CreateDeactivateRouterInterfaceRequest() (request *DeactivateRouterInterfaceRequest) {
	request = &DeactivateRouterInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeactivateRouterInterface", "", "")
	request.Method = requests.POST
	return
}

// CreateDeactivateRouterInterfaceResponse creates a response to parse from DeactivateRouterInterface response
func CreateDeactivateRouterInterfaceResponse() (response *DeactivateRouterInterfaceResponse) {
	response = &DeactivateRouterInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
