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

// ActiveConfigRules invokes the config.ActiveConfigRules API synchronously
func (client *Client) ActiveConfigRules(request *ActiveConfigRulesRequest) (response *ActiveConfigRulesResponse, err error) {
	response = CreateActiveConfigRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ActiveConfigRulesWithChan invokes the config.ActiveConfigRules API asynchronously
func (client *Client) ActiveConfigRulesWithChan(request *ActiveConfigRulesRequest) (<-chan *ActiveConfigRulesResponse, <-chan error) {
	responseChan := make(chan *ActiveConfigRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ActiveConfigRules(request)
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

// ActiveConfigRulesWithCallback invokes the config.ActiveConfigRules API asynchronously
func (client *Client) ActiveConfigRulesWithCallback(request *ActiveConfigRulesRequest, callback func(response *ActiveConfigRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ActiveConfigRulesResponse
		var err error
		defer close(result)
		response, err = client.ActiveConfigRules(request)
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

// ActiveConfigRulesRequest is the request struct for api ActiveConfigRules
type ActiveConfigRulesRequest struct {
	*requests.RpcRequest
	ConfigRuleIds string `position:"Query" name:"ConfigRuleIds"`
}

// ActiveConfigRulesResponse is the response struct for api ActiveConfigRules
type ActiveConfigRulesResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	OperateRuleResult OperateRuleResult `json:"OperateRuleResult" xml:"OperateRuleResult"`
}

// CreateActiveConfigRulesRequest creates a request to invoke ActiveConfigRules API
func CreateActiveConfigRulesRequest() (request *ActiveConfigRulesRequest) {
	request = &ActiveConfigRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2019-01-08", "ActiveConfigRules", "Config", "openAPI")
	request.Method = requests.POST
	return
}

// CreateActiveConfigRulesResponse creates a response to parse from ActiveConfigRules response
func CreateActiveConfigRulesResponse() (response *ActiveConfigRulesResponse) {
	response = &ActiveConfigRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
