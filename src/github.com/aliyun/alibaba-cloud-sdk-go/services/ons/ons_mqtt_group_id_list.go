package ons

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

// OnsMqttGroupIdList invokes the ons.OnsMqttGroupIdList API synchronously
func (client *Client) OnsMqttGroupIdList(request *OnsMqttGroupIdListRequest) (response *OnsMqttGroupIdListResponse, err error) {
	response = CreateOnsMqttGroupIdListResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMqttGroupIdListWithChan invokes the ons.OnsMqttGroupIdList API asynchronously
func (client *Client) OnsMqttGroupIdListWithChan(request *OnsMqttGroupIdListRequest) (<-chan *OnsMqttGroupIdListResponse, <-chan error) {
	responseChan := make(chan *OnsMqttGroupIdListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMqttGroupIdList(request)
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

// OnsMqttGroupIdListWithCallback invokes the ons.OnsMqttGroupIdList API asynchronously
func (client *Client) OnsMqttGroupIdListWithCallback(request *OnsMqttGroupIdListRequest, callback func(response *OnsMqttGroupIdListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMqttGroupIdListResponse
		var err error
		defer close(result)
		response, err = client.OnsMqttGroupIdList(request)
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

// OnsMqttGroupIdListRequest is the request struct for api OnsMqttGroupIdList
type OnsMqttGroupIdListRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// OnsMqttGroupIdListResponse is the response struct for api OnsMqttGroupIdList
type OnsMqttGroupIdListResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	HelpUrl   string                   `json:"HelpUrl" xml:"HelpUrl"`
	Data      DataInOnsMqttGroupIdList `json:"Data" xml:"Data"`
}

// CreateOnsMqttGroupIdListRequest creates a request to invoke OnsMqttGroupIdList API
func CreateOnsMqttGroupIdListRequest() (request *OnsMqttGroupIdListRequest) {
	request = &OnsMqttGroupIdListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMqttGroupIdList", "", "")
	request.Method = requests.POST
	return
}

// CreateOnsMqttGroupIdListResponse creates a response to parse from OnsMqttGroupIdList response
func CreateOnsMqttGroupIdListResponse() (response *OnsMqttGroupIdListResponse) {
	response = &OnsMqttGroupIdListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}