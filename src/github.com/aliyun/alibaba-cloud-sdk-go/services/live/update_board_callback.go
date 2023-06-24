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

// UpdateBoardCallback invokes the live.UpdateBoardCallback API synchronously
func (client *Client) UpdateBoardCallback(request *UpdateBoardCallbackRequest) (response *UpdateBoardCallbackResponse, err error) {
	response = CreateUpdateBoardCallbackResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateBoardCallbackWithChan invokes the live.UpdateBoardCallback API asynchronously
func (client *Client) UpdateBoardCallbackWithChan(request *UpdateBoardCallbackRequest) (<-chan *UpdateBoardCallbackResponse, <-chan error) {
	responseChan := make(chan *UpdateBoardCallbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateBoardCallback(request)
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

// UpdateBoardCallbackWithCallback invokes the live.UpdateBoardCallback API asynchronously
func (client *Client) UpdateBoardCallbackWithCallback(request *UpdateBoardCallbackRequest, callback func(response *UpdateBoardCallbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateBoardCallbackResponse
		var err error
		defer close(result)
		response, err = client.UpdateBoardCallback(request)
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

// UpdateBoardCallbackRequest is the request struct for api UpdateBoardCallback
type UpdateBoardCallbackRequest struct {
	*requests.RpcRequest
	AuthKey        string           `position:"Query" name:"AuthKey"`
	CallbackEnable requests.Integer `position:"Query" name:"CallbackEnable"`
	CallbackEvents string           `position:"Query" name:"CallbackEvents"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	CallbackUri    string           `position:"Query" name:"CallbackUri"`
	AppId          string           `position:"Query" name:"AppId"`
	AuthSwitch     string           `position:"Query" name:"AuthSwitch"`
}

// UpdateBoardCallbackResponse is the response struct for api UpdateBoardCallback
type UpdateBoardCallbackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateBoardCallbackRequest creates a request to invoke UpdateBoardCallback API
func CreateUpdateBoardCallbackRequest() (request *UpdateBoardCallbackRequest) {
	request = &UpdateBoardCallbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateBoardCallback", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateBoardCallbackResponse creates a response to parse from UpdateBoardCallback response
func CreateUpdateBoardCallbackResponse() (response *UpdateBoardCallbackResponse) {
	response = &UpdateBoardCallbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}