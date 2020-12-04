package oam

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

// RevokeRoleFromOperator invokes the oam.RevokeRoleFromOperator API synchronously
// api document: https://help.aliyun.com/api/oam/revokerolefromoperator.html
func (client *Client) RevokeRoleFromOperator(request *RevokeRoleFromOperatorRequest) (response *RevokeRoleFromOperatorResponse, err error) {
	response = CreateRevokeRoleFromOperatorResponse()
	err = client.DoAction(request, response)
	return
}

// RevokeRoleFromOperatorWithChan invokes the oam.RevokeRoleFromOperator API asynchronously
// api document: https://help.aliyun.com/api/oam/revokerolefromoperator.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RevokeRoleFromOperatorWithChan(request *RevokeRoleFromOperatorRequest) (<-chan *RevokeRoleFromOperatorResponse, <-chan error) {
	responseChan := make(chan *RevokeRoleFromOperatorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeRoleFromOperator(request)
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

// RevokeRoleFromOperatorWithCallback invokes the oam.RevokeRoleFromOperator API asynchronously
// api document: https://help.aliyun.com/api/oam/revokerolefromoperator.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RevokeRoleFromOperatorWithCallback(request *RevokeRoleFromOperatorRequest, callback func(response *RevokeRoleFromOperatorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeRoleFromOperatorResponse
		var err error
		defer close(result)
		response, err = client.RevokeRoleFromOperator(request)
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

// RevokeRoleFromOperatorRequest is the request struct for api RevokeRoleFromOperator
type RevokeRoleFromOperatorRequest struct {
	*requests.RpcRequest
	RoleName     string `position:"Query" name:"RoleName"`
	UserType     string `position:"Query" name:"UserType"`
	OperatorName string `position:"Query" name:"OperatorName"`
}

// RevokeRoleFromOperatorResponse is the response struct for api RevokeRoleFromOperator
type RevokeRoleFromOperatorResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
}

// CreateRevokeRoleFromOperatorRequest creates a request to invoke RevokeRoleFromOperator API
func CreateRevokeRoleFromOperatorRequest() (request *RevokeRoleFromOperatorRequest) {
	request = &RevokeRoleFromOperatorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "RevokeRoleFromOperator", "", "")
	request.Method = requests.POST
	return
}

// CreateRevokeRoleFromOperatorResponse creates a response to parse from RevokeRoleFromOperator response
func CreateRevokeRoleFromOperatorResponse() (response *RevokeRoleFromOperatorResponse) {
	response = &RevokeRoleFromOperatorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
