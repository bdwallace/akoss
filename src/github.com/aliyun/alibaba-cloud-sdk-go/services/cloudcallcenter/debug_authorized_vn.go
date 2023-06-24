package cloudcallcenter

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

// DebugAuthorizedVn invokes the cloudcallcenter.DebugAuthorizedVn API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugauthorizedvn.html
func (client *Client) DebugAuthorizedVn(request *DebugAuthorizedVnRequest) (response *DebugAuthorizedVnResponse, err error) {
	response = CreateDebugAuthorizedVnResponse()
	err = client.DoAction(request, response)
	return
}

// DebugAuthorizedVnWithChan invokes the cloudcallcenter.DebugAuthorizedVn API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugauthorizedvn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugAuthorizedVnWithChan(request *DebugAuthorizedVnRequest) (<-chan *DebugAuthorizedVnResponse, <-chan error) {
	responseChan := make(chan *DebugAuthorizedVnResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DebugAuthorizedVn(request)
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

// DebugAuthorizedVnWithCallback invokes the cloudcallcenter.DebugAuthorizedVn API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugauthorizedvn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugAuthorizedVnWithCallback(request *DebugAuthorizedVnRequest, callback func(response *DebugAuthorizedVnResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DebugAuthorizedVnResponse
		var err error
		defer close(result)
		response, err = client.DebugAuthorizedVn(request)
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

// DebugAuthorizedVnRequest is the request struct for api DebugAuthorizedVn
type DebugAuthorizedVnRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	InitialContext string `position:"Query" name:"InitialContext"`
	InstanceId     string `position:"Query" name:"InstanceId"`
}

// DebugAuthorizedVnResponse is the response struct for api DebugAuthorizedVn
type DebugAuthorizedVnResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	Action        string `json:"Action" xml:"Action"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
}

// CreateDebugAuthorizedVnRequest creates a request to invoke DebugAuthorizedVn API
func CreateDebugAuthorizedVnRequest() (request *DebugAuthorizedVnRequest) {
	request = &DebugAuthorizedVnRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DebugAuthorizedVn", "", "")
	request.Method = requests.GET
	return
}

// CreateDebugAuthorizedVnResponse creates a response to parse from DebugAuthorizedVn response
func CreateDebugAuthorizedVnResponse() (response *DebugAuthorizedVnResponse) {
	response = &DebugAuthorizedVnResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
