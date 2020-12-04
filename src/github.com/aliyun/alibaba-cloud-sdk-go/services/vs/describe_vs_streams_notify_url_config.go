package vs

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

// DescribeVsStreamsNotifyUrlConfig invokes the vs.DescribeVsStreamsNotifyUrlConfig API synchronously
func (client *Client) DescribeVsStreamsNotifyUrlConfig(request *DescribeVsStreamsNotifyUrlConfigRequest) (response *DescribeVsStreamsNotifyUrlConfigResponse, err error) {
	response = CreateDescribeVsStreamsNotifyUrlConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVsStreamsNotifyUrlConfigWithChan invokes the vs.DescribeVsStreamsNotifyUrlConfig API asynchronously
func (client *Client) DescribeVsStreamsNotifyUrlConfigWithChan(request *DescribeVsStreamsNotifyUrlConfigRequest) (<-chan *DescribeVsStreamsNotifyUrlConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeVsStreamsNotifyUrlConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVsStreamsNotifyUrlConfig(request)
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

// DescribeVsStreamsNotifyUrlConfigWithCallback invokes the vs.DescribeVsStreamsNotifyUrlConfig API asynchronously
func (client *Client) DescribeVsStreamsNotifyUrlConfigWithCallback(request *DescribeVsStreamsNotifyUrlConfigRequest, callback func(response *DescribeVsStreamsNotifyUrlConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVsStreamsNotifyUrlConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeVsStreamsNotifyUrlConfig(request)
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

// DescribeVsStreamsNotifyUrlConfigRequest is the request struct for api DescribeVsStreamsNotifyUrlConfig
type DescribeVsStreamsNotifyUrlConfigRequest struct {
	*requests.RpcRequest
	ShowLog    string           `position:"Query" name:"ShowLog"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVsStreamsNotifyUrlConfigResponse is the response struct for api DescribeVsStreamsNotifyUrlConfig
type DescribeVsStreamsNotifyUrlConfigResponse struct {
	*responses.BaseResponse
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	LiveStreamsNotifyConfig LiveStreamsNotifyConfig `json:"LiveStreamsNotifyConfig" xml:"LiveStreamsNotifyConfig"`
}

// CreateDescribeVsStreamsNotifyUrlConfigRequest creates a request to invoke DescribeVsStreamsNotifyUrlConfig API
func CreateDescribeVsStreamsNotifyUrlConfigRequest() (request *DescribeVsStreamsNotifyUrlConfigRequest) {
	request = &DescribeVsStreamsNotifyUrlConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeVsStreamsNotifyUrlConfig", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVsStreamsNotifyUrlConfigResponse creates a response to parse from DescribeVsStreamsNotifyUrlConfig response
func CreateDescribeVsStreamsNotifyUrlConfigResponse() (response *DescribeVsStreamsNotifyUrlConfigResponse) {
	response = &DescribeVsStreamsNotifyUrlConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
