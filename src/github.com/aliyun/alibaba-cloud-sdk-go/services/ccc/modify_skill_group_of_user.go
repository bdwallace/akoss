package ccc

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

// ModifySkillGroupOfUser invokes the ccc.ModifySkillGroupOfUser API synchronously
// api document: https://help.aliyun.com/api/ccc/modifyskillgroupofuser.html
func (client *Client) ModifySkillGroupOfUser(request *ModifySkillGroupOfUserRequest) (response *ModifySkillGroupOfUserResponse, err error) {
	response = CreateModifySkillGroupOfUserResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySkillGroupOfUserWithChan invokes the ccc.ModifySkillGroupOfUser API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyskillgroupofuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySkillGroupOfUserWithChan(request *ModifySkillGroupOfUserRequest) (<-chan *ModifySkillGroupOfUserResponse, <-chan error) {
	responseChan := make(chan *ModifySkillGroupOfUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySkillGroupOfUser(request)
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

// ModifySkillGroupOfUserWithCallback invokes the ccc.ModifySkillGroupOfUser API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyskillgroupofuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySkillGroupOfUserWithCallback(request *ModifySkillGroupOfUserRequest, callback func(response *ModifySkillGroupOfUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySkillGroupOfUserResponse
		var err error
		defer close(result)
		response, err = client.ModifySkillGroupOfUser(request)
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

// ModifySkillGroupOfUserRequest is the request struct for api ModifySkillGroupOfUser
type ModifySkillGroupOfUserRequest struct {
	*requests.RpcRequest
	RoleId       *[]string `position:"Query" name:"RoleId"  type:"Repeated"`
	UserId       string    `position:"Query" name:"UserId"`
	SkillLevel   *[]string `position:"Query" name:"SkillLevel"  type:"Repeated"`
	InstanceId   string    `position:"Query" name:"InstanceId"`
	SkillGroupId *[]string `position:"Query" name:"SkillGroupId"  type:"Repeated"`
}

// ModifySkillGroupOfUserResponse is the response struct for api ModifySkillGroupOfUser
type ModifySkillGroupOfUserResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateModifySkillGroupOfUserRequest creates a request to invoke ModifySkillGroupOfUser API
func CreateModifySkillGroupOfUserRequest() (request *ModifySkillGroupOfUserRequest) {
	request = &ModifySkillGroupOfUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifySkillGroupOfUser", "", "")
	return
}

// CreateModifySkillGroupOfUserResponse creates a response to parse from ModifySkillGroupOfUser response
func CreateModifySkillGroupOfUserResponse() (response *ModifySkillGroupOfUserResponse) {
	response = &ModifySkillGroupOfUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
