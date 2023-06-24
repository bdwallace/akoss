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

// RecognizeLicensePlate invokes the ocr.RecognizeLicensePlate API synchronously
func (client *Client) RecognizeLicensePlate(request *RecognizeLicensePlateRequest) (response *RecognizeLicensePlateResponse, err error) {
	response = CreateRecognizeLicensePlateResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeLicensePlateWithChan invokes the ocr.RecognizeLicensePlate API asynchronously
func (client *Client) RecognizeLicensePlateWithChan(request *RecognizeLicensePlateRequest) (<-chan *RecognizeLicensePlateResponse, <-chan error) {
	responseChan := make(chan *RecognizeLicensePlateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeLicensePlate(request)
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

// RecognizeLicensePlateWithCallback invokes the ocr.RecognizeLicensePlate API asynchronously
func (client *Client) RecognizeLicensePlateWithCallback(request *RecognizeLicensePlateRequest, callback func(response *RecognizeLicensePlateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeLicensePlateResponse
		var err error
		defer close(result)
		response, err = client.RecognizeLicensePlate(request)
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

// RecognizeLicensePlateRequest is the request struct for api RecognizeLicensePlate
type RecognizeLicensePlateRequest struct {
	*requests.RpcRequest
	ImageType requests.Integer `position:"Body" name:"ImageType"`
	ImageURL  string           `position:"Body" name:"ImageURL"`
}

// RecognizeLicensePlateResponse is the response struct for api RecognizeLicensePlate
type RecognizeLicensePlateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeLicensePlateRequest creates a request to invoke RecognizeLicensePlate API
func CreateRecognizeLicensePlateRequest() (request *RecognizeLicensePlateRequest) {
	request = &RecognizeLicensePlateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeLicensePlate", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeLicensePlateResponse creates a response to parse from RecognizeLicensePlate response
func CreateRecognizeLicensePlateResponse() (response *RecognizeLicensePlateResponse) {
	response = &RecognizeLicensePlateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}