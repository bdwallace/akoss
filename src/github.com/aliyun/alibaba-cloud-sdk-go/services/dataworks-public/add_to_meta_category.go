package dataworks_public

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

// AddToMetaCategory invokes the dataworks_public.AddToMetaCategory API synchronously
func (client *Client) AddToMetaCategory(request *AddToMetaCategoryRequest) (response *AddToMetaCategoryResponse, err error) {
	response = CreateAddToMetaCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// AddToMetaCategoryWithChan invokes the dataworks_public.AddToMetaCategory API asynchronously
func (client *Client) AddToMetaCategoryWithChan(request *AddToMetaCategoryRequest) (<-chan *AddToMetaCategoryResponse, <-chan error) {
	responseChan := make(chan *AddToMetaCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddToMetaCategory(request)
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

// AddToMetaCategoryWithCallback invokes the dataworks_public.AddToMetaCategory API asynchronously
func (client *Client) AddToMetaCategoryWithCallback(request *AddToMetaCategoryRequest, callback func(response *AddToMetaCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddToMetaCategoryResponse
		var err error
		defer close(result)
		response, err = client.AddToMetaCategory(request)
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

// AddToMetaCategoryRequest is the request struct for api AddToMetaCategory
type AddToMetaCategoryRequest struct {
	*requests.RpcRequest
	TableGuid  string           `position:"Query" name:"TableGuid"`
	CategoryId requests.Integer `position:"Query" name:"CategoryId"`
}

// AddToMetaCategoryResponse is the response struct for api AddToMetaCategory
type AddToMetaCategoryResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           bool   `json:"Data" xml:"Data"`
}

// CreateAddToMetaCategoryRequest creates a request to invoke AddToMetaCategory API
func CreateAddToMetaCategoryRequest() (request *AddToMetaCategoryRequest) {
	request = &AddToMetaCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "AddToMetaCategory", "", "")
	request.Method = requests.POST
	return
}

// CreateAddToMetaCategoryResponse creates a response to parse from AddToMetaCategory response
func CreateAddToMetaCategoryResponse() (response *AddToMetaCategoryResponse) {
	response = &AddToMetaCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
