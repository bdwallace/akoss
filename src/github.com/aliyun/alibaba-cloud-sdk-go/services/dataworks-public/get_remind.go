package dataworks_public

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

// GetRemind invokes the dataworks_public.GetRemind API synchronously
func (client *Client) GetRemind(request *GetRemindRequest) (response *GetRemindResponse, err error) {
	response = CreateGetRemindResponse()
	err = client.DoAction(request, response)
	return
}

// GetRemindWithChan invokes the dataworks_public.GetRemind API asynchronously
func (client *Client) GetRemindWithChan(request *GetRemindRequest) (<-chan *GetRemindResponse, <-chan error) {
	responseChan := make(chan *GetRemindResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRemind(request)
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

// GetRemindWithCallback invokes the dataworks_public.GetRemind API asynchronously
func (client *Client) GetRemindWithCallback(request *GetRemindRequest, callback func(response *GetRemindResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRemindResponse
		var err error
		defer close(result)
		response, err = client.GetRemind(request)
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

// GetRemindRequest is the request struct for api GetRemind
type GetRemindRequest struct {
	*requests.RpcRequest
	RemindId requests.Integer `position:"Body" name:"RemindId"`
}

// GetRemindResponse is the response struct for api GetRemind
type GetRemindResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetRemindRequest creates a request to invoke GetRemind API
func CreateGetRemindRequest() (request *GetRemindRequest) {
	request = &GetRemindRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetRemind", "", "")
	request.Method = requests.POST
	return
}

// CreateGetRemindResponse creates a response to parse from GetRemind response
func CreateGetRemindResponse() (response *GetRemindResponse) {
	response = &GetRemindResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
