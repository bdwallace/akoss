package devops_rdc

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

// UpdateDevopsProject invokes the devops_rdc.UpdateDevopsProject API synchronously
func (client *Client) UpdateDevopsProject(request *UpdateDevopsProjectRequest) (response *UpdateDevopsProjectResponse, err error) {
	response = CreateUpdateDevopsProjectResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDevopsProjectWithChan invokes the devops_rdc.UpdateDevopsProject API asynchronously
func (client *Client) UpdateDevopsProjectWithChan(request *UpdateDevopsProjectRequest) (<-chan *UpdateDevopsProjectResponse, <-chan error) {
	responseChan := make(chan *UpdateDevopsProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDevopsProject(request)
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

// UpdateDevopsProjectWithCallback invokes the devops_rdc.UpdateDevopsProject API asynchronously
func (client *Client) UpdateDevopsProjectWithCallback(request *UpdateDevopsProjectRequest, callback func(response *UpdateDevopsProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDevopsProjectResponse
		var err error
		defer close(result)
		response, err = client.UpdateDevopsProject(request)
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

// UpdateDevopsProjectRequest is the request struct for api UpdateDevopsProject
type UpdateDevopsProjectRequest struct {
	*requests.RpcRequest
	Name        string `position:"Body" name:"Name"`
	Description string `position:"Body" name:"Description"`
	ProjectId   string `position:"Body" name:"ProjectId"`
	OrgId       string `position:"Body" name:"OrgId"`
}

// UpdateDevopsProjectResponse is the response struct for api UpdateDevopsProject
type UpdateDevopsProjectResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       string `json:"Object" xml:"Object"`
}

// CreateUpdateDevopsProjectRequest creates a request to invoke UpdateDevopsProject API
func CreateUpdateDevopsProjectRequest() (request *UpdateDevopsProjectRequest) {
	request = &UpdateDevopsProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "UpdateDevopsProject", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateDevopsProjectResponse creates a response to parse from UpdateDevopsProject response
func CreateUpdateDevopsProjectResponse() (response *UpdateDevopsProjectResponse) {
	response = &UpdateDevopsProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
