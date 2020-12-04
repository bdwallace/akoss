package alimt

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

// GetImageTranslate invokes the alimt.GetImageTranslate API synchronously
func (client *Client) GetImageTranslate(request *GetImageTranslateRequest) (response *GetImageTranslateResponse, err error) {
	response = CreateGetImageTranslateResponse()
	err = client.DoAction(request, response)
	return
}

// GetImageTranslateWithChan invokes the alimt.GetImageTranslate API asynchronously
func (client *Client) GetImageTranslateWithChan(request *GetImageTranslateRequest) (<-chan *GetImageTranslateResponse, <-chan error) {
	responseChan := make(chan *GetImageTranslateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetImageTranslate(request)
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

// GetImageTranslateWithCallback invokes the alimt.GetImageTranslate API asynchronously
func (client *Client) GetImageTranslateWithCallback(request *GetImageTranslateRequest, callback func(response *GetImageTranslateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetImageTranslateResponse
		var err error
		defer close(result)
		response, err = client.GetImageTranslate(request)
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

// GetImageTranslateRequest is the request struct for api GetImageTranslate
type GetImageTranslateRequest struct {
	*requests.RpcRequest
	SourceLanguage string `position:"Body" name:"SourceLanguage"`
	Url            string `position:"Body" name:"Url"`
	Extra          string `position:"Body" name:"Extra"`
	TargetLanguage string `position:"Body" name:"TargetLanguage"`
}

// GetImageTranslateResponse is the response struct for api GetImageTranslate
type GetImageTranslateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetImageTranslateRequest creates a request to invoke GetImageTranslate API
func CreateGetImageTranslateRequest() (request *GetImageTranslateRequest) {
	request = &GetImageTranslateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "GetImageTranslate", "alimt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetImageTranslateResponse creates a response to parse from GetImageTranslate response
func CreateGetImageTranslateResponse() (response *GetImageTranslateResponse) {
	response = &GetImageTranslateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
