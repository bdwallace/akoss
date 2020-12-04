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

// CheckSlrRole invokes the sgw.CheckSlrRole API synchronously
func (client *Client) CheckSlrRole(request *CheckSlrRoleRequest) (response *CheckSlrRoleResponse, err error) {
	response = CreateCheckSlrRoleResponse()
	err = client.DoAction(request, response)
	return
}

// CheckSlrRoleWithChan invokes the sgw.CheckSlrRole API asynchronously
func (client *Client) CheckSlrRoleWithChan(request *CheckSlrRoleRequest) (<-chan *CheckSlrRoleResponse, <-chan error) {
	responseChan := make(chan *CheckSlrRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckSlrRole(request)
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

// CheckSlrRoleWithCallback invokes the sgw.CheckSlrRole API asynchronously
func (client *Client) CheckSlrRoleWithCallback(request *CheckSlrRoleRequest, callback func(response *CheckSlrRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckSlrRoleResponse
		var err error
		defer close(result)
		response, err = client.CheckSlrRole(request)
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

// CheckSlrRoleRequest is the request struct for api CheckSlrRole
type CheckSlrRoleRequest struct {
	*requests.RpcRequest
	RoleName         string           `position:"Query" name:"RoleName"`
	CreateIfNotExist requests.Boolean `position:"Query" name:"CreateIfNotExist"`
	SecurityToken    string           `position:"Query" name:"SecurityToken"`
}

// CheckSlrRoleResponse is the response struct for api CheckSlrRole
type CheckSlrRoleResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	Success            bool   `json:"Success" xml:"Success"`
	Code               string `json:"Code" xml:"Code"`
	Message            string `json:"Message" xml:"Message"`
	Exist              bool   `json:"Exist" xml:"Exist"`
	RequireOldWayCheck bool   `json:"RequireOldWayCheck" xml:"RequireOldWayCheck"`
}

// CreateCheckSlrRoleRequest creates a request to invoke CheckSlrRole API
func CreateCheckSlrRoleRequest() (request *CheckSlrRoleRequest) {
	request = &CheckSlrRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "CheckSlrRole", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckSlrRoleResponse creates a response to parse from CheckSlrRole response
func CreateCheckSlrRoleResponse() (response *CheckSlrRoleResponse) {
	response = &CheckSlrRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
