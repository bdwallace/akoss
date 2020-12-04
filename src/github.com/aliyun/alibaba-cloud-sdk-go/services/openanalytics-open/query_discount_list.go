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

// QueryDiscountList invokes the openanalytics_open.QueryDiscountList API synchronously
func (client *Client) QueryDiscountList(request *QueryDiscountListRequest) (response *QueryDiscountListResponse, err error) {
	response = CreateQueryDiscountListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDiscountListWithChan invokes the openanalytics_open.QueryDiscountList API asynchronously
func (client *Client) QueryDiscountListWithChan(request *QueryDiscountListRequest) (<-chan *QueryDiscountListResponse, <-chan error) {
	responseChan := make(chan *QueryDiscountListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDiscountList(request)
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

// QueryDiscountListWithCallback invokes the openanalytics_open.QueryDiscountList API asynchronously
func (client *Client) QueryDiscountListWithCallback(request *QueryDiscountListRequest, callback func(response *QueryDiscountListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDiscountListResponse
		var err error
		defer close(result)
		response, err = client.QueryDiscountList(request)
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

// QueryDiscountListRequest is the request struct for api QueryDiscountList
type QueryDiscountListRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
}

// QueryDiscountListResponse is the response struct for api QueryDiscountList
type QueryDiscountListResponse struct {
	*responses.BaseResponse
	RequestId  string                        `json:"RequestId" xml:"RequestId"`
	RegionId   string                        `json:"RegionId" xml:"RegionId"`
	PageSize   int                           `json:"PageSize" xml:"PageSize"`
	PageNumber int                           `json:"PageNumber" xml:"PageNumber"`
	TotalCount int                           `json:"TotalCount" xml:"TotalCount"`
	Data       []DataItemInQueryDiscountList `json:"Data" xml:"Data"`
}

// CreateQueryDiscountListRequest creates a request to invoke QueryDiscountList API
func CreateQueryDiscountListRequest() (request *QueryDiscountListRequest) {
	request = &QueryDiscountListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "QueryDiscountList", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDiscountListResponse creates a response to parse from QueryDiscountList response
func CreateQueryDiscountListResponse() (response *QueryDiscountListResponse) {
	response = &QueryDiscountListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
