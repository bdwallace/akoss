package cloudwf

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

// OnoffGroupApRadio invokes the cloudwf.OnoffGroupApRadio API synchronously
// api document: https://help.aliyun.com/api/cloudwf/onoffgroupapradio.html
func (client *Client) OnoffGroupApRadio(request *OnoffGroupApRadioRequest) (response *OnoffGroupApRadioResponse, err error) {
	response = CreateOnoffGroupApRadioResponse()
	err = client.DoAction(request, response)
	return
}

// OnoffGroupApRadioWithChan invokes the cloudwf.OnoffGroupApRadio API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/onoffgroupapradio.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnoffGroupApRadioWithChan(request *OnoffGroupApRadioRequest) (<-chan *OnoffGroupApRadioResponse, <-chan error) {
	responseChan := make(chan *OnoffGroupApRadioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnoffGroupApRadio(request)
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

// OnoffGroupApRadioWithCallback invokes the cloudwf.OnoffGroupApRadio API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/onoffgroupapradio.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnoffGroupApRadioWithCallback(request *OnoffGroupApRadioRequest, callback func(response *OnoffGroupApRadioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnoffGroupApRadioResponse
		var err error
		defer close(result)
		response, err = client.OnoffGroupApRadio(request)
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

// OnoffGroupApRadioRequest is the request struct for api OnoffGroupApRadio
type OnoffGroupApRadioRequest struct {
	*requests.RpcRequest
	ApgroupId requests.Integer `position:"Query" name:"ApgroupId"`
	Disabled  requests.Integer `position:"Query" name:"Disabled"`
}

// OnoffGroupApRadioResponse is the response struct for api OnoffGroupApRadio
type OnoffGroupApRadioResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateOnoffGroupApRadioRequest creates a request to invoke OnoffGroupApRadio API
func CreateOnoffGroupApRadioRequest() (request *OnoffGroupApRadioRequest) {
	request = &OnoffGroupApRadioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "OnoffGroupApRadio", "cloudwf", "openAPI")
	return
}

// CreateOnoffGroupApRadioResponse creates a response to parse from OnoffGroupApRadio response
func CreateOnoffGroupApRadioResponse() (response *OnoffGroupApRadioResponse) {
	response = &OnoffGroupApRadioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
