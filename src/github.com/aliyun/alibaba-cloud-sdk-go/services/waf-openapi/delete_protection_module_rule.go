package waf_openapi

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

// DeleteProtectionModuleRule invokes the waf_openapi.DeleteProtectionModuleRule API synchronously
func (client *Client) DeleteProtectionModuleRule(request *DeleteProtectionModuleRuleRequest) (response *DeleteProtectionModuleRuleResponse, err error) {
	response = CreateDeleteProtectionModuleRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteProtectionModuleRuleWithChan invokes the waf_openapi.DeleteProtectionModuleRule API asynchronously
func (client *Client) DeleteProtectionModuleRuleWithChan(request *DeleteProtectionModuleRuleRequest) (<-chan *DeleteProtectionModuleRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteProtectionModuleRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteProtectionModuleRule(request)
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

// DeleteProtectionModuleRuleWithCallback invokes the waf_openapi.DeleteProtectionModuleRule API asynchronously
func (client *Client) DeleteProtectionModuleRuleWithCallback(request *DeleteProtectionModuleRuleRequest, callback func(response *DeleteProtectionModuleRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteProtectionModuleRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteProtectionModuleRule(request)
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

// DeleteProtectionModuleRuleRequest is the request struct for api DeleteProtectionModuleRule
type DeleteProtectionModuleRuleRequest struct {
	*requests.RpcRequest
	DefenseType string           `position:"Query" name:"DefenseType"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Domain      string           `position:"Query" name:"Domain"`
	Lang        string           `position:"Query" name:"Lang"`
	RuleId      requests.Integer `position:"Query" name:"RuleId"`
}

// DeleteProtectionModuleRuleResponse is the response struct for api DeleteProtectionModuleRule
type DeleteProtectionModuleRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteProtectionModuleRuleRequest creates a request to invoke DeleteProtectionModuleRule API
func CreateDeleteProtectionModuleRuleRequest() (request *DeleteProtectionModuleRuleRequest) {
	request = &DeleteProtectionModuleRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DeleteProtectionModuleRule", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteProtectionModuleRuleResponse creates a response to parse from DeleteProtectionModuleRule response
func CreateDeleteProtectionModuleRuleResponse() (response *DeleteProtectionModuleRuleResponse) {
	response = &DeleteProtectionModuleRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
