package cloudesl

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

// DeleteItem invokes the cloudesl.DeleteItem API synchronously
func (client *Client) DeleteItem(request *DeleteItemRequest) (response *DeleteItemResponse, err error) {
	response = CreateDeleteItemResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteItemWithChan invokes the cloudesl.DeleteItem API asynchronously
func (client *Client) DeleteItemWithChan(request *DeleteItemRequest) (<-chan *DeleteItemResponse, <-chan error) {
	responseChan := make(chan *DeleteItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteItem(request)
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

// DeleteItemWithCallback invokes the cloudesl.DeleteItem API asynchronously
func (client *Client) DeleteItemWithCallback(request *DeleteItemRequest, callback func(response *DeleteItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteItemResponse
		var err error
		defer close(result)
		response, err = client.DeleteItem(request)
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

// DeleteItemRequest is the request struct for api DeleteItem
type DeleteItemRequest struct {
	*requests.RpcRequest
	StoreId     string `position:"Query" name:"StoreId"`
	ItemBarCode string `position:"Query" name:"ItemBarCode"`
}

// DeleteItemResponse is the response struct for api DeleteItem
type DeleteItemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateDeleteItemRequest creates a request to invoke DeleteItem API
func CreateDeleteItemRequest() (request *DeleteItemRequest) {
	request = &DeleteItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2018-08-01", "DeleteItem", "cloudesl", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteItemResponse creates a response to parse from DeleteItem response
func CreateDeleteItemResponse() (response *DeleteItemResponse) {
	response = &DeleteItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
