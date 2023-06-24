package cloudcallcenter

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

// DescribeVnNavigationConfig invokes the cloudcallcenter.DescribeVnNavigationConfig API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/describevnnavigationconfig.html
func (client *Client) DescribeVnNavigationConfig(request *DescribeVnNavigationConfigRequest) (response *DescribeVnNavigationConfigResponse, err error) {
	response = CreateDescribeVnNavigationConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVnNavigationConfigWithChan invokes the cloudcallcenter.DescribeVnNavigationConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/describevnnavigationconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVnNavigationConfigWithChan(request *DescribeVnNavigationConfigRequest) (<-chan *DescribeVnNavigationConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeVnNavigationConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVnNavigationConfig(request)
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

// DescribeVnNavigationConfigWithCallback invokes the cloudcallcenter.DescribeVnNavigationConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/describevnnavigationconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVnNavigationConfigWithCallback(request *DescribeVnNavigationConfigRequest, callback func(response *DescribeVnNavigationConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVnNavigationConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeVnNavigationConfig(request)
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

// DescribeVnNavigationConfigRequest is the request struct for api DescribeVnNavigationConfig
type DescribeVnNavigationConfigRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeVnNavigationConfigResponse is the response struct for api DescribeVnNavigationConfig
type DescribeVnNavigationConfigResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	GreetingConfig       GreetingConfig       `json:"GreetingConfig" xml:"GreetingConfig"`
	UnrecognizingConfig  UnrecognizingConfig  `json:"UnrecognizingConfig" xml:"UnrecognizingConfig"`
	RepeatingConfig      RepeatingConfig      `json:"RepeatingConfig" xml:"RepeatingConfig"`
	AskingBackConfig     AskingBackConfig     `json:"AskingBackConfig" xml:"AskingBackConfig"`
	ComplainingConfig    ComplainingConfig    `json:"ComplainingConfig" xml:"ComplainingConfig"`
	SilenceTimeoutConfig SilenceTimeoutConfig `json:"SilenceTimeoutConfig" xml:"SilenceTimeoutConfig"`
}

// CreateDescribeVnNavigationConfigRequest creates a request to invoke DescribeVnNavigationConfig API
func CreateDescribeVnNavigationConfigRequest() (request *DescribeVnNavigationConfigRequest) {
	request = &DescribeVnNavigationConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DescribeVnNavigationConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeVnNavigationConfigResponse creates a response to parse from DescribeVnNavigationConfig response
func CreateDescribeVnNavigationConfigResponse() (response *DescribeVnNavigationConfigResponse) {
	response = &DescribeVnNavigationConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}