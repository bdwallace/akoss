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

// DeleteCasterEpisode invokes the live.DeleteCasterEpisode API synchronously
func (client *Client) DeleteCasterEpisode(request *DeleteCasterEpisodeRequest) (response *DeleteCasterEpisodeResponse, err error) {
	response = CreateDeleteCasterEpisodeResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCasterEpisodeWithChan invokes the live.DeleteCasterEpisode API asynchronously
func (client *Client) DeleteCasterEpisodeWithChan(request *DeleteCasterEpisodeRequest) (<-chan *DeleteCasterEpisodeResponse, <-chan error) {
	responseChan := make(chan *DeleteCasterEpisodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCasterEpisode(request)
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

// DeleteCasterEpisodeWithCallback invokes the live.DeleteCasterEpisode API asynchronously
func (client *Client) DeleteCasterEpisodeWithCallback(request *DeleteCasterEpisodeRequest, callback func(response *DeleteCasterEpisodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCasterEpisodeResponse
		var err error
		defer close(result)
		response, err = client.DeleteCasterEpisode(request)
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

// DeleteCasterEpisodeRequest is the request struct for api DeleteCasterEpisode
type DeleteCasterEpisodeRequest struct {
	*requests.RpcRequest
	CasterId  string           `position:"Query" name:"CasterId"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	EpisodeId string           `position:"Query" name:"EpisodeId"`
}

// DeleteCasterEpisodeResponse is the response struct for api DeleteCasterEpisode
type DeleteCasterEpisodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CasterId  string `json:"CasterId" xml:"CasterId"`
	EpisodeId string `json:"EpisodeId" xml:"EpisodeId"`
}

// CreateDeleteCasterEpisodeRequest creates a request to invoke DeleteCasterEpisode API
func CreateDeleteCasterEpisodeRequest() (request *DeleteCasterEpisodeRequest) {
	request = &DeleteCasterEpisodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteCasterEpisode", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCasterEpisodeResponse creates a response to parse from DeleteCasterEpisode response
func CreateDeleteCasterEpisodeResponse() (response *DeleteCasterEpisodeResponse) {
	response = &DeleteCasterEpisodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
