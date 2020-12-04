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

// ActiveDevice invokes the aliyuncvc.ActiveDevice API synchronously
func (client *Client) ActiveDevice(request *ActiveDeviceRequest) (response *ActiveDeviceResponse, err error) {
	response = CreateActiveDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// ActiveDeviceWithChan invokes the aliyuncvc.ActiveDevice API asynchronously
func (client *Client) ActiveDeviceWithChan(request *ActiveDeviceRequest) (<-chan *ActiveDeviceResponse, <-chan error) {
	responseChan := make(chan *ActiveDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ActiveDevice(request)
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

// ActiveDeviceWithCallback invokes the aliyuncvc.ActiveDevice API asynchronously
func (client *Client) ActiveDeviceWithCallback(request *ActiveDeviceRequest, callback func(response *ActiveDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ActiveDeviceResponse
		var err error
		defer close(result)
		response, err = client.ActiveDevice(request)
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

// ActiveDeviceRequest is the request struct for api ActiveDevice
type ActiveDeviceRequest struct {
	*requests.RpcRequest
	IP            string `position:"Body" name:"IP"`
	ActiveCode    string `position:"Body" name:"ActiveCode"`
	Mac           string `position:"Body" name:"Mac"`
	DeviceVersion string `position:"Body" name:"DeviceVersion"`
	SN            string `position:"Body" name:"SN"`
}

// ActiveDeviceResponse is the response struct for api ActiveDevice
type ActiveDeviceResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Token     string `json:"Token" xml:"Token"`
}

// CreateActiveDeviceRequest creates a request to invoke ActiveDevice API
func CreateActiveDeviceRequest() (request *ActiveDeviceRequest) {
	request = &ActiveDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "ActiveDevice", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateActiveDeviceResponse creates a response to parse from ActiveDevice response
func CreateActiveDeviceResponse() (response *ActiveDeviceResponse) {
	response = &ActiveDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
