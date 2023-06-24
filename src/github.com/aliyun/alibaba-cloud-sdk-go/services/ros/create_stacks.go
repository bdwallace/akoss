package ros

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

// CreateStacks invokes the ros.CreateStacks API synchronously
// api document: https://help.aliyun.com/api/ros/createstacks.html
func (client *Client) CreateStacks(request *CreateStacksRequest) (response *CreateStacksResponse, err error) {
	response = CreateCreateStacksResponse()
	err = client.DoAction(request, response)
	return
}

// CreateStacksWithChan invokes the ros.CreateStacks API asynchronously
// api document: https://help.aliyun.com/api/ros/createstacks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStacksWithChan(request *CreateStacksRequest) (<-chan *CreateStacksResponse, <-chan error) {
	responseChan := make(chan *CreateStacksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateStacks(request)
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

// CreateStacksWithCallback invokes the ros.CreateStacks API asynchronously
// api document: https://help.aliyun.com/api/ros/createstacks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStacksWithCallback(request *CreateStacksRequest, callback func(response *CreateStacksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateStacksResponse
		var err error
		defer close(result)
		response, err = client.CreateStacks(request)
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

// CreateStacksRequest is the request struct for api CreateStacks
type CreateStacksRequest struct {
	*requests.RoaRequest
}

// CreateStacksResponse is the response struct for api CreateStacks
type CreateStacksResponse struct {
	*responses.BaseResponse
}

// CreateCreateStacksRequest creates a request to invoke CreateStacks API
func CreateCreateStacksRequest() (request *CreateStacksRequest) {
	request = &CreateStacksRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "CreateStacks", "/stacks", "ROS", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateStacksResponse creates a response to parse from CreateStacks response
func CreateCreateStacksResponse() (response *CreateStacksResponse) {
	response = &CreateStacksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
