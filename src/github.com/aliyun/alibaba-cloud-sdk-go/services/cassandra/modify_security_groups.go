package cassandra

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

// ModifySecurityGroups invokes the cassandra.ModifySecurityGroups API synchronously
func (client *Client) ModifySecurityGroups(request *ModifySecurityGroupsRequest) (response *ModifySecurityGroupsResponse, err error) {
	response = CreateModifySecurityGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySecurityGroupsWithChan invokes the cassandra.ModifySecurityGroups API asynchronously
func (client *Client) ModifySecurityGroupsWithChan(request *ModifySecurityGroupsRequest) (<-chan *ModifySecurityGroupsResponse, <-chan error) {
	responseChan := make(chan *ModifySecurityGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySecurityGroups(request)
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

// ModifySecurityGroupsWithCallback invokes the cassandra.ModifySecurityGroups API asynchronously
func (client *Client) ModifySecurityGroupsWithCallback(request *ModifySecurityGroupsRequest, callback func(response *ModifySecurityGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySecurityGroupsResponse
		var err error
		defer close(result)
		response, err = client.ModifySecurityGroups(request)
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

// ModifySecurityGroupsRequest is the request struct for api ModifySecurityGroups
type ModifySecurityGroupsRequest struct {
	*requests.RpcRequest
	ClusterId        string `position:"Query" name:"ClusterId"`
	SecurityGroupIds string `position:"Query" name:"SecurityGroupIds"`
}

// ModifySecurityGroupsResponse is the response struct for api ModifySecurityGroups
type ModifySecurityGroupsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySecurityGroupsRequest creates a request to invoke ModifySecurityGroups API
func CreateModifySecurityGroupsRequest() (request *ModifySecurityGroupsRequest) {
	request = &ModifySecurityGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "ModifySecurityGroups", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySecurityGroupsResponse creates a response to parse from ModifySecurityGroups response
func CreateModifySecurityGroupsResponse() (response *ModifySecurityGroupsResponse) {
	response = &ModifySecurityGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
