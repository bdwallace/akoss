package aliyuncvc

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

// BatchCreateDevice invokes the aliyuncvc.BatchCreateDevice API synchronously
func (client *Client) BatchCreateDevice(request *BatchCreateDeviceRequest) (response *BatchCreateDeviceResponse, err error) {
	response = CreateBatchCreateDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// BatchCreateDeviceWithChan invokes the aliyuncvc.BatchCreateDevice API asynchronously
func (client *Client) BatchCreateDeviceWithChan(request *BatchCreateDeviceRequest) (<-chan *BatchCreateDeviceResponse, <-chan error) {
	responseChan := make(chan *BatchCreateDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchCreateDevice(request)
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

// BatchCreateDeviceWithCallback invokes the aliyuncvc.BatchCreateDevice API asynchronously
func (client *Client) BatchCreateDeviceWithCallback(request *BatchCreateDeviceRequest, callback func(response *BatchCreateDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchCreateDeviceResponse
		var err error
		defer close(result)
		response, err = client.BatchCreateDevice(request)
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

// BatchCreateDeviceRequest is the request struct for api BatchCreateDevice
type BatchCreateDeviceRequest struct {
	*requests.RpcRequest
	SN string `position:"Body" name:"SN"`
}

// BatchCreateDeviceResponse is the response struct for api BatchCreateDevice
type BatchCreateDeviceResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Devices   []Data `json:"Devices" xml:"Devices"`
}

// CreateBatchCreateDeviceRequest creates a request to invoke BatchCreateDevice API
func CreateBatchCreateDeviceRequest() (request *BatchCreateDeviceRequest) {
	request = &BatchCreateDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "BatchCreateDevice", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchCreateDeviceResponse creates a response to parse from BatchCreateDevice response
func CreateBatchCreateDeviceResponse() (response *BatchCreateDeviceResponse) {
	response = &BatchCreateDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
