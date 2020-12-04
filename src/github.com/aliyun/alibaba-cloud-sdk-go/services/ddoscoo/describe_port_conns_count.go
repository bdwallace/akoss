package ddoscoo

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

// DescribePortConnsCount invokes the ddoscoo.DescribePortConnsCount API synchronously
func (client *Client) DescribePortConnsCount(request *DescribePortConnsCountRequest) (response *DescribePortConnsCountResponse, err error) {
	response = CreateDescribePortConnsCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePortConnsCountWithChan invokes the ddoscoo.DescribePortConnsCount API asynchronously
func (client *Client) DescribePortConnsCountWithChan(request *DescribePortConnsCountRequest) (<-chan *DescribePortConnsCountResponse, <-chan error) {
	responseChan := make(chan *DescribePortConnsCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePortConnsCount(request)
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

// DescribePortConnsCountWithCallback invokes the ddoscoo.DescribePortConnsCount API asynchronously
func (client *Client) DescribePortConnsCountWithCallback(request *DescribePortConnsCountRequest, callback func(response *DescribePortConnsCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePortConnsCountResponse
		var err error
		defer close(result)
		response, err = client.DescribePortConnsCount(request)
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

// DescribePortConnsCountRequest is the request struct for api DescribePortConnsCount
type DescribePortConnsCountRequest struct {
	*requests.RpcRequest
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	InstanceIds     *[]string        `position:"Query" name:"InstanceIds"  type:"Repeated"`
	Port            string           `position:"Query" name:"Port"`
}

// DescribePortConnsCountResponse is the response struct for api DescribePortConnsCount
type DescribePortConnsCountResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Conns      int64  `json:"Conns" xml:"Conns"`
	ActConns   int64  `json:"ActConns" xml:"ActConns"`
	InActConns int64  `json:"InActConns" xml:"InActConns"`
	Cps        int64  `json:"Cps" xml:"Cps"`
}

// CreateDescribePortConnsCountRequest creates a request to invoke DescribePortConnsCount API
func CreateDescribePortConnsCountRequest() (request *DescribePortConnsCountRequest) {
	request = &DescribePortConnsCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribePortConnsCount", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePortConnsCountResponse creates a response to parse from DescribePortConnsCount response
func CreateDescribePortConnsCountResponse() (response *DescribePortConnsCountResponse) {
	response = &DescribePortConnsCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
