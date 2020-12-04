package hbase

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

// DescribeMultiZoneAvailableResource invokes the hbase.DescribeMultiZoneAvailableResource API synchronously
func (client *Client) DescribeMultiZoneAvailableResource(request *DescribeMultiZoneAvailableResourceRequest) (response *DescribeMultiZoneAvailableResourceResponse, err error) {
	response = CreateDescribeMultiZoneAvailableResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMultiZoneAvailableResourceWithChan invokes the hbase.DescribeMultiZoneAvailableResource API asynchronously
func (client *Client) DescribeMultiZoneAvailableResourceWithChan(request *DescribeMultiZoneAvailableResourceRequest) (<-chan *DescribeMultiZoneAvailableResourceResponse, <-chan error) {
	responseChan := make(chan *DescribeMultiZoneAvailableResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMultiZoneAvailableResource(request)
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

// DescribeMultiZoneAvailableResourceWithCallback invokes the hbase.DescribeMultiZoneAvailableResource API asynchronously
func (client *Client) DescribeMultiZoneAvailableResourceWithCallback(request *DescribeMultiZoneAvailableResourceRequest, callback func(response *DescribeMultiZoneAvailableResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMultiZoneAvailableResourceResponse
		var err error
		defer close(result)
		response, err = client.DescribeMultiZoneAvailableResource(request)
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

// DescribeMultiZoneAvailableResourceRequest is the request struct for api DescribeMultiZoneAvailableResource
type DescribeMultiZoneAvailableResourceRequest struct {
	*requests.RpcRequest
	ZoneCombination string `position:"Query" name:"ZoneCombination"`
	ChargeType      string `position:"Query" name:"ChargeType"`
}

// DescribeMultiZoneAvailableResourceResponse is the response struct for api DescribeMultiZoneAvailableResource
type DescribeMultiZoneAvailableResourceResponse struct {
	*responses.BaseResponse
	RequestId      string                                             `json:"RequestId" xml:"RequestId"`
	AvailableZones AvailableZonesInDescribeMultiZoneAvailableResource `json:"AvailableZones" xml:"AvailableZones"`
}

// CreateDescribeMultiZoneAvailableResourceRequest creates a request to invoke DescribeMultiZoneAvailableResource API
func CreateDescribeMultiZoneAvailableResourceRequest() (request *DescribeMultiZoneAvailableResourceRequest) {
	request = &DescribeMultiZoneAvailableResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeMultiZoneAvailableResource", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMultiZoneAvailableResourceResponse creates a response to parse from DescribeMultiZoneAvailableResource response
func CreateDescribeMultiZoneAvailableResourceResponse() (response *DescribeMultiZoneAvailableResourceResponse) {
	response = &DescribeMultiZoneAvailableResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
