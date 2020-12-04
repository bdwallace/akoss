package vs

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

// ModifyGroup invokes the vs.ModifyGroup API synchronously
func (client *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
	response = CreateModifyGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGroupWithChan invokes the vs.ModifyGroup API asynchronously
func (client *Client) ModifyGroupWithChan(request *ModifyGroupRequest) (<-chan *ModifyGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGroup(request)
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

// ModifyGroupWithCallback invokes the vs.ModifyGroup API asynchronously
func (client *Client) ModifyGroupWithCallback(request *ModifyGroupRequest, callback func(response *ModifyGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyGroup(request)
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

// ModifyGroupRequest is the request struct for api ModifyGroup
type ModifyGroupRequest struct {
	*requests.RpcRequest
	CaptureVideo     requests.Integer `position:"Query" name:"CaptureVideo"`
	Description      string           `position:"Query" name:"Description"`
	Enabled          requests.Boolean `position:"Query" name:"Enabled"`
	CaptureOssPath   string           `position:"Query" name:"CaptureOssPath"`
	PushDomain       string           `position:"Query" name:"PushDomain"`
	CaptureImage     requests.Integer `position:"Query" name:"CaptureImage"`
	Id               string           `position:"Query" name:"Id"`
	ShowLog          string           `position:"Query" name:"ShowLog"`
	PlayDomain       string           `position:"Query" name:"PlayDomain"`
	OutProtocol      string           `position:"Query" name:"OutProtocol"`
	CaptureInterval  requests.Integer `position:"Query" name:"CaptureInterval"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	InProtocol       string           `position:"Query" name:"InProtocol"`
	LazyPull         requests.Boolean `position:"Query" name:"LazyPull"`
	Name             string           `position:"Query" name:"Name"`
	Callback         string           `position:"Query" name:"Callback"`
	Region           string           `position:"Query" name:"Region"`
	CaptureOssBucket string           `position:"Query" name:"CaptureOssBucket"`
}

// ModifyGroupResponse is the response struct for api ModifyGroup
type ModifyGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateModifyGroupRequest creates a request to invoke ModifyGroup API
func CreateModifyGroupRequest() (request *ModifyGroupRequest) {
	request = &ModifyGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ModifyGroup", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyGroupResponse creates a response to parse from ModifyGroup response
func CreateModifyGroupResponse() (response *ModifyGroupResponse) {
	response = &ModifyGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
