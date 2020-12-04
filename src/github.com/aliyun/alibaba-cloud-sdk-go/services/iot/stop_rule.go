package iot

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

// StopRule invokes the iot.StopRule API synchronously
func (client *Client) StopRule(request *StopRuleRequest) (response *StopRuleResponse, err error) {
	response = CreateStopRuleResponse()
	err = client.DoAction(request, response)
	return
}

// StopRuleWithChan invokes the iot.StopRule API asynchronously
func (client *Client) StopRuleWithChan(request *StopRuleRequest) (<-chan *StopRuleResponse, <-chan error) {
	responseChan := make(chan *StopRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopRule(request)
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

// StopRuleWithCallback invokes the iot.StopRule API asynchronously
func (client *Client) StopRuleWithCallback(request *StopRuleRequest, callback func(response *StopRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopRuleResponse
		var err error
		defer close(result)
		response, err = client.StopRule(request)
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

// StopRuleRequest is the request struct for api StopRule
type StopRuleRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	RuleId        requests.Integer `position:"Query" name:"RuleId"`
}

// StopRuleResponse is the response struct for api StopRule
type StopRuleResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateStopRuleRequest creates a request to invoke StopRule API
func CreateStopRuleRequest() (request *StopRuleRequest) {
	request = &StopRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "StopRule", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopRuleResponse creates a response to parse from StopRule response
func CreateStopRuleResponse() (response *StopRuleResponse) {
	response = &StopRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
