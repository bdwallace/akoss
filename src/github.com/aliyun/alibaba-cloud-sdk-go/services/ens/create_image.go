package ens

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

// CreateImage invokes the ens.CreateImage API synchronously
func (client *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
	response = CreateCreateImageResponse()
	err = client.DoAction(request, response)
	return
}

// CreateImageWithChan invokes the ens.CreateImage API asynchronously
func (client *Client) CreateImageWithChan(request *CreateImageRequest) (<-chan *CreateImageResponse, <-chan error) {
	responseChan := make(chan *CreateImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateImage(request)
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

// CreateImageWithCallback invokes the ens.CreateImage API asynchronously
func (client *Client) CreateImageWithCallback(request *CreateImageRequest, callback func(response *CreateImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateImageResponse
		var err error
		defer close(result)
		response, err = client.CreateImage(request)
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

// CreateImageRequest is the request struct for api CreateImage
type CreateImageRequest struct {
	*requests.RpcRequest
	DeleteAfterImageUpload string `position:"Query" name:"DeleteAfterImageUpload"`
	ImageName              string `position:"Query" name:"ImageName"`
	Product                string `position:"Query" name:"product"`
	Version                string `position:"Query" name:"Version"`
	InstanceId             string `position:"Query" name:"InstanceId"`
}

// CreateImageResponse is the response struct for api CreateImage
type CreateImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
}

// CreateCreateImageRequest creates a request to invoke CreateImage API
func CreateCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateImage", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateImageResponse creates a response to parse from CreateImage response
func CreateCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
