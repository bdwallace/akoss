package cloudcallcenter

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

// ListSkillGroupsOfUser invokes the cloudcallcenter.ListSkillGroupsOfUser API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listskillgroupsofuser.html
func (client *Client) ListSkillGroupsOfUser(request *ListSkillGroupsOfUserRequest) (response *ListSkillGroupsOfUserResponse, err error) {
	response = CreateListSkillGroupsOfUserResponse()
	err = client.DoAction(request, response)
	return
}

// ListSkillGroupsOfUserWithChan invokes the cloudcallcenter.ListSkillGroupsOfUser API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listskillgroupsofuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSkillGroupsOfUserWithChan(request *ListSkillGroupsOfUserRequest) (<-chan *ListSkillGroupsOfUserResponse, <-chan error) {
	responseChan := make(chan *ListSkillGroupsOfUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSkillGroupsOfUser(request)
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

// ListSkillGroupsOfUserWithCallback invokes the cloudcallcenter.ListSkillGroupsOfUser API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listskillgroupsofuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSkillGroupsOfUserWithCallback(request *ListSkillGroupsOfUserRequest, callback func(response *ListSkillGroupsOfUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSkillGroupsOfUserResponse
		var err error
		defer close(result)
		response, err = client.ListSkillGroupsOfUser(request)
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

// ListSkillGroupsOfUserRequest is the request struct for api ListSkillGroupsOfUser
type ListSkillGroupsOfUserRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	UserId     string `position:"Query" name:"UserId"`
}

// ListSkillGroupsOfUserResponse is the response struct for api ListSkillGroupsOfUser
type ListSkillGroupsOfUserResponse struct {
	*responses.BaseResponse
	RequestId      string                             `json:"RequestId" xml:"RequestId"`
	Success        bool                               `json:"Success" xml:"Success"`
	Code           string                             `json:"Code" xml:"Code"`
	Message        string                             `json:"Message" xml:"Message"`
	HttpStatusCode int                                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	SkillLevels    SkillLevelsInListSkillGroupsOfUser `json:"SkillLevels" xml:"SkillLevels"`
}

// CreateListSkillGroupsOfUserRequest creates a request to invoke ListSkillGroupsOfUser API
func CreateListSkillGroupsOfUserRequest() (request *ListSkillGroupsOfUserRequest) {
	request = &ListSkillGroupsOfUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListSkillGroupsOfUser", "", "")
	request.Method = requests.POST
	return
}

// CreateListSkillGroupsOfUserResponse creates a response to parse from ListSkillGroupsOfUser response
func CreateListSkillGroupsOfUserResponse() (response *ListSkillGroupsOfUserResponse) {
	response = &ListSkillGroupsOfUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
