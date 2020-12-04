package openanalytics_open

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

// InitializeRegion invokes the openanalytics_open.InitializeRegion API synchronously
func (client *Client) InitializeRegion(request *InitializeRegionRequest) (response *InitializeRegionResponse, err error) {
	response = CreateInitializeRegionResponse()
	err = client.DoAction(request, response)
	return
}

// InitializeRegionWithChan invokes the openanalytics_open.InitializeRegion API asynchronously
func (client *Client) InitializeRegionWithChan(request *InitializeRegionRequest) (<-chan *InitializeRegionResponse, <-chan error) {
	responseChan := make(chan *InitializeRegionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitializeRegion(request)
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

// InitializeRegionWithCallback invokes the openanalytics_open.InitializeRegion API asynchronously
func (client *Client) InitializeRegionWithCallback(request *InitializeRegionRequest, callback func(response *InitializeRegionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitializeRegionResponse
		var err error
		defer close(result)
		response, err = client.InitializeRegion(request)
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

// InitializeRegionRequest is the request struct for api InitializeRegion
type InitializeRegionRequest struct {
	*requests.RpcRequest
	ExternalUid          string           `position:"Body" name:"ExternalUid"`
	InitPassword         string           `position:"Body" name:"InitPassword"`
	ExternalAliyunUid    string           `position:"Body" name:"ExternalAliyunUid"`
	UseRandomPassword    requests.Boolean `position:"Body" name:"UseRandomPassword"`
	EnableKMS            requests.Boolean `position:"Body" name:"EnableKMS"`
	ExternalBizAliyunUid string           `position:"Body" name:"ExternalBizAliyunUid"`
}

// InitializeRegionResponse is the response struct for api InitializeRegion
type InitializeRegionResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	RegionId  string  `json:"RegionId" xml:"RegionId"`
	Account   Account `json:"Account" xml:"Account"`
}

// CreateInitializeRegionRequest creates a request to invoke InitializeRegion API
func CreateInitializeRegionRequest() (request *InitializeRegionRequest) {
	request = &InitializeRegionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "InitializeRegion", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInitializeRegionResponse creates a response to parse from InitializeRegion response
func CreateInitializeRegionResponse() (response *InitializeRegionResponse) {
	response = &InitializeRegionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
