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

// StartCasterScene invokes the live.StartCasterScene API synchronously
func (client *Client) StartCasterScene(request *StartCasterSceneRequest) (response *StartCasterSceneResponse, err error) {
	response = CreateStartCasterSceneResponse()
	err = client.DoAction(request, response)
	return
}

// StartCasterSceneWithChan invokes the live.StartCasterScene API asynchronously
func (client *Client) StartCasterSceneWithChan(request *StartCasterSceneRequest) (<-chan *StartCasterSceneResponse, <-chan error) {
	responseChan := make(chan *StartCasterSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartCasterScene(request)
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

// StartCasterSceneWithCallback invokes the live.StartCasterScene API asynchronously
func (client *Client) StartCasterSceneWithCallback(request *StartCasterSceneRequest, callback func(response *StartCasterSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartCasterSceneResponse
		var err error
		defer close(result)
		response, err = client.StartCasterScene(request)
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

// StartCasterSceneRequest is the request struct for api StartCasterScene
type StartCasterSceneRequest struct {
	*requests.RpcRequest
	CasterId string           `position:"Query" name:"CasterId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	SceneId  string           `position:"Query" name:"SceneId"`
}

// StartCasterSceneResponse is the response struct for api StartCasterScene
type StartCasterSceneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	StreamUrl string `json:"StreamUrl" xml:"StreamUrl"`
}

// CreateStartCasterSceneRequest creates a request to invoke StartCasterScene API
func CreateStartCasterSceneRequest() (request *StartCasterSceneRequest) {
	request = &StartCasterSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "StartCasterScene", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartCasterSceneResponse creates a response to parse from StartCasterScene response
func CreateStartCasterSceneResponse() (response *StartCasterSceneResponse) {
	response = &StartCasterSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
