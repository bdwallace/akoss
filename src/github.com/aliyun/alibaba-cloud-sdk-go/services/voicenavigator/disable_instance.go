package voicenavigator

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

// DisableInstance invokes the voicenavigator.DisableInstance API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/disableinstance.html
func (client *Client) DisableInstance(request *DisableInstanceRequest) (response *DisableInstanceResponse, err error) {
	response = CreateDisableInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DisableInstanceWithChan invokes the voicenavigator.DisableInstance API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/disableinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableInstanceWithChan(request *DisableInstanceRequest) (<-chan *DisableInstanceResponse, <-chan error) {
	responseChan := make(chan *DisableInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableInstance(request)
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

// DisableInstanceWithCallback invokes the voicenavigator.DisableInstance API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/disableinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableInstanceWithCallback(request *DisableInstanceRequest, callback func(response *DisableInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableInstanceResponse
		var err error
		defer close(result)
		response, err = client.DisableInstance(request)
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

// DisableInstanceRequest is the request struct for api DisableInstance
type DisableInstanceRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DisableInstanceResponse is the response struct for api DisableInstance
type DisableInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateDisableInstanceRequest creates a request to invoke DisableInstance API
func CreateDisableInstanceRequest() (request *DisableInstanceRequest) {
	request = &DisableInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "DisableInstance", "voicebot", "openAPI")
	return
}

// CreateDisableInstanceResponse creates a response to parse from DisableInstance response
func CreateDisableInstanceResponse() (response *DisableInstanceResponse) {
	response = &DisableInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
