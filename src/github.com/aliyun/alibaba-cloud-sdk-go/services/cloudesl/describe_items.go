package cloudesl

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

// DescribeItems invokes the cloudesl.DescribeItems API synchronously
func (client *Client) DescribeItems(request *DescribeItemsRequest) (response *DescribeItemsResponse, err error) {
	response = CreateDescribeItemsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeItemsWithChan invokes the cloudesl.DescribeItems API asynchronously
func (client *Client) DescribeItemsWithChan(request *DescribeItemsRequest) (<-chan *DescribeItemsResponse, <-chan error) {
	responseChan := make(chan *DescribeItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeItems(request)
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

// DescribeItemsWithCallback invokes the cloudesl.DescribeItems API asynchronously
func (client *Client) DescribeItemsWithCallback(request *DescribeItemsRequest, callback func(response *DescribeItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeItemsResponse
		var err error
		defer close(result)
		response, err = client.DescribeItems(request)
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

// DescribeItemsRequest is the request struct for api DescribeItems
type DescribeItemsRequest struct {
	*requests.RpcRequest
	StoreId     string           `position:"Query" name:"StoreId"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	ItemId      requests.Integer `position:"Query" name:"ItemId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ItemBarCode string           `position:"Query" name:"ItemBarCode"`
	BePromotion requests.Boolean `position:"Query" name:"BePromotion"`
	ItemTitle   string           `position:"Query" name:"ItemTitle"`
	ShelfCode   string           `position:"Query" name:"ShelfCode"`
	SkuId       string           `position:"Query" name:"SkuId"`
}

// DescribeItemsResponse is the response struct for api DescribeItems
type DescribeItemsResponse struct {
	*responses.BaseResponse
	RequestId  string               `json:"RequestId" xml:"RequestId"`
	Success    bool                 `json:"Success" xml:"Success"`
	Message    string               `json:"Message" xml:"Message"`
	ErrorCode  string               `json:"ErrorCode" xml:"ErrorCode"`
	TotalCount int                  `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                  `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                  `json:"PageSize" xml:"PageSize"`
	Items      ItemsInDescribeItems `json:"Items" xml:"Items"`
}

// CreateDescribeItemsRequest creates a request to invoke DescribeItems API
func CreateDescribeItemsRequest() (request *DescribeItemsRequest) {
	request = &DescribeItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2018-08-01", "DescribeItems", "cloudesl", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeItemsResponse creates a response to parse from DescribeItems response
func CreateDescribeItemsResponse() (response *DescribeItemsResponse) {
	response = &DescribeItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}