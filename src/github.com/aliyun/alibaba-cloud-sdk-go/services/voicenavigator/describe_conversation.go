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

// DescribeConversation invokes the voicenavigator.DescribeConversation API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/describeconversation.html
func (client *Client) DescribeConversation(request *DescribeConversationRequest) (response *DescribeConversationResponse, err error) {
	response = CreateDescribeConversationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConversationWithChan invokes the voicenavigator.DescribeConversation API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/describeconversation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConversationWithChan(request *DescribeConversationRequest) (<-chan *DescribeConversationResponse, <-chan error) {
	responseChan := make(chan *DescribeConversationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConversation(request)
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

// DescribeConversationWithCallback invokes the voicenavigator.DescribeConversation API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/describeconversation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConversationWithCallback(request *DescribeConversationRequest, callback func(response *DescribeConversationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConversationResponse
		var err error
		defer close(result)
		response, err = client.DescribeConversation(request)
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

// DescribeConversationRequest is the request struct for api DescribeConversation
type DescribeConversationRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	InstanceId     string `position:"Query" name:"InstanceId"`
}

// DescribeConversationResponse is the response struct for api DescribeConversation
type DescribeConversationResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	CallingNumber        string `json:"CallingNumber" xml:"CallingNumber"`
	BeginTime            int64  `json:"BeginTime" xml:"BeginTime"`
	EndTime              int64  `json:"EndTime" xml:"EndTime"`
	TransferredToAgent   bool   `json:"TransferredToAgent" xml:"TransferredToAgent"`
	SkillGroupId         string `json:"SkillGroupId" xml:"SkillGroupId"`
	UserUtteranceCount   int    `json:"UserUtteranceCount" xml:"UserUtteranceCount"`
	EffectiveAnswerCount int    `json:"EffectiveAnswerCount" xml:"EffectiveAnswerCount"`
	ConversationId       string `json:"ConversationId" xml:"ConversationId"`
}

// CreateDescribeConversationRequest creates a request to invoke DescribeConversation API
func CreateDescribeConversationRequest() (request *DescribeConversationRequest) {
	request = &DescribeConversationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "DescribeConversation", "voicebot", "openAPI")
	return
}

// CreateDescribeConversationResponse creates a response to parse from DescribeConversation response
func CreateDescribeConversationResponse() (response *DescribeConversationResponse) {
	response = &DescribeConversationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
