package vod

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

// GetVideoPlayAuth invokes the vod.GetVideoPlayAuth API synchronously
func (client *Client) GetVideoPlayAuth(request *GetVideoPlayAuthRequest) (response *GetVideoPlayAuthResponse, err error) {
	response = CreateGetVideoPlayAuthResponse()
	err = client.DoAction(request, response)
	return
}

// GetVideoPlayAuthWithChan invokes the vod.GetVideoPlayAuth API asynchronously
func (client *Client) GetVideoPlayAuthWithChan(request *GetVideoPlayAuthRequest) (<-chan *GetVideoPlayAuthResponse, <-chan error) {
	responseChan := make(chan *GetVideoPlayAuthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVideoPlayAuth(request)
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

// GetVideoPlayAuthWithCallback invokes the vod.GetVideoPlayAuth API asynchronously
func (client *Client) GetVideoPlayAuthWithCallback(request *GetVideoPlayAuthRequest, callback func(response *GetVideoPlayAuthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVideoPlayAuthResponse
		var err error
		defer close(result)
		response, err = client.GetVideoPlayAuth(request)
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

// GetVideoPlayAuthRequest is the request struct for api GetVideoPlayAuth
type GetVideoPlayAuthRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ReAuthInfo           string           `position:"Query" name:"ReAuthInfo"`
	PlayConfig           string           `position:"Query" name:"PlayConfig"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string           `position:"Query" name:"VideoId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AuthInfoTimeout      requests.Integer `position:"Query" name:"AuthInfoTimeout"`
}

// GetVideoPlayAuthResponse is the response struct for api GetVideoPlayAuth
type GetVideoPlayAuthResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	PlayAuth  string    `json:"PlayAuth" xml:"PlayAuth"`
	VideoMeta VideoMeta `json:"VideoMeta" xml:"VideoMeta"`
}

// CreateGetVideoPlayAuthRequest creates a request to invoke GetVideoPlayAuth API
func CreateGetVideoPlayAuthRequest() (request *GetVideoPlayAuthRequest) {
	request = &GetVideoPlayAuthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetVideoPlayAuth", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVideoPlayAuthResponse creates a response to parse from GetVideoPlayAuth response
func CreateGetVideoPlayAuthResponse() (response *GetVideoPlayAuthResponse) {
	response = &GetVideoPlayAuthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
