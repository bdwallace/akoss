package cloudapi

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

// ModifyLogConfig invokes the cloudapi.ModifyLogConfig API synchronously
// api document: https://help.aliyun.com/api/cloudapi/modifylogconfig.html
func (client *Client) ModifyLogConfig(request *ModifyLogConfigRequest) (response *ModifyLogConfigResponse, err error) {
	response = CreateModifyLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLogConfigWithChan invokes the cloudapi.ModifyLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/modifylogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLogConfigWithChan(request *ModifyLogConfigRequest) (<-chan *ModifyLogConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLogConfig(request)
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

// ModifyLogConfigWithCallback invokes the cloudapi.ModifyLogConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/modifylogconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLogConfigWithCallback(request *ModifyLogConfigRequest, callback func(response *ModifyLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLogConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyLogConfig(request)
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

// ModifyLogConfigRequest is the request struct for api ModifyLogConfig
type ModifyLogConfigRequest struct {
	*requests.RpcRequest
	SlsLogStore   string `position:"Query" name:"SlsLogStore"`
	SlsProject    string `position:"Query" name:"SlsProject"`
	LogType       string `position:"Query" name:"LogType"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// ModifyLogConfigResponse is the response struct for api ModifyLogConfig
type ModifyLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLogConfigRequest creates a request to invoke ModifyLogConfig API
func CreateModifyLogConfigRequest() (request *ModifyLogConfigRequest) {
	request = &ModifyLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyLogConfig", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLogConfigResponse creates a response to parse from ModifyLogConfig response
func CreateModifyLogConfigResponse() (response *ModifyLogConfigResponse) {
	response = &ModifyLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
