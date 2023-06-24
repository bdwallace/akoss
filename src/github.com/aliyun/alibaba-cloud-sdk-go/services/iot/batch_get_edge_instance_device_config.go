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

// BatchGetEdgeInstanceDeviceConfig invokes the iot.BatchGetEdgeInstanceDeviceConfig API synchronously
func (client *Client) BatchGetEdgeInstanceDeviceConfig(request *BatchGetEdgeInstanceDeviceConfigRequest) (response *BatchGetEdgeInstanceDeviceConfigResponse, err error) {
	response = CreateBatchGetEdgeInstanceDeviceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// BatchGetEdgeInstanceDeviceConfigWithChan invokes the iot.BatchGetEdgeInstanceDeviceConfig API asynchronously
func (client *Client) BatchGetEdgeInstanceDeviceConfigWithChan(request *BatchGetEdgeInstanceDeviceConfigRequest) (<-chan *BatchGetEdgeInstanceDeviceConfigResponse, <-chan error) {
	responseChan := make(chan *BatchGetEdgeInstanceDeviceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchGetEdgeInstanceDeviceConfig(request)
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

// BatchGetEdgeInstanceDeviceConfigWithCallback invokes the iot.BatchGetEdgeInstanceDeviceConfig API asynchronously
func (client *Client) BatchGetEdgeInstanceDeviceConfigWithCallback(request *BatchGetEdgeInstanceDeviceConfigRequest, callback func(response *BatchGetEdgeInstanceDeviceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchGetEdgeInstanceDeviceConfigResponse
		var err error
		defer close(result)
		response, err = client.BatchGetEdgeInstanceDeviceConfig(request)
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

// BatchGetEdgeInstanceDeviceConfigRequest is the request struct for api BatchGetEdgeInstanceDeviceConfig
type BatchGetEdgeInstanceDeviceConfigRequest struct {
	*requests.RpcRequest
	IotIds        *[]string `position:"Query" name:"IotIds"  type:"Repeated"`
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
	InstanceId    string    `position:"Query" name:"InstanceId"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
}

// BatchGetEdgeInstanceDeviceConfigResponse is the response struct for api BatchGetEdgeInstanceDeviceConfig
type BatchGetEdgeInstanceDeviceConfigResponse struct {
	*responses.BaseResponse
	RequestId        string         `json:"RequestId" xml:"RequestId"`
	Success          bool           `json:"Success" xml:"Success"`
	Code             string         `json:"Code" xml:"Code"`
	ErrorMessage     string         `json:"ErrorMessage" xml:"ErrorMessage"`
	DeviceConfigList []DeviceConfig `json:"DeviceConfigList" xml:"DeviceConfigList"`
}

// CreateBatchGetEdgeInstanceDeviceConfigRequest creates a request to invoke BatchGetEdgeInstanceDeviceConfig API
func CreateBatchGetEdgeInstanceDeviceConfigRequest() (request *BatchGetEdgeInstanceDeviceConfigRequest) {
	request = &BatchGetEdgeInstanceDeviceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchGetEdgeInstanceDeviceConfig", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchGetEdgeInstanceDeviceConfigResponse creates a response to parse from BatchGetEdgeInstanceDeviceConfig response
func CreateBatchGetEdgeInstanceDeviceConfigResponse() (response *BatchGetEdgeInstanceDeviceConfigResponse) {
	response = &BatchGetEdgeInstanceDeviceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
