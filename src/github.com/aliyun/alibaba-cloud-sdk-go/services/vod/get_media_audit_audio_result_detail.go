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

// GetMediaAuditAudioResultDetail invokes the vod.GetMediaAuditAudioResultDetail API synchronously
func (client *Client) GetMediaAuditAudioResultDetail(request *GetMediaAuditAudioResultDetailRequest) (response *GetMediaAuditAudioResultDetailResponse, err error) {
	response = CreateGetMediaAuditAudioResultDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetMediaAuditAudioResultDetailWithChan invokes the vod.GetMediaAuditAudioResultDetail API asynchronously
func (client *Client) GetMediaAuditAudioResultDetailWithChan(request *GetMediaAuditAudioResultDetailRequest) (<-chan *GetMediaAuditAudioResultDetailResponse, <-chan error) {
	responseChan := make(chan *GetMediaAuditAudioResultDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMediaAuditAudioResultDetail(request)
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

// GetMediaAuditAudioResultDetailWithCallback invokes the vod.GetMediaAuditAudioResultDetail API asynchronously
func (client *Client) GetMediaAuditAudioResultDetailWithCallback(request *GetMediaAuditAudioResultDetailRequest, callback func(response *GetMediaAuditAudioResultDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMediaAuditAudioResultDetailResponse
		var err error
		defer close(result)
		response, err = client.GetMediaAuditAudioResultDetail(request)
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

// GetMediaAuditAudioResultDetailRequest is the request struct for api GetMediaAuditAudioResultDetail
type GetMediaAuditAudioResultDetailRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
}

// GetMediaAuditAudioResultDetailResponse is the response struct for api GetMediaAuditAudioResultDetail
type GetMediaAuditAudioResultDetailResponse struct {
	*responses.BaseResponse
	RequestId                   string                      `json:"RequestId" xml:"RequestId"`
	MediaAuditAudioResultDetail MediaAuditAudioResultDetail `json:"MediaAuditAudioResultDetail" xml:"MediaAuditAudioResultDetail"`
}

// CreateGetMediaAuditAudioResultDetailRequest creates a request to invoke GetMediaAuditAudioResultDetail API
func CreateGetMediaAuditAudioResultDetailRequest() (request *GetMediaAuditAudioResultDetailRequest) {
	request = &GetMediaAuditAudioResultDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetMediaAuditAudioResultDetail", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMediaAuditAudioResultDetailResponse creates a response to parse from GetMediaAuditAudioResultDetail response
func CreateGetMediaAuditAudioResultDetailResponse() (response *GetMediaAuditAudioResultDetailResponse) {
	response = &GetMediaAuditAudioResultDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
