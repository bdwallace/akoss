package cloudauth

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

// CreateVerifySDK invokes the cloudauth.CreateVerifySDK API synchronously
func (client *Client) CreateVerifySDK(request *CreateVerifySDKRequest) (response *CreateVerifySDKResponse, err error) {
	response = CreateCreateVerifySDKResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVerifySDKWithChan invokes the cloudauth.CreateVerifySDK API asynchronously
func (client *Client) CreateVerifySDKWithChan(request *CreateVerifySDKRequest) (<-chan *CreateVerifySDKResponse, <-chan error) {
	responseChan := make(chan *CreateVerifySDKResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVerifySDK(request)
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

// CreateVerifySDKWithCallback invokes the cloudauth.CreateVerifySDK API asynchronously
func (client *Client) CreateVerifySDKWithCallback(request *CreateVerifySDKRequest, callback func(response *CreateVerifySDKResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVerifySDKResponse
		var err error
		defer close(result)
		response, err = client.CreateVerifySDK(request)
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

// CreateVerifySDKRequest is the request struct for api CreateVerifySDK
type CreateVerifySDKRequest struct {
	*requests.RpcRequest
	AppUrl   string `position:"Query" name:"AppUrl"`
	Platform string `position:"Query" name:"Platform"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
}

// CreateVerifySDKResponse is the response struct for api CreateVerifySDK
type CreateVerifySDKResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateVerifySDKRequest creates a request to invoke CreateVerifySDK API
func CreateCreateVerifySDKRequest() (request *CreateVerifySDKRequest) {
	request = &CreateVerifySDKRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2018-09-16", "CreateVerifySDK", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVerifySDKResponse creates a response to parse from CreateVerifySDK response
func CreateCreateVerifySDKResponse() (response *CreateVerifySDKResponse) {
	response = &CreateVerifySDKResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}