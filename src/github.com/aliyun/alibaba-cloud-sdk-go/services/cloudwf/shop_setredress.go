package cloudwf

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

// ShopSetredress invokes the cloudwf.ShopSetredress API synchronously
// api document: https://help.aliyun.com/api/cloudwf/shopsetredress.html
func (client *Client) ShopSetredress(request *ShopSetredressRequest) (response *ShopSetredressResponse, err error) {
	response = CreateShopSetredressResponse()
	err = client.DoAction(request, response)
	return
}

// ShopSetredressWithChan invokes the cloudwf.ShopSetredress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopsetredress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopSetredressWithChan(request *ShopSetredressRequest) (<-chan *ShopSetredressResponse, <-chan error) {
	responseChan := make(chan *ShopSetredressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ShopSetredress(request)
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

// ShopSetredressWithCallback invokes the cloudwf.ShopSetredress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopsetredress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopSetredressWithCallback(request *ShopSetredressRequest, callback func(response *ShopSetredressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ShopSetredressResponse
		var err error
		defer close(result)
		response, err = client.ShopSetredress(request)
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

// ShopSetredressRequest is the request struct for api ShopSetredress
type ShopSetredressRequest struct {
	*requests.RpcRequest
	Workday     string           `position:"Query" name:"Workday"`
	Filterclose requests.Integer `position:"Query" name:"Filterclose"`
	Minstoptime requests.Integer `position:"Query" name:"Minstoptime"`
	Holiday     string           `position:"Query" name:"Holiday"`
	Hnum        string           `position:"Query" name:"Hnum"`
	Sid         requests.Integer `position:"Query" name:"Sid"`
	Clerk       requests.Integer `position:"Query" name:"Clerk"`
	Filterstate requests.Integer `position:"Query" name:"Filterstate"`
	Wnum        string           `position:"Query" name:"Wnum"`
	State       requests.Integer `position:"Query" name:"State"`
	Crowdfixed  requests.Integer `position:"Query" name:"Crowdfixed"`
	Maxstoptime requests.Integer `position:"Query" name:"Maxstoptime"`
}

// ShopSetredressResponse is the response struct for api ShopSetredress
type ShopSetredressResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateShopSetredressRequest creates a request to invoke ShopSetredress API
func CreateShopSetredressRequest() (request *ShopSetredressRequest) {
	request = &ShopSetredressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ShopSetredress", "cloudwf", "openAPI")
	return
}

// CreateShopSetredressResponse creates a response to parse from ShopSetredress response
func CreateShopSetredressResponse() (response *ShopSetredressResponse) {
	response = &ShopSetredressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
