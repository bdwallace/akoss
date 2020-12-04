package xtrace

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

// UpdateSampling invokes the xtrace.UpdateSampling API synchronously
func (client *Client) UpdateSampling(request *UpdateSamplingRequest) (response *UpdateSamplingResponse, err error) {
	response = CreateUpdateSamplingResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSamplingWithChan invokes the xtrace.UpdateSampling API asynchronously
func (client *Client) UpdateSamplingWithChan(request *UpdateSamplingRequest) (<-chan *UpdateSamplingResponse, <-chan error) {
	responseChan := make(chan *UpdateSamplingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSampling(request)
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

// UpdateSamplingWithCallback invokes the xtrace.UpdateSampling API asynchronously
func (client *Client) UpdateSamplingWithCallback(request *UpdateSamplingRequest, callback func(response *UpdateSamplingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSamplingResponse
		var err error
		defer close(result)
		response, err = client.UpdateSampling(request)
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

// UpdateSamplingRequest is the request struct for api UpdateSampling
type UpdateSamplingRequest struct {
	*requests.RpcRequest
	Sampling    string `position:"Query" name:"Sampling"`
	ProxyUserId string `position:"Query" name:"ProxyUserId"`
}

// UpdateSamplingResponse is the response struct for api UpdateSampling
type UpdateSamplingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateUpdateSamplingRequest creates a request to invoke UpdateSampling API
func CreateUpdateSamplingRequest() (request *UpdateSamplingRequest) {
	request = &UpdateSamplingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("xtrace", "2019-08-08", "UpdateSampling", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateSamplingResponse creates a response to parse from UpdateSampling response
func CreateUpdateSamplingResponse() (response *UpdateSamplingResponse) {
	response = &UpdateSamplingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
