package codeup

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

// GetCodeupOrganization invokes the codeup.GetCodeupOrganization API synchronously
func (client *Client) GetCodeupOrganization(request *GetCodeupOrganizationRequest) (response *GetCodeupOrganizationResponse, err error) {
	response = CreateGetCodeupOrganizationResponse()
	err = client.DoAction(request, response)
	return
}

// GetCodeupOrganizationWithChan invokes the codeup.GetCodeupOrganization API asynchronously
func (client *Client) GetCodeupOrganizationWithChan(request *GetCodeupOrganizationRequest) (<-chan *GetCodeupOrganizationResponse, <-chan error) {
	responseChan := make(chan *GetCodeupOrganizationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCodeupOrganization(request)
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

// GetCodeupOrganizationWithCallback invokes the codeup.GetCodeupOrganization API asynchronously
func (client *Client) GetCodeupOrganizationWithCallback(request *GetCodeupOrganizationRequest, callback func(response *GetCodeupOrganizationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCodeupOrganizationResponse
		var err error
		defer close(result)
		response, err = client.GetCodeupOrganization(request)
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

// GetCodeupOrganizationRequest is the request struct for api GetCodeupOrganization
type GetCodeupOrganizationRequest struct {
	*requests.RoaRequest
	OrganizationId       string `position:"Query" name:"OrganizationId"`
	SubUserId            string `position:"Query" name:"SubUserId"`
	OrganizationIdentity string `position:"Path" name:"OrganizationIdentity"`
	AccessToken          string `position:"Query" name:"AccessToken"`
}

// GetCodeupOrganizationResponse is the response struct for api GetCodeupOrganization
type GetCodeupOrganizationResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateGetCodeupOrganizationRequest creates a request to invoke GetCodeupOrganization API
func CreateGetCodeupOrganizationRequest() (request *GetCodeupOrganizationRequest) {
	request = &GetCodeupOrganizationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "GetCodeupOrganization", "/api/v4/organization/[OrganizationIdentity]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetCodeupOrganizationResponse creates a response to parse from GetCodeupOrganization response
func CreateGetCodeupOrganizationResponse() (response *GetCodeupOrganizationResponse) {
	response = &GetCodeupOrganizationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
