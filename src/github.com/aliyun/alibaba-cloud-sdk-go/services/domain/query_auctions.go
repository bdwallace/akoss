package domain

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

// QueryAuctions invokes the domain.QueryAuctions API synchronously
// api document: https://help.aliyun.com/api/domain/queryauctions.html
func (client *Client) QueryAuctions(request *QueryAuctionsRequest) (response *QueryAuctionsResponse, err error) {
	response = CreateQueryAuctionsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAuctionsWithChan invokes the domain.QueryAuctions API asynchronously
// api document: https://help.aliyun.com/api/domain/queryauctions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAuctionsWithChan(request *QueryAuctionsRequest) (<-chan *QueryAuctionsResponse, <-chan error) {
	responseChan := make(chan *QueryAuctionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAuctions(request)
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

// QueryAuctionsWithCallback invokes the domain.QueryAuctions API asynchronously
// api document: https://help.aliyun.com/api/domain/queryauctions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAuctionsWithCallback(request *QueryAuctionsRequest, callback func(response *QueryAuctionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAuctionsResponse
		var err error
		defer close(result)
		response, err = client.QueryAuctions(request)
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

// QueryAuctionsRequest is the request struct for api QueryAuctions
type QueryAuctionsRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	CurrentPage requests.Integer `position:"Body" name:"CurrentPage"`
	Status      string           `position:"Body" name:"Status"`
}

// QueryAuctionsResponse is the response struct for api QueryAuctions
type QueryAuctionsResponse struct {
	*responses.BaseResponse
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int             `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int             `json:"CurrentPageNum" xml:"CurrentPageNum"`
	PageSize       int             `json:"PageSize" xml:"PageSize"`
	TotalPageNum   int             `json:"TotalPageNum" xml:"TotalPageNum"`
	Data           []AuctionDetail `json:"Data" xml:"Data"`
}

// CreateQueryAuctionsRequest creates a request to invoke QueryAuctions API
func CreateQueryAuctionsRequest() (request *QueryAuctionsRequest) {
	request = &QueryAuctionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "QueryAuctions", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryAuctionsResponse creates a response to parse from QueryAuctions response
func CreateQueryAuctionsResponse() (response *QueryAuctionsResponse) {
	response = &QueryAuctionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
