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

// DescribeBoardEvents invokes the live.DescribeBoardEvents API synchronously
func (client *Client) DescribeBoardEvents(request *DescribeBoardEventsRequest) (response *DescribeBoardEventsResponse, err error) {
	response = CreateDescribeBoardEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBoardEventsWithChan invokes the live.DescribeBoardEvents API asynchronously
func (client *Client) DescribeBoardEventsWithChan(request *DescribeBoardEventsRequest) (<-chan *DescribeBoardEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeBoardEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBoardEvents(request)
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

// DescribeBoardEventsWithCallback invokes the live.DescribeBoardEvents API asynchronously
func (client *Client) DescribeBoardEventsWithCallback(request *DescribeBoardEventsRequest, callback func(response *DescribeBoardEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBoardEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeBoardEvents(request)
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

// DescribeBoardEventsRequest is the request struct for api DescribeBoardEvents
type DescribeBoardEventsRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	BoardId   string           `position:"Query" name:"BoardId"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	AppId     string           `position:"Query" name:"AppId"`
}

// DescribeBoardEventsResponse is the response struct for api DescribeBoardEvents
type DescribeBoardEventsResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Events    []Event `json:"Events" xml:"Events"`
}

// CreateDescribeBoardEventsRequest creates a request to invoke DescribeBoardEvents API
func CreateDescribeBoardEventsRequest() (request *DescribeBoardEventsRequest) {
	request = &DescribeBoardEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeBoardEvents", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBoardEventsResponse creates a response to parse from DescribeBoardEvents response
func CreateDescribeBoardEventsResponse() (response *DescribeBoardEventsResponse) {
	response = &DescribeBoardEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
