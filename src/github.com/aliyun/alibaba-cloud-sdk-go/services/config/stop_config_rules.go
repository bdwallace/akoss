package config

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

// StopConfigRules invokes the config.StopConfigRules API synchronously
func (client *Client) StopConfigRules(request *StopConfigRulesRequest) (response *StopConfigRulesResponse, err error) {
	response = CreateStopConfigRulesResponse()
	err = client.DoAction(request, response)
	return
}

// StopConfigRulesWithChan invokes the config.StopConfigRules API asynchronously
func (client *Client) StopConfigRulesWithChan(request *StopConfigRulesRequest) (<-chan *StopConfigRulesResponse, <-chan error) {
	responseChan := make(chan *StopConfigRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopConfigRules(request)
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

// StopConfigRulesWithCallback invokes the config.StopConfigRules API asynchronously
func (client *Client) StopConfigRulesWithCallback(request *StopConfigRulesRequest, callback func(response *StopConfigRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopConfigRulesResponse
		var err error
		defer close(result)
		response, err = client.StopConfigRules(request)
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

// StopConfigRulesRequest is the request struct for api StopConfigRules
type StopConfigRulesRequest struct {
	*requests.RpcRequest
	ConfigRuleIds string `position:"Query" name:"ConfigRuleIds"`
}

// StopConfigRulesResponse is the response struct for api StopConfigRules
type StopConfigRulesResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	OperateRuleResult OperateRuleResult `json:"OperateRuleResult" xml:"OperateRuleResult"`
}

// CreateStopConfigRulesRequest creates a request to invoke StopConfigRules API
func CreateStopConfigRulesRequest() (request *StopConfigRulesRequest) {
	request = &StopConfigRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "StopConfigRules", "Config", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopConfigRulesResponse creates a response to parse from StopConfigRules response
func CreateStopConfigRulesResponse() (response *StopConfigRulesResponse) {
	response = &StopConfigRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
