package sae

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

// ListAppEvents invokes the sae.ListAppEvents API synchronously
func (client *Client) ListAppEvents(request *ListAppEventsRequest) (response *ListAppEventsResponse, err error) {
	response = CreateListAppEventsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppEventsWithChan invokes the sae.ListAppEvents API asynchronously
func (client *Client) ListAppEventsWithChan(request *ListAppEventsRequest) (<-chan *ListAppEventsResponse, <-chan error) {
	responseChan := make(chan *ListAppEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppEvents(request)
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

// ListAppEventsWithCallback invokes the sae.ListAppEvents API asynchronously
func (client *Client) ListAppEventsWithCallback(request *ListAppEventsRequest, callback func(response *ListAppEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppEventsResponse
		var err error
		defer close(result)
		response, err = client.ListAppEvents(request)
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

// ListAppEventsRequest is the request struct for api ListAppEvents
type ListAppEventsRequest struct {
	*requests.RoaRequest
	Reason      string           `position:"Query" name:"Reason"`
	ObjectKind  string           `position:"Query" name:"ObjectKind"`
	AppId       string           `position:"Query" name:"AppId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ObjectName  string           `position:"Query" name:"ObjectName"`
	Namespace   string           `position:"Query" name:"Namespace"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	EventType   string           `position:"Query" name:"EventType"`
}

// ListAppEventsResponse is the response struct for api ListAppEvents
type ListAppEventsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListAppEventsRequest creates a request to invoke ListAppEvents API
func CreateListAppEventsRequest() (request *ListAppEventsRequest) {
	request = &ListAppEventsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "ListAppEvents", "/pop/v1/sam/app/listAppEvents", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListAppEventsResponse creates a response to parse from ListAppEvents response
func CreateListAppEventsResponse() (response *ListAppEventsResponse) {
	response = &ListAppEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
