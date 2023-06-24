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

// GetJobStatusByCallId invokes the cloudcallcenter.GetJobStatusByCallId API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getjobstatusbycallid.html
func (client *Client) GetJobStatusByCallId(request *GetJobStatusByCallIdRequest) (response *GetJobStatusByCallIdResponse, err error) {
	response = CreateGetJobStatusByCallIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetJobStatusByCallIdWithChan invokes the cloudcallcenter.GetJobStatusByCallId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getjobstatusbycallid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJobStatusByCallIdWithChan(request *GetJobStatusByCallIdRequest) (<-chan *GetJobStatusByCallIdResponse, <-chan error) {
	responseChan := make(chan *GetJobStatusByCallIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJobStatusByCallId(request)
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

// GetJobStatusByCallIdWithCallback invokes the cloudcallcenter.GetJobStatusByCallId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getjobstatusbycallid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJobStatusByCallIdWithCallback(request *GetJobStatusByCallIdRequest, callback func(response *GetJobStatusByCallIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJobStatusByCallIdResponse
		var err error
		defer close(result)
		response, err = client.GetJobStatusByCallId(request)
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

// GetJobStatusByCallIdRequest is the request struct for api GetJobStatusByCallId
type GetJobStatusByCallIdRequest struct {
	*requests.RpcRequest
	CallId     string `position:"Query" name:"CallId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetJobStatusByCallIdResponse is the response struct for api GetJobStatusByCallId
type GetJobStatusByCallIdResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Job            Job    `json:"Job" xml:"Job"`
}

// CreateGetJobStatusByCallIdRequest creates a request to invoke GetJobStatusByCallId API
func CreateGetJobStatusByCallIdRequest() (request *GetJobStatusByCallIdRequest) {
	request = &GetJobStatusByCallIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetJobStatusByCallId", "", "")
	request.Method = requests.POST
	return
}

// CreateGetJobStatusByCallIdResponse creates a response to parse from GetJobStatusByCallId response
func CreateGetJobStatusByCallIdResponse() (response *GetJobStatusByCallIdResponse) {
	response = &GetJobStatusByCallIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
