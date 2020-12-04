package green

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

// SendVerifyCodeToEmail invokes the green.SendVerifyCodeToEmail API synchronously
func (client *Client) SendVerifyCodeToEmail(request *SendVerifyCodeToEmailRequest) (response *SendVerifyCodeToEmailResponse, err error) {
	response = CreateSendVerifyCodeToEmailResponse()
	err = client.DoAction(request, response)
	return
}

// SendVerifyCodeToEmailWithChan invokes the green.SendVerifyCodeToEmail API asynchronously
func (client *Client) SendVerifyCodeToEmailWithChan(request *SendVerifyCodeToEmailRequest) (<-chan *SendVerifyCodeToEmailResponse, <-chan error) {
	responseChan := make(chan *SendVerifyCodeToEmailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendVerifyCodeToEmail(request)
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

// SendVerifyCodeToEmailWithCallback invokes the green.SendVerifyCodeToEmail API asynchronously
func (client *Client) SendVerifyCodeToEmailWithCallback(request *SendVerifyCodeToEmailRequest, callback func(response *SendVerifyCodeToEmailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendVerifyCodeToEmailResponse
		var err error
		defer close(result)
		response, err = client.SendVerifyCodeToEmail(request)
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

// SendVerifyCodeToEmailRequest is the request struct for api SendVerifyCodeToEmail
type SendVerifyCodeToEmailRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
	Email    string `position:"Query" name:"Email"`
}

// SendVerifyCodeToEmailResponse is the response struct for api SendVerifyCodeToEmail
type SendVerifyCodeToEmailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSendVerifyCodeToEmailRequest creates a request to invoke SendVerifyCodeToEmail API
func CreateSendVerifyCodeToEmailRequest() (request *SendVerifyCodeToEmailRequest) {
	request = &SendVerifyCodeToEmailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "SendVerifyCodeToEmail", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSendVerifyCodeToEmailResponse creates a response to parse from SendVerifyCodeToEmail response
func CreateSendVerifyCodeToEmailResponse() (response *SendVerifyCodeToEmailResponse) {
	response = &SendVerifyCodeToEmailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
