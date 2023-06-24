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

// QuerySubscribeRelation invokes the iot.QuerySubscribeRelation API synchronously
func (client *Client) QuerySubscribeRelation(request *QuerySubscribeRelationRequest) (response *QuerySubscribeRelationResponse, err error) {
	response = CreateQuerySubscribeRelationResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySubscribeRelationWithChan invokes the iot.QuerySubscribeRelation API asynchronously
func (client *Client) QuerySubscribeRelationWithChan(request *QuerySubscribeRelationRequest) (<-chan *QuerySubscribeRelationResponse, <-chan error) {
	responseChan := make(chan *QuerySubscribeRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySubscribeRelation(request)
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

// QuerySubscribeRelationWithCallback invokes the iot.QuerySubscribeRelation API asynchronously
func (client *Client) QuerySubscribeRelationWithCallback(request *QuerySubscribeRelationRequest, callback func(response *QuerySubscribeRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySubscribeRelationResponse
		var err error
		defer close(result)
		response, err = client.QuerySubscribeRelation(request)
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

// QuerySubscribeRelationRequest is the request struct for api QuerySubscribeRelation
type QuerySubscribeRelationRequest struct {
	*requests.RpcRequest
	Type          string `position:"Query" name:"Type"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// QuerySubscribeRelationResponse is the response struct for api QuerySubscribeRelation
type QuerySubscribeRelationResponse struct {
	*responses.BaseResponse
	RequestId               string   `json:"RequestId" xml:"RequestId"`
	Success                 bool     `json:"Success" xml:"Success"`
	Code                    string   `json:"Code" xml:"Code"`
	ErrorMessage            string   `json:"ErrorMessage" xml:"ErrorMessage"`
	ProductKey              string   `json:"ProductKey" xml:"ProductKey"`
	Type                    string   `json:"Type" xml:"Type"`
	DeviceDataFlag          bool     `json:"DeviceDataFlag" xml:"DeviceDataFlag"`
	DeviceLifeCycleFlag     bool     `json:"DeviceLifeCycleFlag" xml:"DeviceLifeCycleFlag"`
	DeviceStatusChangeFlag  bool     `json:"DeviceStatusChangeFlag" xml:"DeviceStatusChangeFlag"`
	DeviceTopoLifeCycleFlag bool     `json:"DeviceTopoLifeCycleFlag" xml:"DeviceTopoLifeCycleFlag"`
	FoundDeviceListFlag     bool     `json:"FoundDeviceListFlag" xml:"FoundDeviceListFlag"`
	OtaEventFlag            bool     `json:"OtaEventFlag" xml:"OtaEventFlag"`
	ThingHistoryFlag        bool     `json:"ThingHistoryFlag" xml:"ThingHistoryFlag"`
	MnsConfiguration        string   `json:"MnsConfiguration" xml:"MnsConfiguration"`
	DeviceTagFlag           bool     `json:"DeviceTagFlag" xml:"DeviceTagFlag"`
	OtaVersionFlag          bool     `json:"OtaVersionFlag" xml:"OtaVersionFlag"`
	OtaJobFlag              bool     `json:"OtaJobFlag" xml:"OtaJobFlag"`
	ConsumerGroupIds        []string `json:"ConsumerGroupIds" xml:"ConsumerGroupIds"`
}

// CreateQuerySubscribeRelationRequest creates a request to invoke QuerySubscribeRelation API
func CreateQuerySubscribeRelationRequest() (request *QuerySubscribeRelationRequest) {
	request = &QuerySubscribeRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySubscribeRelation", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySubscribeRelationResponse creates a response to parse from QuerySubscribeRelation response
func CreateQuerySubscribeRelationResponse() (response *QuerySubscribeRelationResponse) {
	response = &QuerySubscribeRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}