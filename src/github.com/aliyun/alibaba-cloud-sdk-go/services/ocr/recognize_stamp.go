package ocr

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

// RecognizeStamp invokes the ocr.RecognizeStamp API synchronously
func (client *Client) RecognizeStamp(request *RecognizeStampRequest) (response *RecognizeStampResponse, err error) {
	response = CreateRecognizeStampResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeStampWithChan invokes the ocr.RecognizeStamp API asynchronously
func (client *Client) RecognizeStampWithChan(request *RecognizeStampRequest) (<-chan *RecognizeStampResponse, <-chan error) {
	responseChan := make(chan *RecognizeStampResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeStamp(request)
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

// RecognizeStampWithCallback invokes the ocr.RecognizeStamp API asynchronously
func (client *Client) RecognizeStampWithCallback(request *RecognizeStampRequest, callback func(response *RecognizeStampResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeStampResponse
		var err error
		defer close(result)
		response, err = client.RecognizeStamp(request)
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

// RecognizeStampRequest is the request struct for api RecognizeStamp
type RecognizeStampRequest struct {
	*requests.RpcRequest
	ImageType requests.Integer `position:"Body" name:"ImageType"`
	ImageURL  string           `position:"Body" name:"ImageURL"`
}

// RecognizeStampResponse is the response struct for api RecognizeStamp
type RecognizeStampResponse struct {
	*responses.BaseResponse
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Data      DataInRecognizeStamp `json:"Data" xml:"Data"`
}

// CreateRecognizeStampRequest creates a request to invoke RecognizeStamp API
func CreateRecognizeStampRequest() (request *RecognizeStampRequest) {
	request = &RecognizeStampRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeStamp", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeStampResponse creates a response to parse from RecognizeStamp response
func CreateRecognizeStampResponse() (response *RecognizeStampResponse) {
	response = &RecognizeStampResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}