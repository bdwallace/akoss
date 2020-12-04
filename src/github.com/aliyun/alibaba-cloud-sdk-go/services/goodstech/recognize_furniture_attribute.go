package goodstech

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

// RecognizeFurnitureAttribute invokes the goodstech.RecognizeFurnitureAttribute API synchronously
func (client *Client) RecognizeFurnitureAttribute(request *RecognizeFurnitureAttributeRequest) (response *RecognizeFurnitureAttributeResponse, err error) {
	response = CreateRecognizeFurnitureAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeFurnitureAttributeWithChan invokes the goodstech.RecognizeFurnitureAttribute API asynchronously
func (client *Client) RecognizeFurnitureAttributeWithChan(request *RecognizeFurnitureAttributeRequest) (<-chan *RecognizeFurnitureAttributeResponse, <-chan error) {
	responseChan := make(chan *RecognizeFurnitureAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeFurnitureAttribute(request)
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

// RecognizeFurnitureAttributeWithCallback invokes the goodstech.RecognizeFurnitureAttribute API asynchronously
func (client *Client) RecognizeFurnitureAttributeWithCallback(request *RecognizeFurnitureAttributeRequest, callback func(response *RecognizeFurnitureAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeFurnitureAttributeResponse
		var err error
		defer close(result)
		response, err = client.RecognizeFurnitureAttribute(request)
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

// RecognizeFurnitureAttributeRequest is the request struct for api RecognizeFurnitureAttribute
type RecognizeFurnitureAttributeRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// RecognizeFurnitureAttributeResponse is the response struct for api RecognizeFurnitureAttribute
type RecognizeFurnitureAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeFurnitureAttributeRequest creates a request to invoke RecognizeFurnitureAttribute API
func CreateRecognizeFurnitureAttributeRequest() (request *RecognizeFurnitureAttributeRequest) {
	request = &RecognizeFurnitureAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("goodstech", "2019-12-30", "RecognizeFurnitureAttribute", "goodstech", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeFurnitureAttributeResponse creates a response to parse from RecognizeFurnitureAttribute response
func CreateRecognizeFurnitureAttributeResponse() (response *RecognizeFurnitureAttributeResponse) {
	response = &RecognizeFurnitureAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
