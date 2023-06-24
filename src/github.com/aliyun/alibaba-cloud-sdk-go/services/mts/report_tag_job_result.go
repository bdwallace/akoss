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

// ReportTagJobResult invokes the mts.ReportTagJobResult API synchronously
func (client *Client) ReportTagJobResult(request *ReportTagJobResultRequest) (response *ReportTagJobResultResponse, err error) {
	response = CreateReportTagJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// ReportTagJobResultWithChan invokes the mts.ReportTagJobResult API asynchronously
func (client *Client) ReportTagJobResultWithChan(request *ReportTagJobResultRequest) (<-chan *ReportTagJobResultResponse, <-chan error) {
	responseChan := make(chan *ReportTagJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportTagJobResult(request)
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

// ReportTagJobResultWithCallback invokes the mts.ReportTagJobResult API asynchronously
func (client *Client) ReportTagJobResultWithCallback(request *ReportTagJobResultRequest, callback func(response *ReportTagJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportTagJobResultResponse
		var err error
		defer close(result)
		response, err = client.ReportTagJobResult(request)
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

// ReportTagJobResultRequest is the request struct for api ReportTagJobResult
type ReportTagJobResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Result               string           `position:"Query" name:"Result"`
	JobId                string           `position:"Query" name:"JobId"`
	Tag                  string           `position:"Query" name:"Tag"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ReportTagJobResultResponse is the response struct for api ReportTagJobResult
type ReportTagJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateReportTagJobResultRequest creates a request to invoke ReportTagJobResult API
func CreateReportTagJobResultRequest() (request *ReportTagJobResultRequest) {
	request = &ReportTagJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ReportTagJobResult", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReportTagJobResultResponse creates a response to parse from ReportTagJobResult response
func CreateReportTagJobResultResponse() (response *ReportTagJobResultResponse) {
	response = &ReportTagJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
