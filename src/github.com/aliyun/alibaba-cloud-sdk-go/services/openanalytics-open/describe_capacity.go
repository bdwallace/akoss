package openanalytics_open

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

// DescribeCapacity invokes the openanalytics_open.DescribeCapacity API synchronously
func (client *Client) DescribeCapacity(request *DescribeCapacityRequest) (response *DescribeCapacityResponse, err error) {
	response = CreateDescribeCapacityResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCapacityWithChan invokes the openanalytics_open.DescribeCapacity API asynchronously
func (client *Client) DescribeCapacityWithChan(request *DescribeCapacityRequest) (<-chan *DescribeCapacityResponse, <-chan error) {
	responseChan := make(chan *DescribeCapacityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCapacity(request)
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

// DescribeCapacityWithCallback invokes the openanalytics_open.DescribeCapacity API asynchronously
func (client *Client) DescribeCapacityWithCallback(request *DescribeCapacityRequest, callback func(response *DescribeCapacityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCapacityResponse
		var err error
		defer close(result)
		response, err = client.DescribeCapacity(request)
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

// DescribeCapacityRequest is the request struct for api DescribeCapacity
type DescribeCapacityRequest struct {
	*requests.RpcRequest
	ExternalBizAliyunId string `position:"Body" name:"ExternalBizAliyunId"`
}

// DescribeCapacityResponse is the response struct for api DescribeCapacity
type DescribeCapacityResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	RegionId  string `json:"RegionId" xml:"RegionId"`
	Storage   int    `json:"Storage" xml:"Storage"`
}

// CreateDescribeCapacityRequest creates a request to invoke DescribeCapacity API
func CreateDescribeCapacityRequest() (request *DescribeCapacityRequest) {
	request = &DescribeCapacityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "DescribeCapacity", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCapacityResponse creates a response to parse from DescribeCapacity response
func CreateDescribeCapacityResponse() (response *DescribeCapacityResponse) {
	response = &DescribeCapacityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
