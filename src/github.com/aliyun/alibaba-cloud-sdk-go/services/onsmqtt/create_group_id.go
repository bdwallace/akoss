package onsmqtt

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

// CreateGroupId invokes the onsmqtt.CreateGroupId API synchronously
// api document: https://help.aliyun.com/api/onsmqtt/creategroupid.html
func (client *Client) CreateGroupId(request *CreateGroupIdRequest) (response *CreateGroupIdResponse, err error) {
	response = CreateCreateGroupIdResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGroupIdWithChan invokes the onsmqtt.CreateGroupId API asynchronously
// api document: https://help.aliyun.com/api/onsmqtt/creategroupid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGroupIdWithChan(request *CreateGroupIdRequest) (<-chan *CreateGroupIdResponse, <-chan error) {
	responseChan := make(chan *CreateGroupIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGroupId(request)
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

// CreateGroupIdWithCallback invokes the onsmqtt.CreateGroupId API asynchronously
// api document: https://help.aliyun.com/api/onsmqtt/creategroupid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGroupIdWithCallback(request *CreateGroupIdRequest, callback func(response *CreateGroupIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGroupIdResponse
		var err error
		defer close(result)
		response, err = client.CreateGroupId(request)
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

// CreateGroupIdRequest is the request struct for api CreateGroupId
type CreateGroupIdRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// CreateGroupIdResponse is the response struct for api CreateGroupId
type CreateGroupIdResponse struct {
	*responses.BaseResponse
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateGroupIdRequest creates a request to invoke CreateGroupId API
func CreateCreateGroupIdRequest() (request *CreateGroupIdRequest) {
	request = &CreateGroupIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OnsMqtt", "2020-04-20", "CreateGroupId", "onsmqtt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateGroupIdResponse creates a response to parse from CreateGroupId response
func CreateCreateGroupIdResponse() (response *CreateGroupIdResponse) {
	response = &CreateGroupIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
