package arms

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

// SearchAlertContactGroup invokes the arms.SearchAlertContactGroup API synchronously
func (client *Client) SearchAlertContactGroup(request *SearchAlertContactGroupRequest) (response *SearchAlertContactGroupResponse, err error) {
	response = CreateSearchAlertContactGroupResponse()
	err = client.DoAction(request, response)
	return
}

// SearchAlertContactGroupWithChan invokes the arms.SearchAlertContactGroup API asynchronously
func (client *Client) SearchAlertContactGroupWithChan(request *SearchAlertContactGroupRequest) (<-chan *SearchAlertContactGroupResponse, <-chan error) {
	responseChan := make(chan *SearchAlertContactGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchAlertContactGroup(request)
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

// SearchAlertContactGroupWithCallback invokes the arms.SearchAlertContactGroup API asynchronously
func (client *Client) SearchAlertContactGroupWithCallback(request *SearchAlertContactGroupRequest, callback func(response *SearchAlertContactGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchAlertContactGroupResponse
		var err error
		defer close(result)
		response, err = client.SearchAlertContactGroup(request)
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

// SearchAlertContactGroupRequest is the request struct for api SearchAlertContactGroup
type SearchAlertContactGroupRequest struct {
	*requests.RpcRequest
	ContactId        requests.Integer `position:"Query" name:"ContactId"`
	IsDetail         requests.Boolean `position:"Query" name:"IsDetail"`
	ContactGroupName string           `position:"Query" name:"ContactGroupName"`
	ProxyUserId      string           `position:"Query" name:"ProxyUserId"`
	ContactName      string           `position:"Query" name:"ContactName"`
	ContactGroupIds  string           `position:"Query" name:"ContactGroupIds"`
}

// SearchAlertContactGroupResponse is the response struct for api SearchAlertContactGroup
type SearchAlertContactGroupResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	ContactGroups []ContactGroup `json:"ContactGroups" xml:"ContactGroups"`
}

// CreateSearchAlertContactGroupRequest creates a request to invoke SearchAlertContactGroup API
func CreateSearchAlertContactGroupRequest() (request *SearchAlertContactGroupRequest) {
	request = &SearchAlertContactGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "SearchAlertContactGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateSearchAlertContactGroupResponse creates a response to parse from SearchAlertContactGroup response
func CreateSearchAlertContactGroupResponse() (response *SearchAlertContactGroupResponse) {
	response = &SearchAlertContactGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
