package cdn

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

// ModifyPathCacheExpiredConfig invokes the cdn.ModifyPathCacheExpiredConfig API synchronously
func (client *Client) ModifyPathCacheExpiredConfig(request *ModifyPathCacheExpiredConfigRequest) (response *ModifyPathCacheExpiredConfigResponse, err error) {
	response = CreateModifyPathCacheExpiredConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPathCacheExpiredConfigWithChan invokes the cdn.ModifyPathCacheExpiredConfig API asynchronously
func (client *Client) ModifyPathCacheExpiredConfigWithChan(request *ModifyPathCacheExpiredConfigRequest) (<-chan *ModifyPathCacheExpiredConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyPathCacheExpiredConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPathCacheExpiredConfig(request)
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

// ModifyPathCacheExpiredConfigWithCallback invokes the cdn.ModifyPathCacheExpiredConfig API asynchronously
func (client *Client) ModifyPathCacheExpiredConfigWithCallback(request *ModifyPathCacheExpiredConfigRequest, callback func(response *ModifyPathCacheExpiredConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPathCacheExpiredConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyPathCacheExpiredConfig(request)
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

// ModifyPathCacheExpiredConfigRequest is the request struct for api ModifyPathCacheExpiredConfig
type ModifyPathCacheExpiredConfigRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	Weight        string           `position:"Query" name:"Weight"`
	CacheContent  string           `position:"Query" name:"CacheContent"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	TTL           string           `position:"Query" name:"TTL"`
	ConfigID      string           `position:"Query" name:"ConfigID"`
}

// ModifyPathCacheExpiredConfigResponse is the response struct for api ModifyPathCacheExpiredConfig
type ModifyPathCacheExpiredConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyPathCacheExpiredConfigRequest creates a request to invoke ModifyPathCacheExpiredConfig API
func CreateModifyPathCacheExpiredConfigRequest() (request *ModifyPathCacheExpiredConfigRequest) {
	request = &ModifyPathCacheExpiredConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "ModifyPathCacheExpiredConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyPathCacheExpiredConfigResponse creates a response to parse from ModifyPathCacheExpiredConfig response
func CreateModifyPathCacheExpiredConfigResponse() (response *ModifyPathCacheExpiredConfigResponse) {
	response = &ModifyPathCacheExpiredConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
