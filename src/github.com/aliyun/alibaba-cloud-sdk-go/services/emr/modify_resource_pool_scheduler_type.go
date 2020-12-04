package emr

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

// ModifyResourcePoolSchedulerType invokes the emr.ModifyResourcePoolSchedulerType API synchronously
// api document: https://help.aliyun.com/api/emr/modifyresourcepoolschedulertype.html
func (client *Client) ModifyResourcePoolSchedulerType(request *ModifyResourcePoolSchedulerTypeRequest) (response *ModifyResourcePoolSchedulerTypeResponse, err error) {
	response = CreateModifyResourcePoolSchedulerTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyResourcePoolSchedulerTypeWithChan invokes the emr.ModifyResourcePoolSchedulerType API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyresourcepoolschedulertype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyResourcePoolSchedulerTypeWithChan(request *ModifyResourcePoolSchedulerTypeRequest) (<-chan *ModifyResourcePoolSchedulerTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyResourcePoolSchedulerTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyResourcePoolSchedulerType(request)
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

// ModifyResourcePoolSchedulerTypeWithCallback invokes the emr.ModifyResourcePoolSchedulerType API asynchronously
// api document: https://help.aliyun.com/api/emr/modifyresourcepoolschedulertype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyResourcePoolSchedulerTypeWithCallback(request *ModifyResourcePoolSchedulerTypeRequest, callback func(response *ModifyResourcePoolSchedulerTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyResourcePoolSchedulerTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyResourcePoolSchedulerType(request)
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

// ModifyResourcePoolSchedulerTypeRequest is the request struct for api ModifyResourcePoolSchedulerType
type ModifyResourcePoolSchedulerTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	SchedulerType   string           `position:"Query" name:"SchedulerType"`
}

// ModifyResourcePoolSchedulerTypeResponse is the response struct for api ModifyResourcePoolSchedulerType
type ModifyResourcePoolSchedulerTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyResourcePoolSchedulerTypeRequest creates a request to invoke ModifyResourcePoolSchedulerType API
func CreateModifyResourcePoolSchedulerTypeRequest() (request *ModifyResourcePoolSchedulerTypeRequest) {
	request = &ModifyResourcePoolSchedulerTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyResourcePoolSchedulerType", "emr", "openAPI")
	return
}

// CreateModifyResourcePoolSchedulerTypeResponse creates a response to parse from ModifyResourcePoolSchedulerType response
func CreateModifyResourcePoolSchedulerTypeResponse() (response *ModifyResourcePoolSchedulerTypeResponse) {
	response = &ModifyResourcePoolSchedulerTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
