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

// DeleteLiveRealtimeLogDelivery invokes the live.DeleteLiveRealtimeLogDelivery API synchronously
func (client *Client) DeleteLiveRealtimeLogDelivery(request *DeleteLiveRealtimeLogDeliveryRequest) (response *DeleteLiveRealtimeLogDeliveryResponse, err error) {
	response = CreateDeleteLiveRealtimeLogDeliveryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveRealtimeLogDeliveryWithChan invokes the live.DeleteLiveRealtimeLogDelivery API asynchronously
func (client *Client) DeleteLiveRealtimeLogDeliveryWithChan(request *DeleteLiveRealtimeLogDeliveryRequest) (<-chan *DeleteLiveRealtimeLogDeliveryResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveRealtimeLogDeliveryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveRealtimeLogDelivery(request)
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

// DeleteLiveRealtimeLogDeliveryWithCallback invokes the live.DeleteLiveRealtimeLogDelivery API asynchronously
func (client *Client) DeleteLiveRealtimeLogDeliveryWithCallback(request *DeleteLiveRealtimeLogDeliveryRequest, callback func(response *DeleteLiveRealtimeLogDeliveryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveRealtimeLogDeliveryResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveRealtimeLogDelivery(request)
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

// DeleteLiveRealtimeLogDeliveryRequest is the request struct for api DeleteLiveRealtimeLogDelivery
type DeleteLiveRealtimeLogDeliveryRequest struct {
	*requests.RpcRequest
	Project    string           `position:"Query" name:"Project"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Region     string           `position:"Query" name:"Region"`
	Logstore   string           `position:"Query" name:"Logstore"`
}

// DeleteLiveRealtimeLogDeliveryResponse is the response struct for api DeleteLiveRealtimeLogDelivery
type DeleteLiveRealtimeLogDeliveryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveRealtimeLogDeliveryRequest creates a request to invoke DeleteLiveRealtimeLogDelivery API
func CreateDeleteLiveRealtimeLogDeliveryRequest() (request *DeleteLiveRealtimeLogDeliveryRequest) {
	request = &DeleteLiveRealtimeLogDeliveryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveRealtimeLogDelivery", "live", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDeleteLiveRealtimeLogDeliveryResponse creates a response to parse from DeleteLiveRealtimeLogDelivery response
func CreateDeleteLiveRealtimeLogDeliveryResponse() (response *DeleteLiveRealtimeLogDeliveryResponse) {
	response = &DeleteLiveRealtimeLogDeliveryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
