package retailcloud

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

// DescribePodLog invokes the retailcloud.DescribePodLog API synchronously
// api document: https://help.aliyun.com/api/retailcloud/describepodlog.html
func (client *Client) DescribePodLog(request *DescribePodLogRequest) (response *DescribePodLogResponse, err error) {
	response = CreateDescribePodLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePodLogWithChan invokes the retailcloud.DescribePodLog API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describepodlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePodLogWithChan(request *DescribePodLogRequest) (<-chan *DescribePodLogResponse, <-chan error) {
	responseChan := make(chan *DescribePodLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePodLog(request)
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

// DescribePodLogWithCallback invokes the retailcloud.DescribePodLog API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describepodlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePodLogWithCallback(request *DescribePodLogRequest, callback func(response *DescribePodLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePodLogResponse
		var err error
		defer close(result)
		response, err = client.DescribePodLog(request)
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

// DescribePodLogRequest is the request struct for api DescribePodLog
type DescribePodLogRequest struct {
	*requests.RpcRequest
	DeployOrderId requests.Integer `position:"Body" name:"DeployOrderId"`
	AppInstId     string           `position:"Body" name:"AppInstId"`
}

// DescribePodLogResponse is the response struct for api DescribePodLog
type DescribePodLogResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribePodLogRequest creates a request to invoke DescribePodLog API
func CreateDescribePodLogRequest() (request *DescribePodLogRequest) {
	request = &DescribePodLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DescribePodLog", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePodLogResponse creates a response to parse from DescribePodLog response
func CreateDescribePodLogResponse() (response *DescribePodLogResponse) {
	response = &DescribePodLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
