package resourcemanager

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

// DeletePolicy invokes the resourcemanager.DeletePolicy API synchronously
// api document: https://help.aliyun.com/api/resourcemanager/deletepolicy.html
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
	response = CreateDeletePolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePolicyWithChan invokes the resourcemanager.DeletePolicy API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/deletepolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePolicyWithChan(request *DeletePolicyRequest) (<-chan *DeletePolicyResponse, <-chan error) {
	responseChan := make(chan *DeletePolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePolicy(request)
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

// DeletePolicyWithCallback invokes the resourcemanager.DeletePolicy API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/deletepolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePolicyWithCallback(request *DeletePolicyRequest, callback func(response *DeletePolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePolicyResponse
		var err error
		defer close(result)
		response, err = client.DeletePolicy(request)
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

// DeletePolicyRequest is the request struct for api DeletePolicy
type DeletePolicyRequest struct {
	*requests.RpcRequest
	PolicyName string `position:"Query" name:"PolicyName"`
}

// DeletePolicyResponse is the response struct for api DeletePolicy
type DeletePolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePolicyRequest creates a request to invoke DeletePolicy API
func CreateDeletePolicyRequest() (request *DeletePolicyRequest) {
	request = &DeletePolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "DeletePolicy", "resourcemanager", "openAPI")
	return
}

// CreateDeletePolicyResponse creates a response to parse from DeletePolicy response
func CreateDeletePolicyResponse() (response *DeletePolicyResponse) {
	response = &DeletePolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
