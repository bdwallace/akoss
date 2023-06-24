package retailcloud

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

// DescribePodEvents invokes the retailcloud.DescribePodEvents API synchronously
// api document: https://help.aliyun.com/api/retailcloud/describepodevents.html
func (client *Client) DescribePodEvents(request *DescribePodEventsRequest) (response *DescribePodEventsResponse, err error) {
	response = CreateDescribePodEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePodEventsWithChan invokes the retailcloud.DescribePodEvents API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describepodevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePodEventsWithChan(request *DescribePodEventsRequest) (<-chan *DescribePodEventsResponse, <-chan error) {
	responseChan := make(chan *DescribePodEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePodEvents(request)
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

// DescribePodEventsWithCallback invokes the retailcloud.DescribePodEvents API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describepodevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePodEventsWithCallback(request *DescribePodEventsRequest, callback func(response *DescribePodEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePodEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribePodEvents(request)
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

// DescribePodEventsRequest is the request struct for api DescribePodEvents
type DescribePodEventsRequest struct {
	*requests.RpcRequest
	DeployOrderId requests.Integer `position:"Query" name:"DeployOrderId"`
	AppInstId     string           `position:"Query" name:"AppInstId"`
}

// DescribePodEventsResponse is the response struct for api DescribePodEvents
type DescribePodEventsResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribePodEventsRequest creates a request to invoke DescribePodEvents API
func CreateDescribePodEventsRequest() (request *DescribePodEventsRequest) {
	request = &DescribePodEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DescribePodEvents", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePodEventsResponse creates a response to parse from DescribePodEvents response
func CreateDescribePodEventsResponse() (response *DescribePodEventsResponse) {
	response = &DescribePodEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
