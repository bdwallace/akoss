package iot

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

// CreateSubscribeRelation invokes the iot.CreateSubscribeRelation API synchronously
func (client *Client) CreateSubscribeRelation(request *CreateSubscribeRelationRequest) (response *CreateSubscribeRelationResponse, err error) {
	response = CreateCreateSubscribeRelationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSubscribeRelationWithChan invokes the iot.CreateSubscribeRelation API asynchronously
func (client *Client) CreateSubscribeRelationWithChan(request *CreateSubscribeRelationRequest) (<-chan *CreateSubscribeRelationResponse, <-chan error) {
	responseChan := make(chan *CreateSubscribeRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSubscribeRelation(request)
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

// CreateSubscribeRelationWithCallback invokes the iot.CreateSubscribeRelation API asynchronously
func (client *Client) CreateSubscribeRelationWithCallback(request *CreateSubscribeRelationRequest, callback func(response *CreateSubscribeRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSubscribeRelationResponse
		var err error
		defer close(result)
		response, err = client.CreateSubscribeRelation(request)
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

// CreateSubscribeRelationRequest is the request struct for api CreateSubscribeRelation
type CreateSubscribeRelationRequest struct {
	*requests.RpcRequest
	OtaEventFlag            requests.Boolean `position:"Query" name:"OtaEventFlag"`
	DeviceTopoLifeCycleFlag requests.Boolean `position:"Query" name:"DeviceTopoLifeCycleFlag"`
	DeviceLifeCycleFlag     requests.Boolean `position:"Query" name:"DeviceLifeCycleFlag"`
	Type                    string           `position:"Query" name:"Type"`
	IotInstanceId           string           `position:"Query" name:"IotInstanceId"`
	DeviceStatusChangeFlag  requests.Boolean `position:"Query" name:"DeviceStatusChangeFlag"`
	OtaVersionFlag          requests.Boolean `position:"Query" name:"OtaVersionFlag"`
	DeviceTagFlag           requests.Boolean `position:"Query" name:"DeviceTagFlag"`
	ConsumerGroupIds        *[]string        `position:"Query" name:"ConsumerGroupIds"  type:"Repeated"`
	ProductKey              string           `position:"Query" name:"ProductKey"`
	ThingHistoryFlag        requests.Boolean `position:"Query" name:"ThingHistoryFlag"`
	FoundDeviceListFlag     requests.Boolean `position:"Query" name:"FoundDeviceListFlag"`
	OtaJobFlag              requests.Boolean `position:"Query" name:"OtaJobFlag"`
	ApiProduct              string           `position:"Body" name:"ApiProduct"`
	DeviceDataFlag          requests.Boolean `position:"Query" name:"DeviceDataFlag"`
	ApiRevision             string           `position:"Body" name:"ApiRevision"`
	MnsConfiguration        string           `position:"Query" name:"MnsConfiguration"`
}

// CreateSubscribeRelationResponse is the response struct for api CreateSubscribeRelation
type CreateSubscribeRelationResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateCreateSubscribeRelationRequest creates a request to invoke CreateSubscribeRelation API
func CreateCreateSubscribeRelationRequest() (request *CreateSubscribeRelationRequest) {
	request = &CreateSubscribeRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateSubscribeRelation", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSubscribeRelationResponse creates a response to parse from CreateSubscribeRelation response
func CreateCreateSubscribeRelationResponse() (response *CreateSubscribeRelationResponse) {
	response = &CreateSubscribeRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
