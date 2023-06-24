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

// DisableVnInstance invokes the cloudcallcenter.DisableVnInstance API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/disablevninstance.html
func (client *Client) DisableVnInstance(request *DisableVnInstanceRequest) (response *DisableVnInstanceResponse, err error) {
	response = CreateDisableVnInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DisableVnInstanceWithChan invokes the cloudcallcenter.DisableVnInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/disablevninstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableVnInstanceWithChan(request *DisableVnInstanceRequest) (<-chan *DisableVnInstanceResponse, <-chan error) {
	responseChan := make(chan *DisableVnInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableVnInstance(request)
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

// DisableVnInstanceWithCallback invokes the cloudcallcenter.DisableVnInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/disablevninstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableVnInstanceWithCallback(request *DisableVnInstanceRequest, callback func(response *DisableVnInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableVnInstanceResponse
		var err error
		defer close(result)
		response, err = client.DisableVnInstance(request)
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

// DisableVnInstanceRequest is the request struct for api DisableVnInstance
type DisableVnInstanceRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DisableVnInstanceResponse is the response struct for api DisableVnInstance
type DisableVnInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateDisableVnInstanceRequest creates a request to invoke DisableVnInstance API
func CreateDisableVnInstanceRequest() (request *DisableVnInstanceRequest) {
	request = &DisableVnInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DisableVnInstance", "", "")
	request.Method = requests.GET
	return
}

// CreateDisableVnInstanceResponse creates a response to parse from DisableVnInstance response
func CreateDisableVnInstanceResponse() (response *DisableVnInstanceResponse) {
	response = &DisableVnInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
