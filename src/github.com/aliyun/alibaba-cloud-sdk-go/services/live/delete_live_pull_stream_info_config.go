package live

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

// DeleteLivePullStreamInfoConfig invokes the live.DeleteLivePullStreamInfoConfig API synchronously
func (client *Client) DeleteLivePullStreamInfoConfig(request *DeleteLivePullStreamInfoConfigRequest) (response *DeleteLivePullStreamInfoConfigResponse, err error) {
	response = CreateDeleteLivePullStreamInfoConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLivePullStreamInfoConfigWithChan invokes the live.DeleteLivePullStreamInfoConfig API asynchronously
func (client *Client) DeleteLivePullStreamInfoConfigWithChan(request *DeleteLivePullStreamInfoConfigRequest) (<-chan *DeleteLivePullStreamInfoConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteLivePullStreamInfoConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLivePullStreamInfoConfig(request)
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

// DeleteLivePullStreamInfoConfigWithCallback invokes the live.DeleteLivePullStreamInfoConfig API asynchronously
func (client *Client) DeleteLivePullStreamInfoConfigWithCallback(request *DeleteLivePullStreamInfoConfigRequest, callback func(response *DeleteLivePullStreamInfoConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLivePullStreamInfoConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteLivePullStreamInfoConfig(request)
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

// DeleteLivePullStreamInfoConfigRequest is the request struct for api DeleteLivePullStreamInfoConfig
type DeleteLivePullStreamInfoConfigRequest struct {
	*requests.RpcRequest
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	StreamName    string           `position:"Query" name:"StreamName"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteLivePullStreamInfoConfigResponse is the response struct for api DeleteLivePullStreamInfoConfig
type DeleteLivePullStreamInfoConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLivePullStreamInfoConfigRequest creates a request to invoke DeleteLivePullStreamInfoConfig API
func CreateDeleteLivePullStreamInfoConfigRequest() (request *DeleteLivePullStreamInfoConfigRequest) {
	request = &DeleteLivePullStreamInfoConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLivePullStreamInfoConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLivePullStreamInfoConfigResponse creates a response to parse from DeleteLivePullStreamInfoConfig response
func CreateDeleteLivePullStreamInfoConfigResponse() (response *DeleteLivePullStreamInfoConfigResponse) {
	response = &DeleteLivePullStreamInfoConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
