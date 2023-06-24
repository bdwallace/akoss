package ccc

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

// TaskPreparing invokes the ccc.TaskPreparing API synchronously
// api document: https://help.aliyun.com/api/ccc/taskpreparing.html
func (client *Client) TaskPreparing(request *TaskPreparingRequest) (response *TaskPreparingResponse, err error) {
	response = CreateTaskPreparingResponse()
	err = client.DoAction(request, response)
	return
}

// TaskPreparingWithChan invokes the ccc.TaskPreparing API asynchronously
// api document: https://help.aliyun.com/api/ccc/taskpreparing.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskPreparingWithChan(request *TaskPreparingRequest) (<-chan *TaskPreparingResponse, <-chan error) {
	responseChan := make(chan *TaskPreparingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TaskPreparing(request)
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

// TaskPreparingWithCallback invokes the ccc.TaskPreparing API asynchronously
// api document: https://help.aliyun.com/api/ccc/taskpreparing.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TaskPreparingWithCallback(request *TaskPreparingRequest, callback func(response *TaskPreparingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TaskPreparingResponse
		var err error
		defer close(result)
		response, err = client.TaskPreparing(request)
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

// TaskPreparingRequest is the request struct for api TaskPreparing
type TaskPreparingRequest struct {
	*requests.RpcRequest
	JobId           string           `position:"Query" name:"JobId"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	InstanceOwnerId requests.Integer `position:"Query" name:"InstanceOwnerId"`
}

// TaskPreparingResponse is the response struct for api TaskPreparing
type TaskPreparingResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
}

// CreateTaskPreparingRequest creates a request to invoke TaskPreparing API
func CreateTaskPreparingRequest() (request *TaskPreparingRequest) {
	request = &TaskPreparingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "TaskPreparing", "", "")
	return
}

// CreateTaskPreparingResponse creates a response to parse from TaskPreparing response
func CreateTaskPreparingResponse() (response *TaskPreparingResponse) {
	response = &TaskPreparingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}