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

// ListRoleCellsByRoleName invokes the oam.ListRoleCellsByRoleName API synchronously
// api document: https://help.aliyun.com/api/oam/listrolecellsbyrolename.html
func (client *Client) ListRoleCellsByRoleName(request *ListRoleCellsByRoleNameRequest) (response *ListRoleCellsByRoleNameResponse, err error) {
	response = CreateListRoleCellsByRoleNameResponse()
	err = client.DoAction(request, response)
	return
}

// ListRoleCellsByRoleNameWithChan invokes the oam.ListRoleCellsByRoleName API asynchronously
// api document: https://help.aliyun.com/api/oam/listrolecellsbyrolename.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRoleCellsByRoleNameWithChan(request *ListRoleCellsByRoleNameRequest) (<-chan *ListRoleCellsByRoleNameResponse, <-chan error) {
	responseChan := make(chan *ListRoleCellsByRoleNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRoleCellsByRoleName(request)
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

// ListRoleCellsByRoleNameWithCallback invokes the oam.ListRoleCellsByRoleName API asynchronously
// api document: https://help.aliyun.com/api/oam/listrolecellsbyrolename.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRoleCellsByRoleNameWithCallback(request *ListRoleCellsByRoleNameRequest, callback func(response *ListRoleCellsByRoleNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRoleCellsByRoleNameResponse
		var err error
		defer close(result)
		response, err = client.ListRoleCellsByRoleName(request)
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

// ListRoleCellsByRoleNameRequest is the request struct for api ListRoleCellsByRoleName
type ListRoleCellsByRoleNameRequest struct {
	*requests.RpcRequest
	RoleName  string           `position:"Query" name:"RoleName"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	PageIndex requests.Integer `position:"Query" name:"PageIndex"`
}

// ListRoleCellsByRoleNameResponse is the response struct for api ListRoleCellsByRoleName
type ListRoleCellsByRoleNameResponse struct {
	*responses.BaseResponse
	Code     string                        `json:"Code" xml:"Code"`
	Message  string                        `json:"Message" xml:"Message"`
	PageInfo PageInfo                      `json:"PageInfo" xml:"PageInfo"`
	Data     DataInListRoleCellsByRoleName `json:"Data" xml:"Data"`
}

// CreateListRoleCellsByRoleNameRequest creates a request to invoke ListRoleCellsByRoleName API
func CreateListRoleCellsByRoleNameRequest() (request *ListRoleCellsByRoleNameRequest) {
	request = &ListRoleCellsByRoleNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "ListRoleCellsByRoleName", "", "")
	request.Method = requests.POST
	return
}

// CreateListRoleCellsByRoleNameResponse creates a response to parse from ListRoleCellsByRoleName response
func CreateListRoleCellsByRoleNameResponse() (response *ListRoleCellsByRoleNameResponse) {
	response = &ListRoleCellsByRoleNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
