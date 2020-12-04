package openanalytics_open

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

// GetJobLog invokes the openanalytics_open.GetJobLog API synchronously
func (client *Client) GetJobLog(request *GetJobLogRequest) (response *GetJobLogResponse, err error) {
	response = CreateGetJobLogResponse()
	err = client.DoAction(request, response)
	return
}

// GetJobLogWithChan invokes the openanalytics_open.GetJobLog API asynchronously
func (client *Client) GetJobLogWithChan(request *GetJobLogRequest) (<-chan *GetJobLogResponse, <-chan error) {
	responseChan := make(chan *GetJobLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJobLog(request)
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

// GetJobLogWithCallback invokes the openanalytics_open.GetJobLog API asynchronously
func (client *Client) GetJobLogWithCallback(request *GetJobLogRequest, callback func(response *GetJobLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJobLogResponse
		var err error
		defer close(result)
		response, err = client.GetJobLog(request)
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

// GetJobLogRequest is the request struct for api GetJobLog
type GetJobLogRequest struct {
	*requests.RpcRequest
	JobId  string `position:"Body" name:"JobId"`
	VcName string `position:"Body" name:"VcName"`
}

// GetJobLogResponse is the response struct for api GetJobLog
type GetJobLogResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetJobLogRequest creates a request to invoke GetJobLog API
func CreateGetJobLogRequest() (request *GetJobLogRequest) {
	request = &GetJobLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "GetJobLog", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetJobLogResponse creates a response to parse from GetJobLog response
func CreateGetJobLogResponse() (response *GetJobLogResponse) {
	response = &GetJobLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
