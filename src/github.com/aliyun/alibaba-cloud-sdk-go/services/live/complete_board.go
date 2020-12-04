package live

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

// CompleteBoard invokes the live.CompleteBoard API synchronously
func (client *Client) CompleteBoard(request *CompleteBoardRequest) (response *CompleteBoardResponse, err error) {
	response = CreateCompleteBoardResponse()
	err = client.DoAction(request, response)
	return
}

// CompleteBoardWithChan invokes the live.CompleteBoard API asynchronously
func (client *Client) CompleteBoardWithChan(request *CompleteBoardRequest) (<-chan *CompleteBoardResponse, <-chan error) {
	responseChan := make(chan *CompleteBoardResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CompleteBoard(request)
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

// CompleteBoardWithCallback invokes the live.CompleteBoard API asynchronously
func (client *Client) CompleteBoardWithCallback(request *CompleteBoardRequest, callback func(response *CompleteBoardResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CompleteBoardResponse
		var err error
		defer close(result)
		response, err = client.CompleteBoard(request)
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

// CompleteBoardRequest is the request struct for api CompleteBoard
type CompleteBoardRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
	BoardId string           `position:"Query" name:"BoardId"`
}

// CompleteBoardResponse is the response struct for api CompleteBoard
type CompleteBoardResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCompleteBoardRequest creates a request to invoke CompleteBoard API
func CreateCompleteBoardRequest() (request *CompleteBoardRequest) {
	request = &CompleteBoardRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CompleteBoard", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCompleteBoardResponse creates a response to parse from CompleteBoard response
func CreateCompleteBoardResponse() (response *CompleteBoardResponse) {
	response = &CompleteBoardResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
