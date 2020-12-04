package outboundbot

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

// WithdrawScriptReview invokes the outboundbot.WithdrawScriptReview API synchronously
// api document: https://help.aliyun.com/api/outboundbot/withdrawscriptreview.html
func (client *Client) WithdrawScriptReview(request *WithdrawScriptReviewRequest) (response *WithdrawScriptReviewResponse, err error) {
	response = CreateWithdrawScriptReviewResponse()
	err = client.DoAction(request, response)
	return
}

// WithdrawScriptReviewWithChan invokes the outboundbot.WithdrawScriptReview API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/withdrawscriptreview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) WithdrawScriptReviewWithChan(request *WithdrawScriptReviewRequest) (<-chan *WithdrawScriptReviewResponse, <-chan error) {
	responseChan := make(chan *WithdrawScriptReviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.WithdrawScriptReview(request)
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

// WithdrawScriptReviewWithCallback invokes the outboundbot.WithdrawScriptReview API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/withdrawscriptreview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) WithdrawScriptReviewWithCallback(request *WithdrawScriptReviewRequest, callback func(response *WithdrawScriptReviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *WithdrawScriptReviewResponse
		var err error
		defer close(result)
		response, err = client.WithdrawScriptReview(request)
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

// WithdrawScriptReviewRequest is the request struct for api WithdrawScriptReview
type WithdrawScriptReviewRequest struct {
	*requests.RpcRequest
	ScriptId   string `position:"Query" name:"ScriptId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// WithdrawScriptReviewResponse is the response struct for api WithdrawScriptReview
type WithdrawScriptReviewResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Script         Script `json:"Script" xml:"Script"`
}

// CreateWithdrawScriptReviewRequest creates a request to invoke WithdrawScriptReview API
func CreateWithdrawScriptReviewRequest() (request *WithdrawScriptReviewRequest) {
	request = &WithdrawScriptReviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "WithdrawScriptReview", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateWithdrawScriptReviewResponse creates a response to parse from WithdrawScriptReview response
func CreateWithdrawScriptReviewResponse() (response *WithdrawScriptReviewResponse) {
	response = &WithdrawScriptReviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
