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

// GetSurvey invokes the ccc.GetSurvey API synchronously
// api document: https://help.aliyun.com/api/ccc/getsurvey.html
func (client *Client) GetSurvey(request *GetSurveyRequest) (response *GetSurveyResponse, err error) {
	response = CreateGetSurveyResponse()
	err = client.DoAction(request, response)
	return
}

// GetSurveyWithChan invokes the ccc.GetSurvey API asynchronously
// api document: https://help.aliyun.com/api/ccc/getsurvey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSurveyWithChan(request *GetSurveyRequest) (<-chan *GetSurveyResponse, <-chan error) {
	responseChan := make(chan *GetSurveyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSurvey(request)
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

// GetSurveyWithCallback invokes the ccc.GetSurvey API asynchronously
// api document: https://help.aliyun.com/api/ccc/getsurvey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSurveyWithCallback(request *GetSurveyRequest, callback func(response *GetSurveyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSurveyResponse
		var err error
		defer close(result)
		response, err = client.GetSurvey(request)
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

// GetSurveyRequest is the request struct for api GetSurvey
type GetSurveyRequest struct {
	*requests.RpcRequest
	SurveyId   string `position:"Query" name:"SurveyId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	ScenarioId string `position:"Query" name:"ScenarioId"`
}

// GetSurveyResponse is the response struct for api GetSurvey
type GetSurveyResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Survey         Survey `json:"Survey" xml:"Survey"`
}

// CreateGetSurveyRequest creates a request to invoke GetSurvey API
func CreateGetSurveyRequest() (request *GetSurveyRequest) {
	request = &GetSurveyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetSurvey", "", "")
	return
}

// CreateGetSurveyResponse creates a response to parse from GetSurvey response
func CreateGetSurveyResponse() (response *GetSurveyResponse) {
	response = &GetSurveyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
