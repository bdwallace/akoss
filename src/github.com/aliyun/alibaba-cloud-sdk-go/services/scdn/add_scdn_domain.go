package scdn

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

// AddScdnDomain invokes the scdn.AddScdnDomain API synchronously
func (client *Client) AddScdnDomain(request *AddScdnDomainRequest) (response *AddScdnDomainResponse, err error) {
	response = CreateAddScdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// AddScdnDomainWithChan invokes the scdn.AddScdnDomain API asynchronously
func (client *Client) AddScdnDomainWithChan(request *AddScdnDomainRequest) (<-chan *AddScdnDomainResponse, <-chan error) {
	responseChan := make(chan *AddScdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddScdnDomain(request)
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

// AddScdnDomainWithCallback invokes the scdn.AddScdnDomain API asynchronously
func (client *Client) AddScdnDomainWithCallback(request *AddScdnDomainRequest, callback func(response *AddScdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddScdnDomainResponse
		var err error
		defer close(result)
		response, err = client.AddScdnDomain(request)
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

// AddScdnDomainRequest is the request struct for api AddScdnDomain
type AddScdnDomainRequest struct {
	*requests.RpcRequest
	Sources         string           `position:"Query" name:"Sources"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	Scope           string           `position:"Query" name:"Scope"`
	OwnerAccount    string           `position:"Query" name:"OwnerAccount"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	CheckUrl        string           `position:"Query" name:"CheckUrl"`
}

// AddScdnDomainResponse is the response struct for api AddScdnDomain
type AddScdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddScdnDomainRequest creates a request to invoke AddScdnDomain API
func CreateAddScdnDomainRequest() (request *AddScdnDomainRequest) {
	request = &AddScdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "AddScdnDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateAddScdnDomainResponse creates a response to parse from AddScdnDomain response
func CreateAddScdnDomainResponse() (response *AddScdnDomainResponse) {
	response = &AddScdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
