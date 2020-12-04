package mts

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

// SubmitVideoGifJob invokes the mts.SubmitVideoGifJob API synchronously
func (client *Client) SubmitVideoGifJob(request *SubmitVideoGifJobRequest) (response *SubmitVideoGifJobResponse, err error) {
	response = CreateSubmitVideoGifJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitVideoGifJobWithChan invokes the mts.SubmitVideoGifJob API asynchronously
func (client *Client) SubmitVideoGifJobWithChan(request *SubmitVideoGifJobRequest) (<-chan *SubmitVideoGifJobResponse, <-chan error) {
	responseChan := make(chan *SubmitVideoGifJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitVideoGifJob(request)
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

// SubmitVideoGifJobWithCallback invokes the mts.SubmitVideoGifJob API asynchronously
func (client *Client) SubmitVideoGifJobWithCallback(request *SubmitVideoGifJobRequest, callback func(response *SubmitVideoGifJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitVideoGifJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitVideoGifJob(request)
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

// SubmitVideoGifJobRequest is the request struct for api SubmitVideoGifJob
type SubmitVideoGifJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UserData             string           `position:"Query" name:"UserData"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	VideoGifConfig       string           `position:"Query" name:"VideoGifConfig"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Input                string           `position:"Query" name:"Input"`
}

// SubmitVideoGifJobResponse is the response struct for api SubmitVideoGifJob
type SubmitVideoGifJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitVideoGifJobRequest creates a request to invoke SubmitVideoGifJob API
func CreateSubmitVideoGifJobRequest() (request *SubmitVideoGifJobRequest) {
	request = &SubmitVideoGifJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitVideoGifJob", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitVideoGifJobResponse creates a response to parse from SubmitVideoGifJob response
func CreateSubmitVideoGifJobResponse() (response *SubmitVideoGifJobResponse) {
	response = &SubmitVideoGifJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
