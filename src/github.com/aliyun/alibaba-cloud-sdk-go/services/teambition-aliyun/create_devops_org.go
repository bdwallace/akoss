package teambition_aliyun

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

// CreateDevopsOrg invokes the teambition_aliyun.CreateDevopsOrg API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/createdevopsorg.html
func (client *Client) CreateDevopsOrg(request *CreateDevopsOrgRequest) (response *CreateDevopsOrgResponse, err error) {
	response = CreateCreateDevopsOrgResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDevopsOrgWithChan invokes the teambition_aliyun.CreateDevopsOrg API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/createdevopsorg.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDevopsOrgWithChan(request *CreateDevopsOrgRequest) (<-chan *CreateDevopsOrgResponse, <-chan error) {
	responseChan := make(chan *CreateDevopsOrgResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDevopsOrg(request)
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

// CreateDevopsOrgWithCallback invokes the teambition_aliyun.CreateDevopsOrg API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/createdevopsorg.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDevopsOrgWithCallback(request *CreateDevopsOrgRequest, callback func(response *CreateDevopsOrgResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDevopsOrgResponse
		var err error
		defer close(result)
		response, err = client.CreateDevopsOrg(request)
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

// CreateDevopsOrgRequest is the request struct for api CreateDevopsOrg
type CreateDevopsOrgRequest struct {
	*requests.RpcRequest
	OrgName            string           `position:"Body" name:"OrgName"`
	Source             string           `position:"Body" name:"Source"`
	RealPk             string           `position:"Body" name:"RealPk"`
	DesiredMemberCount requests.Integer `position:"Body" name:"DesiredMemberCount"`
}

// CreateDevopsOrgResponse is the response struct for api CreateDevopsOrg
type CreateDevopsOrgResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       string `json:"Object" xml:"Object"`
}

// CreateCreateDevopsOrgRequest creates a request to invoke CreateDevopsOrg API
func CreateCreateDevopsOrgRequest() (request *CreateDevopsOrgRequest) {
	request = &CreateDevopsOrgRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "CreateDevopsOrg", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDevopsOrgResponse creates a response to parse from CreateDevopsOrg response
func CreateCreateDevopsOrgResponse() (response *CreateDevopsOrgResponse) {
	response = &CreateDevopsOrgResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
