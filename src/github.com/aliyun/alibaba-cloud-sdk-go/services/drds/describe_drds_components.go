package drds

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

// DescribeDrdsComponents invokes the drds.DescribeDrdsComponents API synchronously
func (client *Client) DescribeDrdsComponents(request *DescribeDrdsComponentsRequest) (response *DescribeDrdsComponentsResponse, err error) {
	response = CreateDescribeDrdsComponentsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsComponentsWithChan invokes the drds.DescribeDrdsComponents API asynchronously
func (client *Client) DescribeDrdsComponentsWithChan(request *DescribeDrdsComponentsRequest) (<-chan *DescribeDrdsComponentsResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsComponentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsComponents(request)
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

// DescribeDrdsComponentsWithCallback invokes the drds.DescribeDrdsComponents API asynchronously
func (client *Client) DescribeDrdsComponentsWithCallback(request *DescribeDrdsComponentsRequest, callback func(response *DescribeDrdsComponentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsComponentsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsComponents(request)
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

// DescribeDrdsComponentsRequest is the request struct for api DescribeDrdsComponents
type DescribeDrdsComponentsRequest struct {
	*requests.RpcRequest
	CommodityCode string `position:"Query" name:"CommodityCode"`
}

// DescribeDrdsComponentsResponse is the response struct for api DescribeDrdsComponents
type DescribeDrdsComponentsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	DrdsComponents DrdsComponents `json:"DrdsComponents" xml:"DrdsComponents"`
}

// CreateDescribeDrdsComponentsRequest creates a request to invoke DescribeDrdsComponents API
func CreateDescribeDrdsComponentsRequest() (request *DescribeDrdsComponentsRequest) {
	request = &DescribeDrdsComponentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsComponents", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDrdsComponentsResponse creates a response to parse from DescribeDrdsComponents response
func CreateDescribeDrdsComponentsResponse() (response *DescribeDrdsComponentsResponse) {
	response = &DescribeDrdsComponentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}