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

// DeleteCasterSceneConfig invokes the live.DeleteCasterSceneConfig API synchronously
func (client *Client) DeleteCasterSceneConfig(request *DeleteCasterSceneConfigRequest) (response *DeleteCasterSceneConfigResponse, err error) {
	response = CreateDeleteCasterSceneConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCasterSceneConfigWithChan invokes the live.DeleteCasterSceneConfig API asynchronously
func (client *Client) DeleteCasterSceneConfigWithChan(request *DeleteCasterSceneConfigRequest) (<-chan *DeleteCasterSceneConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteCasterSceneConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCasterSceneConfig(request)
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

// DeleteCasterSceneConfigWithCallback invokes the live.DeleteCasterSceneConfig API asynchronously
func (client *Client) DeleteCasterSceneConfigWithCallback(request *DeleteCasterSceneConfigRequest, callback func(response *DeleteCasterSceneConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCasterSceneConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteCasterSceneConfig(request)
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

// DeleteCasterSceneConfigRequest is the request struct for api DeleteCasterSceneConfig
type DeleteCasterSceneConfigRequest struct {
	*requests.RpcRequest
	Type     string           `position:"Query" name:"Type"`
	CasterId string           `position:"Query" name:"CasterId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	SceneId  string           `position:"Query" name:"SceneId"`
}

// DeleteCasterSceneConfigResponse is the response struct for api DeleteCasterSceneConfig
type DeleteCasterSceneConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCasterSceneConfigRequest creates a request to invoke DeleteCasterSceneConfig API
func CreateDeleteCasterSceneConfigRequest() (request *DeleteCasterSceneConfigRequest) {
	request = &DeleteCasterSceneConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteCasterSceneConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCasterSceneConfigResponse creates a response to parse from DeleteCasterSceneConfig response
func CreateDeleteCasterSceneConfigResponse() (response *DeleteCasterSceneConfigResponse) {
	response = &DeleteCasterSceneConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
