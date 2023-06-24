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

// CreateBoard invokes the live.CreateBoard API synchronously
func (client *Client) CreateBoard(request *CreateBoardRequest) (response *CreateBoardResponse, err error) {
	response = CreateCreateBoardResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBoardWithChan invokes the live.CreateBoard API asynchronously
func (client *Client) CreateBoardWithChan(request *CreateBoardRequest) (<-chan *CreateBoardResponse, <-chan error) {
	responseChan := make(chan *CreateBoardResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBoard(request)
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

// CreateBoardWithCallback invokes the live.CreateBoard API asynchronously
func (client *Client) CreateBoardWithCallback(request *CreateBoardRequest, callback func(response *CreateBoardResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBoardResponse
		var err error
		defer close(result)
		response, err = client.CreateBoard(request)
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

// CreateBoardRequest is the request struct for api CreateBoard
type CreateBoardRequest struct {
	*requests.RpcRequest
	AppUid  string           `position:"Query" name:"AppUid"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// CreateBoardResponse is the response struct for api CreateBoard
type CreateBoardResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	BoardId   string `json:"BoardId" xml:"BoardId"`
}

// CreateCreateBoardRequest creates a request to invoke CreateBoard API
func CreateCreateBoardRequest() (request *CreateBoardRequest) {
	request = &CreateBoardRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CreateBoard", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBoardResponse creates a response to parse from CreateBoard response
func CreateCreateBoardResponse() (response *CreateBoardResponse) {
	response = &CreateBoardResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}