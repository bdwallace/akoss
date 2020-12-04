package cloudcallcenter

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

// ReleaseCabService invokes the cloudcallcenter.ReleaseCabService API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/releasecabservice.html
func (client *Client) ReleaseCabService(request *ReleaseCabServiceRequest) (response *ReleaseCabServiceResponse, err error) {
	response = CreateReleaseCabServiceResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseCabServiceWithChan invokes the cloudcallcenter.ReleaseCabService API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/releasecabservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseCabServiceWithChan(request *ReleaseCabServiceRequest) (<-chan *ReleaseCabServiceResponse, <-chan error) {
	responseChan := make(chan *ReleaseCabServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseCabService(request)
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

// ReleaseCabServiceWithCallback invokes the cloudcallcenter.ReleaseCabService API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/releasecabservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseCabServiceWithCallback(request *ReleaseCabServiceRequest, callback func(response *ReleaseCabServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseCabServiceResponse
		var err error
		defer close(result)
		response, err = client.ReleaseCabService(request)
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

// ReleaseCabServiceRequest is the request struct for api ReleaseCabService
type ReleaseCabServiceRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// ReleaseCabServiceResponse is the response struct for api ReleaseCabService
type ReleaseCabServiceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateReleaseCabServiceRequest creates a request to invoke ReleaseCabService API
func CreateReleaseCabServiceRequest() (request *ReleaseCabServiceRequest) {
	request = &ReleaseCabServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ReleaseCabService", "", "")
	request.Method = requests.POST
	return
}

// CreateReleaseCabServiceResponse creates a response to parse from ReleaseCabService response
func CreateReleaseCabServiceResponse() (response *ReleaseCabServiceResponse) {
	response = &ReleaseCabServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
