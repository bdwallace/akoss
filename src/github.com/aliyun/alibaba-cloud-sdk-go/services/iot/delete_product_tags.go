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

// DeleteProductTags invokes the iot.DeleteProductTags API synchronously
func (client *Client) DeleteProductTags(request *DeleteProductTagsRequest) (response *DeleteProductTagsResponse, err error) {
	response = CreateDeleteProductTagsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteProductTagsWithChan invokes the iot.DeleteProductTags API asynchronously
func (client *Client) DeleteProductTagsWithChan(request *DeleteProductTagsRequest) (<-chan *DeleteProductTagsResponse, <-chan error) {
	responseChan := make(chan *DeleteProductTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteProductTags(request)
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

// DeleteProductTagsWithCallback invokes the iot.DeleteProductTags API asynchronously
func (client *Client) DeleteProductTagsWithCallback(request *DeleteProductTagsRequest, callback func(response *DeleteProductTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteProductTagsResponse
		var err error
		defer close(result)
		response, err = client.DeleteProductTags(request)
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

// DeleteProductTagsRequest is the request struct for api DeleteProductTags
type DeleteProductTagsRequest struct {
	*requests.RpcRequest
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
	ProductTagKey *[]string `position:"Query" name:"ProductTagKey"  type:"Repeated"`
	ProductKey    string    `position:"Query" name:"ProductKey"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
}

// DeleteProductTagsResponse is the response struct for api DeleteProductTags
type DeleteProductTagsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
}

// CreateDeleteProductTagsRequest creates a request to invoke DeleteProductTags API
func CreateDeleteProductTagsRequest() (request *DeleteProductTagsRequest) {
	request = &DeleteProductTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteProductTags", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteProductTagsResponse creates a response to parse from DeleteProductTags response
func CreateDeleteProductTagsResponse() (response *DeleteProductTagsResponse) {
	response = &DeleteProductTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
