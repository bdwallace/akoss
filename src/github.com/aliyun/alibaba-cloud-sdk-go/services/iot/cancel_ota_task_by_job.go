package iot

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

// CancelOTATaskByJob invokes the iot.CancelOTATaskByJob API synchronously
func (client *Client) CancelOTATaskByJob(request *CancelOTATaskByJobRequest) (response *CancelOTATaskByJobResponse, err error) {
	response = CreateCancelOTATaskByJobResponse()
	err = client.DoAction(request, response)
	return
}

// CancelOTATaskByJobWithChan invokes the iot.CancelOTATaskByJob API asynchronously
func (client *Client) CancelOTATaskByJobWithChan(request *CancelOTATaskByJobRequest) (<-chan *CancelOTATaskByJobResponse, <-chan error) {
	responseChan := make(chan *CancelOTATaskByJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelOTATaskByJob(request)
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

// CancelOTATaskByJobWithCallback invokes the iot.CancelOTATaskByJob API asynchronously
func (client *Client) CancelOTATaskByJobWithCallback(request *CancelOTATaskByJobRequest, callback func(response *CancelOTATaskByJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelOTATaskByJobResponse
		var err error
		defer close(result)
		response, err = client.CancelOTATaskByJob(request)
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

// CancelOTATaskByJobRequest is the request struct for api CancelOTATaskByJob
type CancelOTATaskByJobRequest struct {
	*requests.RpcRequest
	CancelScheduledTask  requests.Boolean `position:"Query" name:"CancelScheduledTask"`
	JobId                string           `position:"Query" name:"JobId"`
	IotInstanceId        string           `position:"Query" name:"IotInstanceId"`
	CancelQueuedTask     requests.Boolean `position:"Query" name:"CancelQueuedTask"`
	CancelInProgressTask requests.Boolean `position:"Query" name:"CancelInProgressTask"`
	CancelNotifiedTask   requests.Boolean `position:"Query" name:"CancelNotifiedTask"`
	ApiProduct           string           `position:"Body" name:"ApiProduct"`
	ApiRevision          string           `position:"Body" name:"ApiRevision"`
}

// CancelOTATaskByJobResponse is the response struct for api CancelOTATaskByJob
type CancelOTATaskByJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateCancelOTATaskByJobRequest creates a request to invoke CancelOTATaskByJob API
func CreateCancelOTATaskByJobRequest() (request *CancelOTATaskByJobRequest) {
	request = &CancelOTATaskByJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CancelOTATaskByJob", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelOTATaskByJobResponse creates a response to parse from CancelOTATaskByJob response
func CreateCancelOTATaskByJobResponse() (response *CancelOTATaskByJobResponse) {
	response = &CancelOTATaskByJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
