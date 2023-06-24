package imgsearch

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

// CreateImageDb invokes the imgsearch.CreateImageDb API synchronously
func (client *Client) CreateImageDb(request *CreateImageDbRequest) (response *CreateImageDbResponse, err error) {
	response = CreateCreateImageDbResponse()
	err = client.DoAction(request, response)
	return
}

// CreateImageDbWithChan invokes the imgsearch.CreateImageDb API asynchronously
func (client *Client) CreateImageDbWithChan(request *CreateImageDbRequest) (<-chan *CreateImageDbResponse, <-chan error) {
	responseChan := make(chan *CreateImageDbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateImageDb(request)
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

// CreateImageDbWithCallback invokes the imgsearch.CreateImageDb API asynchronously
func (client *Client) CreateImageDbWithCallback(request *CreateImageDbRequest, callback func(response *CreateImageDbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateImageDbResponse
		var err error
		defer close(result)
		response, err = client.CreateImageDb(request)
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

// CreateImageDbRequest is the request struct for api CreateImageDb
type CreateImageDbRequest struct {
	*requests.RpcRequest
	Name string `position:"Body" name:"Name"`
}

// CreateImageDbResponse is the response struct for api CreateImageDb
type CreateImageDbResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateImageDbRequest creates a request to invoke CreateImageDb API
func CreateCreateImageDbRequest() (request *CreateImageDbRequest) {
	request = &CreateImageDbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imgsearch", "2020-03-20", "CreateImageDb", "imgsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateImageDbResponse creates a response to parse from CreateImageDb response
func CreateCreateImageDbResponse() (response *CreateImageDbResponse) {
	response = &CreateImageDbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
