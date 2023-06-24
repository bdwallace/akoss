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

// DescribeWebsiteInstance invokes the green.DescribeWebsiteInstance API synchronously
func (client *Client) DescribeWebsiteInstance(request *DescribeWebsiteInstanceRequest) (response *DescribeWebsiteInstanceResponse, err error) {
	response = CreateDescribeWebsiteInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebsiteInstanceWithChan invokes the green.DescribeWebsiteInstance API asynchronously
func (client *Client) DescribeWebsiteInstanceWithChan(request *DescribeWebsiteInstanceRequest) (<-chan *DescribeWebsiteInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeWebsiteInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebsiteInstance(request)
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

// DescribeWebsiteInstanceWithCallback invokes the green.DescribeWebsiteInstance API asynchronously
func (client *Client) DescribeWebsiteInstanceWithCallback(request *DescribeWebsiteInstanceRequest, callback func(response *DescribeWebsiteInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebsiteInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebsiteInstance(request)
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

// DescribeWebsiteInstanceRequest is the request struct for api DescribeWebsiteInstance
type DescribeWebsiteInstanceRequest struct {
	*requests.RpcRequest
	TotalCount  requests.Integer `position:"Query" name:"TotalCount"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Lang        string           `position:"Query" name:"Lang"`
}

// DescribeWebsiteInstanceResponse is the response struct for api DescribeWebsiteInstance
type DescribeWebsiteInstanceResponse struct {
	*responses.BaseResponse
	RequestId           string            `json:"RequestId" xml:"RequestId"`
	PageSize            int               `json:"PageSize" xml:"PageSize"`
	CurrentPage         int               `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount          int               `json:"TotalCount" xml:"TotalCount"`
	WebsiteInstanceList []WebsiteInstance `json:"WebsiteInstanceList" xml:"WebsiteInstanceList"`
}

// CreateDescribeWebsiteInstanceRequest creates a request to invoke DescribeWebsiteInstance API
func CreateDescribeWebsiteInstanceRequest() (request *DescribeWebsiteInstanceRequest) {
	request = &DescribeWebsiteInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeWebsiteInstance", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeWebsiteInstanceResponse creates a response to parse from DescribeWebsiteInstance response
func CreateDescribeWebsiteInstanceResponse() (response *DescribeWebsiteInstanceResponse) {
	response = &DescribeWebsiteInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
