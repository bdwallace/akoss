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

// QueryCoverJobList invokes the mts.QueryCoverJobList API synchronously
func (client *Client) QueryCoverJobList(request *QueryCoverJobListRequest) (response *QueryCoverJobListResponse, err error) {
	response = CreateQueryCoverJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCoverJobListWithChan invokes the mts.QueryCoverJobList API asynchronously
func (client *Client) QueryCoverJobListWithChan(request *QueryCoverJobListRequest) (<-chan *QueryCoverJobListResponse, <-chan error) {
	responseChan := make(chan *QueryCoverJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCoverJobList(request)
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

// QueryCoverJobListWithCallback invokes the mts.QueryCoverJobList API asynchronously
func (client *Client) QueryCoverJobListWithCallback(request *QueryCoverJobListRequest, callback func(response *QueryCoverJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCoverJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryCoverJobList(request)
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

// QueryCoverJobListRequest is the request struct for api QueryCoverJobList
type QueryCoverJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NextPageToken              string           `position:"Query" name:"NextPageToken"`
	StartOfJobCreatedTimeRange string           `position:"Query" name:"StartOfJobCreatedTimeRange"`
	CoverJobIds                string           `position:"Query" name:"CoverJobIds"`
	State                      string           `position:"Query" name:"State"`
	EndOfJobCreatedTimeRange   string           `position:"Query" name:"EndOfJobCreatedTimeRange"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	MaximumPageSize            requests.Integer `position:"Query" name:"MaximumPageSize"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId                 string           `position:"Query" name:"PipelineId"`
}

// QueryCoverJobListResponse is the response struct for api QueryCoverJobList
type QueryCoverJobListResponse struct {
	*responses.BaseResponse
	RequestId     string                         `json:"RequestId" xml:"RequestId"`
	NextPageToken string                         `json:"NextPageToken" xml:"NextPageToken"`
	NonExistIds   NonExistIdsInQueryCoverJobList `json:"NonExistIds" xml:"NonExistIds"`
	CoverJobList  CoverJobList                   `json:"CoverJobList" xml:"CoverJobList"`
}

// CreateQueryCoverJobListRequest creates a request to invoke QueryCoverJobList API
func CreateQueryCoverJobListRequest() (request *QueryCoverJobListRequest) {
	request = &QueryCoverJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryCoverJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryCoverJobListResponse creates a response to parse from QueryCoverJobList response
func CreateQueryCoverJobListResponse() (response *QueryCoverJobListResponse) {
	response = &QueryCoverJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
