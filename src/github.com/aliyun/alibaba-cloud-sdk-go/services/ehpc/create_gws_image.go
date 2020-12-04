package ehpc

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

// CreateGWSImage invokes the ehpc.CreateGWSImage API synchronously
// api document: https://help.aliyun.com/api/ehpc/creategwsimage.html
func (client *Client) CreateGWSImage(request *CreateGWSImageRequest) (response *CreateGWSImageResponse, err error) {
	response = CreateCreateGWSImageResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGWSImageWithChan invokes the ehpc.CreateGWSImage API asynchronously
// api document: https://help.aliyun.com/api/ehpc/creategwsimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGWSImageWithChan(request *CreateGWSImageRequest) (<-chan *CreateGWSImageResponse, <-chan error) {
	responseChan := make(chan *CreateGWSImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGWSImage(request)
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

// CreateGWSImageWithCallback invokes the ehpc.CreateGWSImage API asynchronously
// api document: https://help.aliyun.com/api/ehpc/creategwsimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGWSImageWithCallback(request *CreateGWSImageRequest, callback func(response *CreateGWSImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGWSImageResponse
		var err error
		defer close(result)
		response, err = client.CreateGWSImage(request)
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

// CreateGWSImageRequest is the request struct for api CreateGWSImage
type CreateGWSImageRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
}

// CreateGWSImageResponse is the response struct for api CreateGWSImage
type CreateGWSImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ImageId   string `json:"ImageId" xml:"ImageId"`
}

// CreateCreateGWSImageRequest creates a request to invoke CreateGWSImage API
func CreateCreateGWSImageRequest() (request *CreateGWSImageRequest) {
	request = &CreateGWSImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "CreateGWSImage", "", "")
	request.Method = requests.GET
	return
}

// CreateCreateGWSImageResponse creates a response to parse from CreateGWSImage response
func CreateCreateGWSImageResponse() (response *CreateGWSImageResponse) {
	response = &CreateGWSImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
