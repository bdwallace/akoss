package green

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

// DescribeKeyword invokes the green.DescribeKeyword API synchronously
func (client *Client) DescribeKeyword(request *DescribeKeywordRequest) (response *DescribeKeywordResponse, err error) {
	response = CreateDescribeKeywordResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeKeywordWithChan invokes the green.DescribeKeyword API asynchronously
func (client *Client) DescribeKeywordWithChan(request *DescribeKeywordRequest) (<-chan *DescribeKeywordResponse, <-chan error) {
	responseChan := make(chan *DescribeKeywordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeKeyword(request)
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

// DescribeKeywordWithCallback invokes the green.DescribeKeyword API asynchronously
func (client *Client) DescribeKeywordWithCallback(request *DescribeKeywordRequest, callback func(response *DescribeKeywordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeKeywordResponse
		var err error
		defer close(result)
		response, err = client.DescribeKeyword(request)
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

// DescribeKeywordRequest is the request struct for api DescribeKeyword
type DescribeKeywordRequest struct {
	*requests.RpcRequest
	SourceIp     string           `position:"Query" name:"SourceIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	Keyword      string           `position:"Query" name:"Keyword"`
	TotalCount   requests.Integer `position:"Query" name:"TotalCount"`
	KeywordLibId requests.Integer `position:"Query" name:"KeywordLibId"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
}

// DescribeKeywordResponse is the response struct for api DescribeKeyword
type DescribeKeywordResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	TotalCount  int       `json:"TotalCount" xml:"TotalCount"`
	PageSize    int       `json:"PageSize" xml:"PageSize"`
	CurrentPage int       `json:"CurrentPage" xml:"CurrentPage"`
	KeywordList []Keyword `json:"KeywordList" xml:"KeywordList"`
}

// CreateDescribeKeywordRequest creates a request to invoke DescribeKeyword API
func CreateDescribeKeywordRequest() (request *DescribeKeywordRequest) {
	request = &DescribeKeywordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeKeyword", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeKeywordResponse creates a response to parse from DescribeKeyword response
func CreateDescribeKeywordResponse() (response *DescribeKeywordResponse) {
	response = &DescribeKeywordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
