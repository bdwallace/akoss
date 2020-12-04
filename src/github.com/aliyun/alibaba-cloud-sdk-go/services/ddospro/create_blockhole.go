package ddospro

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

// CreateBlockhole invokes the ddospro.CreateBlockhole API synchronously
// api document: https://help.aliyun.com/api/ddospro/createblockhole.html
func (client *Client) CreateBlockhole(request *CreateBlockholeRequest) (response *CreateBlockholeResponse, err error) {
	response = CreateCreateBlockholeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBlockholeWithChan invokes the ddospro.CreateBlockhole API asynchronously
// api document: https://help.aliyun.com/api/ddospro/createblockhole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBlockholeWithChan(request *CreateBlockholeRequest) (<-chan *CreateBlockholeResponse, <-chan error) {
	responseChan := make(chan *CreateBlockholeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBlockhole(request)
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

// CreateBlockholeWithCallback invokes the ddospro.CreateBlockhole API asynchronously
// api document: https://help.aliyun.com/api/ddospro/createblockhole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBlockholeWithCallback(request *CreateBlockholeRequest, callback func(response *CreateBlockholeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBlockholeResponse
		var err error
		defer close(result)
		response, err = client.CreateBlockhole(request)
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

// CreateBlockholeRequest is the request struct for api CreateBlockhole
type CreateBlockholeRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Vip             string           `position:"Query" name:"Vip"`
	BlockZone       string           `position:"Query" name:"BlockZone"`
	BlockTime       requests.Integer `position:"Query" name:"BlockTime"`
}

// CreateBlockholeResponse is the response struct for api CreateBlockhole
type CreateBlockholeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateBlockholeRequest creates a request to invoke CreateBlockhole API
func CreateCreateBlockholeRequest() (request *CreateBlockholeRequest) {
	request = &CreateBlockholeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "CreateBlockhole", "", "")
	return
}

// CreateCreateBlockholeResponse creates a response to parse from CreateBlockhole response
func CreateCreateBlockholeResponse() (response *CreateBlockholeResponse) {
	response = &CreateBlockholeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
