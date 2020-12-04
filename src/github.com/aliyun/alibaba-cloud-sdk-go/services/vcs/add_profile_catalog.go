package vcs

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

// AddProfileCatalog invokes the vcs.AddProfileCatalog API synchronously
func (client *Client) AddProfileCatalog(request *AddProfileCatalogRequest) (response *AddProfileCatalogResponse, err error) {
	response = CreateAddProfileCatalogResponse()
	err = client.DoAction(request, response)
	return
}

// AddProfileCatalogWithChan invokes the vcs.AddProfileCatalog API asynchronously
func (client *Client) AddProfileCatalogWithChan(request *AddProfileCatalogRequest) (<-chan *AddProfileCatalogResponse, <-chan error) {
	responseChan := make(chan *AddProfileCatalogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddProfileCatalog(request)
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

// AddProfileCatalogWithCallback invokes the vcs.AddProfileCatalog API asynchronously
func (client *Client) AddProfileCatalogWithCallback(request *AddProfileCatalogRequest, callback func(response *AddProfileCatalogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddProfileCatalogResponse
		var err error
		defer close(result)
		response, err = client.AddProfileCatalog(request)
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

// AddProfileCatalogRequest is the request struct for api AddProfileCatalog
type AddProfileCatalogRequest struct {
	*requests.RpcRequest
	IsvSubId        string           `position:"Body" name:"IsvSubId"`
	ParentCatalogId requests.Integer `position:"Body" name:"ParentCatalogId"`
	CorpId          string           `position:"Body" name:"CorpId"`
	CatalogName     string           `position:"Body" name:"CatalogName"`
}

// AddProfileCatalogResponse is the response struct for api AddProfileCatalog
type AddProfileCatalogResponse struct {
	*responses.BaseResponse
	Code      string                  `json:"Code" xml:"Code"`
	Message   string                  `json:"Message" xml:"Message"`
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Data      DataInAddProfileCatalog `json:"Data" xml:"Data"`
}

// CreateAddProfileCatalogRequest creates a request to invoke AddProfileCatalog API
func CreateAddProfileCatalogRequest() (request *AddProfileCatalogRequest) {
	request = &AddProfileCatalogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "AddProfileCatalog", "", "")
	request.Method = requests.POST
	return
}

// CreateAddProfileCatalogResponse creates a response to parse from AddProfileCatalog response
func CreateAddProfileCatalogResponse() (response *AddProfileCatalogResponse) {
	response = &AddProfileCatalogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
