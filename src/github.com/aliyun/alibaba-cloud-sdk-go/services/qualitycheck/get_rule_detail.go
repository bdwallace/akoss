package qualitycheck

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

// GetRuleDetail invokes the qualitycheck.GetRuleDetail API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/getruledetail.html
func (client *Client) GetRuleDetail(request *GetRuleDetailRequest) (response *GetRuleDetailResponse, err error) {
	response = CreateGetRuleDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetRuleDetailWithChan invokes the qualitycheck.GetRuleDetail API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getruledetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRuleDetailWithChan(request *GetRuleDetailRequest) (<-chan *GetRuleDetailResponse, <-chan error) {
	responseChan := make(chan *GetRuleDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRuleDetail(request)
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

// GetRuleDetailWithCallback invokes the qualitycheck.GetRuleDetail API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getruledetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRuleDetailWithCallback(request *GetRuleDetailRequest, callback func(response *GetRuleDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRuleDetailResponse
		var err error
		defer close(result)
		response, err = client.GetRuleDetail(request)
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

// GetRuleDetailRequest is the request struct for api GetRuleDetail
type GetRuleDetailRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetRuleDetailResponse is the response struct for api GetRuleDetail
type GetRuleDetailResponse struct {
	*responses.BaseResponse
	RequestId string              `json:"RequestId" xml:"RequestId"`
	Success   bool                `json:"Success" xml:"Success"`
	Code      string              `json:"Code" xml:"Code"`
	Message   string              `json:"Message" xml:"Message"`
	Data      DataInGetRuleDetail `json:"Data" xml:"Data"`
}

// CreateGetRuleDetailRequest creates a request to invoke GetRuleDetail API
func CreateGetRuleDetailRequest() (request *GetRuleDetailRequest) {
	request = &GetRuleDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetRuleDetail", "", "")
	return
}

// CreateGetRuleDetailResponse creates a response to parse from GetRuleDetail response
func CreateGetRuleDetailResponse() (response *GetRuleDetailResponse) {
	response = &GetRuleDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
