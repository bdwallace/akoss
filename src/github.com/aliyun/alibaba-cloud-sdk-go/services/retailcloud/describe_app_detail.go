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

// DescribeAppDetail invokes the retailcloud.DescribeAppDetail API synchronously
// api document: https://help.aliyun.com/api/retailcloud/describeappdetail.html
func (client *Client) DescribeAppDetail(request *DescribeAppDetailRequest) (response *DescribeAppDetailResponse, err error) {
	response = CreateDescribeAppDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppDetailWithChan invokes the retailcloud.DescribeAppDetail API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describeappdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppDetailWithChan(request *DescribeAppDetailRequest) (<-chan *DescribeAppDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeAppDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAppDetail(request)
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

// DescribeAppDetailWithCallback invokes the retailcloud.DescribeAppDetail API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describeappdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAppDetailWithCallback(request *DescribeAppDetailRequest, callback func(response *DescribeAppDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeAppDetail(request)
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

// DescribeAppDetailRequest is the request struct for api DescribeAppDetail
type DescribeAppDetailRequest struct {
	*requests.RpcRequest
	AppId requests.Integer `position:"Query" name:"AppId"`
}

// DescribeAppDetailResponse is the response struct for api DescribeAppDetail
type DescribeAppDetailResponse struct {
	*responses.BaseResponse
	Code       int64  `json:"Code" xml:"Code"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	Result     Result `json:"Result" xml:"Result"`
}

// CreateDescribeAppDetailRequest creates a request to invoke DescribeAppDetail API
func CreateDescribeAppDetailRequest() (request *DescribeAppDetailRequest) {
	request = &DescribeAppDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DescribeAppDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeAppDetailResponse creates a response to parse from DescribeAppDetail response
func CreateDescribeAppDetailResponse() (response *DescribeAppDetailResponse) {
	response = &DescribeAppDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
