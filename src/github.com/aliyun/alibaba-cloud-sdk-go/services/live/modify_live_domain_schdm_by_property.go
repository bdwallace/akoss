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

// ModifyLiveDomainSchdmByProperty invokes the live.ModifyLiveDomainSchdmByProperty API synchronously
func (client *Client) ModifyLiveDomainSchdmByProperty(request *ModifyLiveDomainSchdmByPropertyRequest) (response *ModifyLiveDomainSchdmByPropertyResponse, err error) {
	response = CreateModifyLiveDomainSchdmByPropertyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLiveDomainSchdmByPropertyWithChan invokes the live.ModifyLiveDomainSchdmByProperty API asynchronously
func (client *Client) ModifyLiveDomainSchdmByPropertyWithChan(request *ModifyLiveDomainSchdmByPropertyRequest) (<-chan *ModifyLiveDomainSchdmByPropertyResponse, <-chan error) {
	responseChan := make(chan *ModifyLiveDomainSchdmByPropertyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLiveDomainSchdmByProperty(request)
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

// ModifyLiveDomainSchdmByPropertyWithCallback invokes the live.ModifyLiveDomainSchdmByProperty API asynchronously
func (client *Client) ModifyLiveDomainSchdmByPropertyWithCallback(request *ModifyLiveDomainSchdmByPropertyRequest, callback func(response *ModifyLiveDomainSchdmByPropertyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLiveDomainSchdmByPropertyResponse
		var err error
		defer close(result)
		response, err = client.ModifyLiveDomainSchdmByProperty(request)
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

// ModifyLiveDomainSchdmByPropertyRequest is the request struct for api ModifyLiveDomainSchdmByProperty
type ModifyLiveDomainSchdmByPropertyRequest struct {
	*requests.RpcRequest
	Property   string           `position:"Query" name:"Property"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyLiveDomainSchdmByPropertyResponse is the response struct for api ModifyLiveDomainSchdmByProperty
type ModifyLiveDomainSchdmByPropertyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLiveDomainSchdmByPropertyRequest creates a request to invoke ModifyLiveDomainSchdmByProperty API
func CreateModifyLiveDomainSchdmByPropertyRequest() (request *ModifyLiveDomainSchdmByPropertyRequest) {
	request = &ModifyLiveDomainSchdmByPropertyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ModifyLiveDomainSchdmByProperty", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLiveDomainSchdmByPropertyResponse creates a response to parse from ModifyLiveDomainSchdmByProperty response
func CreateModifyLiveDomainSchdmByPropertyResponse() (response *ModifyLiveDomainSchdmByPropertyResponse) {
	response = &ModifyLiveDomainSchdmByPropertyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
