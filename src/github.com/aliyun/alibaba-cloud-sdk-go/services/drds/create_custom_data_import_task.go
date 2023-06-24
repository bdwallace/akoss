package drds

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

// CreateCustomDataImportTask invokes the drds.CreateCustomDataImportTask API synchronously
func (client *Client) CreateCustomDataImportTask(request *CreateCustomDataImportTaskRequest) (response *CreateCustomDataImportTaskResponse, err error) {
	response = CreateCreateCustomDataImportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCustomDataImportTaskWithChan invokes the drds.CreateCustomDataImportTask API asynchronously
func (client *Client) CreateCustomDataImportTaskWithChan(request *CreateCustomDataImportTaskRequest) (<-chan *CreateCustomDataImportTaskResponse, <-chan error) {
	responseChan := make(chan *CreateCustomDataImportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCustomDataImportTask(request)
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

// CreateCustomDataImportTaskWithCallback invokes the drds.CreateCustomDataImportTask API asynchronously
func (client *Client) CreateCustomDataImportTaskWithCallback(request *CreateCustomDataImportTaskRequest, callback func(response *CreateCustomDataImportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCustomDataImportTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateCustomDataImportTask(request)
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

// CreateCustomDataImportTaskRequest is the request struct for api CreateCustomDataImportTask
type CreateCustomDataImportTaskRequest struct {
	*requests.RpcRequest
	ImportParam string `position:"Query" name:"ImportParam"`
}

// CreateCustomDataImportTaskResponse is the response struct for api CreateCustomDataImportTask
type CreateCustomDataImportTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
}

// CreateCreateCustomDataImportTaskRequest creates a request to invoke CreateCustomDataImportTask API
func CreateCreateCustomDataImportTaskRequest() (request *CreateCustomDataImportTaskRequest) {
	request = &CreateCustomDataImportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "CreateCustomDataImportTask", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCustomDataImportTaskResponse creates a response to parse from CreateCustomDataImportTask response
func CreateCreateCustomDataImportTaskResponse() (response *CreateCustomDataImportTaskResponse) {
	response = &CreateCustomDataImportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}