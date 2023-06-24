package smartag

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

// AssociateFlowLog invokes the smartag.AssociateFlowLog API synchronously
func (client *Client) AssociateFlowLog(request *AssociateFlowLogRequest) (response *AssociateFlowLogResponse, err error) {
	response = CreateAssociateFlowLogResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateFlowLogWithChan invokes the smartag.AssociateFlowLog API asynchronously
func (client *Client) AssociateFlowLogWithChan(request *AssociateFlowLogRequest) (<-chan *AssociateFlowLogResponse, <-chan error) {
	responseChan := make(chan *AssociateFlowLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateFlowLog(request)
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

// AssociateFlowLogWithCallback invokes the smartag.AssociateFlowLog API asynchronously
func (client *Client) AssociateFlowLogWithCallback(request *AssociateFlowLogRequest, callback func(response *AssociateFlowLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateFlowLogResponse
		var err error
		defer close(result)
		response, err = client.AssociateFlowLog(request)
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

// AssociateFlowLogRequest is the request struct for api AssociateFlowLog
type AssociateFlowLogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	FlowLogId            string           `position:"Query" name:"FlowLogId"`
}

// AssociateFlowLogResponse is the response struct for api AssociateFlowLog
type AssociateFlowLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateFlowLogRequest creates a request to invoke AssociateFlowLog API
func CreateAssociateFlowLogRequest() (request *AssociateFlowLogRequest) {
	request = &AssociateFlowLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "AssociateFlowLog", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateFlowLogResponse creates a response to parse from AssociateFlowLog response
func CreateAssociateFlowLogResponse() (response *AssociateFlowLogResponse) {
	response = &AssociateFlowLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
