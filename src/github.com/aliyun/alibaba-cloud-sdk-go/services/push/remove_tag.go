package push

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

// RemoveTag invokes the push.RemoveTag API synchronously
func (client *Client) RemoveTag(request *RemoveTagRequest) (response *RemoveTagResponse, err error) {
	response = CreateRemoveTagResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveTagWithChan invokes the push.RemoveTag API asynchronously
func (client *Client) RemoveTagWithChan(request *RemoveTagRequest) (<-chan *RemoveTagResponse, <-chan error) {
	responseChan := make(chan *RemoveTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveTag(request)
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

// RemoveTagWithCallback invokes the push.RemoveTag API asynchronously
func (client *Client) RemoveTagWithCallback(request *RemoveTagRequest, callback func(response *RemoveTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveTagResponse
		var err error
		defer close(result)
		response, err = client.RemoveTag(request)
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

// RemoveTagRequest is the request struct for api RemoveTag
type RemoveTagRequest struct {
	*requests.RpcRequest
	TagName string           `position:"Query" name:"TagName"`
	AppKey  requests.Integer `position:"Query" name:"AppKey"`
}

// RemoveTagResponse is the response struct for api RemoveTag
type RemoveTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveTagRequest creates a request to invoke RemoveTag API
func CreateRemoveTagRequest() (request *RemoveTagRequest) {
	request = &RemoveTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "RemoveTag", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveTagResponse creates a response to parse from RemoveTag response
func CreateRemoveTagResponse() (response *RemoveTagResponse) {
	response = &RemoveTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
