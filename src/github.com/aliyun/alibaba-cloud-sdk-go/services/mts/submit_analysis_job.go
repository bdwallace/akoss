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

// SubmitAnalysisJob invokes the mts.SubmitAnalysisJob API synchronously
func (client *Client) SubmitAnalysisJob(request *SubmitAnalysisJobRequest) (response *SubmitAnalysisJobResponse, err error) {
	response = CreateSubmitAnalysisJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitAnalysisJobWithChan invokes the mts.SubmitAnalysisJob API asynchronously
func (client *Client) SubmitAnalysisJobWithChan(request *SubmitAnalysisJobRequest) (<-chan *SubmitAnalysisJobResponse, <-chan error) {
	responseChan := make(chan *SubmitAnalysisJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitAnalysisJob(request)
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

// SubmitAnalysisJobWithCallback invokes the mts.SubmitAnalysisJob API asynchronously
func (client *Client) SubmitAnalysisJobWithCallback(request *SubmitAnalysisJobRequest, callback func(response *SubmitAnalysisJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitAnalysisJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitAnalysisJob(request)
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

// SubmitAnalysisJobRequest is the request struct for api SubmitAnalysisJob
type SubmitAnalysisJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AnalysisConfig       string           `position:"Query" name:"AnalysisConfig"`
	UserData             string           `position:"Query" name:"UserData"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             string           `position:"Query" name:"Priority"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Input                string           `position:"Query" name:"Input"`
}

// SubmitAnalysisJobResponse is the response struct for api SubmitAnalysisJob
type SubmitAnalysisJobResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	AnalysisJob AnalysisJob `json:"AnalysisJob" xml:"AnalysisJob"`
}

// CreateSubmitAnalysisJobRequest creates a request to invoke SubmitAnalysisJob API
func CreateSubmitAnalysisJobRequest() (request *SubmitAnalysisJobRequest) {
	request = &SubmitAnalysisJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitAnalysisJob", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitAnalysisJobResponse creates a response to parse from SubmitAnalysisJob response
func CreateSubmitAnalysisJobResponse() (response *SubmitAnalysisJobResponse) {
	response = &SubmitAnalysisJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
